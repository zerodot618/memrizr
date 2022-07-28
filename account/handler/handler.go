package handler

import (
	"net/http"

	"memrizr/model"

	"github.com/gin-gonic/gin"
)

// Handler 保存处理程序运行所需的服务
type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

// Config 初始化 handler 包所需的配置数据
type Config struct {
	R            *gin.Engine
	UserService  model.UserService
	TokenService model.TokenService
	BaseURL      string
}

// NewHandler 初始化需要注入的路由及初始数据
// 不返回，因为它直接处理 gin 引擎的引用
func NewHandler(c *Config) {
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}

	// g := c.R.Group("/api/account")
	g := c.R.Group(c.BaseURL)

	g.GET("/me", h.Me)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "space peoples wewewe",
		})
	})
}

func (h *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

func (h *Handler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's me",
	})
}

// Image handler
func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

// DeleteImage handler
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteImage",
	})
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
