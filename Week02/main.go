package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func init() {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		"root",
		"root",
		"localhost",
		"mydb",
		true,
		"Local")

	var err error
	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		panic("failed to init db")
	}

}

var db *gorm.DB

// dao 层
type UserModel struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func getUserInRepository(uid uint) (*UserModel, error) {
	var user = UserModel{
		Id: uid,
	}
	err := db.First(&user).Error
	if err != nil {
		return nil, errors.Wrapf(err, "[user repo] get user info")
	}
	return &user, err

}

// service 层
func getUserInService(uid uint) (*UserModel, error) {
	//因为service层没有其他逻辑，并不会生成错误，无需wrap，所以直接返回
	return getUserInRepository(uid)

	/*
	 如果service 层需要判断user的信息，可以这么写:

	 user,err := getUserInRepository(uid)
	 if err != nil{
	 	return nil,errors.Wrapf(err,"failed to get user in db")
	 }
	 if len(user.Name) > 20 {
	 	return nil,errors.Wrapf(err,"invalid user name length ")
	 }
	 return user,err
	*/

}

// controller 层
type Req struct {
	Uid uint `form:"uid"`
}

func getUser(c *gin.Context) {
	var req Req
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    1,
			"message": "bind parameter error",
		})
		return
	}

	user, err := getUserInService(req.Uid)
	if err != nil {
		// 日志放在controller层打
		log.Printf("failed to get user: %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": err.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    user,
	})
}

func main() {

	server := gin.Default()
	server.GET("/api/getUser", getUser)
	server.Run("localhost:8080")
}
