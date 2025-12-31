package MagicKaitoid

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	b := c.Bot()
	channelID := int64(-1003598783228)

	for i := 2; i < 20; i++ {

		msg := &tele.Message{
			ID:   i,
			Chat: &tele.Chat{ID: channelID},
		}

		// 🔒 AGAR VIDEO / RASM / FAYL bo‘lsa — SKIP
		if msg.Video != nil || msg.Photo != nil || msg.Document != nil {
			continue
		}

		_, err := b.Copy(c.Sender(), msg)
		if err != nil {
			continue
		}
	}

	return c.Send("❌ choparlaga mukun emas")
}
