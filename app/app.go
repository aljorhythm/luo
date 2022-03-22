package app

import "github.com/google/uuid"

type App struct {
}

func (a App) EnterItem(item Item) (uuid.UUID, error) {
	id, _ := uuid.NewUUID()
	return id, nil
}

func (a App) RetrieveItem(id uuid.UUID) Item {
	return Item{}
}
