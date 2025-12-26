package anmelaruzb

import (
	tele "gopkg.in/telebot.v4"
	"namelaruzb_bot/kodi/anmelaruzb/Abadiylikqoriqchisi"
	"namelaruzb_bot/kodi/anmelaruzb/Afsonaviyilohyilnomasi"
	"namelaruzb_bot/kodi/anmelaruzb/Afsonaviyqahramonlar"
	"namelaruzb_bot/kodi/anmelaruzb/Aliyabazanmegarustilidanoskarashmaqiladi"
	"namelaruzb_bot/kodi/anmelaruzb/AprelYolgoni"
	"namelaruzb_bot/kodi/anmelaruzb/ArifuretaDunyodagiengkuchli"
	"namelaruzb_bot/kodi/anmelaruzb/ArraOdam"
	"namelaruzb_bot/kodi/anmelaruzb/Badargaqilinganqahramon"
	"namelaruzb_bot/kodi/anmelaruzb/Baholovchiningmashhur"
	"namelaruzb_bot/kodi/anmelaruzb/Barmoqlaruchidagisevgi"
	"namelaruzb_bot/kodi/anmelaruzb/BasketbolKuroko"
	"namelaruzb_bot/kodi/anmelaruzb/BekamiKostYashashimUshun"
	"namelaruzb_bot/kodi/anmelaruzb/Birsoatliqizcha"
	"namelaruzb_bot/kodi/anmelaruzb/Birzarbliodam"
	"namelaruzb_bot/kodi/anmelaruzb/Bleach"
	"namelaruzb_bot/kodi/anmelaruzb/Boshqadunyobirzumda"
	"namelaruzb_bot/kodi/anmelaruzb/Boshqadunyodanmuammolibolalar"
	"namelaruzb_bot/kodi/anmelaruzb/Bosningqizivauning"
	"namelaruzb_bot/kodi/anmelaruzb/Bucchgiri"
	"namelaruzb_bot/kodi/anmelaruzb/Chegaraortida"
	"namelaruzb_bot/kodi/anmelaruzb/Cheksizdendagron"
	"namelaruzb_bot/kodi/anmelaruzb/Cheksizlikgacha"
	"namelaruzb_bot/kodi/anmelaruzb/DMC"
	"namelaruzb_bot/kodi/anmelaruzb/DahoShifokorningSoyadagi"
	"namelaruzb_bot/kodi/anmelaruzb/Daholaruchunozga"
	"namelaruzb_bot/kodi/anmelaruzb/DastlabkiDrift"
	"namelaruzb_bot/kodi/anmelaruzb/Davolashsehridan"
	"namelaruzb_bot/kodi/anmelaruzb/Davolovchiqahramon"
	"namelaruzb_bot/kodi/anmelaruzb/Daydiitlarningbuyugi"
	"namelaruzb_bot/kodi/anmelaruzb/Detektivallaqchon"
	"namelaruzb_bot/kodi/anmelaruzb/DoktorElizahayotlarni"
	"namelaruzb_bot/kodi/anmelaruzb/Doktorstoun"
	"namelaruzb_bot/kodi/anmelaruzb/Domekano"
	"namelaruzb_bot/kodi/anmelaruzb/Dororo"
	"namelaruzb_bot/kodi/anmelaruzb/Dostimningsinglisibezovta"
	"namelaruzb_bot/kodi/anmelaruzb/DragonRajaS"
	"namelaruzb_bot/kodi/anmelaruzb/Egzartisboshqadunyodaqayta"
	"namelaruzb_bot/kodi/anmelaruzb/ElitaSinfi"
	"namelaruzb_bot/kodi/anmelaruzb/Engqudratlipartiya"
	"namelaruzb_bot/kodi/anmelaruzb/Engzaifyirtqich"
	"namelaruzb_bot/kodi/anmelaruzb/Fojiagasababchiblganmalika"
	"namelaruzb_bot/kodi/anmelaruzb/Franksdagisevgi"
	"namelaruzb_bot/kodi/anmelaruzb/FrirenSonggimanzilga"
	"namelaruzb_bot/kodi/anmelaruzb/GachiakutaQasos"
	"namelaruzb_bot/kodi/anmelaruzb/Gertsogningqizigashaxsiyoqituvchiboldim"
	"namelaruzb_bot/kodi/anmelaruzb/Gildiyaadminstratoribolib"
	"namelaruzb_bot/kodi/anmelaruzb/Grimgaldagikulgi"
	"namelaruzb_bot/kodi/anmelaruzb/Hikaruningsongiyozi"
	"namelaruzb_bot/kodi/anmelaruzb/Horimiya"
	"namelaruzb_bot/kodi/anmelaruzb/Hukmdor"
	"namelaruzb_bot/kodi/anmelaruzb/Hyouka"
	"namelaruzb_bot/kodi/anmelaruzb/Iblislarhukmdoriqoshimcha"
	"namelaruzb_bot/kodi/anmelaruzb/Iblislarmaktabigahush"
	"namelaruzb_bot/kodi/anmelaruzb/Ilohiyqilichmaktabining"
	"namelaruzb_bot/kodi/anmelaruzb/Jahannamjannati"
	"namelaruzb_bot/kodi/anmelaruzb/Jodugarlarjangi"
	"namelaruzb_bot/kodi/anmelaruzb/JonliSana"
	"namelaruzb_bot/kodi/anmelaruzb/JosusXOilasi"
	"namelaruzb_bot/kodi/anmelaruzb/KayjuRaqam"
	"namelaruzb_bot/kodi/anmelaruzb/Kelajakkundaligi"
	"namelaruzb_bot/kodi/anmelaruzb/KimMeniMalikaQildi"
	"namelaruzb_bot/kodi/anmelaruzb/KokZindon"
	"namelaruzb_bot/kodi/anmelaruzb/Kumushqirolning"
	"namelaruzb_bot/kodi/anmelaruzb/LideylDunyosi"
	"namelaruzb_bot/kodi/anmelaruzb/Lookism"
	"namelaruzb_bot/kodi/anmelaruzb/LordiArmiyasiningeng"
	"namelaruzb_bot/kodi/anmelaruzb/MabudlarHohishiBilan"
	"namelaruzb_bot/kodi/anmelaruzb/Maktabtomonidantan"
	"namelaruzb_bot/kodi/anmelaruzb/Manorgimchakman"
	"namelaruzb_bot/kodi/anmelaruzb/MasamuneQasosi"
	"namelaruzb_bot/kodi/anmelaruzb/Mendamillionhayot"
	"namelaruzb_bot/kodi/anmelaruzb/Menengkuchlisarguzashtchi"
	"namelaruzb_bot/kodi/anmelaruzb/Mengalaktikalararo"
	"namelaruzb_bot/kodi/anmelaruzb/MeniQilich"
	"namelaruzb_bot/kodi/anmelaruzb/Meniahramonlikakademiyam"
	"namelaruzb_bot/kodi/anmelaruzb/MeningQotillikMaqomim"
	"namelaruzb_bot/kodi/anmelaruzb/Meniqizchamnafaqatgozal"
	"namelaruzb_bot/kodi/anmelaruzb/Menmuvaffaqiyatsiz"
	"namelaruzb_bot/kodi/anmelaruzb/MinograApokalipsis"
	"namelaruzb_bot/kodi/anmelaruzb/MoviyQuticha"
	"namelaruzb_bot/kodi/anmelaruzb/Naruto"
	"namelaruzb_bot/kodi/anmelaruzb/Negahammameni"
	"namelaruzb_bot/kodi/anmelaruzb/Nikohuzuklarihaqida"
	"namelaruzb_bot/kodi/anmelaruzb/NinjavaYakudza"
	"namelaruzb_bot/kodi/anmelaruzb/NomsizXotira"
	"namelaruzb_bot/kodi/anmelaruzb/OchkozBerserk"
	"namelaruzb_bot/kodi/anmelaruzb/Ochkozbosek"
	"namelaruzb_bot/kodi/anmelaruzb/Oddiyinsondanqahramonlikkacha"
	"namelaruzb_bot/kodi/anmelaruzb/Olimkundaligi"
	"namelaruzb_bot/kodi/anmelaruzb/Olmasqirolningkundalikhayoti"
	"namelaruzb_bot/kodi/anmelaruzb/Oltibargqahramonlari"
	"namelaruzb_bot/kodi/anmelaruzb/OltinVaqt"
	"namelaruzb_bot/kodi/anmelaruzb/Omadsizning"
	"namelaruzb_bot/kodi/anmelaruzb/OnmyoQaytaTugilishi"
	"namelaruzb_bot/kodi/anmelaruzb/Ortayoshli"
	"namelaruzb_bot/kodi/anmelaruzb/OsmondagiJanglar"
	"namelaruzb_bot/kodi/anmelaruzb/OvozShakli"
	"namelaruzb_bot/kodi/anmelaruzb/OxirgiTelbaBosspaydo"
	"namelaruzb_bot/kodi/anmelaruzb/OyindagiEngBoyodam"
	"namelaruzb_bot/kodi/anmelaruzb/Oyingbilanbirgavideo"
	"namelaruzb_bot/kodi/anmelaruzb/Oyinsizhayotyoq"
	"namelaruzb_bot/kodi/anmelaruzb/Oysayohatiyangi"
	"namelaruzb_bot/kodi/anmelaruzb/Ozgadunyoda"
	"namelaruzb_bot/kodi/anmelaruzb/OzgadunyodaNoyob"
	"namelaruzb_bot/kodi/anmelaruzb/Ozgadunyodafermerlik"
	"namelaruzb_bot/kodi/anmelaruzb/Ozgadunyodayolgizhujum"
	"namelaruzb_bot/kodi/anmelaruzb/Ozjonigaqasd"
	"namelaruzb_bot/kodi/anmelaruzb/ParazitHayotSaboqlari"
	"namelaruzb_bot/kodi/anmelaruzb/PariDumihaqidaafsona"
	"namelaruzb_bot/kodi/anmelaruzb/Polatqaladagikabaneri"
	"namelaruzb_bot/kodi/anmelaruzb/QahramonBolishX"
	"namelaruzb_bot/kodi/anmelaruzb/Qahramonningqaytishi"
	"namelaruzb_bot/kodi/anmelaruzb/Qalqonqahromoni"
	"namelaruzb_bot/kodi/anmelaruzb/Qaytatugilganaristokratningmisli"
	"namelaruzb_bot/kodi/anmelaruzb/Qilichsanationline"
	"namelaruzb_bot/kodi/anmelaruzb/Qiroloyini"
	"namelaruzb_bot/kodi/anmelaruzb/QobilyatsizNana"
	"namelaruzb_bot/kodi/anmelaruzb/Qogirchoqlarsirki"
	"namelaruzb_bot/kodi/anmelaruzb/QoraKlever"
	"namelaruzb_bot/kodi/anmelaruzb/Qorachaqiruvchi"
	"namelaruzb_bot/kodi/anmelaruzb/Qoraoq"
	"namelaruzb_bot/kodi/anmelaruzb/Qoshnifarishta"
	"namelaruzb_bot/kodi/anmelaruzb/QotilAkame"
	"namelaruzb_bot/kodi/anmelaruzb/QudratYangiHikoya"
	"namelaruzb_bot/kodi/anmelaruzb/QudratliSongiDushman"
	"namelaruzb_bot/kodi/anmelaruzb/QuyonQiz"
	"namelaruzb_bot/kodi/anmelaruzb/Ragnarok"
	"namelaruzb_bot/kodi/anmelaruzb/Rainbow"
	"namelaruzb_bot/kodi/anmelaruzb/ReZero"
	"namelaruzb_bot/kodi/anmelaruzb/Realizmqahramoniqirollikni"
	"namelaruzb_bot/kodi/anmelaruzb/Sakamotokunlari"
	"namelaruzb_bot/kodi/anmelaruzb/Sakkizinchiogilqoysangizlarchi"
	"namelaruzb_bot/kodi/anmelaruzb/Saksonolti"
	"namelaruzb_bot/kodi/anmelaruzb/SarguzashtchilarRestorani"
	"namelaruzb_bot/kodi/anmelaruzb/SehrYaratuvchi"
	"namelaruzb_bot/kodi/anmelaruzb/Senioshqozon"
	"namelaruzb_bot/kodi/anmelaruzb/SenuchunOlmas"
	"namelaruzb_bot/kodi/anmelaruzb/Sevgidebatalganshart"
	"namelaruzb_bot/kodi/anmelaruzb/ShamolniBoysundirish"
	"namelaruzb_bot/kodi/anmelaruzb/Shangrilachegarasi"
	"namelaruzb_bot/kodi/anmelaruzb/Sharlota"
	"namelaruzb_bot/kodi/anmelaruzb/Shikastlanishniistamasdimshuning"
	"namelaruzb_bot/kodi/anmelaruzb/Shilliqsifatida"
	"namelaruzb_bot/kodi/anmelaruzb/Shirinrenkarnatsiya"
	"namelaruzb_bot/kodi/anmelaruzb/SirlarHukmdori"
	"namelaruzb_bot/kodi/anmelaruzb/SkeletRitsarozgadunyoda"
	"namelaruzb_bot/kodi/anmelaruzb/Songiserafim"
	"namelaruzb_bot/kodi/anmelaruzb/Soqolimniolib"
	"namelaruzb_bot/kodi/anmelaruzb/Soyadakotarilish"
	"namelaruzb_bot/kodi/anmelaruzb/SuvSehrgari"
	"namelaruzb_bot/kodi/anmelaruzb/Tahlilqilishqobiliyatigaega"
	"namelaruzb_bot/kodi/anmelaruzb/TajribasizSenpai"
	"namelaruzb_bot/kodi/anmelaruzb/TaqdirBuyukTartibMutlaqIblislarJabhasi"
	"namelaruzb_bot/kodi/anmelaruzb/TaqdirJangKechasi"
	"namelaruzb_bot/kodi/anmelaruzb/Taxtmuxri"
	"namelaruzb_bot/kodi/anmelaruzb/Tayoqvaqilich"
	"namelaruzb_bot/kodi/anmelaruzb/Titanlarjangi"
	"namelaruzb_bot/kodi/anmelaruzb/TokyoGul"
	"namelaruzb_bot/kodi/anmelaruzb/Tungiboyqushlarkuyi"
	"namelaruzb_bot/kodi/anmelaruzb/Uqizyolgiz"
	"namelaruzb_bot/kodi/anmelaruzb/UyatchangQahramonvaQotilMalikalar"
	"namelaruzb_bot/kodi/anmelaruzb/UysizMabud"
	"namelaruzb_bot/kodi/anmelaruzb/Uzoqpaladin"
	"namelaruzb_bot/kodi/anmelaruzb/UzuklarHukumdor"
	"namelaruzb_bot/kodi/anmelaruzb/Vanitasxotiralari"
	"namelaruzb_bot/kodi/anmelaruzb/Vayronbolganmojizalar"
	"namelaruzb_bot/kodi/anmelaruzb/VioletEvergarden"
	"namelaruzb_bot/kodi/anmelaruzb/Voleybol"
	"namelaruzb_bot/kodi/anmelaruzb/Xarobalarqiroligi"
	"namelaruzb_bot/kodi/anmelaruzb/XunukAmmoKuchli"
	"namelaruzb_bot/kodi/anmelaruzb/XushboyGulViqorBilanGulaydi"
	"namelaruzb_bot/kodi/anmelaruzb/Yanabirnarsa"
	"namelaruzb_bot/kodi/anmelaruzb/YangiDarvoza"
	"namelaruzb_bot/kodi/anmelaruzb/YangiSaga"
	"namelaruzb_bot/kodi/anmelaruzb/YaponiyagaXush"
	"namelaruzb_bot/kodi/anmelaruzb/YettiOlimGunohlari"
	"namelaruzb_bot/kodi/anmelaruzb/Yettinchishahzodasifatida"
	"namelaruzb_bot/kodi/anmelaruzb/Yettinchiumrnibetashvishyashayotgan"
	"namelaruzb_bot/kodi/anmelaruzb/YovuzlikdarajasiMenyashirin"
	"namelaruzb_bot/kodi/anmelaruzb/YozgiUrushimaTuneliOldidagiHayrlashuv"
	"namelaruzb_bot/kodi/anmelaruzb/YozukuraOilasi"
	"namelaruzb_bot/kodi/anmelaruzb/Yugureningabadiyligi"
	"namelaruzb_bot/kodi/anmelaruzb/YulduzFarzandi"
	"namelaruzb_bot/kodi/anmelaruzb/Zombi100"
	"namelaruzb_bot/kodi/anmelaruzb/ZulmatFarzandi"
	"namelaruzb_bot/kodi/anmelaruzb/ZulmatIblisi"
	"namelaruzb_bot/kodi/anmelaruzb/bizbirga"
	"namelaruzb_bot/kodi/anmelaruzb/dahoshahzodanimamlakat"
	_default "namelaruzb_bot/kodi/anmelaruzb/default"
	"namelaruzb_bot/kodi/anmelaruzb/iblislarqotili"
	"namelaruzb_bot/kodi/anmelaruzb/qipqizilragna"
	"namelaruzb_bot/kodi/anmelaruzb/tahlilashqobilyati"
	"namelaruzb_bot/kodi/anmelaruzb/tanyangurushyilnomalari"
	"namelaruzb_bot/kodi/anmelaruzb/tokiyoqasoskorlari"
	"namelaruzb_bot/kodi/anmelaruzb/toplabhayotimni"
	"namelaruzb_bot/kodi/anmelaruzb/vanpis"
	"namelaruzb_bot/kodi/anmelaruzb/yolgizlikdadarajakotaish"
)

