package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testcase struct {
	item     Item
	testName string
}

func TestAppEnterRetrieveItem(t *testing.T) {

	cases := []testcase{
		{
			testName: "fool test",
			item:     Item{},
		},

		{
			testName: "an item",
			item: Item{
				Url: "http://somewhere",
			},
		},
	}

	app := NewApp()

	for _, aCase := range cases {
		t.Run(aCase.testName, func(t *testing.T) {
			item := aCase.item
			itemId, err := app.EnterItem(item)

			assert.NoError(t, err)

			retrievedItem := app.RetrieveItem(itemId)
			assert.Equal(t, item, retrievedItem)
		})
	}
}
