package _default

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	b := c.Bot()
	channelID := int64(-1002974691269)
	messageIDs := []int{2, 4}

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
