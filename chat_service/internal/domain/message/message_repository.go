package message

type MessageRepository interface {
	Save(message Message) error
}
