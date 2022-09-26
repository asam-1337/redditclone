package v1

import (
	"fmt"
	"github.com/asam-1337/reddit-clone.git/internal/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) CreatePost(c *gin.Context) {
	log.Printf("url: %s, method: %s\n", c.Request.URL, c.Request.Method)

	val, ok := c.Get(userIDCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "userID does not exist in context")
		return
	}

	log.Printf("createPost: get context params userID: %s", val)

	input := &entity.Post{}
	err := c.BindJSON(input)
	fmt.Println()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("createPost: ", input)

	userID, ok := val.(string)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "invalid type of userID")
		return
	}

	post, err := h.services.Posts.CreatePost(input, userID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Println("create Post")

	c.JSON(http.StatusOK, post)
}

func (h *Handler) GetPostByID(c *gin.Context) {
	postID := c.Param("post_id")

	post, err := h.services.Posts.GetPostByID(postID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) GetPostsByUsername(c *gin.Context) {
	username := c.Param("username")

	posts, err := h.services.Posts.GetPostsByUsername(username)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, &posts)
}

func (h *Handler) GetPostsByCategory(c *gin.Context) {
	category := c.Param("category")

	posts, err := h.services.Posts.GetPostsByCategory(category)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, &posts)
}

func (h *Handler) GetAllPosts(c *gin.Context) {
	posts, err := h.services.Posts.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, &posts)
}

func (h *Handler) DeletePost(c *gin.Context) {
	postID := c.Param("post_id")

	err := h.services.Posts.DeletePost(postID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "post not found")
		return
	}

	newErrorResponse(c, http.StatusOK, "success")
}

func (h *Handler) CreateComment(c *gin.Context) {
	val, ok := c.Get(userIDCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "userID does not exist in context")
		return
	}

	userID, ok := val.(string)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "invalid type of userID")
		return
	}

	postID := c.Param("post_id")
	input := map[string]string{
		"comment": "",
	}

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	post, err := h.services.Posts.CreateComment(userID, postID, input["comment"])
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) GetUpvote(c *gin.Context) {
	val, ok := c.Get(userIDCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "userID does not exist in context")
		return
	}

	userID, ok := val.(string)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "invalid type of userID")
		return
	}

	postID := c.Param("post_id")

	input := &entity.Vote{
		UserId: userID,
		Vote:   1,
	}

	post, err := h.services.Vote(postID, input)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) GetDownvote(c *gin.Context) {
	val, ok := c.Get(userIDCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "userID does not exist in context")
		return
	}

	userID, ok := val.(string)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "invalid type of userID")
		return
	}

	postID := c.Param("post_id")

	input := &entity.Vote{
		UserId: userID,
		Vote:   -1,
	}

	post, err := h.services.Vote(postID, input)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) GetUnvote(c *gin.Context) {
	val, ok := c.Get(userIDCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "userID does not exist in context")
		return
	}

	userID, ok := val.(string)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "invalid type of userID")
		return
	}

	postID := c.Param("post_id")

	post, err := h.services.Unvote(userID, postID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}
