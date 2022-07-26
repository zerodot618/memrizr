package handler

import (
	"log"
	"memrizr/model"
	"memrizr/model/apperrors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signupReq 注册请求结构体
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup 注册
func (h *Handler) Signup(c *gin.Context) {
	var req signupReq

	// 检验请求数据
	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	ctx := c.Request.Context()
	err := h.UserService.Signup(ctx, u)

	if err != nil {
		log.Printf("Faild to sign up user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// 创建 token pair 并转化为字符串
	tokens, err := h.TokenService.NewTokenPairFromUser(ctx, u, "")

	if err != nil {
		log.Printf("Faild to create tokens for user: %v\n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}
