package sdk

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	t.Run("a new client is created", func(t *testing.T) {
		client := Connect("vulncheck_token")
		if client == nil {
			t.Error("client is nil")
		}

		assert.NotNil(t, client)
	})
}

func TestSetAuthHeader(t *testing.T) {
	t.Run("auth header is set", func(t *testing.T) {
		client := Connect("vulncheck_token")
		req := httptest.NewRequest("GET", TestURL+"/index", nil)
		client.SetAuthHeader(req)

		assert.Equal(t, "Bearer vulncheck_token", req.Header.Get("Authorization"))
	})
}

func TestGetBaseURL(t *testing.T) {
	t.Run("base url is set", func(t *testing.T) {
		client := Connect("vulncheck_token")
		assert.Equal(t, ProdURL, client.GetBaseURL())
	})
}
