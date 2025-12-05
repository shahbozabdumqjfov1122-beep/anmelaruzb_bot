package Menu

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	text := c.Text()

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btnNaruto := menu.Text("Naruto")
	btnIblislarqotili := menu.Text("Iblislar qotili")
	btnFranksdagisevgi := menu.Text("Franksdagi sevgi")
	btnVanpis := menu.Text("Vanpis")
	btnTokiyoqasoskorlari := menu.Text("Tokiyo qasoskorlari")
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
	//btnUyatchangQahramonvaQotilMalikalar := menu.Text("Uyatchang Qahramon va Qotil Malikalar")
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
	btnAbadiylikqoriqchisi := menu.Text("Abadiylik qoriqchisi")
	btnYanabirnarsasorasamboladimi := menu.Text("Yana bir narsa soʻrasam boʻladimi")
	btnOlimkundaligi := menu.Text("O'lim kundaligi")
	btnVoleybol := menu.Text("Voleybol")
	btnShilliqsifatida := menu.Text("Shilliq sifatida qayta tug'ilganim haqida")
	btnTokyoGul := menu.Text("Tokyo Gul")
	btnHorimiya := menu.Text("Horimiya")
	btnSarguzashtchilarRestorani := menu.Text("Sarguzashtchilar Restorani")
	btnOvozShakli := menu.Text("Ovoz Shakli")
	btnQuyonQiz := menu.Text("Quyon Qiz")
	btnBirzarbliodam := menu.Text("Bir zarbli odam")
	btnHukmdor := menu.Text("Hukmdor")
	btnQilichsanationline := menu.Text("Qilich sanati online")
	btnDMC := menu.Text("DMC")
	btnOchkozBerserk := menu.Text("Ochko'z Berserk")
	btnShamolniBoysundirish := menu.Text("Shamolni Bo'ysundirish")
	btnDororo := menu.Text("Dororo")
	btnArraOdam := menu.Text("Arra Odam")
	btnPolatqaladagikabaneri := menu.Text("Poʻlat qal'adagi kabaneri")
	btnRainbow := menu.Text("Rainbow")
	btnQogirchoqlarsirki := menu.Text("Qo'g'irchoqlar sirki")
	btnYulduzFarzandi := menu.Text("Yulduz Farzandi")
	btnElitaSinfi := menu.Text("Elita Sinfi")
	btnTajribasizSenpai := menu.Text("Tajribasiz Senpai")
	btnQudratliSongiDushman := menu.Text("Qudratli Soʻngi Dushman")
	btnYangiDarvoza := menu.Text("Yangi Darvoza")
	btnCheksizlikgachaLv9999 := menu.Text("Cheksizlikgacha Lv9999")
	btnJosusXOilasi := menu.Text("Josus X Oilasi")
	btnRagnarokRekordi := menu.Text("Ragnarok Rekordi")
	btnSenuchunOlmas := menu.Text("Sen uchun O'lmas")
	btnMasamuneQasosi := menu.Text("Masamune Qasosi")
	btnKimMeniMalikaQildi := menu.Text("Kim Meni Malika Qildi")
	btnKokZindon := menu.Text("Ko'k Zindon")
	btnParazitHayotSaboqlari := menu.Text("Parazit - Hayot Saboqlari")
	btnUysizMabud := menu.Text("Uysiz Ma'bud")
	btnPariDumihaqidaafsona := menu.Text("Pari Dumi haqida afsona")
	btnKayjuRaqam := menu.Text("Kayju 8-Raqam")
	btnSoqolimniolib := menu.Text("Soqolimni olib, yuqori maktab qizini uyimga olib keldim")
	btnYettiOlimGunohlari := menu.Text("Yetti O'lim Gunohlari")
	btnOmadsizningqaytatugilishi := menu.Text("Omadsizning qayta tug'ilishi")
	btnQiroloyini := menu.Text("Qirol oʻyini")
	btnVioletEvergarden := menu.Text("Violet Evergarden")
	btnOzgadunyoda := menu.Text("Oʻzga dunyoda darajamni koʻtarib, bu dunyoda ham tengsiz boʻldim")
	btnHyouka := menu.Text("Hyouka")
	btnUqizyolgiz := menu.Text("U qiz yolgiz")
	btnBirsoatliqizcha := menu.Text("Bir soatli qizcha")
	btnJonliSana := menu.Text("Jonli Sana")

	menu.Reply(
		menu.Row(btnVanpis, btnNaruto),
		menu.Row(btnFranksdagisevgi, btnIblislarqotili),
		menu.Row(btnTokiyoqasoskorlari, btnXarobalarqiroligi),
		menu.Row(btnDahoshahzodanimamlakat),
		menu.Row(btnSoyadakotarilish),
		menu.Row(btnTitanlarhujumi),
		menu.Row(btnJodugarlarjangi),
		menu.Row(btnSharlota),
		menu.Row(btnQalqonqahromoni),
		menu.Row(btnQoshnifarishta),
		menu.Row(btnAliyabazanmegarustilidanoskarashmaqiladi),
		menu.Row(btnYolgizlikdadarajakotaish),
		menu.Row(btnDavolovchiqahramon),
		menu.Row(btnSaksonolti),
		menu.Row(btnSongiserafim),
		menu.Row(btnOchkozbersek),
		menu.Row(btnQotilAkame),
		menu.Row(btnZombi100, btnNomsizXotira),
		menu.Row(btnQoraKlever),
		menu.Row(btnJahannamjannati, btnBleach),
		menu.Row(btnOlmasqirolningkundalikhayoti, btnQoraoq),
		menu.Row(btnMenmuvaffaqiyatsiz, btnSakamotokunlari),
		menu.Row(btnTungiBoyqushlarKuyi, btnZulmatiblisi),
		menu.Row(btnQahramonBolishX, btnYozukuraOilasi),
		menu.Row(btnZulmatFarzandi, btnQudratYangiHikoya),
		menu.Row(btnYangiSaga),
		menu.Row(btnVanitasxotiralari),
		//menu.Row(btnUyatchangQahramonvaQotilMalikalar),
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
		menu.Row(btnDomekano),
		menu.Row(btnChegara),
		menu.Row(btnBucchgiri),
		menu.Row(btnAprel),
		menu.Row(btnAbadiylikqoriqchisi),
		menu.Row(btnAfsona),
		menu.Row(btnYettinchiumrni),
		menu.Row(btnYanabirnarsasorasamboladimi),
		menu.Row(btnOlimkundaligi),
		menu.Row(btnVoleybol),
		menu.Row(btnShilliqsifatida),
		menu.Row(btnHorimiya),
		menu.Row(btnTokyoGul),
		menu.Row(btnSarguzashtchilarRestorani),
		menu.Row(btnOvozShakli, btnQuyonQiz),
		menu.Row(btnBirzarbliodam, btnHukmdor),
		menu.Row(btnQilichsanationline, btnDMC),
		menu.Row(btnOchkozBerserk, btnJosusXOilasi),
		menu.Row(btnShamolniBoysundirish),
		menu.Row(btnDororo, btnArraOdam),
		menu.Row(btnPolatqaladagikabaneri),
		menu.Row(btnRainbow, btnQogirchoqlarsirki),
		menu.Row(btnYulduzFarzandi, btnElitaSinfi),
		menu.Row(btnTajribasizSenpai, btnYangiDarvoza),
		menu.Row(btnQudratliSongiDushman),
		menu.Row(btnCheksizlikgachaLv9999),
		menu.Row(btnRagnarokRekordi),
		menu.Row(btnSenuchunOlmas),
		menu.Row(btnMasamuneQasosi),
		menu.Row(btnKimMeniMalikaQildi),
		menu.Row(btnKokZindon, btnUysizMabud),
		menu.Row(btnParazitHayotSaboqlari),
		menu.Row(btnPariDumihaqidaafsona),
		menu.Row(btnKayjuRaqam, btnHyouka),
		menu.Row(btnSoqolimniolib),
		menu.Row(btnYettiOlimGunohlari),
		menu.Row(btnOmadsizningqaytatugilishi),
		menu.Row(btnQiroloyini, btnVioletEvergarden),
		menu.Row(btnOzgadunyoda),
		menu.Row(btnUqizyolgiz, btnBirsoatliqizcha),
		menu.Row(btnUqizyolgiz, btnJonliSana),
	)

	switch text {

	case "Animelar", "/menu":
		return c.Send("asosiy kanal -https://t.me/Anmelaruzb ", menu)
	default:
		return c.Send("Iltimos, /menu dan danlang.")
	}
}
