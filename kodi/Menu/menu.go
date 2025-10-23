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
	btnQahramonBolishX := menu.Text("Qahramon Boʻlish X")
	btnYozukuraOilasi := menu.Text("Yozukura Oilasi")
	btnZulmatFarzandi := menu.Text("Zulmat Farzandi")
	btnQudratYangiHikoya := menu.Text("Qudrat! Yangi Hikoya")
	btnYangiSaga := menu.Text("Yangi Saga")
	btnUyatchangQahramonvaQotilMalikalar := menu.Text("Uyatchang Qahramon va Qotil Malikalar")
	btnDoktorstoun := menu.Text("Doktor stoun")
	btnTaxtmuxri := menu.Text("Taxt muxri")
	btnBadarga := menu.Text("Badargʻa qilingan qahramon")
	btnBoshqadunyo := menu.Text("Boshqa dunyodan muammoli bolalar")
	btnDrift := menu.Text("Taxt Dastlabki Drift")
	btnDragonRaja := menu.Text("Dragon Raja")
	btnXushboy := menu.Text("Xushboʻy Gul Viqor Bilan Gulaydi")
	btnXunuk := menu.Text("Xunuk Ammo Kuchli: Busamen G‘alaba Jangi")
	btnUzuklar := menu.Text("Uzuklar Hukumdori: Rohhirm Urushi")
	btnTaqdir := menu.Text("Taqdir Jang Kechasi")
	btnTahlil := menu.Text("Tahlil Qilish Qobiliyati Bilan Aristokrat Sifatida Qayta Tugʻildim")
	btnSuv := menu.Text("Suv Sehrgari")
	btnSirlar := menu.Text("Sirlar Hukmdori")
	btnSeniOshqozon := menu.Text("Seni oshqozon osti bezingni yemoqchi man")
	btnNana := menu.Text("Qobilyatsiz Nana")
	btnOzgaDunyo := menu.Text("Ozga dunyoda yolgiz hujum")
	btnOsmondagi := menu.Text("Osmondagi Janglar")
	btnOltin := menu.Text("Oltin Vaqt")
	btnMinogra := menu.Text("Minogra Apokalipsis")
	btnMenGalaktika := menu.Text("Men galaktikalar aro imperiyaning yovuz lordiman")
	btnMabudlar := menu.Text("Mabudlar Hohishi Bilan")
	btnLookism := menu.Text("Lookism")
	btnGertsog := menu.Text("Gertsogning qiziga shaxsiy o'qituvchi boʻldim")
	btnGachiakuta := menu.Text("Gachiakuta Qasos")
	btnDomekano := menu.Text("Domekano")
	btnChegara := menu.Text("Chegara ortida")
	btnBucchgiri := menu.Text("Bucchgiri")
	btnAprel := menu.Text("AprelYolgoni")
	btnAfsona := menu.Text("Afsonaviy ilohy ilnomasi")
	btnYettinchiumrni := menu.Text("Yettinchi umrni betashvish yashayotgan yovuz ayol")
	btnAbadiylikqoriqchisi := menu.Text("Abadiylikqoriqchisi")

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
		menu.Row(btnQahramonBolishX, btnYozukuraOilasi),
		menu.Row(btnZulmatFarzandi, btnQudratYangiHikoya),
		menu.Row(btnYangiSaga, btnUyatchangQahramonvaQotilMalikalar),
		menu.Row(btnDoktorstoun, btnTaxtmuxri),
		menu.Row(btnBadarga, btnBoshqadunyo),
		menu.Row(btnDrift, btnDragonRaja),
		menu.Row(btnXushboy, btnXunuk),
		menu.Row(btnUzuklar, btnTaqdir),
		menu.Row(btnTahlil),
		menu.Row(btnSirlar, btnSuv),
		menu.Row(btnSeniOshqozon),
		menu.Row(btnNana, btnOzgaDunyo),
		menu.Row(btnOsmondagi, btnOltin),
		menu.Row(btnMinogra, btnGachiakuta),
		menu.Row(btnMenGalaktika),
		menu.Row(btnMabudlar, btnLookism),
		menu.Row(btnGertsog),
		menu.Row(btnDomekano, btnChegara),
		menu.Row(btnBucchgiri, btnAprel),
		menu.Row(btnAfsona, btnAbadiylikqoriqchisi),
		menu.Row(btnYettinchiumrni),
	)

	switch text {

	case "Animelar", "/menu":
		return c.Send("asosiy kanal -https://t.me/Anmelaruzb ", menu)
	default:
		return c.Send("Iltimos, /menu dan danlang.")
	}
}
