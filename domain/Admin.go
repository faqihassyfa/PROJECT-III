package domain

import "time"

type Admin struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Clothes struct {
	ID        int
	Name      string
	Price     int
	Stock     int
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdminUseCase interface {
	CreateProducts(data Clothes)
	GetHistory([]Clothes, int)
}

type AdminData interface {
	GetHistoryData() []Clothes
}
