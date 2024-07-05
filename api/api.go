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
	// Register and Login endpoints
	router.POST("/register", st.Register)
	router.POST("/login", st.Login)

	// Protected routes with JWT middleware
	user := router.Group("/user").Use(st.ValidateJWT())
	user.GET("/get/:id", st.GetUser)
	user.PUT("/update/:id", st.UpdateUser)
	user.DELETE("/delete/:id", st.GetUser)

	userprofile := router.Group("/userprofile").Use(st.ValidateJWT())
	userprofile.GET("/get/:id", st.GetUserProfile)
	userprofile.PUT("/update/:id", st.UpdateUserProfile)

	community := router.Group("/community").Use(st.ValidateJWT())
	community.POST("/create", st.CreateGroup)
	community.GET("/get/:groupId", st.GetGroup)
	community.PUT("/update/:groupId", st.UpdateGroup)
	community.DELETE("/delete/:groupId", st.DeleteGroup)
	community.GET("/getAll", st.GetAllGroups)
	community.POST("/joinGroup/:groupId/:userId", st.JoinGroupUser)
	community.PUT("/leaveGroup/:groupId/:userId", st.LeaveGroupUser)
	community.PUT("/updateRole", st.UpdateGroupMember)
	community.POST("/createPost", st.CreatePost)
	community.PUT("/updatePost/:postId", st.UpdatePost)
	community.GET("/getPost/:postId", st.GetPost)
	community.DELETE("/deletePost/:postId", st.DeletePost)
	community.GET("/getGroupPost/:groupId/:postId", st.GetGroupPost)
	community.POST("/createPostComment", st.CreatePostComment)
	community.GET("/getPostComment/:postId/:commentId", st.GetPostComment)

	impact := router.Group("/impactCalculator").Use(st.ValidateJWT())
	impact.POST("/create", st.CreateFootprint)
	impact.GET("/get/userImpact/:id", st.GetUserImpact)
	impact.GET("/get/groupImpact/:id", st.GetGroupImpact)
	impact.GET("/getUsers", st.GetLeaderBoardUsers)
	impact.GET("/getGroups", st.GetLeaderBoardGroups)
	impact.POST("/createDonation", st.CreateDonation)
	impact.GET("/getDonation/:cause", st.GetDonations)

	habit := router.Group("/habit").Use(st.ValidateJWT())
	habit.POST("/create", st.CreateHabit)
	habit.GET("/get/habit/:id", st.GetHabit)
	habit.PUT("/update", st.UpdateHabit)
	habit.DELETE("/delete/:id", st.DeleteHabit)
	habit.GET("/get/userhabits/:id", st.GetUserHabits)
	habit.POST("/create/habitLog", st.CreateHabitLog)
	habit.GET("/get/suggestion", st.GetHabitSuggestions)

	return router
}
