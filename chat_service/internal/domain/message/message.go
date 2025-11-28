package message

import (
	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID
	ReceiverId uuid.UUID
}
