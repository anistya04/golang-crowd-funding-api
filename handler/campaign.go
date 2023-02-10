package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"vuegolang/campaign"
	"vuegolang/helper"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) FindAll(c *gin.Context) {

	campaigns, err := h.campaignService.GetAll()
	fmt.Println(err)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaigns", http.StatusBadRequest, "error", gin.H{"error": err})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success to upload avatar", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)
	return
}
