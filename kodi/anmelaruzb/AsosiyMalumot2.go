package anmelaruzb

import (
	tele "gopkg.in/telebot.v4"
	"namelaruzb_bot/kodi/anmelaruzb/Abadiylikqoriqchisi"
	"namelaruzb_bot/kodi/anmelaruzb/Afsonaviyilohyilnomasi"
	"namelaruzb_bot/kodi/anmelaruzb/Aliyabazanmegarustilidanoskarashmaqiladi"
	"namelaruzb_bot/kodi/anmelaruzb/AprelYolgoni"
	"namelaruzb_bot/kodi/anmelaruzb/Badargaqilinganqahramon"
	"namelaruzb_bot/kodi/anmelaruzb/Birzarbliodam"
	"namelaruzb_bot/kodi/anmelaruzb/Bleach"
	"namelaruzb_bot/kodi/anmelaruzb/Boshqadunyodanmuammolibolalar"
	"namelaruzb_bot/kodi/anmelaruzb/Bucchgiri"
	"namelaruzb_bot/kodi/anmelaruzb/Chegaraortida"
	"namelaruzb_bot/kodi/anmelaruzb/DastlabkiDrift"
	"namelaruzb_bot/kodi/anmelaruzb/Davolovchiqahramon"
	"namelaruzb_bot/kodi/anmelaruzb/Doktorstoun"
	"namelaruzb_bot/kodi/anmelaruzb/Domekano"
	"namelaruzb_bot/kodi/anmelaruzb/DragonRajaS"
	"namelaruzb_bot/kodi/anmelaruzb/Franksdagisevgi"
	"namelaruzb_bot/kodi/anmelaruzb/GachiakutaQasos"
	"namelaruzb_bot/kodi/anmelaruzb/Gertsogningqizigashaxsiyoqituvchiboldim"
	"namelaruzb_bot/kodi/anmelaruzb/Horimiya"
	"namelaruzb_bot/kodi/anmelaruzb/Hukmdor"
	"namelaruzb_bot/kodi/anmelaruzb/Jahannamjannati"
	"namelaruzb_bot/kodi/anmelaruzb/Jodugarlarjangi"
	"namelaruzb_bot/kodi/anmelaruzb/Lookism"
	"namelaruzb_bot/kodi/anmelaruzb/MabudlarHohishiBilan"
	"namelaruzb_bot/kodi/anmelaruzb/Mengalaktikalararo"
	"namelaruzb_bot/kodi/anmelaruzb/Menmuvaffaqiyatsiz"
	"namelaruzb_bot/kodi/anmelaruzb/MinograApokalipsis"
	"namelaruzb_bot/kodi/anmelaruzb/Naruto"
	"namelaruzb_bot/kodi/anmelaruzb/NomsizXotira"
	"namelaruzb_bot/kodi/anmelaruzb/Ochkozbosek"
	"namelaruzb_bot/kodi/anmelaruzb/Olimkundaligi"
	"namelaruzb_bot/kodi/anmelaruzb/Olmasqirolningkundalikhayoti"
	"namelaruzb_bot/kodi/anmelaruzb/OltinVaqt"
	"namelaruzb_bot/kodi/anmelaruzb/OsmondagiJanglar"
	"namelaruzb_bot/kodi/anmelaruzb/OvozShakli"
	"namelaruzb_bot/kodi/anmelaruzb/Ozgadunyodayolgizhujum"
	"namelaruzb_bot/kodi/anmelaruzb/QahramonBolishX"
	"namelaruzb_bot/kodi/anmelaruzb/Qalqonqahromoni"
	"namelaruzb_bot/kodi/anmelaruzb/Qilichsanationline"
	"namelaruzb_bot/kodi/anmelaruzb/QobilyatsizNana"
	"namelaruzb_bot/kodi/anmelaruzb/QoraKlever"
	"namelaruzb_bot/kodi/anmelaruzb/Qoraoq"
	"namelaruzb_bot/kodi/anmelaruzb/Qoshnifarishta"
	"namelaruzb_bot/kodi/anmelaruzb/QotilAkame"
	"namelaruzb_bot/kodi/anmelaruzb/QudratYangiHikoya"
	"namelaruzb_bot/kodi/anmelaruzb/QuyonQiz"
	"namelaruzb_bot/kodi/anmelaruzb/Sakamotokunlari"
	"namelaruzb_bot/kodi/anmelaruzb/Saksonolti"
	"namelaruzb_bot/kodi/anmelaruzb/SarguzashtchilarRestorani"
	"namelaruzb_bot/kodi/anmelaruzb/Senioshqozon"
	"namelaruzb_bot/kodi/anmelaruzb/Sharlota"
	"namelaruzb_bot/kodi/anmelaruzb/Shilliqsifatida"
	"namelaruzb_bot/kodi/anmelaruzb/SirlarHukmdori"
	"namelaruzb_bot/kodi/anmelaruzb/Songiserafim"
	"namelaruzb_bot/kodi/anmelaruzb/Soyadakotarilish"
	"namelaruzb_bot/kodi/anmelaruzb/SuvSehrgari"
	"namelaruzb_bot/kodi/anmelaruzb/TaqdirJangKechasi"
	"namelaruzb_bot/kodi/anmelaruzb/Taxtmuxri"
	"namelaruzb_bot/kodi/anmelaruzb/Titanlarjangi"
	"namelaruzb_bot/kodi/anmelaruzb/TokyoGul"
	"namelaruzb_bot/kodi/anmelaruzb/Tungiboyqushlarkuyi"
	"namelaruzb_bot/kodi/anmelaruzb/UyatchangQahramonvaQotilMalikalar"
	"namelaruzb_bot/kodi/anmelaruzb/UzuklarHukumdor"
	"namelaruzb_bot/kodi/anmelaruzb/Vanitasxotiralari"
	"namelaruzb_bot/kodi/anmelaruzb/Voleybol"
	"namelaruzb_bot/kodi/anmelaruzb/Xarobalarqiroligi"
	"namelaruzb_bot/kodi/anmelaruzb/XunukAmmoKuchli"
	"namelaruzb_bot/kodi/anmelaruzb/XushboyGulViqorBilanGulaydi"
	"namelaruzb_bot/kodi/anmelaruzb/Yanabirnarsa"
	"namelaruzb_bot/kodi/anmelaruzb/YangiSaga"
	"namelaruzb_bot/kodi/anmelaruzb/Yettinchiumrnibetashvishyashayotgan"
	"namelaruzb_bot/kodi/anmelaruzb/YozgiUrushimaTuneliOldidagiHayrlashuv"
	"namelaruzb_bot/kodi/anmelaruzb/YozukuraOilasi"
	"namelaruzb_bot/kodi/anmelaruzb/Zombi100"
	"namelaruzb_bot/kodi/anmelaruzb/ZulmatFarzandi"
	"namelaruzb_bot/kodi/anmelaruzb/ZulmatIblisi"
	"namelaruzb_bot/kodi/anmelaruzb/dahoshahzodanimamlakat"
	"namelaruzb_bot/kodi/anmelaruzb/iblislarqotili"
	"namelaruzb_bot/kodi/anmelaruzb/tahlilashqobilyati"
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
	case "3", "Franksdagi sevgi":
		return Franksdagisevgi.Home(c)
	case "4", "Daho shahzodani mamlakatni qutqargani haqida":
		return dahoshahzodanimamlakat.Home(c)
	case "5", "Xarobalar qiroligi":
		return Xarobalarqiroligi.Home(c)
	case "6", "Soyada kotarilish":
		return Soyadakotarilish.Home(c)
	case "7", "Vanpis":
		return vanpis.Home(c)
	case "8", "Tokiyo qasoskorlari":
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
	case "30", "Sakamoto Kunlari":
		return Sakamotokunlari.Home(c)
	case "31", "Tungi Boyqushlar Kuyi":
		return Tungiboyqushlarkuyi.Home(c)
	case "32", "Zulmat Iblisi":
		return ZulmatIblisi.Home(c)
	case "33", "Qahramon Boʻlish X":
		return QahramonBolishX.Home(c)
	case "34", "Yozukura Oilasi":
		return YozukuraOilasi.Home(c)
	case "35", "Zulmat Farzandi":
		return ZulmatFarzandi.Home(c)
	case "36", "Qudrat! Yangi Hikoya":
		return QudratYangiHikoya.Home(c)
	case "37", "Yangi Saga":
		return YangiSaga.Home(c)
	case "38", "Uyatchang Qahramon va Qotil Malikalar":
		return UyatchangQahramonvaQotilMalikalar.Home(c)
	case "39", "Doktor stoun":
		return Doktorstoun.Home(c)
	case "40", "Taxt muxri":
		return Taxtmuxri.Home(c)
	case "41", "Badargʻa qilingan qahramon":
		return Badargaqilinganqahramon.Home(c)
	case "42", "Boshqa dunyodan muammoli bolalar":
		return Boshqadunyodanmuammolibolalar.Home(c)
	case "43", "Taxt Dastlabki Drift":
		return DastlabkiDrift.Home(c)
	case "44", "Dragon Raja":
		return DragonRajaS.Home(c)
	case "45", "Xushboʻy Gul Viqor Bilan Gulaydi":
		return XushboyGulViqorBilanGulaydi.Home(c)
	case "46", "Xunuk Ammo Kuchli: Busamen G‘alaba Jangi":
		return XunukAmmoKuchli.Home(c)
	case "47", "Uzuklar Hukumdori: Rohhirm Urushi":
		return UzuklarHukumdor.Home(c)
	case "48", "Taqdir Jang Kechasi":
		return TaqdirJangKechasi.Home(c)
	case "49", "Tahlil Qilish Qobiliyati Bilan Aristokrat Sifatida Qayta Tugʻildim":
		return tahlilashqobilyati.Home(c)
	case "50", "Suv Sehrgari":
		return SuvSehrgari.Home(c)
	case "51", "Sirlar Hukmdori":
		return SirlarHukmdori.Home(c)
	case "52", "Seni oshqozon osti bezingni yemoqchi man":
		return Senioshqozon.Home(c)
	case "53", "Qobilyatsiz Nana":
		return QobilyatsizNana.Home(c)
	case "54", "Ozga dunyoda yolgiz hujum":
		return Ozgadunyodayolgizhujum.Home(c)
	case "55", "Osmondagi Janglar":
		return OsmondagiJanglar.Home(c)
	case "56", "Oltin Vaqt":
		return OltinVaqt.Home(c)
	case "57", "Minogra Apokalipsis":
		return MinograApokalipsis.Home(c)
	case "58", "Men galaktikalar aro imperiyaning yovuz lordiman":
		return Mengalaktikalararo.Home(c)
	case "59", "Mabudlar Hohishi Bilan":
		return MabudlarHohishiBilan.Home(c)
	case "60", "Lookism":
		return Lookism.Home(c)
	case "61", "Gertsogning qiziga shaxsiy o'qituvchi boʻldim":
		return Gertsogningqizigashaxsiyoqituvchiboldim.Home(c)
	case "62", "Gachiakuta Qasos":
		return GachiakutaQasos.Home(c)
	case "63", "Domekano":
		return Domekano.Home(c)
	case "64", "Chegara ortida":
		return Chegaraortida.Home(c)
	case "65", "Bucchgiri":
		return Bucchgiri.Home(c)
	case "66", "AprelYolgoni":
		return AprelYolgoni.Home(c)
	case "67", "Afsonaviy ilohy yilnomasi":
		return Afsonaviyilohyilnomasi.Home(c)
	case "68", "Yettinchi umrni betashvish yashayotgan yovuz ayol":
		return Yettinchiumrnibetashvishyashayotgan.Home(c)
	case "69", "Abadiylik qoriqchisi":
		return Abadiylikqoriqchisi.Home(c)
	case "70", "Yana bir narsa soʻrasam boʻladimi":
		return Yanabirnarsa.Home(c)
	case "71", "O'lim kundaligi":
		return Olimkundaligi.Home(c)
	case "72", "Voleybol":
		return Voleybol.Home(c)
	case "73", "Shilliq sifatida qayta tug'ilganim haqida":
		return Shilliqsifatida.Home(c)
	case "74", "Tokyo Gul":
		return TokyoGul.Home(c)
	case "75", "Horimiya":
		return Horimiya.Home(c)
	case "76", "Sarguzashtchilar Restorani":
		return SarguzashtchilarRestorani.Home(c)
	case "77", "Ovoz Shakli":
		return OvozShakli.Home(c)
	case "78", "Quyon Qiz":
		return QuyonQiz.Home(c)
	case "79", "Bir zarbli odam":
		return Birzarbliodam.Home(c)
	case "80", "Hukmdor":
		return Hukmdor.Home(c)
	case "81", "Yozgi Urushima Tuneli Oldidagi Hayrlashuv":
		return YozgiUrushimaTuneliOldidagiHayrlashuv.Home(c)
	case "82", "Qilich sanati online":
		return Qilichsanationline.Home(c)

	//
	//
	//
	default:
		return c.Send("Noma'lum buyruq. Iltimos, /menu dan foydalaning.")
	}
}
