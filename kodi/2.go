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
	{ID: -1003095728984, Username: "beshariq_365"}, // 1-kanal
	{ID: -1003050934981, Username: "Anmelaruzb"},   // 2-kanal
	{ID: -1002328747274, Username: "uzb_FCB_fans"}, // 3-kanal
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
	btnRASMLAR := menu.Text("📷 Rasmlar")
	menu.Reply(
		menu.Row(btnRASMLAR, btnID),
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
		case "🖋️ anme izlash", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20":
			return anmelaruzb.Home(c)
		case "Naruto", "Iblislar qotili", "86", "Franksdagisevgi", "Vanpis", "Tokiyoqasoskorlari", "Xarobalar qiroligi", "Daho shahzodani mamlakatni qutqargani haqida",
			"Qalqon qahromoni", "Soyada kotarilish", "Titanlar hujumi", "Jodugarlar jangi", "Sharlota", "Qoshni farishta", "Aliya bazan mega rustilida nos karashma qiladi", "Davolovchi qahramon", "yolg'izlikda daraja ko'taish":
			return anmelaruzb.Home(c)
		case "Qotil Akame", "Ochkoz bersek", "Qora o'q", "O‘lmas qirolning kundalik hayoti", "Zombi 100", "Nomsiz Xotira", "So'ngi serafim", "Qora Klever", "Bleach", "Jahannam jannati", "Vanitas xotiralari":
			return anmelaruzb.Home(c)
		case "Animelar", "/menu":
			return Menu.Home(c)
		case "🧩 help", "/help":
			return Help.Home(c)
		default:
			return _default.Home(c)
		}
	})

	log.Println("🤖 Bot ishga tushdi...")
	b.Start()
}
