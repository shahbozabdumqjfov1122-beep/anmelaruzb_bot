package Help

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	text := c.Text()

	switch text {

	case "🧩 help", "/help":
		return c.Send("🧠 Yordam bo‘limi\n\n" +
			"Bizning bot turli xizmatlarni taqdim etadi va sizga yordam berishga tayyor!\n\n" +
			"📞 *Agar muammo yoki takliflaringiz bo‘lsa*, quyidagi adminlar bilan bog‘laning:\n" +
			"👤 @TM_ESPORTS\n" +
			"👤 @Animelaruzb_admin\n\n" +
			"✨ Har doim siz bilan birga — Biz bot jamoasi! 🤖")

	default:
		return c.Send("Noma'lum buyruq. Iltimos, menyudan foydalaning.")
	}
}
