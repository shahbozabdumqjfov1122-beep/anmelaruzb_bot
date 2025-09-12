package Titanlarjangi

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	b := c.Bot()
	channelID := int64(-1003043779549)
	messageIDs := []int{}
	for i := 4; i < 100; i++ {
		messageIDs = append(messageIDs, i)
	}

	for _, msgID := range messageIDs {

		msg := &tele.Message{
			ID:   msgID,
			Chat: &tele.Chat{ID: channelID},
		}

		_, err := b.Copy(c.Sender(), msg)
		if err != nil {
			// ❌ Xato bo‘lsa o‘sha xabarni tashlab keyingisiga o‘tadi
			continue
		}
	}

	return c.Send("✅ Videolar yuborish tugadi")
}
