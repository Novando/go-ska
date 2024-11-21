package uuid

import (
	"github.com/google/uuid"
)

func Undashed(id uuid.UUID) string {
	src := id.String()
	return (src[0:8] + src[9:13] + src[14:18] + src[19:23] + src[24:])
}

func GenerateUUID() uuid.UUID { return uuid.Must(uuid.NewV7()) }
