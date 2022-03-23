package output

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type mockData struct{}

func TestOutputV1(t *testing.T) {
	data := mockData{}
	output := outputV1{}
	actual := string(output.print(data))
	expected := `
     {
"version": 1
}
`
	require.JSONEq(t, expected, actual)
}
