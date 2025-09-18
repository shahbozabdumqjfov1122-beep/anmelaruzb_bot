package QoraKlever

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	b := c.Bot()
	channelID := int64(-1002809090010)
	messageIDs := []int{}
	for i := 2; i < 200; i++ {
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
