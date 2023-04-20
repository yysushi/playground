package example

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yysushi/playground/go/vektra/mockery/mocks"
)

func TestExample(t *testing.T) {
	mockedDoer := mocks.NewMockDoer(t)
	mockedDoer.EXPECT().Do("hoge").RunAndReturn(func(s string) string {
		src := []byte(s)
		return hex.EncodeToString(src)
	}).Once()
	c := Character{
		Source: "hoge",
		Doer:   mockedDoer,
	}
	_ = c
	assert.Equal(t, c.Output(), "686f6765")
	// assert.Equal(t, c.Output(), "686f6765")
}
