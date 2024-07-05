package handler

import (
 pb "api_gateway/generated/impact"
 "net/http"

 "github.com/gin-gonic/gin"
)

func (h *Handler) CreateFootprint(ctx *gin.Context) {

 req := pb.CarbonFootprint{}

 if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

 resp, err := h.ImpactCalculator.CreateFootprint(ctx, &req)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, err.Error())
  return
 }

 ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserImpact(ctx *gin.Context) {
 id := ctx.Param("id")

 req := pb.UserId{UserId: id}

 resp, err := h.ImpactCalculator.GetUserImpact(ctx, &req)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, err.Error())
  return
 }

 ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetGroupImpact(ctx *gin.Context) {
 id := ctx.Param("id")

 req := pb.GroupId{GroupId: id}

 resp, err := h.ImpactCalculator.GetGroupImpact(ctx, &req)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, err.Error())
  return
 }

 ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetLeaderBoardUsers(ctx *gin.Context) {

    var req pb.LeaderBoard

    resp, err := h.ImpactCalculator.GetLeaderBoardUsers(ctx, &req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, resp)
}


func (h *Handler) GetLeaderBoardGroups(ctx *gin.Context) {

    var req pb.LeaderBoard

    resp, err := h.ImpactCalculator.GetLeaderBoardGroups(ctx, &req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateDonation(ctx *gin.Context) {

 req := pb.Donation{}

 if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

 resp, err := h.ImpactCalculator.CreateDonation(ctx, &req)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, err.Error())
  return
 }

 ctx.JSON(http.StatusOK, resp)
}


func (h *Handler) GetDonations(ctx *gin.Context) {
 cause := ctx.Param("cause")

 req := pb.DonationCause{Cause: cause}
 resp, err := h.ImpactCalculator.GetDonations(ctx, &req)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, err.Error())
  return
 }

 ctx.JSON(http.StatusOK, resp)
}
