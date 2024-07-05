package handler

import (
	"api_gateway/generated/community"
	"api_gateway/generated/habit"
	"api_gateway/generated/impact"
	"api_gateway/generated/user"
)

type Handler struct {
	UserService user.UserServiceClient
	Community community.CommunityServiceClient
	HabitTracker habit.HabitTrackerClient
	ImpactCalculator impact.ImpactClient
}