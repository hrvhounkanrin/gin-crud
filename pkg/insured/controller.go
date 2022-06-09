package insured

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type DbHandler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &DbHandler{
		DB: db,
	}

	routes := r.Group("/insured")
	routes.POST("/", h.AddInsured)
	routes.GET("/", h.GetInsureds)
	routes.GET("/:id", h.GetInsured)
	routes.PUT("/:id", h.UpdateInsured)
	routes.DELETE("/:id", h.DeleteInsured)
}
