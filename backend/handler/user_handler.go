package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/your_project/domain/user"
	"github.com/your_project/usecase"
)

var userUsecase = usecase.UserUsecase{
	Repo: user.NewRepository(),
}

func GetUsers(c *gin.Context) {
	users, err := userUsecase.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "失敗"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// func CreateUser(c *gin.Context) {
// 	var u user.User
// 	if err := c.ShouldBindJSON(&u); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "不正なデータ"})
// 		return
// 	}
// 	if err := userUsecase.CreateUser(&u); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "作成に失敗"})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, u)
// }
