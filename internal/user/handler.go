package user

import (
	"github.com/D1Y0RBEKORIFJONOV/rest-api-project/internal/handlers"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

const (
	UsersUrl = "/users"
	UserUrl  = "/users/:id"
)

func NewHandler() handlers.Handlers {
	return &handler{}
}

func (h *handler) Register(router *gin.Engine, user interface{}) {
	router.GET(UsersUrl, user.(Users).GetUsers)
	router.GET(UserUrl, user.(Users).GetUserById)
	router.POST(UsersUrl, user.(Users).CreateUser)
}

func (users Users) CreateUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}
	err = CreateUser(&users, &user)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, &user)
}

func (users Users) GetUsers(c *gin.Context) {
	if len(users.Users) == 0 {
		c.IndentedJSON(http.StatusNoContent, gin.H{
			"Message": "empty",
		})
	}
	c.IndentedJSON(http.StatusOK, &users)
}
func (users Users) GetUserById(c *gin.Context) {
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].ID == id_int {
			c.IndentedJSON(http.StatusOK, &users.Users[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"Message": "not found!",
	})
}
