package server

type webhook struct {
	Id        string `json:"id"`
	Version   string `json:"version"`
	Timestamp int    `json:"timestamp"`
	Messages  []struct {
		Id       string `json:"id"`
		ChatId   string `json:"chat_id"`
		UserId   int    `json:"user_id"`
		AuthorId int    `json:"author_id"`
		Created  int    `json:"created"`
		Type     string `json:"type"`
		Content  struct {
			Text string `json:"text"`
		} `json:"content,omitempty"`
	} `json:"messages,omitempty"`
}

func (w webhook) String() (result string) {
	for _, message := range w.Messages {
		result += "Тип сообщения: " + message.Type + "\n" + message.Content.Text + "\n"
	}
	return
}
