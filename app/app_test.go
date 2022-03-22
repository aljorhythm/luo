package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppEnterRetrieveItem(t *testing.T) {
	app := App{}
	item := Item{}
	itemId, err := app.EnterItem(item)

	assert.NoError(t, err)

	retrievedItem := app.RetrieveItem(itemId)
	assert.Equal(t, item, retrievedItem)
}
