package main

import (
	"github.com/D1Y0RBEKORIFJONOV/rest-api-project/internal/user"

	"github.com/gin-gonic/gin"
	_ "github.com/julienschmidt/httprouter"
)

func main() {
	users, err := user.ReadUser()
	if err != nil {
		panic(err)
	}
	router := gin.Default()

	handler := user.NewHandler()

	handler.Register(router, users)
	router.Run()
}
