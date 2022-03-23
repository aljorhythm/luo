package app

import "github.com/google/uuid"

type app struct {
	items map[uuid.UUID]Item
}

type App interface {
	EnterItem(item Item) (uuid.UUID, error)
	RetrieveItem(id uuid.UUID) Item
}

func (a app) EnterItem(item Item) (uuid.UUID, error) {
	id, _ := uuid.NewUUID()
	a.items[id] = item
	return id, nil
}

func (a app) RetrieveItem(id uuid.UUID) Item {
	return a.items[id]
}

func NewApp() App {
	return &app{
		items: map[uuid.UUID]Item{},
	}
}
