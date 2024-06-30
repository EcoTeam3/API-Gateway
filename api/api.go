package api

import (
	"api_gateway/api/handler"
	"api_gateway/genproto/userService"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewRouter(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	userService := userService.NewUserServiceClient(conn)

	handler := handler.NewHandler(userService)

	user := router.Group("/user")
	user.GET("/get/:id", handler.GetUser)
	user.PUT("/update/:id", handler.UpdateUser)
	user.DELETE("/delete/:id", handler.GetUser)



	userprofile := router.Group("/userprofile")
	userprofile.GET("/get/:id", handler.GetUserProfile)
	userprofile.PUT("/update/:id", handler.UpdateUserProfile)


	return router
}