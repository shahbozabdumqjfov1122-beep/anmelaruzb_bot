package Zombi100

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	b := c.Bot()
	channelID := int64(-1003260007356)
	messageIDs := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	for _, msgID := range messageIDs {
		msg := &tele.Message{
			ID:   msgID,
			Chat: &tele.Chat{ID: channelID},
		}
		_, err := b.Copy(c.Sender(), msg)
		if err != nil {
			return c.Send("❌ Video olishda xatolik: " + err.Error())
		}
	}
	return nil
}
