package kodi

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"namelaruzb_bot/kodi/Help"
	"namelaruzb_bot/kodi/Menu"
	"namelaruzb_bot/kodi/anmelaruzb"
	"sort"
	"strconv"
	"strings"

	"net/url"
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
var (
	scheduledPosts = make(map[int]*ScheduledPost)
	scheduleMutex  sync.RWMutex
	scheduleAutoID = 1
)

type ChannelInfo struct {
	ID     int64
	Name   string
	Invite string
}

type ScheduledPost struct {
	ID      int
	AdminID int64 // Kim rejalashtirganini bilish uchun
	Content ContentItem

	SendTime time.Time
	ChatIDs  []int64
}

type ContentItem struct {
	Kind   string // text | photo | video
	FileID string // rasm/video uchun
	Text   string // matn yoki caption
}

var myChannels = []ChannelInfo{
	{ID: -1003050934981, Name: "anmelaruzb", Invite: "https://t.me/anmelaruzb"},
	//{ID: -1003369138926, Name: "AniDarkX", Invite: "https://t.me/AniDarkX"},
	//{ID: -1003557426309, Name: "nakrutkaurush", Invite: "https://t.me/nakrutkaurush"},
	{ID: -1003411861509, Name: "Maxfiy Kanal", Invite: "https://t.me/+C0qmcf4ZHY83NmNi"},
	//{ID: -1003588929805, Name: "Maxfiy Kanal", Invite: "https://t.me/+CPtYbpger5U0YjNi"},
	{ID: -1003540484817, Name: "Maxfiy Kanal", Invite: "https://t.me/+mbBXFN4zFHAyMDQy"},
	{ID: -1003532028606, Name: "Maxfiy Kanal", Invite: "https://t.me/+4u-B783Cgvs5YWNi"},
	//{ID: -1003323161290, Name: "Manga Uzb", Invite: "https://t.me/Manga_uzbekcha26"},
	//{ID: -1003276785399, Name: "animelaruzbektilid3", Invite: "https://t.me/animelaruzbektilid3"},
	//{ID: -1003316396409, Name: "anmelar_chat", Invite: "https://t.me/anmelar_chat"},
	//{ID: -1003227139819, Name: "Maxfiy Kanal", Invite: "https://t.me/+O3K3g71yc2cwYThi"},
}
var (
	// Mavjud o'zgaruvchilaringiz...

	// 🆕 Mana buni qo'shing:
	userState = make(map[int64]string)
)
var (
	userActive  = make(map[int64]time.Time)
	userJoined  = make(map[int64]time.Time)
	statsMutex  sync.RWMutex
	searchStats = make(map[string]int)
)

var (
	vipUsers = make(map[int64]bool)
	vipMutex sync.RWMutex
)

func saveVips() {
	vipMutex.RLock()
	data, _ := json.Marshal(vipUsers)
	vipMutex.RUnlock()
	_ = os.WriteFile("vips.json", data, 0644)
}

func loadVips() {
	file, err := os.ReadFile("vips.json")
	if err != nil {
		return
	}
	vipMutex.Lock()
	_ = json.Unmarshal(file, &vipUsers)
	vipMutex.Unlock()
}

func parseRelativeTime(input string) (time.Time, error) {
	now := time.Now()
	input = strings.TrimSpace(strings.ToLower(input))
	if len(input) < 2 {
		return time.Time{}, fmt.Errorf("invalid input")
	}

	unit := input[len(input)-1:] // oxirgi harf: s/m/h/d
	numStr := input[:len(input)-1]

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return time.Time{}, err
	}

	var dur time.Duration
	switch unit {
	case "s":
		dur = time.Duration(num) * time.Second
	case "m":
		dur = time.Duration(num) * time.Minute
	case "h":
		dur = time.Duration(num) * time.Hour
	case "d":
		dur = time.Duration(num) * 24 * time.Hour
	default:
		return time.Time{}, fmt.Errorf("unknown unit")
	}

	return now.Add(dur), nil
}

