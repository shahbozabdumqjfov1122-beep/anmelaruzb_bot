package ko_di

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"namelaruzb_bot/kodi/Help"
	"namelaruzb_bot/kodi/Menu"
	"namelaruzb_bot/kodi/anmelaruzb"
	"namelaruzb_bot/kodi/anmelaruzb/default"
	"sort"

	"os"
	"sync"
	"time"

	tele "gopkg.in/telebot.v4"
)

var (
	pendingRequests = make(map[int64]map[int64]bool)
	requestMutex    sync.RWMutex

	// Reklama uchun yangi o'zgaruvchilar
	adminState     = make(map[int64]string)        // Admin holatini saqlash
	adminWaitingAd = make(map[int64]*tele.Message) // Yuborilishi kerak bo'lgan xabarni saqlash
)

type ChannelInfo struct {
	ID     int64
	Name   string
	Invite string
}

var myChannels = []ChannelInfo{
	{ID: -1003050934981, Name: "anmelaruzb", Invite: "https://t.me/anmelaruzb"},
	//{ID: -1003316396409, Name: "anmelar_chat", Invite: "https://t.me/anmelar_chat"},
	{ID: -1003323161290, Name: "Manga Uzb", Invite: "https://t.me/Manga_uzbekcha26"},
	{ID: -1003411861509, Name: "Maxfiy Kanal", Invite: "https://t.me/+C0qmcf4ZHY83NmNi"},
}

var (
	userActive  = make(map[int64]time.Time)
	userJoined  = make(map[int64]time.Time)
	statsMutex  sync.RWMutex
	searchStats = make(map[string]int)
)

func updateUserActivity(userID int64) {
	statsMutex.Lock()
	now := time.Now()
	userActive[userID] = now

	isNew := false
	if _, ok := userJoined[userID]; !ok {
		userJoined[userID] = now
		isNew = true
	}
	statsMutex.Unlock()

	if isNew {
		saveStats()
	}
}

func saveStats() {
	statsMutex.RLock()
	data := struct {
		UserJoined  map[int64]time.Time `json:"userJoined"`
		UserActive  map[int64]time.Time `json:"userActive"`
		SearchStats map[string]int      `json:"searchStats"`
	}{
		UserJoined:  userJoined,
		UserActive:  userActive,
		SearchStats: searchStats,
	}
	statsMutex.RUnlock()

	file, _ := json.MarshalIndent(data, "", "  ")
	_ = os.WriteFile("stats.json", file, 0644)
}

func loadStats() {
	file, err := os.ReadFile("stats.json")
	if err != nil {
		return
	}
	var data struct {
		UserJoined  map[int64]time.Time `json:"userJoined"`
		UserActive  map[int64]time.Time `json:"userActive"`
		SearchStats map[string]int      `json:"searchStats"`
	}
	if err := json.Unmarshal(file, &data); err == nil {
		statsMutex.Lock()
		userJoined = data.UserJoined
		userActive = data.UserActive
		searchStats = data.SearchStats
		statsMutex.Unlock()
	}
}

var admins = map[int64]bool{7518992824: true}

func isAdmin(userID int64) bool { return admins[userID] }

func notAllowedChannels(b *tele.Bot, userID int64) []ChannelInfo {
	var missing []ChannelInfo
	for _, ch := range myChannels {
		chat := &tele.Chat{ID: ch.ID}
		member, err := b.ChatMemberOf(chat, &tele.User{ID: userID})
		if err == nil && (member.Role == tele.Member || member.Role == tele.Administrator || member.Role == tele.Creator) {
			continue
		}
		requestMutex.RLock()
		userReqs := pendingRequests[userID]
		hasRequested := userReqs != nil && userReqs[ch.ID]
		requestMutex.RUnlock()
		if !hasRequested {
			missing = append(missing, ch)
		}
	}
	return missing
}

