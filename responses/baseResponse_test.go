package responses

import (
	"github.com/stretchr/testify/assert"
	"go-simple-blog/contracts/statusCodes"
	"go-simple-blog/entities"
	"testing"
)

func TestResponse(t *testing.T) {
	t.Run("create user response", func(t *testing.T) {
		response := New(PostResponse{}, true, statusCodes.Error, "ok", entities.Post{})
		assert.True(t, response.GetStatus())
		assert.Equal(t, statusCodes.Error, response.GetStatusCode())
		assert.Equal(t, "ok", response.GetMessage())
		response.SetMessage("ko")
		assert.Equal(t, "ko", response.GetMessage())
		assert.True(t, response.ErrorIs(statusCodes.Error))
		assert.False(t, response.ErrorIs(statusCodes.Success))
		assert.False(t, response.IsFailed())
		assert.IsType(t, entities.Post{}, response.GetData())
		assert.IsType(t, PostResponse{}, response)
		assert.Equal(t, response.ToMap()["data"], response.GetData())
	})

	t.Run("create auth response", func(t *testing.T) {
		response := New(UserResponse{}, true, statusCodes.Error, "ok", entities.Post{})
		assert.True(t, response.GetStatus())
		assert.Equal(t, statusCodes.Error, response.GetStatusCode())
		assert.Equal(t, "ok", response.GetMessage())
		assert.IsType(t, entities.Post{}, response.GetData())
		assert.IsType(t, UserResponse{}, response)
	})

	t.Run("create auth response with status false", func(t *testing.T) {
		response := New(UserResponse{}, false, statusCodes.Error, "ok", entities.User{})
		assert.False(t, response.GetStatus())
		assert.True(t, response.IsFailed())
		assert.Equal(t, statusCodes.Error, response.GetStatusCode())
		assert.Equal(t, "ok", response.GetMessage())
		assert.IsType(t, entities.User{}, response.GetData())
		assert.IsType(t, UserResponse{}, response)
	})

}
