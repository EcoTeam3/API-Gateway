package handler

import (
	"api_gateway/generated/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	u := &user.User{}
	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	resp, err := h.UserService.CheckUser(c, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if !resp.Status{
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}

	res, err := h.UserService.CreateUser(c, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	if res.Status {
		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
	}
}
