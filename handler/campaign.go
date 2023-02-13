package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vuegolang/campaign"
	"vuegolang/helper"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user-id"))
	if err != nil {
		response := helper.ApiResponse("Failed to get campaigns", http.StatusUnprocessableEntity, "error", gin.H{"error": err})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaigns, err := h.campaignService.GetCampaigns(userId)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaigns", http.StatusBadRequest, "error", gin.H{"error": err})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success get campaigns", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
	return
}
