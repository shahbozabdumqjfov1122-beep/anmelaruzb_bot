package Menu

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	text := c.Text()

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnNaruto := menu.Text("Naruto")
	btnIblislarqotili := menu.Text("Iblislar qotili")
	btnFranksdagisevgi := menu.Text("Franksdagisevgi")
	btnVanpis := menu.Text("Vanpis")
	btnTokiyoqasoskorlari := menu.Text("Tokiyoqasoskorlari")
	btnXarobalarqiroligi := menu.Text("Xarobalar qiroligi")
	btnDahoshahzodanimamlakat := menu.Text("Daho shahzodani mamlakatni qutqargani haqida")
	btnSoyadakotarilish := menu.Text("Soyada kotarilish")
	btnTitanlarhujumi := menu.Text("Titanlar hujumi")
	btnJodugarlarjangi := menu.Text("Jodugarlar jangi")
	btnQalqonqahromoni := menu.Text("Qalqon qahromoni")
	btnSharlota := menu.Text("Sharlota")
	btnQoshnifarishta := menu.Text("Qoshni farishta")
	btnAliyabazanmegarustilidanoskarashmaqiladi := menu.Text("Aliya bazan mega rustilida nos karashma qiladi")
	btnYolgizlikdadarajakotaish := menu.Text("yolg'izlikda daraja ko'taish")
	btnDavolovchiqahramon := menu.Text("Davolovchi qahramon")
	btnSongiserafim := menu.Text("So'ngi serafim")
	btnSaksonolti := menu.Text("86")
	btnOchkozbersek := menu.Text("Ochkoz bersek")
	btnQotilAkame := menu.Text("Qotil Akame")
	btnZombi100 := menu.Text("Zombi 100")
	btnNomsizXotira := menu.Text("Nomsiz Xotira")
	btnQoraKlever := menu.Text("Qora Klever")
	btnVanitasxotiralari := menu.Text("Vanitas xotiralari")
	btnJahannamjannati := menu.Text("Jahannam jannati")
	btnBleach := menu.Text("Bleach")
	btnOlmasqirolningkundalikhayoti := menu.Text("O‘lmas qirolning kundalik hayoti")
	btnQoraoq := menu.Text("Qora o'q")
	btnMenmuvaffaqiyatsiz := menu.Text("Men Muvaffaqiyatsiz...")
	btnSakamotokunlari := menu.Text("Sakamoto Kunlari")
	btnTungiBoyqushlarKuyi := menu.Text("Tungi Boyqushlar Kuyi")
	btnZulmatiblisi := menu.Text("Zulmat Iblisi")
	btnQahramonBoʻlishX := menu.Text("Qahramon Boʻlish X")
	btnYozukuraOilasi := menu.Text("Yozukura Oilasi")
	btnZulmatFarzandi := menu.Text("Zulmat Farzandi")
	btnQudratYangiHikoya := menu.Text("Qudrat! Yangi Hikoya")
	btnYangiSaga := menu.Text("Yangi Saga")
	btnUyatchangQahramonvaQotilMalikalar := menu.Text("Uyatchang Qahramon va Qotil Malikalar")
	btnDoktorstoun := menu.Text("Doktor stoun")
	btnTaxtmuxri := menu.Text("Taxt muxri")

	menu.Reply(
		menu.Row(btnNaruto, btnIblislarqotili),
		menu.Row(btnFranksdagisevgi, btnVanpis),
		menu.Row(btnTokiyoqasoskorlari, btnXarobalarqiroligi),
		menu.Row(btnDahoshahzodanimamlakat),
		menu.Row(btnSoyadakotarilish),
		menu.Row(btnTitanlarhujumi, btnJodugarlarjangi),
		menu.Row(btnQalqonqahromoni, btnSharlota),
		menu.Row(btnQoshnifarishta, btnAliyabazanmegarustilidanoskarashmaqiladi),
		menu.Row(btnYolgizlikdadarajakotaish, btnDavolovchiqahramon),
		menu.Row(btnSongiserafim, btnSaksonolti),
		menu.Row(btnOchkozbersek, btnQotilAkame),
		menu.Row(btnZombi100, btnNomsizXotira),
		menu.Row(btnQoraKlever, btnVanitasxotiralari),
		menu.Row(btnJahannamjannati, btnBleach),
		menu.Row(btnOlmasqirolningkundalikhayoti, btnQoraoq),
		menu.Row(btnMenmuvaffaqiyatsiz, btnSakamotokunlari),
		menu.Row(btnTungiBoyqushlarKuyi, btnZulmatiblisi),
		menu.Row(btnQahramonBoʻlishX, btnYozukuraOilasi),
		menu.Row(btnZulmatFarzandi, btnQudratYangiHikoya),
		menu.Row(btnYangiSaga, btnUyatchangQahramonvaQotilMalikalar),
		menu.Row(btnDoktorstoun, btnTaxtmuxri),
	)

	switch text {

	case "Animelar", "/menu":
		return c.Send("asosiy kanal -https://t.me/Anmelaruzb ", menu)
	default:
		return c.Send("Iltimos, /menu dan danlang.")
	}
}
