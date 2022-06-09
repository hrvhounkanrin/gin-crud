package insured

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/setondji357/gin-crud/pkg/common/models"
)

type UpdateInsuredDto struct {
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
}

func (h DbHandler) UpdateInsured(c *gin.Context) {
	id := c.Param("id")
	body := UpdateInsuredDto{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var insured models.Insured

	if result := h.DB.First(&insured, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	insured.LastName = body.LastName
	insured.FirstName = body.FirstName
	insured.DateOfBirth = body.DateOfBirth
	insured.PhoneNumber = body.PhoneNumber

	h.DB.Save(&insured)

	c.JSON(http.StatusOK, &insured)
}