func schedulePost(bot *tele.Bot, post *ScheduledPost) {
	// 1. Kutish vaqtini hisoblaymiz
	delay := time.Until(post.SendTime)
	if delay <= 0 {
		delay = 1 * time.Second
	}

	// 2. Taymerni ishga tushiramiz
	time.AfterFunc(delay, func() {
		// --- BEKOR QILINGANLIGINI TEKSHIRISH ---
		scheduleMutex.Lock()
		_, exists := scheduledPosts[post.ID]
		scheduleMutex.Unlock()

		// Agar admin "Bekor qilish" tugmasini bosgan bo'lsa, post mapdan o'chirilgan bo'ladi
		if !exists {
			fmt.Printf("🚫 Post (ID: %d) bekor qilingan, yuborish to'xtatildi.\n", post.ID)
			return
		}
		// ---------------------------------------

		count := 0
		total := len(post.ChatIDs)

		fmt.Printf("\n🚀 Post yuborish boshlandi. Jami: %d kishiga\n", total)

		// 3. Foydalanuvchilarga yuborish loopi
		for _, chatID := range post.ChatIDs {
			// Yuborish jarayonida ham post bekor qilinganini tekshirish (ixtiyoriy, juda katta bazalar uchun)
			err := sendScheduledContent(bot, chatID, post.Content)
			if err == nil {
				count++
			} else {
				fmt.Printf("⚠️ Xato (ChatID %d): %v\n", chatID, err)
			}

			// Telegram limiti: ~30 msg/sec
			time.Sleep(35 * time.Millisecond)
		}

		// 4. ADMINGA HISOBOT YUBORISH
		reportMsg := fmt.Sprintf(
			"✅ <b>Rejalashtirilgan post tugatildi!</b>\n\n"+
				"📊 <b>Natija:</b>\n"+
				"🟢 Muvaffaqiyatli: <code>%d</code>\n"+
				"🔴 Xatolik: <code>%d</code>\n"+
				"🏁 Jami: <code>%d</code>",
			count, total-count, total,
		)

		fmt.Printf("📢 Adminga hisobot yuborilmoqda... (AdminID: %d)\n", post.AdminID)

		// Admin ID 0 bo'lmasligi kerak (yuqoridagi handleAll da to'g'irladik)
		_, err := bot.Send(tele.ChatID(post.AdminID), reportMsg, tele.ModeHTML)
		if err != nil {
			fmt.Printf("❌ Admin xabari ketmadi (AdminID %d): %v\n", post.AdminID, err)
		}

		// 5. Xotirani tozalash
		scheduleMutex.Lock()
		delete(scheduledPosts, post.ID)
		scheduleMutex.Unlock()

		fmt.Println("🏁 Jarayon yakunlandi.")
	})
}

