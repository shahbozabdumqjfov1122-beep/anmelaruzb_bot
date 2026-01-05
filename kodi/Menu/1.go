package Menu

import tele "gopkg.in/telebot.v4"

func Home(c tele.Context) error {
	text := c.Text()

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}
	btn1 := menu.Text("Vanpis")
	btn2 := menu.Text("Naruto")
	btn3 := menu.Text("Iblislar qotili")
	btn4 := menu.Text("Franksdagi sevgi")
	btn5 := menu.Text("Daho shahzodani mamlakatni qutqargani haqida")
	btn6 := menu.Text("Xarobalar qiroligi")
	btn7 := menu.Text("Soyada kotarilish")
	btn8 := menu.Text("Tokiyo qasoskorlari")
	btn9 := menu.Text("Titanlar hujumi")
	btn10 := menu.Text("Jodugarlar jangi")
	btn11 := menu.Text("Qalqon qahromoni")
	btn12 := menu.Text("Sharlota")
	btn13 := menu.Text("Qoshni farishta")
	btn14 := menu.Text("yolg'izlikda daraja ko'taish")
	btn15 := menu.Text("Aliya bazan mega rustilida nos karashma qiladi")
	btn16 := menu.Text("Davolovchi qahramon")
	btn17 := menu.Text("So'ngi serafim")
	btn18 := menu.Text("Basketbol Kuroko")
	btn19 := menu.Text("Ochkoz bersek")
	btn20 := menu.Text("Zombi 100")
	btn21 := menu.Text("Nomsiz Xotira")
	btn22 := menu.Text("Qotil Akame")
	btn23 := menu.Text("Jahannam jannati")
	btn24 := menu.Text("Qora Klever")
	btn25 := menu.Text("Vanitas xotiralari")
	btn26 := menu.Text("Bleach")
	btn27 := menu.Text("O‘lmas qirolning kundalik hayoti")
	btn28 := menu.Text("Qora o'q")
	btn29 := menu.Text("Men Muvaffaqiyatsiz...")
	btn30 := menu.Text("Sakamoto Kunlari")
	btn31 := menu.Text("Tungi Boyqushlar Kuyi")
	btn32 := menu.Text("Zulmat Iblisi")
	btn33 := menu.Text("Qahramon Boʻlish X")
	btn34 := menu.Text("Yozukura Oilasi")
	btn35 := menu.Text("Zulmat Farzandi")
	btn36 := menu.Text("Qudrat! Yangi Hikoya")
	btn37 := menu.Text("Yangi Saga")
	btn38 := menu.Text("Uyatchang Qahramon va Qotil Malikalar")
	btn39 := menu.Text("Doktor stoun")
	btn40 := menu.Text("Taxt muxri")
	btn41 := menu.Text("Badargʻa qilingan qahramon")
	btn42 := menu.Text("Boshqa dunyodan muammoli bolalar")
	btn43 := menu.Text("Taxt Dastlabki Drift")
	btn44 := menu.Text("Dragon Raja")
	btn45 := menu.Text("Xushboʻy Gul Viqor Bilan Gulaydi")
	btn46 := menu.Text("Xunuk Ammo Kuchli: Busamen G‘alaba Jangi")
	btn47 := menu.Text("Uzuklar Hukumdori: Rohhirm Urushi")
	btn48 := menu.Text("Taqdir Jang Kechasi")
	btn49 := menu.Text("Tahlil Qilish Qobiliyati Bilan Aristokrat Sifatida Qayta Tugʻildim")
	btn50 := menu.Text("Suv Sehrgari")
	btn51 := menu.Text("Sirlar Hukmdori")
	btn52 := menu.Text("Seni oshqozon osti bezingni yemoqchi man")
	btn53 := menu.Text("Qobilyatsiz Nana")
	btn54 := menu.Text("Ozga dunyoda yolgiz hujum")
	btn55 := menu.Text("Osmondagi Janglar")
	btn56 := menu.Text("Oltin Vaqt")
	btn57 := menu.Text("Minogra Apokalipsis")
	btn58 := menu.Text("Men galaktikalar aro imperiyaning yovuz lordiman")
	btn59 := menu.Text("Mabudlar Hohishi Bilan")
	btn60 := menu.Text("Lookism")
	btn61 := menu.Text("Gertsogning qiziga shaxsiy o'qituvchi boʻldim")
	btn62 := menu.Text("Gachiakuta Qasos")
	btn63 := menu.Text("Domekano")
	btn64 := menu.Text("Chegara ortida")
	btn65 := menu.Text("Bucchgiri")
	btn66 := menu.Text("AprelYolgoni")
	btn67 := menu.Text("Afsonaviy ilohy yilnomasi")
	btn68 := menu.Text("Yettinchi umrni betashvish yashayotgan yovuz ayol")
	btn69 := menu.Text("Abadiylik qoriqchisi")
	btn70 := menu.Text("Yana bir narsa soʻrasam boʻladimi")
	btn71 := menu.Text("O'lim kundaligi")
	btn72 := menu.Text("Voleybol")
	btn73 := menu.Text("Shilliq sifatida qayta tug'ilganim haqida")
	btn74 := menu.Text("Tokyo Gul")
	btn75 := menu.Text("Horimiya")
	btn76 := menu.Text("Sarguzashtchilar Restorani")
	btn77 := menu.Text("Ovoz Shakli")
	btn78 := menu.Text("Quyon Qiz")
	btn79 := menu.Text("Bir zarbli odam")
	btn80 := menu.Text("Hukmdor")
	btn81 := menu.Text("Yozgi Urushima Tuneli Oldidagi Hayrlashuv")
	btn82 := menu.Text("Qilich sanati online")
	btn83 := menu.Text("DMC")
	btn84 := menu.Text("Ochko'z Berserk")
	btn85 := menu.Text("Shamolni Bo'ysundirish")
	btn86 := menu.Text("86")
	btn87 := menu.Text("Dororo")
	btn88 := menu.Text("Arra Odam")
	btn89 := menu.Text("Poʻlat qal'adagi kabaneri")
	btn90 := menu.Text("Rainbow")
	btn91 := menu.Text("Qo'g'irchoqlar sirki")
	btn92 := menu.Text("Yulduz Farzandi")
	btn93 := menu.Text("Elita Sinfi")
	btn94 := menu.Text("Tajribasiz Senpai")
	btn95 := menu.Text("Qudratli Soʻngi Dushman")
	btn96 := menu.Text("Yangi Darvoza")
	btn97 := menu.Text("Cheksizlikgacha Lv9999")
	btn98 := menu.Text("Josus X Oilasi")
	btn99 := menu.Text("Ragnarok Rekordi")
	btn100 := menu.Text("Sen uchun O'lmas")
	btn101 := menu.Text("Masamune Qasosi")
	btn102 := menu.Text("Kim Meni Malika Qildi")
	btn103 := menu.Text("Koʻk Zindon")
	btn104 := menu.Text("Parazit - Hayot Saboqlari")
	btn105 := menu.Text("Uysiz Ma'bud")
	btn106 := menu.Text("Pari Dumi haqida afsona")
	btn107 := menu.Text("Kayju 8-Raqam")
	btn108 := menu.Text("Soqolimni olib, yuqori maktab qizini uyimga olib keldim")
	btn109 := menu.Text("Yetti O'lim Gunohlari")
	btn110 := menu.Text("Omadsizning qayta tug'ilishi")
	btn111 := menu.Text("Qirol oʻyini")
	btn112 := menu.Text("Violet Evergarden")
	btn113 := menu.Text("Oʻzga dunyoda darajamni koʻtarib, bu dunyoda ham tengsiz boʻldim")
	btn114 := menu.Text("Hyouka")
	btn115 := menu.Text("U qiz yolgiz")
	btn116 := menu.Text("Bir soatli qizcha")
	btn117 := menu.Text("Jonli Sana")
	btn118 := menu.Text("qip-qizil ragna")
	btn119 := menu.Text("Biz birga bo'lsak, sevgimiz har qanday to'siqni ortda qoldiradi")
	btn120 := menu.Text("Moviy Quticha")
	btn121 := menu.Text("Hikaruning songi yozi")
	btn122 := menu.Text("Taqdir: Buyuk Tartib Mutlaq Iblislar Jabhasi")
	btn123 := menu.Text("Sevgi deb atalgan shart")
	btn124 := menu.Text("Afsonaviy qahramonlar va Ruhlar malikasi qizi sifatida qayta tug‘ildim")
	btn125 := menu.Text("Lideyl Dunyosi")
	btn126 := menu.Text("Skelet Ritsar o‘zga dunyoda")
	btn127 := menu.Text("Qora chaqiruvchi")
	btn128 := menu.Text("Davolash sehridan foydalanishni noto'g'ri usuli")
	btn129 := menu.Text("Lordi Armiyasining eng kuchli Sehrgari...")
	btn130 := menu.Text("Tahlil qilish qobiliyatiga ega aristokrat bo'lib qayta tug'ilish")
	btn131 := menu.Text("Ilohiy qilich maktabining Iblis qilich egasi")
	btn132 := menu.Text("Meni Qilich bo'lib qayta tug'ilishim haqida")
	btn133 := menu.Text("O'z joniga qasd qiluvchilar o'zga dunyoda")
	btn134 := menu.Text("Eng qudratli partiya tomonidan o'limgacha tarbiyalangan Ossan ...")
	btn135 := menu.Text("Baholovchining mashhur bo'lmagan ishi aslida eng Kuchlisi hisoblanadi")
	btn136 := menu.Text("Sehr Yaratuvchi Boshqa dunyoda qanday qilib sehr yaratish mumkin")
	btn137 := menu.Text("O'rta yoshli erkakning zodagon qiziga aylanishi")
	btn138 := menu.Text("Yaponiyaga Xush Kelibsiz, Elf Xonim!")
	btn139 := menu.Text("Grimgaldagi kulgi va illuziya")
	btn140 := menu.Text("Nega hamma meni dunyomni unutdi?")
	btn141 := menu.Text("80.000 oltin tanga to'plab hayotimni qayta qurdim")
	btn142 := menu.Text("Cheksiz dendagron")
	btn143 := menu.Text("Oy sayohati yangi dunyoga olib keladi")
	btn144 := menu.Text("Man o'rgimchakman! Ho'sh shunga nima qibti?")
	btn145 := menu.Text("O'zga dunyoda fermerlik hayotim")
	btn146 := menu.Text("Iblislar maktabiga hush kelibsiz")
	btn147 := menu.Text("Daholar uchun oʻzga dunyoda yashash ham muammo emas")
	btn148 := menu.Text("Shirin renkarnatsiya")
	btn149 := menu.Text("Menda million hayot bor")
	btn150 := menu.Text("Maktab tomonidan tan olinmagan iblislar hukmdori")
	btn151 := menu.Text("Boshqa dunyo bir zumda o'lim kuchiga dosh bera olmaydi")
	btn152 := menu.Text("Kumush qirolning qayta tugʻilishi")
	btn153 := menu.Text("Oʻzga dunyoda Noyob mahoratim")
	btn154 := menu.Text("Egzartis boshqa dunyoda qayta tug'ilib eng kuchli bo'lishga intiladi")
	btn155 := menu.Text("Nikoh uzuklari haqida afsona")
	btn156 := menu.Text("tanyang urush yilnomalari")
	btn157 := menu.Text("Yovuzlik darajasi 99: Men yashirin xo'jayin bo'lishim mumkin, lekin men jin xo'jayini emasman")
	btn158 := menu.Text("Doktor Eliza: hayotlarni saqlab qoluvchi malika")
	btn159 := menu.Text("Eng zaif yirtqich hayvon")
	btn160 := menu.Text("Fojiaga sababchi bõlgan malika xalq uchun qõlidan kelgan barcha ishni qiladi")
	btn161 := menu.Text("Iblislar hukmdori qoʻshimcha ishda!")
	btn162 := menu.Text("Oying bilan birga video o'yin")
	btn163 := menu.Text("Bekami Ko'st Yashashim Ushun Hamma Narsa Qildim")
	btn164 := menu.Text("Uzoq paladin")
	btn165 := menu.Text("Realizm qahramoni qirollikni qayta qurishi")
	btn166 := menu.Text("Olti barg qahramonlari")
	btn167 := menu.Text("Daho Shifokorning Soyadagi Yangi Hayoti")
	btn168 := menu.Text("Qayta tugʻilgan aristokratning misli koʻrilmagan sarguzashtlari")
	btn169 := menu.Text("Sakkizinchi o'g'il qo'ysangizlarchi?")
	btn170 := menu.Text("O'yindagi Eng Boy odam")
	btn171 := menu.Text("O'yinsiz hayot yo'q")
	btn172 := menu.Text("Arifureta: Dunyodagi eng kuchli hunarmand")
	btn173 := menu.Text("Gildiya adminstratori bo'lib ortiqcha ishlashni xohlamaganim uchun ishdan ketmoqchiman")
	btn174 := menu.Text("Oddiy insondan qahramonlikkacha")
	btn175 := menu.Text("Tayoq va qilich")
	btn176 := menu.Text("Re:Zero")
	btn177 := menu.Text("Onmyo Qayta Tug‘ilishi: Hayolot Olami")
	btn178 := menu.Text("Yettinchi shahzoda sifatida qayta tug'ildim va endi sehrimni istaganimcha kuchaytiraman!")
	btn179 := menu.Text("Mening Qotillik Maqomim Qahramonlik Maqomidan Yaxshiroq")
	btn180 := menu.Text("Yugurening abadiyligi")
	btn181 := menu.Text("Daydi itlarning buyugi")
	btn182 := menu.Text("Meni qizcham nafaqat go'zal")
	btn183 := menu.Text("Meni qahramonlik akademiyam")
	btn184 := menu.Text("Barmoqlar uchidagi sevgi")
	btn185 := menu.Text("Shangri-la chegarasi")
	btn186 := menu.Text("Kelajak kundaligi")
	btn187 := menu.Text("Men eng kuchli sarguzashtchi bo'lish uchun har doim mashq qildim")
	btn188 := menu.Text("Do'stimning singlisi bezovta qilyapti")
	btn189 := menu.Text("Shikastlanishni istamasdim shuning uchun himoyamni kuchaytirdim")
	btn190 := menu.Text("Oxirgi Telba Boss paydo bo'ldi")
	btn191 := menu.Text("Vayron bo'lgan mo'jizalar mamlakati")
	btn192 := menu.Text("Qahramonning qaytishi")
	btn193 := menu.Text("Ninja va Yakudza")
	btn194 := menu.Text("Bosning qizi va uning Enagasi")
	btn195 := menu.Text("Detektiv allaqchon o'lgan")
	btn196 := menu.Text("Friren - Soʻnggi manzilga kuzatuvchi")
	btn197 := menu.Text("Exo Terror")
	btn198 := menu.Text("Arknayts: Prelyudiya tongga tomon")
	btn199 := menu.Text("Takt Opus. Taqdir")
	btn200 := menu.Text("Orient")
	btn201 := menu.Text("Magic Kaito Kid ID")
	btn202 := menu.Text("Chitose-kun va ramuna shishasi")

	menu.Reply(
		menu.Row(btn1, btn2),
		menu.Row(btn3),
		menu.Row(btn4),
		menu.Row(btn5),
		menu.Row(btn6),
		menu.Row(btn7),
		menu.Row(btn8),
		menu.Row(btn9),
		menu.Row(btn10),
		menu.Row(btn11),
		menu.Row(btn12),
		menu.Row(btn13),
		menu.Row(btn14),
		menu.Row(btn15),
		menu.Row(btn16),
		menu.Row(btn17),
		menu.Row(btn18),
		menu.Row(btn19),
		menu.Row(btn20),
		menu.Row(btn21),
		menu.Row(btn22),
		menu.Row(btn23),
		menu.Row(btn24),
		menu.Row(btn25),
		menu.Row(btn26),
		menu.Row(btn27),
		menu.Row(btn28),
		menu.Row(btn29),
		menu.Row(btn30),
		menu.Row(btn31),
		menu.Row(btn32),
		menu.Row(btn33),
		menu.Row(btn34),
		menu.Row(btn35),
		menu.Row(btn36),
		menu.Row(btn37),
		menu.Row(btn38),
		menu.Row(btn39),
		menu.Row(btn40),
		menu.Row(btn41),
		menu.Row(btn42),
		menu.Row(btn43),
		menu.Row(btn44),
		menu.Row(btn45),
		menu.Row(btn46),
		menu.Row(btn47),
		menu.Row(btn48),
		menu.Row(btn49),
		menu.Row(btn50),
		menu.Row(btn51),
		menu.Row(btn52),
		menu.Row(btn53),
		menu.Row(btn54),
		menu.Row(btn55),
		menu.Row(btn56),
		menu.Row(btn57),
		menu.Row(btn58),
		menu.Row(btn59),
		menu.Row(btn60),
		menu.Row(btn61),
		menu.Row(btn62),
		menu.Row(btn63),
		menu.Row(btn64),
		menu.Row(btn65),
		menu.Row(btn66),
		menu.Row(btn67),
		menu.Row(btn68),
		menu.Row(btn69),
		menu.Row(btn70),
		menu.Row(btn71),
		menu.Row(btn72),
		menu.Row(btn73),
		menu.Row(btn74),
		menu.Row(btn75),
		menu.Row(btn76),
		menu.Row(btn77),
		menu.Row(btn78),
		menu.Row(btn79),
		menu.Row(btn80),
		menu.Row(btn81),
		menu.Row(btn82),
		menu.Row(btn83),
		menu.Row(btn84),
		menu.Row(btn85),
		menu.Row(btn86),
		menu.Row(btn87),
		menu.Row(btn88),
		menu.Row(btn89),
		menu.Row(btn90),
		menu.Row(btn91),
		menu.Row(btn92),
		menu.Row(btn93),
		menu.Row(btn94),
		menu.Row(btn95),
		menu.Row(btn96),
		menu.Row(btn97),
		menu.Row(btn98),
		menu.Row(btn99),
		menu.Row(btn100),
		menu.Row(btn101),
		menu.Row(btn102),
		menu.Row(btn103),
		menu.Row(btn104),
		menu.Row(btn105),
		menu.Row(btn106),
		menu.Row(btn107),
		menu.Row(btn108),
		menu.Row(btn109),
		menu.Row(btn110),
		menu.Row(btn111),
		menu.Row(btn112),
		menu.Row(btn113),
		menu.Row(btn114),
		menu.Row(btn115),
		menu.Row(btn116),
		menu.Row(btn117),
		menu.Row(btn118),
		menu.Row(btn119),
		menu.Row(btn120),
		menu.Row(btn121),
		menu.Row(btn122),
		menu.Row(btn123),
		menu.Row(btn124),
		menu.Row(btn125),
		menu.Row(btn126),
		menu.Row(btn127),
		menu.Row(btn128),
		menu.Row(btn129),
		menu.Row(btn130),
		menu.Row(btn131),
		menu.Row(btn132),
		menu.Row(btn133),
		menu.Row(btn134),
		menu.Row(btn135),
		menu.Row(btn136),
		menu.Row(btn137),
		menu.Row(btn138),
		menu.Row(btn139),
		menu.Row(btn140),
		menu.Row(btn141),
		menu.Row(btn142),
		menu.Row(btn143),
		menu.Row(btn144),
		menu.Row(btn145),
		menu.Row(btn146),
		menu.Row(btn147),
		menu.Row(btn148),
		menu.Row(btn149),
		menu.Row(btn150),
		menu.Row(btn151),
		menu.Row(btn152),
		menu.Row(btn153),
		menu.Row(btn154),
		menu.Row(btn155),
		menu.Row(btn156),
		menu.Row(btn157),
		menu.Row(btn158),
		menu.Row(btn159),
		menu.Row(btn160),
		menu.Row(btn161),
		menu.Row(btn162),
		menu.Row(btn163),
		menu.Row(btn164),
		menu.Row(btn165),
		menu.Row(btn166),
		menu.Row(btn167),
		menu.Row(btn168),
		menu.Row(btn169),
		menu.Row(btn170),
		menu.Row(btn171),
		menu.Row(btn172),
		menu.Row(btn173),
		menu.Row(btn174),
		menu.Row(btn175),
		menu.Row(btn176),
		menu.Row(btn177),
		menu.Row(btn178),
		menu.Row(btn179),
		menu.Row(btn180),
		menu.Row(btn181),
		menu.Row(btn182),
		menu.Row(btn183),
		menu.Row(btn184),
		menu.Row(btn185),
		menu.Row(btn186),
		menu.Row(btn187),
		menu.Row(btn188),
		menu.Row(btn189),
		menu.Row(btn190),
		menu.Row(btn191),
		menu.Row(btn192),
		menu.Row(btn193),
		menu.Row(btn194),
		menu.Row(btn195),
		menu.Row(btn196),
		menu.Row(btn197),
		menu.Row(btn198),
		menu.Row(btn199),
		menu.Row(btn200),
		menu.Row(btn201),
		menu.Row(btn202),
	)
	switch text {

	case "Animelar", "/menu":
		return c.Send("asosiy kanal -https://t.me/Anmelaruzb ", menu)
	default:
		return c.Send("Iltimos, /menu dan danlang.")
	}
}
