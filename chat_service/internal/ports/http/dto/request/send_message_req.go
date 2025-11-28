package request

import "github.com/google/uuid"

type SendMessageReqDto struct {
	UserID uuid.UUID `json:"userId"`
	Body   string    `json:"body"`
}
