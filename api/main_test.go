package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
	"udemy-course/db/sqlc"
	"udemy-course/utils"
)

func newTestServer(t *testing.T, store sqlc.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(33),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
