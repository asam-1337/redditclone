package v1

import (
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) RootPage(c *gin.Context) {
	err := h.Tmpl.ExecuteTemplate(c.Writer, "../static/index.html", nil)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) SignUp(c *gin.Context) {
	var input entity.User
	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	err = h.services.Authorization.CreateUser(&input)
}

func (h *Handler) Login(c *gin.Context) {

}
