package vanpis

import (
	tele "gopkg.in/telebot.v4"
)

func Home(c tele.Context) error {
	b := c.Bot()
	channelID := int64(-1002957299408)
	messageIDs := []int{}
	for i := 559; i < 2000; i++ {
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
