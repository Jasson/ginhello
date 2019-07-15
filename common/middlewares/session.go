package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// SessionsMiddleware .
func SessionsMiddleware(c *gin.Context) gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "localhost:32774", "", []byte("secret"))
	return sessions.Sessions("mysession", store)
}
