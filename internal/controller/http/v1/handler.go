package v1

import (
	"github.com/asam-1337/reddit-clone.git/internal/service"
	"github.com/gin-gonic/gin"
	"html/template"
)

type Handler struct {
	services *service.Service
	Tmpl     *template.Template
}

func NewHandler(services *service.Service, tmpl *template.Template) *Handler {
	return &Handler{
		services: services,
		Tmpl:     tmpl,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", h.RootPage)
	api := router.Group("/api")
	{
		api.POST("/register")
		api.POST("/login")
		api.GET("/user/:user_login")

		posts := api.Group("/posts")
		{
			posts.GET("/")
			posts.POST("/")
			posts.GET("/:category_name")
		}

		post := api.Group("/post")
		{
			post.GET("/:post_id")
			post.POST("/:post_id")
			post.DELETE("/:post_id")
			post.DELETE("/:post_id/:comment_id")
			post.GET("/:post_id/upvote")
			post.GET("/:post_id/downvote")
		}
	}

	return router
}
