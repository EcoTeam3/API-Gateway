package handler

import (
	"api_gateway/generated/community"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGroup(c *gin.Context) {
	group := &community.Group{}
	c.ShouldBindJSON(group)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := h.Community.CreateGroup(ctx, group)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetGroup(c *gin.Context) {
	groupId := c.Param("groupId")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := h.Community.GetGroup(ctx, &community.GroupId{GroupId: groupId})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateGroup(c *gin.Context) {
	groupId := c.Param("groupId")
	group := &community.Group{GroupId: groupId}
	err := c.ShouldBindJSON(group)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	ctx, canccel := context.WithTimeout(context.Background(), time.Second)
	defer canccel()
	resp, err := h.Community.UpdateGroup(ctx, group)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteGroup(c *gin.Context) {
	groupId := c.Param("groupId")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := h.Community.DeleteGroup(ctx, &community.GroupId{GroupId: groupId})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllGroups(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := h.Community.GetAllGroups(ctx, &community.Req{})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) JoinGroupUser(c *gin.Context) {
	groupId := c.Query("groupId")
	userId := c.Param("userId")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := h.Community.JoinGroupUser(ctx, &community.JoinLeave{GroupId: groupId, UserId: userId})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) LeaveGroupUser(c *gin.Context) {
	groupId := c.Query("groupId")
	userId := c.Param("userId")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := h.Community.LeaveGroupUser(ctx, &community.JoinLeave{GroupId: groupId, UserId: userId})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateGroupMember(c *gin.Context) {
	userRole := &community.UserRole{}
	err := c.ShouldBindJSON(userRole)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := h.Community.UpdateGroupMeber(ctx, userRole)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *Handler) CreatePost(c *gin.Context){
	post := &community.Post{}
	err := c.ShouldBindJSON(post)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := h.Community.CreatePost(ctx, post)
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
