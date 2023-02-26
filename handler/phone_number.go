package handler

import (
	"gin-gorm/common"
	"gin-gorm/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhoneNumberHandler struct {
	DB *gorm.DB
}

func (h *PhoneNumberHandler) HandleInsertPhoneNumber(ctx *gin.Context) {
	var phoneNumber entity.PhoneNumber
	err := ctx.ShouldBind(&phoneNumber)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	tx := h.DB.Create(&phoneNumber)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest,
			common.Response{
				Status:  false,
				Message: tx.Error.Error(),
				Data:    nil,
			})
		return
	}

	ctx.JSON(http.StatusOK,
		common.Response{
			Status:  true,
			Message: "phone number inserted successfully",
			Data:    phoneNumber,
		})
}
