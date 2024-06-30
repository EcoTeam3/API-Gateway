package handler

import "api_gateway/genproto/userService"

type Handler struct {
	UserService userService.UserServiceClient
}

func NewHandler(UserService userService.UserServiceClient) *Handler {
	return &Handler{UserService: UserService}
}