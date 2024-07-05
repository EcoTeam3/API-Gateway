package api

import (
	"api_gateway/api/handler"
	"api_gateway/generated/community"
	"api_gateway/generated/habit"
	"api_gateway/generated/impact"
	"api_gateway/generated/user"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewRouter(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	userService := user.NewUserServiceClient(conn)
	communityService := community.NewCommunityServiceClient(conn)
	habitTracker := habit.NewHabitTrackerClient(conn)
	ImpactCalculator := impact.NewImpactClient(conn)

	st := handler.Handler{UserService: userService, Community: communityService, HabitTracker: habitTracker, ImpactCalculator: ImpactCalculator}
	
	user := router.Group("/user")
	user.GET("/get/:id", st.GetUser)
	user.PUT("/update/:id", st.UpdateUser)
	user.DELETE("/delete/:id", st.GetUser)

	userprofile := router.Group("/userprofile")
	userprofile.GET("/get/:id", st.GetUserProfile)
	userprofile.PUT("/update/:id", st.UpdateUserProfile)

	community := router.Group("/community")
	community.POST("/create", st.CreateGroup)
	community.GET("/get/:id", st.GetGroup)
	community.PUT("/update/:id", st.UpdateGroup)
	community.DELETE("/delete/:id", st.DeleteGroup)
	community.GET("/get", st.GetAllGroups)
	community.POST("/")

	impact := router.Group("/impactCalculator")
	impact.POST("/create", st.CreateFootprint)
	impact.GET("/get/userImpact/:id", st.GetUserImpact)
	impact.GET("/get/groupImpact/:id", st.GetGroupImpact)
	impact.GET("/getUsers", st.GetLeaderBoardUsers)
	impact.GET("/getGroups", st.GetLeaderBoardGroups)
	impact.POST("/createDonation", st.CreateDonation)
	impact.GET("/getDonation/:cause", st.GetDonations)

	habit := router.Group("/habit")
	habit.POST("/create", st.CreateHabit)
	habit.GET("/get/habit/:id", st.GetHabit)
	habit.PUT("/update", st.UpdateHabit)
	habit.DELETE("/delete/:id", st.DeleteHabit)
	habit.GET("/get/userhabits/:id", st.GetUserHabits)
	habit.POST("/create/habitLog", st.CreateHabitLog)
	habit.GET("/get/suggestion", st.GetHabitSuggestions)

	return router
}
