package entities

type UserRole uint8

const (
	AdminRole UserRole = iota
	SystemRole
	FanRole
)