func sendScheduledContent(bot *tele.Bot, chatID int64, item ContentItem) error {
	recipient := tele.ChatID(chatID)
	var err error

	switch item.Kind {
	case "text":
		_, err = bot.Send(recipient, item.Text)
	case "photo":
		_, err = bot.Send(recipient, &tele.Photo{
			File:    tele.File{FileID: item.FileID},
			Caption: item.Text,
		})
	case "video":
		_, err = bot.Send(recipient, &tele.Video{
			File:    tele.File{FileID: item.FileID},
			Caption: item.Text,
		})
	default:
		return fmt.Errorf("noma'lum kontent turi: %s", item.Kind)
	}
	return err
}

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
	vipMutex.RLock()
	isVip := vipUsers[userID]
	vipMutex.RUnlock()

	if isVip {
		return nil
	}

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

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	menu.Reply(
		menu.Row(menu.Text("Animelar"),
			menu.Text("🔍 Qidiruv"),
		),
		menu.Row(
			menu.Text("🖼 Rasm orqali qidirish"),
			menu.Text("🧩 help"),

			//menu.Text("📹 Video orqali qidirish"), // Yangi tugma qo'shildi
		),
		menu.Row(
		//menu.Text("🔍 Qidiruv"),
		//menu.Text("🧩 help"),
		),
	)

	loadVips()
	adminMenu := &tele.ReplyMarkup{}
	btnSchedule := adminMenu.Data("⏰ Rejalashtirilgan post", "admin_schedule")
	btnSettings := adminMenu.Data("⚙️ Sozlamalar", "admin_settings")
	btnBroadcast := adminMenu.Data("📢 Reklama", "admin_broadcast")
	btnStats := adminMenu.Data("📊 Statistika", "admin_stats")
	btnVip := adminMenu.Data("🌟 VIP Boshqaruv", "admin_vip_main")
	// VIP Menyu
	vipSubMenu := &tele.ReplyMarkup{}
	btnAddVip := vipSubMenu.Data("➕ Qo'shish", "vip_add")
	btnDelVip := vipSubMenu.Data("➖ O'chirish", "vip_del")
	btnListVip := vipSubMenu.Data("📜 Ro'yxat", "vip_list")
	btnBackAdmin := vipSubMenu.Data("⬅️ Orqaga", "back_admin")
	// ===== ADMIN PANEL =====
	b.Handle("/admin", func(c tele.Context) error {
		if !isAdmin(c.Sender().ID) {
			return nil
		}

		adminMenu.Inline(
			adminMenu.Row(btnSchedule, btnSettings),
			adminMenu.Row(btnBroadcast, btnStats),
			adminMenu.Row(btnVip),
		)

		return c.Send("👨‍💻 *Admin Panel*", adminMenu, tele.ModeMarkdown)
	})
	b.Handle(&btnVip, func(c tele.Context) error {
		vipSubMenu.Inline(
			vipSubMenu.Row(btnAddVip, btnDelVip),
			vipSubMenu.Row(btnListVip),
			vipSubMenu.Row(btnBackAdmin),
		)

		return c.Edit("🌟 *VIP foydalanuvchilarni boshqarish*", vipSubMenu, tele.ModeMarkdown)
	})
	b.Handle(&btnBackAdmin, func(c tele.Context) error {
		adminMenu.Inline(
			adminMenu.Row(btnSchedule, btnSettings),
			adminMenu.Row(btnBroadcast, btnStats),
			adminMenu.Row(btnVip),
		)

		return c.Edit("👨‍💻 Admin Panel", adminMenu, tele.ModeMarkdown)
	})
	b.Handle(&btnSchedule, func(c tele.Context) error {
		adminState[c.Sender().ID] = "wait_schedule_time"
		return c.Send(
			"⏰ *Post yuborish vaqtini kiriting*\n\n"+
				"Masalan:\n"+
				"`10s` – 10 soniya\n"+
				"`10m` – 10 minut\n"+
				"`3h` – 3 soat\n"+
				"`2d` – 2 kun",
			tele.ModeMarkdown,
		)
	})
	b.Handle(&btnAddVip, func(c tele.Context) error {
		adminState[c.Sender().ID] = "wait_vip_add"
		return c.Send("🆔 VIP qilmoqchi bo'lgan foydalanuvchi ID sini yuboring:")
	})
	b.Handle(&btnDelVip, func(c tele.Context) error {
		adminState[c.Sender().ID] = "wait_vip_del"
		return c.Send("🆔 VIP-dan olib tashlamoqchi bo'lgan foydalanuvchi ID sini yuboring:")
	})
	b.Handle(&btnListVip, func(c tele.Context) error {
		vipMutex.RLock()
		text := "🌟 VIP Foydalanuvchilar ro'yxati:\n\n"
		for id := range vipUsers {
			text += fmt.Sprintf("👤 ` %d `\n", id)
		}
		vipMutex.RUnlock()
		return c.Send(text, tele.ModeMarkdown)
	})
	b.Handle(&btnBroadcast, func(c tele.Context) error {
		adminState[c.Sender().ID] = "waiting_for_ad"
		return c.Send("📥 Reklama xabarini yuboring. (Rasm, video yoki matn)")
	})
	b.Handle(&btnStats, func(c tele.Context) error {
		return sendStatistics(c)
	})
	b.Handle(&tele.Btn{Unique: "confirm_ad"}, func(c tele.Context) error {
		adMsg := adminWaitingAd[c.Sender().ID]
		if adMsg == nil {
			return c.Edit("❌ Xabar topilmadi.")
		}

		// 1. Foydalanuvchilar ro'yxatini xotiraga nusxalab olamiz (Copying slice)
		var usersToBroadcast []int64
		statsMutex.RLock()
		for userID := range userJoined {
			usersToBroadcast = append(usersToBroadcast, userID)
		}
		statsMutex.RUnlock() // Qulfni darhol ochamiz!

		// 2. Reklamani alohida goroutine ichida yuboramiz
		go func(msg *tele.Message, targets []int64) {
			count := 0
			for _, userID := range targets {
				_, err := b.Copy(tele.ChatID(userID), msg)
				if err == nil {
					count++
				}
				// Telegram cheklovlari (Limit) uchun kichik pauza
				time.Sleep(35 * time.Millisecond)
			}
			b.Send(c.Sender(), fmt.Sprintf("✅ Reklama tugatildi! %d kishiga yuborildi.", count))
		}(adMsg, usersToBroadcast)

		delete(adminState, c.Sender().ID)
		delete(adminWaitingAd, c.Sender().ID) // Xotirani tozalash
		return c.Edit("🚀 Reklama tarqatish boshlandi. Bot boshqa buyruqlarni qabul qilaveradi.")
	})
	b.Handle(&tele.Btn{Unique: "cancel_ad"}, func(c tele.Context) error {
		delete(adminState, c.Sender().ID)
		return c.Edit("❌ Reklama bekor qilindi.")
	})
	b.Handle(tele.OnCallback, func(c tele.Context) error {
		data := c.Callback().Data

		// 1. Qaysi tugma bosilganini aniqlaymiz
		if strings.Contains(data, "cancel_post") {
			// Telebot data formati: "\funique|data"
			parts := strings.Split(data, "|")
			if len(parts) < 2 {
				return c.Respond(&tele.CallbackResponse{Text: "❌ Xatolik: ID topilmadi"})
			}

			// ID ni olamiz
			postID, err := strconv.Atoi(parts[1])
			if err != nil {
				return c.Respond(&tele.CallbackResponse{Text: "❌ Noto'g'ri ID"})
			}

			scheduleMutex.Lock()
			// 2. Post hali mapda bormi (yuborilmaganmi)?
			if _, exists := scheduledPosts[postID]; exists {
				delete(scheduledPosts, postID) // Mapdan o'chirish = Yuborishni to'xtatish
				scheduleMutex.Unlock()

				// Tugmani o'chirib, matnni o'zgartiramiz
				c.Respond(&tele.CallbackResponse{Text: "✅ Muvaffaqiyatli bekor qilindi"})
				return c.Edit("🗑 <b>Rejalashtirilgan post bekor qilindi va o'chirildi!</b>", tele.ModeHTML)
			}
			scheduleMutex.Unlock()

			return c.Respond(&tele.CallbackResponse{Text: "⚠️ Kech qoldingiz yoki post topilmadi!"})
		}

		return c.Respond() // Callbackni yopish (soat belgisi ketishi uchun)
	})
	b.Handle(tele.OnCallback, func(c tele.Context) error {
		data := c.Callback().Data

		// Telebot callback data formati: "\fselect_anime|97"
		if strings.HasPrefix(data, "\fselect_anime") {
			parts := strings.Split(data, "|")
			if len(parts) > 1 {
				animeID := parts[1]

				_ = c.Respond() // Soat belgisini yo'qotish
				_ = c.Delete()  // Qidiruv natijalari menyusini o'chirib tashlash (toza turishi uchun)

				// Foydalanuvchiga ID ni qanday yozish kerakligini ko'rsatamiz
				return c.Send(fmt.Sprintf("%s bu kodni  kiriting", animeID))
			}
		}
		return c.Respond()
	})
	b.Handle(tele.OnText, func(c tele.Context) error {
		text := c.Text()
		userID := c.Sender().ID

		if text == "🖼 Rasm orqali qidirish" {
			userState[userID] = "wait_image_search"
			return c.Send("🔍 Iltimos, qidirmoqchi bo'lgan rasmingizni yuboring...")
		}

		// FAQAT matnli xabarlar uchun default javob
		return c.Send("Uzr bu kod noto'g'ri yoki hozircha mavjud emas!")
	})
	b.Handle(tele.OnPhoto, func(c tele.Context) error {
		userID := c.Sender().ID

		// Faqat holat "wait_image_search" bo'lgandagina ishlaydi
		if userState[userID] == "wait_image_search" {
			photo := c.Message().Photo // Ba'zi versiyalarda bu massiv bo'lishi mumkin

			// Agar indexing xatosi bersa, shunchaki photo.FileID ishlating
			// Agar bermasa, eng oxirgisini oling:
			file, err := b.FileByID(photo.FileID)
			if err != nil {
				return c.Send("❌ Rasmni yuklab olishda xatolik yuz berdi.")
			}

			imageURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", token, file.FilePath)

			// Google Lens'ga URL-ni xavfsiz formatda yuboramiz
			encodedURL := url.QueryEscape(imageURL)
			googleSearchURL := "https://lens.google.com/uploadbyurl?url=" + encodedURL

			inlineMenu := &tele.ReplyMarkup{}
			btnResult := inlineMenu.URL("🔍 Ob'ektni aniqlash (Google Lens)", googleSearchURL)
			inlineMenu.Inline(inlineMenu.Row(btnResult))

			delete(userState, userID) // Holatni tozalash
			return c.Send("✅ Rasm qabul qilindi! Ma'lumot olish uchun tugmani bosing:", inlineMenu)
		}
		return nil
	})
	handleAll := func(c tele.Context) error {
		updateUserActivity(c.Sender().ID)
		userID := c.Sender().ID
		state := adminState[userID]
		text := c.Text()
		msg := c.Message()
		if text == "🔍 Qidiruv" {
			userState[userID] = "wait_anime_search"
			return c.Send("🔍 Anime nomini kiriting (masalan: Narotu):")
		}
		if userState[userID] == "wait_anime_search" && text != "" {
			results := SearchAnime(text) // search.go dagi funksiya

			if len(results) == 0 {
				return anmelaruzb.Home(c)
			}

			inlineMenu := &tele.ReplyMarkup{}
			var rows []tele.Row

			for _, res := range results {
				// Tugma ustida Nomi, bosilganda ID ketsin
				btn := inlineMenu.Data("� "+res.Name, "select_anime", res.ID)
				rows = append(rows, inlineMenu.Row(btn))
			}

			inlineMenu.Inline(rows...)
			delete(userState, userID) // Holatni yopamiz
			return c.Send(" Natijalar (keraklisini tanlang):", inlineMenu)
		}
		if userState[userID] == "wait_image_search" {

			// ❌ Agar rasm yuborilmagan bo‘lsa
			if c.Message().Photo == nil {
				// state’dan chiqamiz
				delete(userState, userID)

				return c.Send("ℹ️ Rasm yuborilmadi. Qidiruv bekor qilindi.\nBot odatdagi rejimda ishlashda davom etadi.")
			}

			// ✅ Agar rasm bo‘lsa — davom etamiz
			photo := c.Message().Photo

			file, err := b.FileByID(photo.FileID)
			if err != nil {
				delete(userState, userID)
				return c.Send("❌ Rasmni yuklab olishda xatolik yuz berdi.")
			}

			imageURL := fmt.Sprintf(
				"https://api.telegram.org/file/bot%s/%s",
				token,
				file.FilePath,
			)

			encodedURL := url.QueryEscape(imageURL)
			googleSearchURL := "https://lens.google.com/uploadbyurl?url=" + encodedURL

			inlineMenu := &tele.ReplyMarkup{}
			btnResult := inlineMenu.URL("🔍 Ob'ektni aniqlash (Google Lens)", googleSearchURL)
			inlineMenu.Inline(inlineMenu.Row(btnResult))

			delete(userState, userID)

			return c.Send("✅ Rasm qabul qilindi! Ma'lumot olish uchun tugmani bosing:", inlineMenu)
		}

		if text == "🖼 Rasm orqali qidirish" {
			userState[userID] = "wait_image_search"
			return c.Send("🔍 Iltimos, qidirmoqchi bo'lgan rasmingizni yuboring...")
		}
		if isAdmin(userID) && state != "" {
			// REKLAMA YUBORISH HOLATI Message
			if state == "waiting_for_ad" {
				// 1. Jarayon boshlanganini tekshirish
				log.Printf(">>> [REKLAMA] User %d xabar yubordi", userID)

				// 2. Kelgan xabar turini aniqlash
				msg := c.Message()
				if msg.Photo != nil {
					log.Printf(">>> [TUR] Rasm keldi. Caption: %s", msg.Caption)
				} else if msg.Video != nil {
					log.Printf(">>> [TUR] Video keldi. Caption: %s", msg.Caption)
				} else {
					log.Printf(">>> [TUR] Matn keldi. Text: %s", msg.Text)
				}

				// 3. Xabarni xotiraga saqlash
				adminWaitingAd[userID] = msg
				log.Println(">>> [OK] Xabar adminWaitingAd xaritasiga saqlandi")

				m := &tele.ReplyMarkup{}
				btnYes := m.Data("✅ Tasdiqlash", "confirm_ad")
				btnNo := m.Data("❌ Bekor qilish", "cancel_ad")
				m.Inline(m.Row(btnYes, btnNo))

				// 4. Nusxalash jarayoni
				_ = c.Send("👇 **Reklama ko'rinishi:**")
				_, err := b.Copy(c.Recipient(), msg)
				if err != nil {
					log.Printf(">>> [XATO] Copy qilishda xatolik: %v", err)
					return c.Send("❌ Xabarni nusxalashda xatolik.")
				}
				log.Println(">>> [OK] Reklama nusxasi adminning o'ziga ko'rsatildi")

				return c.Send("Ushbu xabarni hamma foydalanuvchilarga yuboramizmi?", m)
			}
			// VIP QO'SHISH/O'CHIRISH HOLATI
			if state == "wait_vip_add" || state == "wait_vip_del" {
				var targetID int64
				_, err := fmt.Sscanf(text, "%d", &targetID)
				if err != nil {
					return c.Send("❌ Xato! Faqat raqamlardan iborat ID yuboring.")
				}

				vipMutex.Lock()
				if state == "wait_vip_add" {
					vipUsers[targetID] = true
					c.Send(fmt.Sprintf("✅ %d muvaffaqiyatli VIP qilindi!", targetID))
				} else {
					delete(vipUsers, targetID)
					c.Send(fmt.Sprintf("🗑 %d VIP ro'yxatidan o'chirildi!", targetID))
				}
				vipMutex.Unlock()
				saveVips()
				delete(adminState, userID)
				return nil
			}

			// VAQTNI QABUL QILISH HOLATI (10s, 10m...)
			if state == "wait_schedule_time" {
				if c.Message().Media() != nil {
					return c.Send("❌ Avval vaqtni yuboring! (Masalan: 10m)")
				}

				sendTime, err := parseRelativeTime(text)
				if err != nil {
					return c.Send("❌ Format xato! Misol: `10m`, `1h`...", tele.ModeMarkdown)
				}

				adminState[userID] = "wait_schedule_content"
				scheduleMutex.Lock()
				scheduledPosts[scheduleAutoID] = &ScheduledPost{
					ID:       scheduleAutoID,
					AdminID:  userID,
					SendTime: sendTime,
				}
				scheduleAutoID++
				scheduleMutex.Unlock()
				return c.Send("📨 Endi postni yuboring (matn / rasm / video)")
			}

			// POST KONTENTINI QABUL QILISH HOLATI
			if state == "wait_schedule_content" {
				var post *ScheduledPost
				scheduleMutex.Lock()
				for _, p := range scheduledPosts {
					if p.Content.Kind == "" && p.AdminID == userID {
						post = p
						break
					}
				}
				scheduleMutex.Unlock()

				if post == nil {
					return c.Send("❌ Xatolik: Rejalashtirilgan vaqt topilmadi.")
				}

				// Kontentni aniqlash
				if msg.Photo != nil {
					post.Content = ContentItem{Kind: "photo", FileID: msg.Photo.FileID, Text: msg.Caption}
				} else if msg.Video != nil {
					post.Content = ContentItem{Kind: "video", FileID: msg.Video.FileID, Text: msg.Caption}
				} else if msg.Text != "" {
					post.Content = ContentItem{Kind: "text", Text: msg.Text}
				} else {
					return c.Send("❌ Faqat matn, rasm yoki video!")
				}

				// ... qolgan schedulePost va javob qaytarish kodi ...
				go schedulePost(b, post)
				delete(adminState, userID)
				return c.Send("✅ Rejalashtirildi!")
			}
		}

		// 2. OBUNA TEKSHIRUVI (Oddiy foydalanuvchilar uchun)
		if !isAdmin(userID) {
			missing := notAllowedChannels(b, userID)
			if len(missing) > 0 {
				return sendSubMessage(c, missing)
			}
		}

		// 4. ASOSIY BUYRUQLAR VA KODLAR
		switch text {
		case "/start":
			return c.Send("✅ Xush kelibsiz marhamt /menu dan foydalaning!", menu)
		case "Animelar", "/menu":
			return Menu.Home(c)
		case "🧩 help":
			return Help.Home(c)
		default:
			// Agar hech qanday shartga tushmasa (Masalan: Anime kodi yozilsa)
			return anmelaruzb.Home(c)
		}
	}

	b.Handle(tele.OnText, handleAll)
	b.Handle(tele.OnMedia, handleAll)
	b.Handle(tele.OnPhoto, handleAll)
	b.Handle(tele.OnVideo, handleAll)
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
	text := "<b>❗️ Botdan foydalanish uchun quyidagi kanallarga a'zo bo‘ling yoki so‘rov yuboring:</b>"
	m := &tele.ReplyMarkup{}
	var rows []tele.Row
	for _, ch := range missing {
		rows = append(rows, m.Row(m.URL("📢 "+ch.Name, ch.Invite)))
	}
	rows = append(rows, m.Row(m.Data("✅ Tekshirish", "check_sub")))
	m.Inline(rows...)
	return c.Send(text, m, tele.ModeHTML)
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

//
