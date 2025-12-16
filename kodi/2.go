package ko_di

import (
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"namelaruzb_bot/kodi/Help"
	"namelaruzb_bot/kodi/Menu"
	"namelaruzb_bot/kodi/anmelaruzb"
	"sync"
	"time"

	tele "gopkg.in/telebot.v4"
)

// 🔐 Join request yuborganlarni saqlash
var (
	pendingRequests = make(map[int64]map[int64]bool) // userID -> channelID
	requestMutex    sync.RWMutex
)

type ChannelInfo struct {
	ID     int64
	Name   string
	Invite string
}

var myChannels = []ChannelInfo{
	{ID: -1003050934981, Name: "anmelaruzb", Invite: "https://t.me/anmelaruzb"},
	{ID: -1003323161290, Name: "Manga Uzb", Invite: "https://t.me/Manga_uzbekcha26"},
	{ID: -1003276785399, Name: "Maxfiy Kanal", Invite: "https://t.me/+9bsKINaEOHJiNjUy"},
	{ID: -1003411861509, Name: "Maxfiy Kanal", Invite: "https://t.me/+C0qmcf4ZHY83NmNi"},
}

// 🔍 Majburiy tekshiruv
func notAllowedChannels(b *tele.Bot, userID int64) []ChannelInfo {
	var missing []ChannelInfo

	for _, ch := range myChannels {
		chat := &tele.Chat{ID: ch.ID}
		member, err := b.ChatMemberOf(chat, &tele.User{ID: userID})

		isMember := err == nil &&
			(member.Role == tele.Member ||
				member.Role == tele.Administrator ||
				member.Role == tele.Creator)

		if isMember {
			continue
		}

		// 🔒 maxfiy kanal uchun request tekshiramiz
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
	token := beego.AppConfig.DefaultString("telegram::token", "")
	b, err := tele.NewBot(tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	// 🎛 Menu
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	menu.Reply(
		menu.Row(menu.Text("Animelar")),
		menu.Row(menu.Text("🧩 help")),
	)

	// 📥 Join request tutib olish
	b.Handle(tele.OnChatJoinRequest, func(c tele.Context) error {
		req := c.ChatJoinRequest()
		if req == nil {
			return nil
		}

		requestMutex.Lock()
		if pendingRequests[req.Sender.ID] == nil {
			pendingRequests[req.Sender.ID] = make(map[int64]bool)
		}
		pendingRequests[req.Sender.ID][req.Chat.ID] = true
		requestMutex.Unlock()

		log.Printf("📥 Join request: user=%d chat=%d", req.Sender.ID, req.Chat.ID)
		return nil
	})

	// ▶ START
	b.Handle("/start", func(c tele.Context) error {
		missing := notAllowedChannels(b, c.Sender().ID)
		if len(missing) > 0 {
			return sendSubMessage(c, missing)
		}
		return c.Send("✅ Botga xush kelibsiz!", menu)
	})

	// 🔄 Tekshirish tugmasi
	b.Handle(&tele.Btn{Unique: "check_sub"}, func(c tele.Context) error {
		missing := notAllowedChannels(b, c.Sender().ID)
		if len(missing) == 0 {
			_ = c.Delete()
			return c.Send("✅ Hammasi joyida, davom eting!", menu)
		}
		return c.Respond(&tele.CallbackResponse{
			Text:      "❌ Hali hamma kanal bajarilmadi",
			ShowAlert: true,
		})
	})

	// 📨 HAR BIR XABARDA TEKSHIRUV
	b.Handle(tele.OnText, func(c tele.Context) error {
		missing := notAllowedChannels(b, c.Sender().ID)
		if len(missing) > 0 {
			return sendSubMessage(c, missing)
		}

		switch c.Text() {
		case "🖋️ anme izlash",
			"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
			"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
			"21", "22", "23", "24", "25", "26", "27", "28", "29", "30",
			"31", "32", "33", "34", "35", "36", "37", "38", "39", "40",
			"41", "42", "43", "44", "45", "46", "47", "48", "49", "50",
			"51", "52", "53", "54", "55", "56", "57", "58", "59", "60",
			"61", "62", "63", "64", "65", "66", "67", "68", "69", "70",
			"71", "72", "73", "74", "75", "76", "77", "78", "79", "80",
			"81", "82", "83", "84", "85", "86", "87", "88", "89", "90",
			"91", "92", "93", "94", "95", "96", "97", "98", "99", "100",
			"101", "102", "103", "104", "105", "106", "107", "108", "109", "110",
			"111", "112", "113", "114", "115", "116", "117", "118", "119", "120",
			"121", "122", "123", "124", "125", "126", "127", "128", "129", "130",
			"131", "132", "133", "134", "135", "136", "137", "138", "139", "140",
			"141", "142", "143", "144", "145", "146", "147", "148", "149", "150",
			"151", "152", "153", "154", "155", "156", "157", "158", "159", "160",
			"161", "162", "163", "164", "165", "166", "167", "168", "169", "170",
			"171", "172", "173", "174", "175", "176", "177", "178", "179", "180",
			"181", "182", "183", "184", "185", "186", "187", "188", "189", "190",
			"191", "192", "193", "194", "195", "196", "197", "198", "199", "200":
			// bu yerda kodi bajariladi

			return anmelaruzb.Home(c)
		case "Naruto", "Iblislar qotili", "Franksdagisevgi", "Vanpis", "Franksdagi sevgi",
			"Tokiyoqasoskorlari", "Xarobalar qiroligi", "Daho shahzodani mamlakatni qutqargani haqida",
			"Qirol oʻyini", "Yetti O'lim Gunohlari", "Parazit - Hayot Saboqlari", "Doktor stoun",
			"Sakamoto Kunlari", "Uyatchang Qahramon va Qotil Malikalar", "Tungi Boyqushlar Kuyi",
			"Omadsizning qayta tug'ilishi", "Uysiz Ma'bud", "Ochko'z Berserk", "Qalqon qahromoni",
			"Soyada kotarilish", "Titanlar hujumi", "Jodugarlar jangi", "Sharlota", "Soqolimni olib,",
			"Soqolimni olib, yuqori maktab qizini uyimga olib keldim", "DMC", "Qilich sanati online", "Ovoz Shakli",
			"Yozgi Urushima Tuneli Oldidagi Hayrlashuv", "Hukmdor", "Bir zarbli odam", "Quyon Qiz",
			"Sarguzashtchilar Restorani", "Horimiya", "Hyouka", "U qiz yolgiz", "Josus X Oilasi",
			"Qoshni farishta", "Aliya bazan mega rustilida nos karashma qiladi", "Yangi Saga",
			"Davolovchi qahramon", "Kayju 8-Raqam", "Pari Dumi haqida afsona", "yolg'izlikda daraja ko'taish":
			return anmelaruzb.Home(c)
		case "Qotil Akame", "Ochkoz bersek", "Qora o'q", "O‘lmas qirolning kundalik hayoti", "Zombi 100",
			"Nomsiz Xotira", "Shamolni Bo'ysundirish", "Dororo", "Arra Odam", "So'ngi serafim", "Qora Klever", "Taxt muxri",
			"Bleach", "Zulmat Farzandi", "Qudrat! Yangi Hikoya", "Yozukura Oilasi", "Poʻlat qal'adagi kabaneri",
			"Rainbow", "Qo'g'irchoqlar sirki", "Yulduz Farzandi", "Men Muvaffaqiyatsiz...", "Qahramon Boʻlish X",
			"Zulmat Iblisi", "Jahannam jannati", "Vanitas xotiralari", "Violet Evergarden", "Elita Sinfi",
			"Davolash sehridan foydalanishni noto'g'ri usuli", "Lordi Armiyasining eng kuchli Sehrgari...",
			"Tahlil qilish qobiliyatiga ega aristokrat bo'lib qayta tug'ilish", "Ilohiy qilich maktabining Iblis qilich egasi",
			"Meni Qilich bo'lib qayta tug'ilishim haqida", "O'z joniga qasd qiluvchilar o'zga dunyoda",
			"Sehr Yaratuvchi Boshqa dunyoda qanday qilib sehr yaratish mumkin", "O'rta yoshli erkakning zodagon qiziga aylanishi",
			"Eng qudratli partiya tomonidan o'limgacha tarbiyalangan Ossan ...", "Baholovchining mashhur bo'lmagan ishi aslida eng Kuchlisi hisoblanadi":
			return anmelaruzb.Home(c)
		case "Badargʻa qilingan qahramon",
			"Boshqa dunyodan muammoli bolalar", "Tokyo Gul", "Shilliq sifatida qayta tug'ilganim haqida",
			"Voleybol", "O'lim kundaligi", "Yana bir narsa soʻrasam boʻladimi", "Abadiylik qoriqchisi",
			"Yettinchi umrni betashvish yashayotgan yovuz ayol", "Taxt Dastlabki Drift", "Dragon Raja",
			"Xushboʻy Gul Viqor Bilan Gulaydi", "Xunuk Ammo Kuchli: Busamen G‘alaba Jangi", "Afsonaviy ilohy yilnomasi",
			"Uzuklar Hukumdori: Rohhirm Urushi", "Taqdir Jang Kechasi", "Qora chaqiruvchi",
			"Tahlil Qilish Qobiliyati Bilan Aristokrat Sifatida Qayta Tugʻildim", "Suv Sehrgari", "Sirlar Hukmdori",
			"Seni oshqozon osti bezingni yemoqchi man", "Qobilyatsiz Nana", "Ozga dunyoda yolgiz hujum",
			"Osmondagi Janglar", "Oltin Vaqt", "Minogra Apokalipsis", "Men galaktikalar aro imperiyaning yovuz lordiman",
			"Tajribasiz Senpai", "Cheksizlikgacha Lv9999", "Mabudlar Hohishi Bilan", "Lookism", "Domekano",
			"Gertsogning qiziga shaxsiy o'qituvchi boʻldim", "Gachiakuta Qasos", "Qudratli Soʻngi Dushman", "Yangi Darvoza",
			"Chegara ortida", "Bucchgiri", "AprelYolgoni", "Afsonaviy ilohy ilnomasi", "Sen uchun O'lmas", "qip-qizil ragna",
			"Oʻzga dunyoda darajamni koʻtarib, bu dunyoda ham tengsiz boʻldim", "Josus X Oilasi0", "Jonli Sana",
			"Ragnarok Rekordi", "Masamune Qasosi", "Kim Meni Malika Qildi", "Bir soatli qizcha", "Koʻk Zindon",
			"Biz birga bo'lsak, sevgimiz har qanday to'siqni ortda qoldiradi", "Moviy Quticha", "Hikaruning songi yozi",
			"Taqdir: Buyuk Tartib Mutlaq Iblislar Jabhasi", "Sevgi deb atalgan shart", "Lideyl Dunyosi",
			"Afsonaviy qahramonlar va Ruhlar malikasi qizi sifatida qayta tug‘ildim", "Skelet Ritsar o‘zga dunyoda":

			return anmelaruzb.Home(c)
		case "Yaponiyaga Xush Kelibsiz, Elf Xonim!", "Grimgaldagi kulgi va illuziya", "Nega hamma meni dunyomni unutdi?",
			"80.000 oltin tanga to'plab hayotimni qayta qurdim", "Cheksiz dendagron", "Oy sayohati yangi dunyoga olib keladi",
			"Man o'rgimchakman ! Ho'sh shunga nma qibti?", "O'zga dunyoda fermerlik hayotim", "Iblislar maktabiga hush kelibsiz",
			"Daholar uchun oʻzga dunyoda yashash ham muammo emas", "Menda million hayot bor", "Maktab tomonidan tan olinmagan iblislar hukmdori",
			"Boshqa dunyo bir zumda o'lim kuchiga dosh bera olmaydi", "Kumush qirolning qayta tugʻilishi", "Oʻzga dunyoda Noyob mahoratim",
			"Egzartis boshqa dunyoda qayta tug'ilib eng kuchli bo'lishga intiladi", "Nikoh uzuklari haqida afsona", "tanyang urush yilnomalari",
			"Yovuzlik darajasi 99: Men yashirin xo'jayin bo'lishim mumkin, lekin men jin xo'jayini emasman", "Doktor Eliza: hayotlarni saqlab qoluvchi malika",
			"Eng zaif yirtqich hayvon", "Fojiaga sababchi bõlgan malika xalq uchun qõlidan kelgan barcha ishni qiladi",
			"Iblislar hukmdori qoʻshimcha ishda!", "Oying bilan birga video o'yin", "Bekami Ko'st Yashashim Ushun Hamma Narsa Qildim",
			"Uzoq paladin", "Realizm qahramoni qirollikni qayta qurishi", "Olti barg qahramonlari", "Daho Shifokorning Soyadagi Yangi Hayoti",
			"Qayta tugʻilgan aristokratning misli koʻrilmagan sarguzashtlari", "Sakkizinchi o'g'il, qo'ysangizlarchi?", "O'yindagi Eng Boy odam",
			"O'yinsiz hayot yo'q", "Arifureta: Dunyodagi eng kuchli hunarmand", "Gildiya adminstratori bo'lib ortiqcha ishlashni xohlamaganim uchun ishdan ketmoqchiman",
			"Oddiy insondan qahramonlikkacha", "Tayoq va qilich", "Re:Zero", "Onmyo Qayta Tug‘ilishi: Hayolot Olami",
			"Yettinchi shahzoda sifatida qayta tug'ildim va endi sehrimni istaganimcha kuchaytiraman!", "Mening Qotillik Maqomim Qahramonlik Maqomidan Yaxshiroq",
			"Yugurening abadiyligi", "Daydi itlarning buyugi", "Meni qizcham nafaqat go'zal":

			return anmelaruzb.Home(c)

		case "Animelar":
			return Menu.Home(c)
		case "🧩 help":
			return Help.Home(c)
		default:
			return anmelaruzb.Home(c)
		}
	})

	log.Println("🤖 Bot ishga tushdi")
	b.Start()
}

// 📢 Obuna xabari
func sendSubMessage(c tele.Context, missing []ChannelInfo) error {
	text := "<b>❗ Botdan foydalanish uchun quyidagi kanallarga a'zo bo‘ling yoki so‘rov yuboring:</b>"
	m := &tele.ReplyMarkup{}
	var rows []tele.Row

	for _, ch := range missing {
		rows = append(rows, m.Row(
			m.URL("📢 "+ch.Name, ch.Invite),
		))
	}

	rows = append(rows, m.Row(
		m.Data("✅ Tekshirish", "check_sub"),
	))

	m.Inline(rows...)
	return c.Send(text, m, tele.ModeHTML)
}

//package ko_di
//
//import (
//	beego "github.com/beego/beego/v2/server/web"
//	"log"
//	"namelaruzb_bot/kodi/Help"
//	"namelaruzb_bot/kodi/Menu"
//	"namelaruzb_bot/kodi/anmelaruzb"
//	"sync" // Mutex uchun
//	"time"
//
//	tele "gopkg.in/telebot.v4"
//)
//
//// --- MA'LUMOTLARNI SAQLASH ---
//var (
//	pendingRequests = make(map[int64]bool) // So'rov yuborganlar bazasi
//	requestMutex    sync.RWMutex
//)
//
//// Kanallar ro'yxati (Username ochiq kanal uchun, Username'siz bo'lsa maxfiy deb hisoblaymiz)
//type ChannelInfo struct {
//	ID       int64
//	Username string // Ochiq bo'lsa: "username", Maxfiy bo'lsa: "" (bo'sh)
//	Invite   string // Maxfiy bo'lsa ulanish havolasi
//}
//
//var myChannels = []ChannelInfo{
//	{ID: -1003050934981, Username: "Anmelaruzb", Invite: "https://t.me/Anmelaruzb"},
//	{ID: -1003276785399, Username: "", Invite: "https://t.me/+4UV-HNVBaOs4Zjli"},
//	{ID: -1003323161290, Username: "Manga_uzbekcha26", Invite: "https://t.me/Manga_uzbekcha26"},
//	{ID: -1003411861509, Username: "", Invite: "https://t.me/+C0qmcf4ZHY83NmNi"},
//	// Maxfiy kanal bo'lsa quyidagicha qo'shasiz:
//	// {ID: -10012345678, Username:	 "", Invite: "https://t.me/+AbCdEfGhIj"},
//}
//
//// Obuna yoki so'rovni tekshirish funksiyasi
//func isUserAllowed(b *tele.Bot, user *tele.User) (bool, []ChannelInfo) {
//	// 1. Avval bazadan "Join Request" yuborganini tekshiramiz
//	requestMutex.RLock()
//	allowedByRequest := pendingRequests[user.ID]
//	requestMutex.RUnlock()
//
//	if allowedByRequest {
//		return true, nil
//	}
//
//	var missingChannels []ChannelInfo
//
//	// 2. Kanallarda bormi yoki yo'qligini tekshirish
//	for _, ch := range myChannels {
//		chat := &tele.Chat{ID: ch.ID}
//		member, err := b.ChatMemberOf(chat, user)
//
//		isMember := err == nil && (member.Role == tele.Member || member.Role == tele.Administrator || member.Role == tele.Creator)
//
//		if !isMember {
//			missingChannels = append(missingChannels, ch)
//		}
//	}
//
//	if len(missingChannels) == 0 {
//		return true, nil
//	}
//
//	return false, missingChannels
//}
//
//func Bot() {
//	token := beego.AppConfig.DefaultString("telegram::token", "")
//	pref := tele.Settings{
//		Token:  token,
//		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
//	}
//
//	b, err := tele.NewBot(pref)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	// --- 1. JOIN REQUEST TUTIB OLISH ---
//	b.Handle(tele.OnChatJoinRequest, func(c tele.Context) error {
//		userID := c.Sender().ID
//
//		requestMutex.Lock()
//		pendingRequests[userID] = true
//		requestMutex.Unlock()
//
//		log.Printf("✅ So'rov qabul qilindi: %d", userID)
//		return nil
//	})
//
//	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
//	btnID := menu.Text("🖋️ anme izlash")
//	btnMENU := menu.Text("Animelar")
//	btnHELP := menu.Text("🧩 help")
//	menu.Reply(menu.Row(btnID), menu.Row(btnMENU, btnHELP))
//
//	// --- 2. START BUYRUG'I ---
//	b.Handle("/start", func(c tele.Context) error {
//		allowed, missing := isUserAllowed(b, c.Sender())
//
//		if allowed {
//			return c.Send("✅ Botga xush kelibsiz!", menu)
//		}
//
//		return sendSubMessage(c, missing)
//	})
//
//	// --- 3. TEKSHIRISH TUGMASI ---
//	b.Handle(&tele.Btn{Unique: "check_sub"}, func(c tele.Context) error {
//		allowed, missing := isUserAllowed(b, c.Sender())
//
//		if allowed {
//			return c.Send("✅ Rahmat! Endi foydalanishingiz mumkin.", menu)
//		}
//
//		_ = c.Respond(&tele.CallbackResponse{Text: "⚠️ Hali a'zo emassiz yoki so'rov yubormadingiz!"})
//		return sendSubMessage(c, missing)
//	})
//
//	// --- 4. ASOSIY MATNLAR ---
//	b.Handle(tele.OnText, func(c tele.Context) error {
//		allowed, missing := isUserAllowed(b, c.Sender())
//		if !allowed {
//			return sendSubMessage(c, missing)
//		}
//
//		// Sizning mavjud switch/case mantiqingiz...
//		text := c.Text()
//		switch text {
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
//			"Yugurening abadiyligi", "Daydi itlarning buyugi", "Meni qizcham nafaqat go'zal":
//
//			return anmelaruzb.Home(c)
//
//		case "Animelar", "/menu":
//			return Menu.Home(c)
//		case "🧩 help", "/help":
//			return Help.Home(c)
//		default:
//			// Anime kodlarini tekshirish
//			return anmelaruzb.Home(c)
//		}
//	})
//
//	log.Println("🤖 Bot ishga tushdi....")
//	b.Start()
//}
//
//// Obuna bo'lish xabarini yuborish uchun yordamchi funksiya
//func sendSubMessage(c tele.Context, missing []ChannelInfo) error {
//	msg := "❗ Botdan foydalanish uchun quyidagi kanallarga obuna bo‘ling\n"
//	inlineMenu := &tele.ReplyMarkup{}
//	var rows []tele.Row
//
//	for _, ch := range missing {
//		label := "📢 Kanalga kirish"
//		if ch.Username != "" {
//			label = "📢 @" + ch.Username
//		}
//		btn := inlineMenu.URL(label, ch.Invite)
//		rows = append(rows, inlineMenu.Row(btn))
//	}
//
//	btnCheck := inlineMenu.Data("✅ Tekshirish", "check_sub")
//	rows = append(rows, inlineMenu.Row(btnCheck))
//	inlineMenu.Inline(rows...)
//
//	return c.Send(msg, inlineMenu)
//}

//package ko_di
//
//import (
//	beego "github.com/beego/beego/v2/server/web"
//	"log"
//	"namelaruzb_bot/kodi/Help"
//	"namelaruzb_bot/kodi/Menu"
//	"namelaruzb_bot/kodi/anmelaruzb"
//	_default "namelaruzb_bot/kodi/default"
//	"time"
//
//	tele "gopkg.in/telebot.v4"
//)
//
//var channels = []*tele.Chat{
//	{ID: -1003050934981, Username: "Anmelaruzb"},        /// 1-kanal
//	{ID: -1003276785399, Username: "animelaruzbekcha9"}, /// 2-kanal
//	{ID: -1003323161290, Username: "Manga_uzbekcha26"},  /// 3-kanal
//	//{ID: -1003316396409, Username: "anmelaruzbar_chat"},      // 4-kanal
//	//{ID: -1002328747274, Username: "uzb_FCB_fans"}, // 5-kanal
//}
//
//// Obuna bo‘lmagan kanallarni topish
//func notSubscribedChannels(b *tele.Bot, user *tele.User) []*tele.Chat {
//	var notSubs []*tele.Chat
//	for _, ch := range channels {
//		member, err := b.ChatMemberOf(ch, user)
//		if err != nil {
//			log.Println("Tekshirishda xatolik:", err)
//			notSubs = append(notSubs, ch)
//			continue
//		}
//		if !(member.Role == tele.Member || member.Role == tele.Administrator || member.Role == tele.Creator) {
//			notSubs = append(notSubs, ch)
//		}
//	}
//	return notSubs
//}
//
//func Bot() {
//	token := beego.AppConfig.DefaultString("telegram::token", "")
//	pref := tele.Settings{
//		Token:  token,
//		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
//	}
//
//	b, err := tele.NewBot(pref)
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
//	btnID := menu.Text("🖋️ anme izlash")
//	btnMENU := menu.Text("Animelar")
//	btnHELP := menu.Text("🧩 help")
//	menu.Reply(
//		menu.Row(btnID),
//		menu.Row(btnMENU, btnHELP),
//	)
//
//	b.Handle("/start", func(c tele.Context) error {
//		user := c.Sender()
//		notSubs := notSubscribedChannels(b, user)
//
//		if len(notSubs) == 0 {
//			return c.Send("✅ Botga xush kelibsiz!\n\nBotdan foydalanishingiz mumkin.", menu)
//		}
//		msg := "❗ Botdan foydalanish uchun quyidagi kanallarga obuna bo‘ling:\n"
//		inlineMenu := &tele.ReplyMarkup{}
//		var rows []tele.Row
//
//		for _, ch := range notSubs {
//			btn := inlineMenu.URL("📢 "+ch.Username, "https://t.me/"+ch.Username)
//			rows = append(rows, inlineMenu.Row(btn))
//		}
//		btnCheck := inlineMenu.Data("✅ Tekshirish", "check_sub")
//		rows = append(rows, inlineMenu.Row(btnCheck))
//		inlineMenu.Inline(rows...)
//
//		return c.Send(msg, inlineMenu)
//	})
//	b.Handle(&tele.Btn{Unique: "check_sub"}, func(c tele.Context) error {
//		user := c.Sender()
//		notSubs := notSubscribedChannels(b, user)
//
//		if len(notSubs) == 0 {
//			_ = c.Respond(&tele.CallbackResponse{Text: "✅ Obuna tasdiqlandi!"})
//			return c.Send("👋 Botga xush kelibsiz ✋", menu)
//		}
//
//		msg := "❌ Hali quyidagi kanallarga obuna bo‘lmadingiz:\n"
//		inlineMenu := &tele.ReplyMarkup{}
//		var rows []tele.Row
//
//		for _, ch := range notSubs {
//			btn := inlineMenu.URL("📢 "+ch.Username, "https://t.me/"+ch.Username)
//			rows = append(rows, inlineMenu.Row(btn))
//		}
//		btnCheck := inlineMenu.Data("✅ Tekshirish", "check_sub")
//		rows = append(rows, inlineMenu.Row(btnCheck))
//		inlineMenu.Inline(rows...)
//
//		return c.Send(msg, inlineMenu)
//	})
//
//	b.Handle(tele.OnText, func(c tele.Context) error {
//		user := c.Sender()
//
//		// 🔹 Har safar foydalanuvchi kod yuborsa, kanallarga obuna bo‘lishini tekshiramiz
//		notSubs := notSubscribedChannels(b, user)
//		if len(notSubs) > 0 {
//			msg := "❌ Avval barcha kanallarga obuna bo‘ling:\n"
//			inlineMenu := &tele.ReplyMarkup{}
//			var rows []tele.Row
//			for _, ch := range notSubs {
//				btn := inlineMenu.URL("📢 "+ch.Username, "https://t.me/"+ch.Username)
//				rows = append(rows, inlineMenu.Row(btn))
//			}
//			btnCheck := inlineMenu.Data("✅ Tekshirish", "check_sub")
//			rows = append(rows, inlineMenu.Row(btnCheck))
//			inlineMenu.Inline(rows...)
//
//			log.Printf("❗ Foydalanuvchi %s obuna emas: %+v\n", user.Username, notSubs)
//			return c.Send(msg, inlineMenu)
//		}
//
//		text := c.Text()
//
//		// 🔹 Kodlar va nomlar bo‘yicha switch
//		switch text {
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
//			"Yugurening abadiyligi", "Daydi itlarning buyugi":
//
//			return anmelaruzb.Home(c)
//
//		case "Animelar", "/menu":
//			return Menu.Home(c)
//		case "🧩 help", "/help":
//			return Help.Home(c)
//		default:
//			return _default.Home(c)
//		}
//	})
//
//	log.Println("🤖 Bot ishga tushdi....")
//	b.Start()
//}
