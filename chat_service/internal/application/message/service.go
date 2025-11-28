package message

import "github.com/noo-dev/pet_chat/internal/domain/message"

type MessageUsecase struct {
	messageRepo message.MessageRepository
}

func NewMessageUsecase() *MessageUsecase {
	return &MessageUsecase{}
}

func (s *MessageUsecase) SendMessage(msg message.Message) error {
	return s.messageRepo.Save(msg)
}