func Bot() {
	loadStats()
	token := beego.AppConfig.DefaultString("telegram::token", "")
	b, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	// --- MENYULAR ---
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	menu.Reply(menu.Row(menu.Text("Animelar")), menu.Row(menu.Text("🧩 help")))

	adminMenu := &tele.ReplyMarkup{}
	btnBroadcast := adminMenu.Data("📢 Reklama yuborish", "admin_broadcast")
	btnStats := adminMenu.Data("📊 Statistika", "admin_stats")

	// --- ADMIN HANDLERLARI ---
	b.Handle("/admin", func(c tele.Context) error {
		if !isAdmin(c.Sender().ID) {
			return nil
		}
		adminMenu.Inline(adminMenu.Row(btnBroadcast, btnStats))
		return c.Send("👨‍💻 **Admin Boshqaruv Paneli:**", adminMenu, tele.ModeMarkdown)
	})

	b.Handle(&btnBroadcast, func(c tele.Context) error {
		adminState[c.Sender().ID] = "waiting_for_ad"
		return c.Send("📥 Reklama xabarini yuboring. (Rasm, video yoki matn)")
	})

	b.Handle(&btnStats, func(c tele.Context) error {
		return sendStatistics(c)
	})

	// Reklamani tasdiqlash
	b.Handle(&tele.Btn{Unique: "confirm_ad"}, func(c tele.Context) error {
		adMsg := adminWaitingAd[c.Sender().ID]
		if adMsg == nil {
			return c.Edit("❌ Xabar topilmadi.")
		}

		go func() {
			count := 0
			statsMutex.RLock()
			for userID := range userJoined {
				_, err := b.Copy(tele.ChatID(userID), adMsg)
				if err == nil {
					count++
				}
				time.Sleep(33 * time.Millisecond)
			}
			statsMutex.RUnlock()
			b.Send(c.Sender(), fmt.Sprintf("✅ Reklama tugatildi! %d kishiga yuborildi.", count))
		}()
		delete(adminState, c.Sender().ID)
		return c.Edit("🚀 Reklama tarqatish boshlandi...")
	})

	b.Handle(&tele.Btn{Unique: "cancel_ad"}, func(c tele.Context) error {
		delete(adminState, c.Sender().ID)
		return c.Edit("❌ Reklama bekor qilindi.")
	})

	// --- ASOSIY XABARLARNI QABUL QILISH ---
	handleAll := func(c tele.Context) error {
		updateUserActivity(c.Sender().ID)

		// 1. Admin reklama yuborayotgan bo'lsa
		if isAdmin(c.Sender().ID) && adminState[c.Sender().ID] == "waiting_for_ad" {
			adminWaitingAd[c.Sender().ID] = c.Message()
			m := &tele.ReplyMarkup{}
			btnYes := m.Data("✅ Tasdiqlash", "confirm_ad")
			btnNo := m.Data("❌ Bekor qilish", "cancel_ad")
			m.Inline(m.Row(btnYes, btnNo))

			_ = c.Send("👇 **Reklama ko'rinishi:**")
			_, _ = b.Copy(c.Recipient(), c.Message())
			return c.Send("Ushbu xabarni yuboramizmi?", m)
		}

		// 2. Obuna tekshiruvi (Admin bo'lmasa)
		if !isAdmin(c.Sender().ID) {
			missing := notAllowedChannels(b, c.Sender().ID)
			if len(missing) > 0 {
				return sendSubMessage(c, missing)
			}
		}

		// 3. Matnli buyruqlar
		text := c.Text()
		if text != "" {
			switch text {
			case "/start":
				return c.Send("✅ Xush kelibsiz!", menu)
			case "Animelar", "/menu":
				return Menu.Home(c)
			case "🧩 help":
				return Help.Home(c)
			default:
				return anmelaruzb.Home(c)
			}
		}
		return _default.Home(c)
	}

	// Faqat shu ikkita handler yetarli:
	b.Handle(tele.OnText, handleAll)
	b.Handle(tele.OnMedia, handleAll)

	// Join request uchun
	b.Handle(tele.OnChatJoinRequest, func(c tele.Context) error {
		req := c.ChatJoinRequest()
		requestMutex.Lock()
		if pendingRequests[req.Sender.ID] == nil {
			pendingRequests[req.Sender.ID] = make(map[int64]bool)
		}
		pendingRequests[req.Sender.ID][req.Chat.ID] = true
		requestMutex.Unlock()
		return nil
	})

	// Tekshirish tugmasi
	b.Handle(&tele.Btn{Unique: "check_sub"}, func(c tele.Context) error {
		missing := notAllowedChannels(b, c.Sender().ID)
		if len(missing) == 0 {
			_ = c.Delete()
			return c.Send("✅ Rahmat! Botdan foydalanishingiz mumkin.", menu)
		}
		return c.Respond(&tele.CallbackResponse{Text: "❌ Hali hamma kanal bajarilmadi", ShowAlert: true})
	})

	log.Println("🤖 Bot ishga tushdi")
	b.Start()
}

func sendSubMessage(c tele.Context, missing []ChannelInfo) error {
	text := "<b>❗ Botdan foydalanish uchun quyidagi kanallarga a'zo bo‘ling yoki so‘rov yuboring:</b>"
	m := &tele.ReplyMarkup{}
	var rows []tele.Row
	for _, ch := range missing {
		rows = append(rows, m.Row(m.URL("📢 "+ch.Name, ch.Invite)))
	}
	rows = append(rows, m.Row(m.Data("✅ Tekshirish", "check_sub")))
	m.Inline(rows...)
	return c.Send(text, m, tele.ModeHTML)
}

