package handler

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"os" // مكتبة للتعامل مع النظام
	"time"
)

// استدعاء التوكن من متغيرات البيئة
func getApiURL() string {
	token := os.Getenv("TELEGRAM_TOKEN")
	return "https://api.telegram.org/bot" + token + "/setMessageReaction"
}

var emojis = []string{
	"👍", "👎", "❤️", "🔥", "🥰", "👏", "😁", "🤔", "🤯", "😱", "🤬", "😢", "🎉", "🤩", "🤮", "💩", "🙏", "👌", "🕊️", "🤡",
	"🥱", "🥴", "😍", " whale", "❤️‍🔥", "🌚", "🌭", "💯", "🤣", "⚡", "🍌", "🏆", "💔", "🤨", "😐", "🍓", "🍾", "💋", "🖕", "😈",
	"😴", "😭", "🤓", "👻", "👨‍💻", "👀", "🎃", "🙈", "😇", "😨", "🤝", "✍️", "🤗", "🫡", "🎅", "🎄", "☃️", "💅", "🤪", "🗿",
	"🆒", "💘", "🙊", "🦄", "😘", "💊", "🙉", "😎", "👾", "🤷‍♂️", "🤷", "🤷‍♀️", "😡",
}

type Update struct {
	Message struct {
		MessageID int64 `json:"message_id"`
		Chat      struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		return
	}

	if update.Message.MessageID != 0 {
		rand.Seed(time.Now().UnixNano())
		randomEmoji := emojis[rand.Intn(len(emojis))]
		
		addReaction(update.Message.Chat.ID, update.Message.MessageID, randomEmoji)
	}

	w.WriteHeader(http.StatusOK)
}

func addReaction(chatID int64, messageID int64, emoji string) {
	payload, _ := json.Marshal(map[string]interface{}{
		"chat_id":    chatID,
		"message_id": messageID,
		"reaction": []map[string]string{
			{"type": "emoji", "emoji": emoji},
		},
	})
	http.Post(getApiURL(), "application/json", bytes.NewBuffer(payload))
}
