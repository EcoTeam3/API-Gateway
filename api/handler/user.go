package handler

import (
	pb "api_gateway/generated/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(ctx *gin.Context) {
    
	id := ctx.Param("id")

	req := pb.UserId{UserId: id}

	resp, err := h.UserService.GetUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
    id := ctx.Param("id")

    var req pb.User
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if req.UserId == "" {
        req.UserId = id
    }

    resp, err := h.UserService.UpdateUser(ctx, &req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	req := pb.UserId{UserId: id}

	resp, err := h.UserService.DeleteUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	req := pb.UserId{UserId: id}

	resp, err := h.UserService.GetUserProfile(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateUserProfile(ctx *gin.Context) {
    id := ctx.Param("id")

    var req pb.UserProfile
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if req.UserId == "" {
        req.UserId = id
    }

    resp, err := h.UserService.UpdateUserProfile(ctx, &req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, resp)
}

