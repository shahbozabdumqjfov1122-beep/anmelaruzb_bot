package Franksdagisevgi

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	b := c.Bot()
	channelID := int64(-1003091302822)
	messageIDs := []int{28, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}

	for _, msgID := range messageIDs {
		// Kanal xabarini olish
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
