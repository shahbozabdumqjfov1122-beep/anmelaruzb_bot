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
	{ID: -1003095728984}, // 1-kanal
	{ID: -1003050934981}, // 2-kanal
	{ID: -1002328747274}, // 3-kanal
}

func isSubscribed(b *tele.Bot, user *tele.User) bool {
	for _, ch := range channels {
		member, err := b.ChatMemberOf(ch, user)
		if err != nil {
			log.Println("Tekshirishda xatolik:", err)
			return false
		}
		if !(member.Role == tele.Member || member.Role == tele.Administrator || member.Role == tele.Creator) {
			return false
		}
	}
	return true
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
	inlineMenu := &tele.ReplyMarkup{}
	btnCh1 := inlineMenu.URL("📢 Kanal 1", "https://t.me/beshariq_365")
	btnCh2 := inlineMenu.URL("📢 Kanal 2", "https://t.me/Anmelaruzb")
	btnCh3 := inlineMenu.URL("📢 Kanal 3", "https://t.me/uzb_FCB_fans")
	btnCheck := inlineMenu.Data("✅ Tekshirish", "check_sub")
	inlineMenu.Inline(
		inlineMenu.Row(btnCh1),
		inlineMenu.Row(btnCh2),
		inlineMenu.Row(btnCh3),
		inlineMenu.Row(btnCheck),
	)
	b.Handle("/start", func(c tele.Context) error {
		user := c.Sender()
		if isSubscribed(b, user) {
			return c.Send(" Botga xush kelibsiz \nbotga joylangan anmlerni ko'rish uchun---/menu"+
				"\nbizni botimizdan bironta  anme ni "+
				"\nqidirib topolmasagiz adminlarga murojat qiling\nmurojat uchun bosing---/help  ", menu)
		}
		msg := "❗ Botdan foydalanish uchun avval barcha kanallarga obuna bo‘ling."
		return c.Send(msg, inlineMenu)
	})
	b.Handle(&tele.Btn{Unique: "check_sub"}, func(c tele.Context) error {
		user := c.Sender()

		if isSubscribed(b, user) {
			_ = c.Respond(&tele.CallbackResponse{Text: "✅ Obuna tasdiqlandi!"})
			return c.Send("👋 Botga xush kelibsiz ✋", menu)
		}

		return c.Respond(&tele.CallbackResponse{Text: "❌ Avval barcha kanallarga obuna bo‘ling!"})
	})
	b.Handle(tele.OnText, func(c tele.Context) error {
		user := c.Sender()
		if !isSubscribed(b, user) {
			return c.Send("❌ Avval barcha kanallarga obuna bo‘ling!", inlineMenu)
		}
		text := c.Text()
		switch text {
		case "🖋️ anme izlash", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20":
			return anmelaruzb.Home(c)
		case "Naruto", "Iblislar qotili", "86", "Franksdagisevgi", "Vanpis", "Tokiyoqasoskorlari", "Xarobalar qiroligi", "Daho shahzodani mamlakatni qutqargani haqida",
			"Qalqon qahromoni", "Soyada kotarilish", "Titanlar hujumi", "Jodugarlar jangi", "Sharlota", "Qoshni farishta", "Aliya bazan mega rustilida nos karashma qiladi", "Davolovchi qahramon", "yolg'izlikda daraja ko'taish":
			return anmelaruzb.Home(c)
		case "Qotil Akame", "Ochkoz bersek", "Zombi 100", "Nomsiz Xotira", "So'ngi serafim":
			return anmelaruzb.Home(c)
		case "Animelar", "/menu":
			return Menu.Home(c)
		case "🧩 help", "/help":
			return Help.Home(c)

		default:
			return _default.Home(c)
			/*c.Send("Uzr bu kod noto'g'ri yoki hozircha mavjud emas!" +
			"\n\n Iltimos, menu dan foydalaning 🌎 !")*/
		}
	})
	log.Println("🤖 Bot ishga tushdi...")
	b.Start()
}
