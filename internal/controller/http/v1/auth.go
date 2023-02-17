package v1

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type authInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignUp(c *gin.Context) {
	input := &authInput{}
	err := c.BindJSON(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.CreateUser(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	sess := &entity.Session{
		Username:  input.Username,
		Useragent: c.GetHeader("User-Agent"),
	}
	sessID, err := h.sessManager.Create(sess)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   sessID.ID,
		Expires: time.Now().Add(12 * time.Hour),
	}

	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	input := &authInput{}
	err := c.BindJSON(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.Authenticate(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
