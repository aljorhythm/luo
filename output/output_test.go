package output

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type mockData struct{}

func TestOutput(t *testing.T) {
	data := mockData{}
	output := outputV1{}
	actual := string(output.print(data))
	expected := `
     {}
`
	require.JSONEq(t, expected, actual)
}
