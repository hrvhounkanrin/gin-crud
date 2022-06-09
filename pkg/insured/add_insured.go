package insured

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/setondji357/gin-crud/pkg/common/models"
)

type CreateInsuredDto struct {
	ID          uint   `gorm:"primaryKey, autoIncrement"`
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
}

func (h DbHandler) AddInsured(c *gin.Context) {
	body := CreateInsuredDto{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var insured models.Insured

	insured.LastName = body.LastName
	insured.FirstName = body.FirstName
	insured.DateOfBirth = body.DateOfBirth
	insured.PhoneNumber = body.PhoneNumber

	if result := h.DB.Create(&insured); result.Error != nil {
		fmt.Printf("An error occured: %v\n", result.Error)
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &insured)
}
