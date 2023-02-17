package v1

import (
	"github.com/asam-1337/reddit-clone.git/internal/service"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	services    *service.Service
	sessManager service.SessionManagerInterface
	logrus      *logrus.Entry
}

func NewHandler(services *service.Service, conn redis.Conn) *Handler {
	return &Handler{
		services:    services,
		sessManager: service.NewSessionManager(conn),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("./static/html/*")
	router.StaticFS("/static", http.Dir("./static"))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Root",
		})
	})
	api := router.Group("/api", h.AccessLogMiddleware)
	{
		api.POST("/register", h.SignUp)
		api.POST("/login", h.SignIn)
		api.GET("/user/:username", h.GetPostsByUsername)

		posts := api.Group("/posts")
		{
			posts.GET("/", h.GetAllPosts)
			posts.POST("/", h.AuthMiddleware, h.CreatePost)
			posts.GET("/:category", h.GetPostsByCategory)
		}

		post := api.Group("/post/:post_id")
		{
			post.GET("/", h.GetPostByID)
			post.POST("/", h.CreateComment)
			post.DELETE("/", h.AuthMiddleware, h.DeletePost)
			post.DELETE("/:comment_id", h.AuthMiddleware)

			post.GET("/upvote", h.AuthMiddleware, h.GetUpvote)
			post.GET("/unvote", h.AuthMiddleware, h.GetUnvote)
			post.GET("/downvote", h.AuthMiddleware, h.GetDownvote)
		}
	}

	return router
}
