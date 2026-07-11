package infra

import (
	"nickdloucks/acfl-web/internal/entities"
	googleUuid "github.com/google/uuid"
)

type UuidV7ProviderImpl struct {}

func NewUuidV7ProviderImpl() UuidV7ProviderImpl {
	return UuidV7ProviderImpl{}
}

func (*UuidV7ProviderImpl) NewUuidV7() (entities.UuidV7Str, error) {
	uuid, googleUuidErr := googleUuid.NewV7()
	if googleUuidErr != nil {
		return "", googleUuidErr
	}
	str := uuid.String()
	return entities.UuidV7Str(str), nil
}

