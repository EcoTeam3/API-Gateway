package handler

import (
	"api_gateway/generated/community"
	"api_gateway/generated/user"
)

type Handler struct {
	UserService user.UserServiceClient
	Community community.CommunityServiceClient
}

func NewHandler(UserService user.UserServiceClient) *Handler {
	return &Handler{UserService: UserService}
}