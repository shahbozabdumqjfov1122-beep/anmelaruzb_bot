package anmelaruzb

import (
	tele "gopkg.in/telebot.v4"
	"namelaruzb_bot/kodi/anmelaruzb/Aliyabazanmegarustilidanoskarashmaqiladi"
	"namelaruzb_bot/kodi/anmelaruzb/Bleach"
	"namelaruzb_bot/kodi/anmelaruzb/Davolovchiqahramon"
	"namelaruzb_bot/kodi/anmelaruzb/Franksdagisevgi"
	"namelaruzb_bot/kodi/anmelaruzb/Jahannamjannati"
	"namelaruzb_bot/kodi/anmelaruzb/Jodugarlarjangi"
	"namelaruzb_bot/kodi/anmelaruzb/Menmuvaffaqiyatsiz"
	"namelaruzb_bot/kodi/anmelaruzb/Naruto"
	"namelaruzb_bot/kodi/anmelaruzb/NomsizXotira"
	"namelaruzb_bot/kodi/anmelaruzb/Ochkozbosek"
	"namelaruzb_bot/kodi/anmelaruzb/Olmasqirolningkundalikhayoti"
	"namelaruzb_bot/kodi/anmelaruzb/Qalqonqahromoni"
	"namelaruzb_bot/kodi/anmelaruzb/QoraKlever"
	"namelaruzb_bot/kodi/anmelaruzb/Qoraoq"
	"namelaruzb_bot/kodi/anmelaruzb/Qoshnifarishta"
	"namelaruzb_bot/kodi/anmelaruzb/QotilAkame"
	"namelaruzb_bot/kodi/anmelaruzb/Saksonolti"
	"namelaruzb_bot/kodi/anmelaruzb/Sharlota"
	"namelaruzb_bot/kodi/anmelaruzb/Songiserafim"
	"namelaruzb_bot/kodi/anmelaruzb/Soyadakotarilish"
	"namelaruzb_bot/kodi/anmelaruzb/Titanlarjangi"
	"namelaruzb_bot/kodi/anmelaruzb/Vanitasxotiralari"
	"namelaruzb_bot/kodi/anmelaruzb/Xarobalarqiroligi"
	"namelaruzb_bot/kodi/anmelaruzb/Zombi100"
	"namelaruzb_bot/kodi/anmelaruzb/dahoshahzodanimamlakat"
	"namelaruzb_bot/kodi/anmelaruzb/iblislarqotili"
	"namelaruzb_bot/kodi/anmelaruzb/tokiyoqasoskorlari"
	"namelaruzb_bot/kodi/anmelaruzb/vanpis"
	"namelaruzb_bot/kodi/anmelaruzb/yolgizlikdadarajakotaish"
)

func Home(c tele.Context) error {
	text := c.Text()
	switch text {
	case "🖋️ anme izlash":
		return c.Send("ID-kiriting:")
	case "1", "Naruto":
		return Naruto.Home(c)
	case "2", "Iblislar qotili":
		return iblislarqotili.Home(c)
	case "3", "Franksdagisevgi":
		return Franksdagisevgi.Home(c)
	case "4", "Daho shahzodani mamlakatni qutqargani haqida":
		return dahoshahzodanimamlakat.Home(c)
	case "5", "Xarobalar qiroligi":
		return Xarobalarqiroligi.Home(c)
	case "6", "Soyada kotarilish":
		return Soyadakotarilish.Home(c)
	case "7", "Vanpis":
		return vanpis.Home(c)
	case "8", "Tokiyoqasoskorlari":
		return tokiyoqasoskorlari.Home(c)
	case "9", "Titanlar hujumi":
		return Titanlarjangi.Home(c)
	case "10", "Jodugarlar jangi":
		return Jodugarlarjangi.Home(c)
	case "11", "Qalqon qahromoni":
		return Qalqonqahromoni.Home(c)
	case "12", "Sharlota":
		return Sharlota.Home(c)
	case "13", "Qoshni farishta":
		return Qoshnifarishta.Home(c)
	case "14", "yolg'izlikda daraja ko'taish":
		return yolgizlikdadarajakotaish.Home(c)
	case "15", "Aliya bazan mega rustilida nos karashma qiladi":
		return Aliyabazanmegarustilidanoskarashmaqiladi.Home(c)
	case "16", "Davolovchi qahramon":
		return Davolovchiqahramon.Home(c)
	case "17", "So'ngi serafim":
		return Songiserafim.Home(c)
	case "86":
		return Saksonolti.Home(c)
	case "19", "Ochkoz bersek":
		return Ochkozbosek.Home(c)
	case "20", "Zombi 100":
		return Zombi100.Home(c)
	case "21", "Nomsiz Xotira":
		return NomsizXotira.Home(c)
	case "22", "Qotil Akame":
		return QotilAkame.Home(c)
	case "23", "Jahannam jannati":
		return Jahannamjannati.Home(c)
	case "24", "Qora Klever":
		return QoraKlever.Home(c)
	case "25", "Vanitas xotiralari":
		return Vanitasxotiralari.Home(c)
	case "26", "Bleach":
		return Bleach.Home(c)
	case "27", "O‘lmas qirolning kundalik hayoti":
		return Olmasqirolningkundalikhayoti.Home(c)
	case "28", "Qora o'q":
		return Qoraoq.Home(c)
	case "29", "Men Muvaffaqiyatsiz...":
		return Menmuvaffaqiyatsiz.Home(c)

	//
	//
	//
	default:
		return c.Send("Noma'lum buyruq. Iltimos, /menu dan foydalaning.")
	}
}
