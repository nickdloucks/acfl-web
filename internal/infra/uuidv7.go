package infra

import (
	"nickdloucks/acfl-web/internal/entities"
	googleUuid "github.com/google/uuid"
)

type GoogleUuidV7Provider struct {}

func (GoogleUuidV7Provider) NewUuidV7() entities.UuidV7Str {
	uuid, _ := googleUuid.NewV7()
	str := uuid.String()
	return entities.UuidV7Str(str)
}

