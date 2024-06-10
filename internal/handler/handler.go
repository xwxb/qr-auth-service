package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xwxb/qr-auth-service/internal/consts"
	"github.com/xwxb/qr-auth-service/internal/service"
	"net/http"
)

func GenUUID(c *gin.Context) {
	// Generate a random UUID
	newUUID := uuid.New()

	// Return the UUID as the response
	c.JSON(http.StatusOK, gin.H{
		"uuid": newUUID,
	})
}

// 去 redis 看这个 username 是否已经验证通过，如果通过就把 uuid 标记通过
func VerifyUsername(c *gin.Context) {
	userName, _ := c.Cookie("username")
	varUUId, _ := c.Cookie("uuid")

	pass, _ := service.CheckUserNameFormRedis(userName)
	if pass {
		service.SaveAuthSessionToRedis(varUUId, 5*consts.SecInMinute)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "fail",
	})
}

// 轮询接口，去 redis 看这个 uuid 也就是现在这个登录的 session 是否已经验证通过
func VerifyQRCode(c *gin.Context) {
	varUUID, _ := c.Cookie("uuid")
	pass, _ := service.CheckAuthSessionFromRedis(varUUID)
	if pass { // 这里就不是很安全 todo
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "fail",
	})
}
