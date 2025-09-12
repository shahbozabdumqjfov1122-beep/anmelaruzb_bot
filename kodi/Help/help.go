package Help

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	text := c.Text()

	switch text {

	case "🧩 help", "/help":
		return c.Send("admin---@TM_ESPORTS\nadmin---@TM_ESPORTS2")

	default:
		return c.Send("Noma'lum buyruq. Iltimos, menyudan foydalaning.")
	}
}
