package test

import (
	"GO1/global"
	"GO1/models"
	"GO1/pkg/snowflake"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(router *gin.RouterGroup) {
	testDB()
	testRouter()
	testLogger()
	testSnowflake()
}

func testDB() {
	var users []models.User
	global.DB.Find(&users)
	fmt.Println(users)
}

func testRouter() {
}

func testLogger() {
	global.Logger.Info(fmt.Sprintf("log info:%d", 1))
}

func testSnowflake() {
	id := snowflake.Snowflake{}.GenID()
	fmt.Println(id)
}
