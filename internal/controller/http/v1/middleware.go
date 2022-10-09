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
		log.Println("middleware: get authorization header failed")
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		log.Println("middleware: invalid auth header")
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		log.Println("middleware: ", err.Error())
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userIDCtx, userID)
}
