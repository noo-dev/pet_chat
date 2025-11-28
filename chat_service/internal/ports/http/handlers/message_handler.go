package handlers

import (
	"github.com/labstack/echo/v4"
	msgUsecase "github.com/noo-dev/pet_chat/internal/application/message"
	"github.com/noo-dev/pet_chat/internal/domain/message"
	"github.com/noo-dev/pet_chat/internal/ports/http/dto/request"
)

type MessageHandler struct {
	messageUsecase *msgUsecase.MessageUsecase
}

func NewMessageHandler(mu *msgUsecase.MessageUsecase) *MessageHandler {
	return &MessageHandler{messageUsecase: mu}
}

func (h *MessageHandler) SendMessage(c echo.Context) error {
	var reqBody request.SendMessageReqDto
	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	var msg message.Message
	return h.messageUsecase.SendMessage(msg)
}
