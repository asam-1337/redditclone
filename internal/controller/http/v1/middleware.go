package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "user_id"
	usernameCtx         = "username"
)

func (h *Handler) AuthMiddleware(c *gin.Context) {

	log.Printf("url: %s, method: %s\n", c.Request.URL, c.Request.Method)

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	log.Println("middleware: get authorization header")

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	log.Println("middleware: split authorization header")

	userID, username, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		log.Println("middleware: ", err.Error())
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	log.Println("middleware: parse token")

	c.Set(userIDCtx, userID)
	c.Set(usernameCtx, username)
}
