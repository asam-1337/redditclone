package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "user_id"
	usernameCtx         = "username"
)

func (h *Handler) AuthMiddleware(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		logrus.Println("middleware: get authorization header failed")
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		logrus.Println("middleware: invalid auth header")
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		logrus.Println("middleware: ", err.Error())
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userIDCtx, userID)
}

func (h *Handler) AccessLogMiddleware(c *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"metgod": c.Request.Method,
		"url":    c.Request.URL.Path,
	}).Info(c.Request.URL.Path)
}
