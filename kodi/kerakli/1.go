package kerakli

//
//import (
//	"encoding/json"
//	"fmt"
//	beego "github.com/beego/beego/v2/server/web"
//	"log"
//	"namelaruzb_bot/kodi/Help"
//	"namelaruzb_bot/kodi/Menu"
//	"namelaruzb_bot/kodi/anmelaruzb"
//	_default "namelaruzb_bot/kodi/default"
//	"os"
//	"sort"
//	"sync"
//	"time"
//
//	tele "gopkg.in/telebot.v4"
//)
//
//var (
//	pendingRequests = make(map[int64]map[int64]bool)
//	requestMutex    sync.RWMutex
//)
//
//type ChannelInfo struct {
//	ID     int64
//	Name   string
//	Invite string
//}
//
//var myChannels = []ChannelInfo{
//	{ID: -1003050934981, Name: "anmelaruzb", Invite: "https://t.me/anmelaruzb"},
//	{ID: -1003316396409, Name: "anmelar_chat", Invite: "https://t.me/anmelar_chat"},
//	{ID: -1003323161290, Name: "Manga Uzb", Invite: "https://t.me/Manga_uzbekcha26"},
//	//{ID: -1003276785399, Name: "Maxfiy Kanal", Invite: "https://t.me/+9bsKINaEOHJiNjUy"},
//	{ID: -1003411861509, Name: "Maxfiy Kanal", Invite: "https://t.me/+C0qmcf4ZHY83NmNi"},
//}
//
//var (
//	userActive  = make(map[int64]time.Time)
//	userJoined  = make(map[int64]time.Time)
//	statsMutex  sync.RWMutex
//	searchStats = make(map[string]int)
//)
//
//func updateUserActivity(userID int64) {
//	statsMutex.Lock()
//	now := time.Now()
//	userActive[userID] = now
//
//	isNew := false
//	if _, ok := userJoined[userID]; !ok {
//		userJoined[userID] = now
//		isNew = true
//	}
//	statsMutex.Unlock()
//
//	// Faqat yangi odam qo'shilganda saqlaymiz (faylni qiynamaslik uchun)
//	if isNew {
//		saveStats()
//	}
//}
//func saveStats() {
//	statsMutex.RLock() // O'qish uchun lock
//	data := struct {
//		UserJoined  map[int64]time.Time `json:"userJoined"`
//		UserActive  map[int64]time.Time `json:"userActive"`
//		SearchStats map[string]int      `json:"searchStats"`
//	}{
//		UserJoined:  userJoined,
//		UserActive:  userActive,
//		SearchStats: searchStats,
//	}
//	statsMutex.RUnlock()
//
//	file, err := json.MarshalIndent(data, "", "  ")
//	if err != nil {
//		log.Println("Saqlashda xato:", err)
//		return
//	}
//	_ = os.WriteFile("stats.json", file, 0644)
//}
//
//func loadStats() {
//	file, err := os.ReadFile("stats.json")
//	if err != nil {
//		return // Fayl yo'q bo'lsa normal holat
//	}
//
//	var data struct {
//		UserJoined  map[int64]time.Time `json:"userJoined"`
//		UserActive  map[int64]time.Time `json:"userActive"`
//		SearchStats map[string]int      `json:"searchStats"`
//	}
//
//	if err := json.Unmarshal(file, &data); err != nil {
//		log.Println("Yuklashda xato:", err)
//		return
//	}
//
//	statsMutex.Lock()
//	if data.UserJoined != nil {
//		userJoined = data.UserJoined
//	}
//	if data.UserActive != nil {
//		userActive = data.UserActive
//	}
//	if data.SearchStats != nil {
//		searchStats = data.SearchStats
//	}
//	statsMutex.Unlock()
//}
//
//var admins = map[int64]bool{
//	7518992824: true,
//	//987654321: true,
//}
//
//func isAdmin(userID int64) bool {
//	return admins[userID]
//}
//
//func notAllowedChannels(b *tele.Bot, userID int64) []ChannelInfo {
//	var missing []ChannelInfo
//
//	for _, ch := range myChannels {
//		chat := &tele.Chat{ID: ch.ID}
//		member, err := b.ChatMemberOf(chat, &tele.User{ID: userID})
//
//		isMember := err == nil &&
//			(member.Role == tele.Member ||
//				member.Role == tele.Administrator ||
//				member.Role == tele.Creator)
//
//		if isMember {
//			continue
//		}
//
//		// maxfiy kanal uchun request tekshirish
//		requestMutex.RLock()
//		userReqs := pendingRequests[userID]
//		hasRequested := userReqs != nil && userReqs[ch.ID]
//		requestMutex.RUnlock()
//
//		if !hasRequested {
//			missing = append(missing, ch)
//		}
//	}
//	return missing
//}
//
//func Bot() {
//	loadStats()
//	token := beego.AppConfig.DefaultString("telegram::token", "")
//	b, err := tele.NewBot(tele.Settings{
//		Token:  token,
//		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Menu
//	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
//	row := menu.Row(menu.Text("Animelar"))
//	menu.Reply(row)
//	menu.Reply(menu.Row(menu.Text("🧩 help")))
//
//	// Join request
//	b.Handle(tele.OnChatJoinRequest, func(c tele.Context) error {
//		req := c.ChatJoinRequest()
//		if req == nil {
//			return nil
//		}
//
//		requestMutex.Lock()
//		if pendingRequests[req.Sender.ID] == nil {
//			pendingRequests[req.Sender.ID] = make(map[int64]bool)
//		}
//		pendingRequests[req.Sender.ID][req.Chat.ID] = true
//		requestMutex.Unlock()
//
//		log.Printf("📥 Join request: user=%d chat=%d", req.Sender.ID, req.Chat.ID)
//		return nil
//	})
//
//	// START
//	b.Handle("/start", func(c tele.Context) error {
//		updateUserActivity(c.Sender().ID)
//		missing := notAllowedChannels(b, c.Sender().ID)
//		if len(missing) > 0 {
//			return sendSubMessage(c, missing)
//		}
//		return c.Send("✅ Botga xush kelibsiz!", menu)
//	})
//
//	// Tekshirish tugmasi
//	b.Handle(&tele.Btn{Unique: "check_sub"}, func(c tele.Context) error {
//		updateUserActivity(c.Sender().ID)
//		missing := notAllowedChannels(b, c.Sender().ID)
//		if len(missing) == 0 {
//			_ = c.Delete()
//			return c.Send("✅ Hammasi joyida, davom eting!", menu)
//		}
//		return c.Respond(&tele.CallbackResponse{
//			Text:      "❌ Hali hamma kanal bajarilmadi",
//			ShowAlert: true,
//		})
//	})
//
//	// HAR BIR XABARDA TEKSHIRUV
//	b.Handle(tele.OnText, func(c tele.Context) error {
//		updateUserActivity(c.Sender().ID)
//		missing := notAllowedChannels(b, c.Sender().ID)
//		if len(missing) > 0 {
//			return sendSubMessage(c, missing)
//		}
//
//		switch c.Text() {
//		case "🖋️ anme izlash",
//			"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
//			"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
//			"21", "22", "23", "24", "25", "26", "27", "28", "29", "30",
//			"31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
//			"41", "42", "43", "44", "45", "46", "47", "48", "49", "50",
//			"51", "52", "53", "54", "55", "56", "57", "58", "59", "60",
//			"61", "62", "63", "64", "65", "66", "67", "68", "69", "70",
//			"71", "72", "73", "74", "75", "76", "77", "78", "79", "80",
//			"81", "82", "83", "84", "85", "86", "87", "88", "89", "90",
//			"91", "92", "93", "94", "95", "96", "97", "98", "99", "100",
//			"101", "102", "103", "104", "105", "106", "107", "108", "109", "110",
//			"111", "112", "113", "114", "115", "116", "117", "118", "119", "120",
//			"121", "122", "123", "124", "125", "126", "127", "128", "129", "130",
//			"131", "132", "133", "134", "135", "136", "137", "138", "139", "140",
//			"141", "142", "143", "144", "145", "146", "147", "148", "149", "150",
//			"151", "152", "153", "154", "155", "156", "157", "158", "159", "160",
//			"161", "162", "163", "164", "165", "166", "167", "168", "169", "170",
//			"171", "172", "173", "174", "175", "176", "177", "178", "179", "180",
//			"181", "182", "183", "184", "185", "186", "187", "188", "189", "190",
//			"191", "192", "193", "194", "195", "196", "197", "198", "199", "200":
//			// bu yerda kodi bajariladi
//
//			return anmelaruzb.Home(c)
//		case "Naruto", "Iblislar qotili", "Franksdagisevgi", "Vanpis", "Franksdagi sevgi",
//			"Tokiyoqasoskorlari", "Xarobalar qiroligi", "Daho shahzodani mamlakatni qutqargani haqida",
//			"Qirol oʻyini", "Yetti O'lim Gunohlari", "Parazit - Hayot Saboqlari", "Doktor stoun",
//			"Sakamoto Kunlari", "Uyatchang Qahramon va Qotil Malikalar", "Tungi Boyqushlar Kuyi",
//			"Omadsizning qayta tug'ilishi", "Uysiz Ma'bud", "Ochko'z Berserk", "Qalqon qahromoni",
//			"Soyada kotarilish", "Titanlar hujumi", "Jodugarlar jangi", "Sharlota", "Soqolimni olib,",
//			"Soqolimni olib, yuqori maktab qizini uyimga olib keldim", "DMC", "Qilich sanati online", "Ovoz Shakli",
//			"Yozgi Urushima Tuneli Oldidagi Hayrlashuv", "Hukmdor", "Bir zarbli odam", "Quyon Qiz",
//			"Sarguzashtchilar Restorani", "Horimiya", "Hyouka", "U qiz yolgiz", "Josus X Oilasi",
//			"Qoshni farishta", "Aliya bazan mega rustilida nos karashma qiladi", "Yangi Saga",
//			"Davolovchi qahramon", "Kayju 8-Raqam", "Pari Dumi haqida afsona", "yolg'izlikda daraja ko'taish":
//			return anmelaruzb.Home(c)
//		case "Qotil Akame", "Ochkoz bersek", "Qora o'q", "O‘lmas qirolning kundalik hayoti", "Zombi 100",
//			"Nomsiz Xotira", "Shamolni Bo'ysundirish", "Dororo", "Arra Odam", "So'ngi serafim", "Qora Klever", "Taxt muxri",
//			"Bleach", "Zulmat Farzandi", "Qudrat! Yangi Hikoya", "Yozukura Oilasi", "Poʻlat qal'adagi kabaneri",
//			"Rainbow", "Qo'g'irchoqlar sirki", "Yulduz Farzandi", "Men Muvaffaqiyatsiz...", "Qahramon Boʻlish X",
//			"Zulmat Iblisi", "Jahannam jannati", "Vanitas xotiralari", "Violet Evergarden", "Elita Sinfi",
//			"Davolash sehridan foydalanishni noto'g'ri usuli", "Lordi Armiyasining eng kuchli Sehrgari...",
//			"Tahlil qilish qobiliyatiga ega aristokrat bo'lib qayta tug'ilish", "Ilohiy qilich maktabining Iblis qilich egasi",
//			"Meni Qilich bo'lib qayta tug'ilishim haqida", "O'z joniga qasd qiluvchilar o'zga dunyoda",
//			"Sehr Yaratuvchi Boshqa dunyoda qanday qilib sehr yaratish mumkin", "O'rta yoshli erkakning zodagon qiziga aylanishi",
//			"Eng qudratli partiya tomonidan o'limgacha tarbiyalangan Ossan ...", "Baholovchining mashhur bo'lmagan ishi aslida eng Kuchlisi hisoblanadi":
//			return anmelaruzb.Home(c)
//		case "Badargʻa qilingan qahramon",
//			"Boshqa dunyodan muammoli bolalar", "Tokyo Gul", "Shilliq sifatida qayta tug'ilganim haqida",
//			"Voleybol", "O'lim kundaligi", "Yana bir narsa soʻrasam boʻladimi", "Abadiylik qoriqchisi",
//			"Yettinchi umrni betashvish yashayotgan yovuz ayol", "Taxt Dastlabki Drift", "Dragon Raja",
//			"Xushboʻy Gul Viqor Bilan Gulaydi", "Xunuk Ammo Kuchli: Busamen G‘alaba Jangi", "Afsonaviy ilohy yilnomasi",
//			"Uzuklar Hukumdori: Rohhirm Urushi", "Taqdir Jang Kechasi", "Qora chaqiruvchi",
//			"Tahlil Qilish Qobiliyati Bilan Aristokrat Sifatida Qayta Tugʻildim", "Suv Sehrgari", "Sirlar Hukmdori",
//			"Seni oshqozon osti bezingni yemoqchi man", "Qobilyatsiz Nana", "Ozga dunyoda yolgiz hujum",
//			"Osmondagi Janglar", "Oltin Vaqt", "Minogra Apokalipsis", "Men galaktikalar aro imperiyaning yovuz lordiman",
//			"Tajribasiz Senpai", "Cheksizlikgacha Lv9999", "Mabudlar Hohishi Bilan", "Lookism", "Domekano",
//			"Gertsogning qiziga shaxsiy o'qituvchi boʻldim", "Gachiakuta Qasos", "Qudratli Soʻngi Dushman", "Yangi Darvoza",
//			"Chegara ortida", "Bucchgiri", "AprelYolgoni", "Afsonaviy ilohy ilnomasi", "Sen uchun O'lmas", "qip-qizil ragna",
//			"Oʻzga dunyoda darajamni koʻtarib, bu dunyoda ham tengsiz boʻldim", "Josus X Oilasi0", "Jonli Sana",
//			"Ragnarok Rekordi", "Masamune Qasosi", "Kim Meni Malika Qildi", "Bir soatli qizcha", "Koʻk Zindon",
//			"Biz birga bo'lsak, sevgimiz har qanday to'siqni ortda qoldiradi", "Moviy Quticha", "Hikaruning songi yozi",
//			"Taqdir: Buyuk Tartib Mutlaq Iblislar Jabhasi", "Sevgi deb atalgan shart", "Lideyl Dunyosi",
//			"Afsonaviy qahramonlar va Ruhlar malikasi qizi sifatida qayta tug‘ildim", "Skelet Ritsar o‘zga dunyoda":
//
//			return anmelaruzb.Home(c)
//		case "Yaponiyaga Xush Kelibsiz, Elf Xonim!", "Grimgaldagi kulgi va illuziya", "Nega hamma meni dunyomni unutdi?",
//			"80.000 oltin tanga to'plab hayotimni qayta qurdim", "Cheksiz dendagron", "Oy sayohati yangi dunyoga olib keladi",
//			"Man o'rgimchakman ! Ho'sh shunga nma qibti?", "O'zga dunyoda fermerlik hayotim", "Iblislar maktabiga hush kelibsiz",
//			"Daholar uchun oʻzga dunyoda yashash ham muammo emas", "Menda million hayot bor", "Maktab tomonidan tan olinmagan iblislar hukmdori",
//			"Boshqa dunyo bir zumda o'lim kuchiga dosh bera olmaydi", "Kumush qirolning qayta tugʻilishi", "Oʻzga dunyoda Noyob mahoratim",
//			"Egzartis boshqa dunyoda qayta tug'ilib eng kuchli bo'lishga intiladi", "Nikoh uzuklari haqida afsona", "tanyang urush yilnomalari",
//			"Yovuzlik darajasi 99: Men yashirin xo'jayin bo'lishim mumkin, lekin men jin xo'jayini emasman", "Doktor Eliza: hayotlarni saqlab qoluvchi malika",
//			"Eng zaif yirtqich hayvon", "Fojiaga sababchi bõlgan malika xalq uchun qõlidan kelgan barcha ishni qiladi",
//			"Iblislar hukmdori qoʻshimcha ishda!", "Oying bilan birga video o'yin", "Bekami Ko'st Yashashim Ushun Hamma Narsa Qildim",
//			"Uzoq paladin", "Realizm qahramoni qirollikni qayta qurishi", "Olti barg qahramonlari", "Daho Shifokorning Soyadagi Yangi Hayoti",
//			"Qayta tugʻilgan aristokratning misli koʻrilmagan sarguzashtlari", "Sakkizinchi o'g'il, qo'ysangizlarchi?", "O'yindagi Eng Boy odam",
//			"O'yinsiz hayot yo'q", "Arifureta: Dunyodagi eng kuchli hunarmand", "Gildiya adminstratori bo'lib ortiqcha ishlashni xohlamaganim uchun ishdan ketmoqchiman",
//			"Oddiy insondan qahramonlikkacha", "Tayoq va qilich", "Re:Zero", "Onmyo Qayta Tug‘ilishi: Hayolot Olami",
//			"Yettinchi shahzoda sifatida qayta tug'ildim va endi sehrimni istaganimcha kuchaytiraman!", "Mening Qotillik Maqomim Qahramonlik Maqomidan Yaxshiroq",
//			"Yugurening abadiyligi", "Daydi itlarning buyugi", "Meni qizcham nafaqat go'zal", "Meni qahramonlik akademiyam",
//			"Shangri-la chegarasi", "Barmoqlar uchidagi sevgi", " Kelajak kundaligi", "Men eng kuchli sarguzashtchi bo'lish uchun har doim mashq qildim",
//			"Do'stimning singlisi bezovta qilyapti", "Shikastlanishni istamasdim shuning uchun himoyamni kuchaytirdim",
//			"Oxirgi Telba Boss paydo bo'ldi", "Basketbol Kuroko", "Vayron bo'lgan mo'jizalar mamlakati", "Qahramonning qaytishi", "Bosning qizi va uning Enagasi",
//			"Ninja va Yakudza", "Detektiv allaqchon o'lgan", "Friren - Soʻnggi manzilga kuzatuvchi":
//
//			return anmelaruzb.Home(c)
//
//		case "Animelar", "/menu":
//			return Menu.Home(c)
//		case "🧩 help":
//			return Help.Home(c)
//		case "/stats", "📊 Statistika":
//			if !isAdmin(c.Sender().ID) {
//				return c.Send("❌ Sizda bu komandani ishlatish huquqi yo‘q")
//			}
//			return sendStatistics(c)
//		default:
//			return _default.Home(c)
//		}
//	})
//
//	log.Println("🤖 Bot ishga tushdi")
//	b.Start()
//}
//
//func sendSubMessage(c tele.Context, missing []ChannelInfo) error {
//	text := "<b>❗ Botdan foydalanish uchun quyidagi kanallarga a'zo bo‘ling yoki so‘rov yuboring:</b>"
//	m := &tele.ReplyMarkup{}
//	var rows []tele.Row
//
//	for _, ch := range missing {
//		rows = append(rows, m.Row(
//			m.URL("📢 "+ch.Name, ch.Invite),
//		))
//	}
//
//	rows = append(rows, m.Row(
//		m.Data("✅ Tekshirish", "check_sub"),
//	))
//
//	m.Inline(rows...)
//	return c.Send(text, m, tele.ModeHTML)
//}
//
//func sendStatistics(c tele.Context) error {
//	statsMutex.RLock()
//	defer statsMutex.RUnlock()
//
//	loc := time.FixedZone("Asia/Tashkent", 5*3600)
//	now := time.Now().In(loc)
//
//	// Vaqt chegaralari
//	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
//	startOf7Days := startOfToday.AddDate(0, 0, -6)
//	startOf30Days := startOfToday.AddDate(0, 0, -29)
//
//	// Foydalanuvchi hisoblash
//	total := len(userJoined)
//	active := 0
//	for _, t := range userActive {
//		if now.Sub(t) < 24*time.Hour {
//			active++
//		}
//	}
//	inactive := total - active
//
//	todayNew, last7, last30 := 0, 0, 0
//	for _, t := range userJoined {
//		tInLoc := t.In(loc)
//		if tInLoc.After(startOfToday) {
//			todayNew++
//		}
//		if tInLoc.After(startOf7Days) {
//			last7++
//		}
//		if tInLoc.After(startOf30Days) {
//			last30++
//		}
//	}
//
//	// Top 5 Anime saralash
//	type kv struct {
//		Key   string
//		Value int
//	}
//	var searchList []kv
//	for k, v := range searchStats {
//		searchList = append(searchList, kv{k, v})
//	}
//	sort.Slice(searchList, func(i, j int) bool {
//		return searchList[i].Value > searchList[j].Value
//	})
//
//	topCount := 5
//	if len(searchList) < 5 {
//		topCount = len(searchList)
//	}
//
//	// XABARNI SHAKLLANTIRISH
//	text := "------------------------------------------\n"
//	text += "🏆 **ENG MASHHUR 5 ANIME:**\n"
//	if topCount == 0 {
//		text += "ℹ️ Hozircha ma'lumot yo'q\n"
//	} else {
//		for i := 0; i < topCount; i++ {
//			text += fmt.Sprintf("%d. %s — %d qidiruv\n", i+1, searchList[i].Key, searchList[i].Value)
//		}
//	}
//
//	text += "------------------------------------------\n"
//	text += "🔗 **KANAL OBUNALARI:**\n"
//	for _, ch := range myChannels {
//		text += fmt.Sprintf("✅ [Kanalga o'tish](%s)\n", ch.Invite)
//	}
//
//	text += "------------------------------------------\n"
//	text += "📊 **FOYDALANUVCHILAR:**\n"
//	text += fmt.Sprintf("🟢 Faol foydalanuvchilar: %d\n", active)
//	text += fmt.Sprintf("🚫 Nofaol foydalanuvchilar: %d\n", inactive)
//	text += fmt.Sprintf("👥 Umumiy obunachilar: %d\n", total)
//
//	text += "\n🆕 **OBUNACHILAR (YANGI):**\n"
//	text += fmt.Sprintf("📅 Bugun: %d\n", todayNew)
//	text += fmt.Sprintf("🗓 7 kunlik: %d\n", last7)
//	text += fmt.Sprintf("🗓 30 kunlik: %d\n", last30)
//
//	text += "------------------------------------------\n"
//	text += fmt.Sprintf("ℹ️ Ma'lumotlar yangilangan: \n`%s`", now.Format("02.01.2006 | 15:04:05"))
//
//	return c.Send(text, tele.ModeMarkdown, tele.NoPreview)
//}
