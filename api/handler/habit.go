package handler

import (
	pb "api_gateway/generated/habit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateHabit(ctx *gin.Context) {

	habit := &pb.Habit{}

	if err := ctx.ShouldBindJSON(habit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	resp, err := h.HabitTracker.CreateHabit(ctx, habit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetHabit(ctx *gin.Context) {
	habitId := ctx.Param("id")

	req := &pb.HabitId{HabitId: habitId}

	resp, err := h.HabitTracker.GetHabit(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateHabit(ctx *gin.Context) {
	id := ctx.Param("id")

	var req pb.Habit
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if req.HabitId == "" {
		req.HabitId = id
	}

	resp, err := h.HabitTracker.UpdateHabit(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteHabit(ctx *gin.Context) {
	id := ctx.Param("id")

	req := pb.HabitId{HabitId: id}

	resp, err := h.HabitTracker.DeleteHabit(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserHabits(ctx *gin.Context) {
	id := ctx.Param("id")

	req := pb.UserId{UserId: id}
	resp, err := h.HabitTracker.GetUserHabits(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateHabitLog(ctx *gin.Context) {

	req := pb.HabitLog{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	resp, err := h.HabitTracker.CreateHabitLog(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetHabitLogs(ctx *gin.Context) {
	id := ctx.Param("id")

	req := pb.HabitId{HabitId: id}

	resp, err := h.HabitTracker.GetHabitLogs(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetHabitSuggestions(ctx *gin.Context) {

	role := pb.Req{}

	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.HabitTracker.GetHabitSuggestions(ctx, &role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
