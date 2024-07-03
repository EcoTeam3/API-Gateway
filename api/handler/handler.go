package handler

import "api_gateway/generated/user"

type Handler struct {
	UserService user.UserServiceClient
}

func NewHandler(UserService user.UserServiceClient) *Handler {
	return &Handler{UserService: UserService}
}