func Home(c tele.Context) error {
	text := c.Text()
	switch text {
	case "🖋️ anme izlash":
		return c.Send("ID-kiriting:")
	case "1", "Vanpis":
		return vanpis.Home(c)
	case "2", "Naruto":
		return Naruto.Home(c)
	case "3", "Iblislar qotili":
		return iblislarqotili.Home(c)
	case "4", "Franksdagi sevgi":
		return Franksdagisevgi.Home(c)
	case "5", "Daho shahzodani mamlakatni qutqargani haqida":
		return dahoshahzodanimamlakat.Home(c)
	case "6", "Xarobalar qiroligi":
		return Xarobalarqiroligi.Home(c)
	case "7", "Soyada kotarilish":
		return Soyadakotarilish.Home(c)
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
	case "18", "Basketbol Kuroko":
		return BasketbolKuroko.Home(c)
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
	case "83", "DMC":
		return DMC.Home(c)
	case "84", "Ochko'z Berserk":
		return OchkozBerserk.Home(c)
	case "85", "Shamolni Bo'ysundirish":
		return ShamolniBoysundirish.Home(c)
	case "86":
		return Saksonolti.Home(c)
	case "87", "Dororo":
		return Dororo.Home(c)
	case "88", "Arra Odam":
		return ArraOdam.Home(c)
	case "89", "Poʻlat qal'adagi kabaneri":
		return Polatqaladagikabaneri.Home(c)
	case "90", "Rainbow":
		return Rainbow.Home(c)
	case "91", "Qo'g'irchoqlar sirki":
		return Qogirchoqlarsirki.Home(c)
	case "92", "Yulduz Farzandi":
		return YulduzFarzandi.Home(c)
	case "93", "Elita Sinfi":
		return ElitaSinfi.Home(c)
	case "94", "Tajribasiz Senpai":
		return TajribasizSenpai.Home(c)
	case "95", "Qudratli Soʻngi Dushman":
		return QudratliSongiDushman.Home(c)
	case "96", "Yangi Darvoza":
		return YangiDarvoza.Home(c)
	case "97", "Cheksizlikgacha Lv9999":
		return Cheksizlikgacha.Home(c)
	case "98", "Josus X Oilasi":
		return JosusXOilasi.Home(c)
	case "99", "Ragnarok Rekordi":
		return Ragnarok.Home(c)
	case "100", "Sen uchun O'lmas":
		return SenuchunOlmas.Home(c)
	case "101", "Masamune Qasosi":
		return MasamuneQasosi.Home(c)
	case "102", "Kim Meni Malika Qildi":
		return KimMeniMalikaQildi.Home(c)
	case "103", "Koʻk Zindon":
		return KokZindon.Home(c)
	case "104", "Parazit - Hayot Saboqlari":
		return ParazitHayotSaboqlari.Home(c)
	case "105", "Uysiz Ma'bud":
		return UysizMabud.Home(c)
	case "106", "Pari Dumi haqida afsona":
		return PariDumihaqidaafsona.Home(c)
	case "107", "Kayju 8-Raqam":
		return KayjuRaqam.Home(c)
	case "108", "Soqolimni olib, yuqori maktab qizini uyimga olib keldim":
		return Soqolimniolib.Home(c)
	case "109", "Yetti O'lim Gunohlari":
		return YettiOlimGunohlari.Home(c)
	case "110", "Omadsizning qayta tug'ilishi":
		return Omadsizning.Home(c)
	case "111", "Qirol oʻyini":
		return Qiroloyini.Home(c)
	case "112", "Violet Evergarden":
		return VioletEvergarden.Home(c)
	case "113", "Oʻzga dunyoda darajamni koʻtarib, bu dunyoda ham tengsiz boʻldim":
		return Ozgadunyoda.Home(c)
	case "114", "Hyouka":
		return Hyouka.Home(c)
	case "115", "U qiz yolgiz":
		return Uqizyolgiz.Home(c)
	case "116", "Bir soatli qizcha":
		return Birsoatliqizcha.Home(c)
	case "117", "Jonli Sana":
		return JonliSana.Home(c)
	case "118", "qip-qizil ragna":
		return qipqizilragna.Home(c)
	case "119", "Biz birga bo'lsak, sevgimiz har qanday to'siqni ortda qoldiradi":
		return bizbirga.Home(c)
	case "120", "Moviy Quticha":
		return MoviyQuticha.Home(c)
	case "121", "Hikaruning songi yozi":
		return Hikaruningsongiyozi.Home(c)
	case "122", "Taqdir: Buyuk Tartib Mutlaq Iblislar Jabhasi":
		return TaqdirBuyukTartibMutlaqIblislarJabhasi.Home(c)
	case "123", "Sevgi deb atalgan shart":
		return Sevgidebatalganshart.Home(c)
	case "124", "Afsonaviy qahramonlar va Ruhlar malikasi qizi sifatida qayta tug‘ildim":
		return Afsonaviyqahramonlar.Home(c)
	case "125", "Lideyl Dunyosi":
		return LideylDunyosi.Home(c)
	case "126", "Skelet Ritsar o‘zga dunyoda":
		return SkeletRitsarozgadunyoda.Home(c)
	case "127", "Qora chaqiruvchi":
		return Qorachaqiruvchi.Home(c)
	case "128", "Davolash sehridan foydalanishni noto'g'ri usuli":
		return Davolashsehridan.Home(c)
	case "129", "Lordi Armiyasining eng kuchli Sehrgari...":
		return LordiArmiyasiningeng.Home(c)
	case "130", "Tahlil qilish qobiliyatiga ega aristokrat bo'lib qayta tug'ilish":
		return Tahlilqilishqobiliyatigaega.Home(c)
	case "131", "Ilohiy qilich maktabining Iblis qilich egasi":
		return Ilohiyqilichmaktabining.Home(c)
	case "132", "Meni Qilich bo'lib qayta tug'ilishim haqida":
		return MeniQilich.Home(c)
	case "133", "O'z joniga qasd qiluvchilar o'zga dunyoda":
		return Ozjonigaqasd.Home(c)
	case "134", "Eng qudratli partiya tomonidan o'limgacha tarbiyalangan Ossan ...":
		return Engqudratlipartiya.Home(c)
	case "135", "Baholovchining mashhur bo'lmagan ishi aslida eng Kuchlisi hisoblanadi":
		return Baholovchiningmashhur.Home(c)
	case "136", "Sehr Yaratuvchi Boshqa dunyoda qanday qilib sehr yaratish mumkin":
		return SehrYaratuvchi.Home(c)
	case "137", "O'rta yoshli erkakning zodagon qiziga aylanishi":
		return Ortayoshli.Home(c)
	case "138", "Yaponiyaga Xush Kelibsiz, Elf Xonim!":
		return YaponiyagaXush.Home(c)
	case "139", "Grimgaldagi kulgi va illuziya":
		return Grimgaldagikulgi.Home(c)
	case "140", "Nega hamma meni dunyomni unutdi?":
		return Negahammameni.Home(c)
	case "141", "80.000 oltin tanga to'plab hayotimni qayta qurdim":
		return toplabhayotimni.Home(c)
	case "142", "Cheksiz dendagron":
		return Cheksizdendagron.Home(c)
	case "143", "Oy sayohati yangi dunyoga olib keladi":
		return Oysayohatiyangi.Home(c)
	case "144", "Man o'rgimchakman ! Ho'sh shunga nma qibti?":
		return Manorgimchakman.Home(c)
	case "145", "O'zga dunyoda fermerlik hayotim":
		return Ozgadunyodafermerlik.Home(c)
	case "146", "Iblislar maktabiga hush kelibsiz":
		return Iblislarmaktabigahu.Home(c)
	case "147", "Daholar uchun oʻzga dunyoda yashash ham muammo emas":
		return Daholaruchunozga.Home(c)
	case "148", "Shirin renkarnatsiya":
		return Shirinrenkarnatsi.Home(c)
	case "149", "Menda million hayot bor":
		return Mendamillionhayot.Home(c)
	case "150", "Maktab tomonidan tan olinmagan iblislar hukmdori":
		return Maktabtomonidantan.Home(c)
	case "151", "Boshqa dunyo bir zumda o'lim kuchiga dosh bera olmaydi":
		return Boshqadunyobirzumda.Home(c)
	case "152", "Kumush qirolning qayta tugʻilishi":
		return Kumushqirolning.Home(c)
	case "153", "Oʻzga dunyoda Noyob mahoratim":
		return OzgadunyodaNoyob.Home(c)
	case "154", "Egzartis boshqa dunyoda qayta tug'ilib eng kuchli bo'lishga intiladi":
		return Egzartisboshqadunyodaqayta.Home(c)
	case "155", "Nikoh uzuklari haqida afsona":
		return Nikohuzuklarihaqida.Home(c)
	case "156", "tanyang urush yilnomalari":
		return tanyangurushyilnomalari.Home(c)
	case "157", "Yovuzlik darajasi 99: Men yashirin xo'jayin bo'lishim mumkin, lekin men jin xo'jayini emasman":
		return YovuzlikdarajasiMenyashirin.Home(c)
	case "158", "Doktor Eliza: hayotlarni saqlab qoluvchi malika":
		return DoktorElizahayotlarni.Home(c)
	case "159", "Eng zaif yirtqich hayvon":
		return Engzaifyirtqich.Home(c)
	case "160", "Fojiaga sababchi bõlgan malika xalq uchun qõlidan kelgan barcha ishni qiladi":
		return Fojiagasababchiblganmalik.Home(c)
	case "161", "Iblislar hukmdori qoʻshimcha ishda!":
		return Iblislarhukmdoriqoshimcha.Home(c)
	case "162", "Oying bilan birga video o'yin":
		return Oyingbilanbirgavideo.Home(c)
	case "163", "Bekami Ko'st Yashashim Ushun Hamma Narsa Qildim":
		return BekamiKostYashashimUshu.Home(c)
	case "164", "Uzoq paladin":
		return Uzoqpaladin.Home(c)
	case "165", "Realizm qahramoni qirollikni qayta qurishi":
		return Realizmqahramoniqirollikni.Home(c)
	case "166", "Olti barg qahramonlari":
		return Oltibargqahramonlari.Home(c)
	case "167", "Daho Shifokorning Soyadagi Yangi Hayoti":
		return DahoShifokorningSoyadag.Home(c)
	case "168", "Qayta tugʻilgan aristokratning misli koʻrilmagan sarguzashtlari":
		return Qaytatugilganaristokratningmisli.Home(c)
	case "169", "Sakkizinchi o'g'il qo'ysangizlarchi?":
		return Sakkizinchiogilqoysangizlarchi.Home(c)
	case "170", "O'yindagi Eng Boy odam":
		return OyindagiEngBoyodam.Home(c)
	case "171", "O'yinsiz hayot yo'q":
		return Oyinsizhayotyoq.Home(c)
	case "172", "Arifureta: Dunyodagi eng kuchli hunarmand":
		return ArifuretaDunyodagiengkuchli.Home(c)
	case "173", "Gildiya adminstratori bo'lib ortiqcha ishlashni xohlamaganim uchun ishdan ketmoqchiman":
		return Gildiyaadminstratoribolib.Home(c)
	case "174", "Oddiy insondan qahramonlikkacha":
		return Oddiyinsondanqahramonlikkach.Home(c)
	case "175", "Tayoq va qilich":
		return Tayoqvaqilich.Home(c)
	case "176", "Re:Zero":
		return ReZero.Home(c)
	case "177", "Onmyo Qayta Tug‘ilishi: Hayolot Olami":
		return OnmyoQaytaTugilishi.Home(c)
	case "178", "Yettinchi shahzoda sifatida qayta tug'ildim va endi sehrimni istaganimcha kuchaytiraman!":
		return Yettinchishahzodasifatid.Home(c)
	case "179", "Mening Qotillik Maqomim Qahramonlik Maqomidan Yaxshiroq":
		return MeningQotillikMaqomim.Home(c)
	case "180", "Yugurening abadiyligi":
		return Yugureningabadiyligi.Home(c)
	case "181", "Daydi itlarning buyugi":
		return Daydiitlarningbuyug.Home(c)
	case "182", "Meni qizcham nafaqat go'zal":
		return Meniqizchamnafaqatgozal.Home(c)
	case "183", "Meni qahramonlik akademiyam":
		return Meniahramonlikakademiyam.Home(c)
	case "184", "Barmoqlar uchidagi sevgi":
		return Barmoqlaruchidagisevgi.Home(c)
	case "185", "Shangri-la chegarasi":
		return Shangrilachegarasi.Home(c)
	case "186", "Kelajak kundaligi":
		return Kelajakkundaligi.Home(c)
	case "187", "Men eng kuchli sarguzashtchi bo'lish uchun har doim mashq qildim":
		return Menengkuchlisarguzashtchi.Home(c)
	case "188", "Do'stimning singlisi bezovta qilyapti":
		return Dostimningsinglisibezovta.Home(c)
	case "189", "Shikastlanishni istamasdim shuning uchun himoyamni kuchaytirdim":
		return Shikastlanishniistamasdimshuning.Home(c)
	case "190", "Oxirgi Telba Boss paydo bo'ldi":
		return OxirgiTelbaBosspayd.Home(c)
	case "191", "Vayron bo'lgan mo'jizalar mamlakati":
		return Vayronbolganmojizalar.Home(c)
	case "192", "Qahramonning qaytishi":
		return Qahramonningqaytishi.Home(c)
	case "193", "Ninja va Yakudza":
		return NinjavaYakudza.Home(c)
	case "194", "Bosning qizi va uning Enagasi":
		return Bosningqizivauning.Home(c)
	case "195", "Detektiv allaqchon o'lgan":
		return Detektivallaqchon.Home(c)
	case "196", "Friren - Soʻnggi manzilga kuzatuvchi":
		return FrirenSonggimanzilga.Home(c)

	//
	//
	//
	default:
		return _default.Home(c)
	}
}
