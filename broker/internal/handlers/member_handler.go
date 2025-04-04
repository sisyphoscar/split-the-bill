package handlers

import (
	"broker/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	MemberService services.MemberService
}

type MemberRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func NewMemberHandler(memberService *services.MemberService) *MemberHandler {
	return &MemberHandler{
		MemberService: *memberService,
	}
}

func (h *MemberHandler) CreateMember(c *gin.Context) {
	var req MemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	log.Printf("Creating member with name: %s, email: %s", req.Name, req.Email)

	member, err := h.MemberService.CreateMember(req.Name, req.Email)
	if err != nil {
		log.Printf("Failed to create member: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create member",
		})
		return
	}

	log.Printf("Member created: %v", member)

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data": gin.H{
			"id":    member.Id,
			"name":  member.Name,
			"email": member.Email,
		},
	})
}
