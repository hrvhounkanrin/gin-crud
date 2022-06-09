package insured

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/setondji357/gin-crud/pkg/common/models"
)

func (h DbHandler) DeleteInsured(c *gin.Context) {
	id := c.Param("id")

	var insured models.Insured

	if result := h.DB.First(&insured, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&insured)
	c.Status(http.StatusOK)
}
