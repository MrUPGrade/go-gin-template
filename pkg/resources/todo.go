package resources

import (
	"ginapi/pkg/db/models"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	ResourceBase
}

func (r *TodoController) Get(c *gin.Context) {
	todos := make([]pkg.Todo, 0)
	r.DB.Limit(10).Order("id").Find(&todos)

	c.JSON(200, todos)
}
