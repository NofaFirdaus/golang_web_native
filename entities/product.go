package entities

type Product struct {
	Id          uint
	Name        string
	CategoryId  uint
	Stock       uint
	Description string
	CreatedAt   []byte
	UpdatedAt   []byte
}
