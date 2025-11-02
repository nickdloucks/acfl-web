package usecases

import "nickdloucks/acfl-web/internal/entities"


type UuidV7Provider interface {
	NewUuidV7() entities.UuidV7Str
}