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
	userId, _ := strconv.Atoi(c.Query("user-id"))
	campaigns, err := h.campaignService.GetCampaigns(userId)

	if err != nil {
		response := helper.ApiResponse("Failed to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// just experiment to implement function as parameter
	formattedCampaigns := campaign.FormatCampaignCollection(campaigns, campaign.FormatSingleCampaign)

	response := helper.ApiResponse("Success get campaigns", http.StatusOK, "success", formattedCampaigns)
	c.JSON(http.StatusOK, response)
	return
}
