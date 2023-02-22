package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/miron239/wb/authz"
	"google.golang.org/api/idtoken"
)

type authEnforcement int

const (
	mandatory authEnforcement = iota
	optional
)

func authHandler(clientId string, enforcing authEnforcement) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" || c.Request.URL.Path == "/swagger/doc.json" {
			return
		}
		authH, ok := c.Request.Header["Authorization"]
		var user string
		if !ok && enforcing == mandatory {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "no Authorization header in request")
			return
		} else if !ok && enforcing == optional {
			user = "default-test-user@"
			if callerId := c.Request.Header.Get("CallerId"); callerId != "" {
				user = callerId
			}
		} else {
			token := strings.Replace(authH[0], "Bearer ", "", 1)
			payload, err := idtoken.Validate(context.Background(), token, clientId)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
				return
			}
			user = payload.Claims["email"].(string)
		}
		c.Set("user", user)
		path := strings.Split(strings.Trim(c.Request.URL.Path, "/"), "/")
		c.Set("decisionRequest", &authz.DecisionRequest{
			Method: c.Request.Method,
			Path:   path,
			User:   user,
		})
		c.Next()
	}
}
