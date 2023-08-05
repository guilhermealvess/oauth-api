package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestApiServer(t *testing.T) {
	t.Run("successfully new api server", func(t *testing.T) {
		const (
			name   = "xpto api"
			author = "fulano@example.com"
		)
		apiServer, err := NewApiServer(name, "", author)

		assert.NoError(t, err)
		assert.Equal(t, apiServer.Name, name)
		assert.Equal(t, apiServer.Description, "")
		assert.Equal(t, apiServer.CreatedBy, author)
		assert.Equal(t, "xpto-api", apiServer.Slug)
		assert.Equal(t, true, time.Now().After(apiServer.CreatedAt))
	})

	t.Run("failure new api server", func(t *testing.T) {
		testCases := []struct {
			name   string
			desc   string
			author string
		}{
			{"@Xpto", "", ""},
			{"xpto api", "", ""},
			{"", "", "xpto@gmail.com"},
		}

		for _, tc := range testCases {
			_, err := NewApiServer(tc.name, tc.desc, tc.author)
			assert.Error(t, err)
		}
	})
}
