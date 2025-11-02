package entities

type UuidV7Str string

type UuidV7Provider interface {
	NewUuidV7() UuidV7Str
}