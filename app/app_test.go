package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppEnterRetrieveItem(t *testing.T) {
	app := App{}
	item := Item{}
	err := app.EnterItem(item)

	assert.NoError(t, err)
}
