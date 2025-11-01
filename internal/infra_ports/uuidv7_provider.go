package ports

import (
	"nickdloucks/acfl-web/internal/domain"
)

type UuidV7 interface {
	NewUuidV7()	UuidV7Str
}