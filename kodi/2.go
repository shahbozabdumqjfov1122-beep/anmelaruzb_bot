package ko_di

import (
	beego "github.com/beego/beego/v2/server/web"
	"log"
	"namelaruzb_bot/kodi/Help"
	"namelaruzb_bot/kodi/Menu"
	"namelaruzb_bot/kodi/anmelaruzb"
	_default "namelaruzb_bot/kodi/default"
	"time"

	tele "gopkg.in/telebot.v4"
)

var channels = []*tele.Chat{
	{ID: -1003050934981, Username: "Anmelaruzb"},        // 1-kanal
	{ID: -1003276785399, Username: "animelaruzbekcha9"}, // 2-kanal
	//{ID: -1002328747274, Username: "uzb_FCB_fans"}, // 3-kanal
	//{ID: -1003185463264, Username: "abu_w"},        // 4-kanal
}

// Obuna bo‘lmagan kanallarni topish
func notSubscribedChannels(b *tele.Bot, user *tele.User) []*tele.Chat {
	var notSubs []*tele.Chat
	for _, ch := range channels {
		member, err := b.ChatMemberOf(ch, user)
		if err != nil {
			log.Println("Tekshirishda xatolik:", err)
			notSubs = append(notSubs, ch)
			continue
		}
		if !(member.Role == tele.Member || member.Role == tele.Administrator || member.Role == tele.Creator) {
			notSubs = append(notSubs, ch)
		}
	}
	return notSubs
}

func Bot() {
	token := beego.AppConfig.DefaultString("telegram::token", "")
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnID := menu.Text("🖋️ anme izlash")
	btnMENU := menu.Text("Animelar")
	btnHELP := menu.Text("🧩 help")
	menu.Reply(
		menu.Row(btnID),
		menu.Row(btnMENU, btnHELP),
	)

	b.Handle("/start", func(c tele.Context) error {
		user := c.Sender()
		notSubs := notSubscribedChannels(b, user)

		if len(notSubs) == 0 {
			return c.Send("✅ Botga xush kelibsiz!\n\nBotdan foydalanishingiz mumkin.", menu)
		}
		msg := "❗ Botdan foydalanish uchun quyidagi kanallarga obuna bo‘ling:\n"
		inlineMenu := &tele.ReplyMarkup{}
		var rows []tele.Row

		for _, ch := range notSubs {
			btn := inlineMenu.URL("📢 "+ch.Username, "https://t.me/"+ch.Username)
			rows = append(rows, inlineMenu.Row(btn))
		}
		btnCheck := inlineMenu.Data("✅ Tekshirish", "check_sub")
		rows = append(rows, inlineMenu.Row(btnCheck))
		inlineMenu.Inline(rows...)

		return c.Send(msg, inlineMenu)
	})
	b.Handle(&tele.Btn{Unique: "check_sub"}, func(c tele.Context) error {
		user := c.Sender()
		notSubs := notSubscribedChannels(b, user)

		if len(notSubs) == 0 {
			_ = c.Respond(&tele.CallbackResponse{Text: "✅ Obuna tasdiqlandi!"})
			return c.Send("👋 Botga xush kelibsiz ✋", menu)
		}

		msg := "❌ Hali quyidagi kanallarga obuna bo‘lmadingiz:\n"
		inlineMenu := &tele.ReplyMarkup{}
		var rows []tele.Row

		for _, ch := range notSubs {
			btn := inlineMenu.URL("📢 "+ch.Username, "https://t.me/"+ch.Username)
			rows = append(rows, inlineMenu.Row(btn))
		}
		btnCheck := inlineMenu.Data("✅ Tekshirish", "check_sub")
		rows = append(rows, inlineMenu.Row(btnCheck))
		inlineMenu.Inline(rows...)

		return c.Send(msg, inlineMenu)
	})
	b.Handle(tele.OnText, func(c tele.Context) error {
		user := c.Sender()
		if len(notSubscribedChannels(b, user)) > 0 {
			return c.Send("❌ Avval barcha kanallarga obuna bo‘ling!", nil)
		}

		text := c.Text()
		switch text {
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

			return anmelaruzb.Home(c)
		case "Naruto", "Iblislar qotili", "Franksdagi sevgi", "Vanpis",
			"Tokiyo qasoskorlari", "Xarobalar qiroligi", "Daho shahzodani mamlakatni qutqargani haqida",
			"Doktor stoun", "Sakamoto Kunlari", "Uyatchang Qahramon va Qotil Malikalar", "Tungi Boyqushlar Kuyi",
			"Qalqon qahromoni", "Soyada kotarilish", "Titanlar hujumi", "Jodugarlar jangi", "Sharlota",
			"Qoshni farishta", "Aliya bazan mega rustilida nos karashma qiladi", "Yangi Saga", "Davolovchi qahramon",
			"yolg'izlikda daraja ko'taish":
			return anmelaruzb.Home(c)
		case "Qotil Akame", "Ochkoz bersek", "Qora o'q", "O‘lmas qirolning kundalik hayoti",
			"Zombi 100", "Nomsiz Xotira", "So'ngi serafim", "Qora Klever", "Taxt muxri", "Bleach",
			"Zulmat Farzandi", "Qudrat! Yangi Hikoya", "Yozukura Oilasi", "Men Muvaffaqiyatsiz...",
			"Qahramon Boʻlish X", "Zulmat Iblisi", "Jahannam jannati", "Vanitas xotiralari", "Arra Odam", "Poʻlat qal'adagi kabaneri",
			"Rainbow", "Qo'g'irchoqlar sirki":
			return anmelaruzb.Home(c)
		case "Badargʻa qilingan qahramon",
			"Boshqa dunyodan muammoli bolalar", "Taxt Dastlabki Drift", "Dragon Raja", "Xushboʻy Gul Viqor Bilan Gulaydi",
			"Xunuk Ammo Kuchli: Busamen G‘alaba Jangi", "Uzuklar Hukumdori: Rohhirm Urushi", "Taqdir Jang Kechasi",
			"Tahlil Qilish Qobiliyati Bilan Aristokrat Sifatida Qayta Tugʻildim", "Suv Sehrgari", "Sirlar Hukmdori",
			"Seni oshqozon osti bezingni yemoqchi man", "Qobilyatsiz Nana", "Ozga dunyoda yolgiz hujum", "Osmondagi Janglar", "Oltin Vaqt",
			"Minogra Apokalipsis", "Men galaktikalar aro imperiyaning yovuz lordiman", "Mabudlar Hohishi Bilan",
			"Lookism", "Gertsogning qiziga shaxsiy o'qituvchi boʻldim", "Gachiakuta Qasos", "Domekano", "Chegara ortida",
			"Bucchgiri", "AprelYolgoni", "Abadiylik qoriqchisi", "Yettinchi umrni betashvish yashayotgan yovuz ayol", "Afsonaviy ilohy ilnomasi", "O'lim kundaligi", "Voleybol", "Tokyo Gul",
			"Horimiya", "Ovoz Shakli", "Quyon Qiz", "Hukmdor", "Bir zarbli odam", "Sarguzashtchilar Restorani", "Shilliq sifatida qayta tug'ilganim haqida",
			"Qilich sanati online", "Yana bir narsa soʻrasam boʻladimi", "DMC", "Ochko'z Berserk", "Shamolni Bo'ysundirish", "Dororo":
			return anmelaruzb.Home(c)

		case "Animelar", "/menu":
			return Menu.Home(c)
		case "🧩 help", "/help":
			return Help.Home(c)
		default:
			return _default.Home(c)
		}
	})

	log.Println("🤖 Bot ishga tushdi....")
	b.Start()
}
