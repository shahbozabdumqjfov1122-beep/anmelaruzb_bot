package Lookism

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	b := c.Bot()
	channelID := int64(-1003287478561)
	messageIDs := []int{2, 3, 4, 5, 6, 7, 8, 9}

	for _, msgID := range messageIDs {
		msg := &tele.Message{
			ID:   msgID,
			Chat: &tele.Chat{ID: channelID},
		}

		// Xabarni forward emas, copy qilish (faqat video chiqadi, yozuvsiz)
		_, err := b.Copy(c.Sender(), msg)
		if err != nil {
			return c.Send("❌ Video olishda xatolik: " + err.Error())
		}
	}

	return nil
}