func handleAllMessages(b *tele.Bot, c tele.Context) error {
	updateUserActivity(c.Sender().ID)

	// ADMIN REKLAMA KUTAYOTGAN BO'LSA
	if isAdmin(c.Sender().ID) && adminState[c.Sender().ID] == "waiting_for_ad" {
		adminWaitingAd[c.Sender().ID] = c.Message()

		confirmMarkup := &tele.ReplyMarkup{}
		btnConfirm := confirmMarkup.Data("✅ Tasdiqlash", "confirm_ad")
		btnCancel := confirmMarkup.Data("❌ Bekor qilish", "cancel_ad")
		confirmMarkup.Inline(confirmMarkup.Row(btnConfirm, btnCancel))

		_ = c.Send("👇 **Reklama ko'rinishi:**")
		_, _ = b.Copy(c.Recipient(), c.Message())
		return c.Send("Yuqoridagi xabarni hamma foydalanuvchilarga yuboramizmi?", confirmMarkup)
	}

	// OBUNA TEKSHIRUVI
	missing := notAllowedChannels(b, c.Sender().ID)
	if len(missing) > 0 {
		return sendSubMessage(c, missing)
	}

	// BUYRUQLAR (Faqat matn bo'lsa ishlaydi)
	text := c.Text()
	if text != "" {
		switch text {
		case "Animelar", "/menu":
			return Menu.Home(c)
		case "🧩 help":
			return Help.Home(c)
		default:
			return anmelaruzb.Home(c)
		}
	}
	return nil
}

func sendStatistics(c tele.Context) error {
	statsMutex.RLock()
	defer statsMutex.RUnlock()

	loc := time.FixedZone("Asia/Tashkent", 5*3600)
	now := time.Now().In(loc)

	// Vaqt chegaralari
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	startOf7Days := startOfToday.AddDate(0, 0, -6)
	startOf30Days := startOfToday.AddDate(0, 0, -29)

	// Foydalanuvchi hisoblash
	total := len(userJoined)
	active := 0
	for _, t := range userActive {
		if now.Sub(t) < 24*time.Hour {
			active++
		}
	}
	inactive := total - active

	todayNew, last7, last30 := 0, 0, 0
	for _, t := range userJoined {
		tInLoc := t.In(loc)
		if tInLoc.After(startOfToday) {
			todayNew++
		}
		if tInLoc.After(startOf7Days) {
			last7++
		}
		if tInLoc.After(startOf30Days) {
			last30++
		}
	}

	// Top 5 Anime saralash
	type kv struct {
		Key   string
		Value int
	}
	var searchList []kv
	for k, v := range searchStats {
		searchList = append(searchList, kv{k, v})
	}
	sort.Slice(searchList, func(i, j int) bool {
		return searchList[i].Value > searchList[j].Value
	})

	topCount := 5
	if len(searchList) < 5 {
		topCount = len(searchList)
	}

	// XABARNI SHAKLLANTIRISH
	text := "------------------------------------------\n"
	text += "🏆 **ENG MASHHUR 5 ANIME:**\n"
	if topCount == 0 {
		text += "ℹ️ Hozircha ma'lumot yo'q\n"
	} else {
		for i := 0; i < topCount; i++ {
			text += fmt.Sprintf("%d. %s — %d qidiruv\n", i+1, searchList[i].Key, searchList[i].Value)
		}
	}

	text += "------------------------------------------\n"
	text += "🔗 **KANAL OBUNALARI:**\n"
	for _, ch := range myChannels {
		text += fmt.Sprintf("✅ [Kanalga o'tish](%s)\n", ch.Invite)
	}

	text += "------------------------------------------\n"
	text += "📊 **FOYDALANUVCHILAR:**\n"
	text += fmt.Sprintf("🟢 Faol foydalanuvchilar: %d\n", active)
	text += fmt.Sprintf("🚫 Nofaol foydalanuvchilar: %d\n", inactive)
	text += fmt.Sprintf("👥 Umumiy obunachilar: %d\n", total)

	text += "\n🆕 **OBUNACHILAR (YANGI):**\n"
	text += fmt.Sprintf("📅 Bugun: %d\n", todayNew)
	text += fmt.Sprintf("🗓 7 kunlik: %d\n", last7)
	text += fmt.Sprintf("🗓 30 kunlik: %d\n", last30)

	text += "------------------------------------------\n"
	text += fmt.Sprintf("ℹ️ Ma'lumotlar yangilangan: \n`%s`", now.Format("02.01.2006 | 15:04:05"))

	return c.Send(text, tele.ModeMarkdown, tele.NoPreview)
}
