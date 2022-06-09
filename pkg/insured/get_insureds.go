package insured

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/setondji357/gin-crud/pkg/common/models"
)

func (h DbHandler) GetInsureds(c *gin.Context) {
	var insureds []models.Insured

	if result := h.DB.Find(&insureds); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &insureds)
}
