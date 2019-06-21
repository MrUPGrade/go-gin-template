package resources

import (
	models "ginapi/pkg/db/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoController struct {
	ResourceBase
}

func (r *TodoController) Get(c *gin.Context) {
	todos := make([]models.Todo, 0)
	r.DB.Limit(10).Order("id").Find(&todos)

	c.JSON(http.StatusOK, todos)
}

func (r *TodoController) Post(c *gin.Context) {
	todo := new(models.Todo)
	if err := c.ShouldBindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	r.DB.Create(todo)

	c.JSON(http.StatusOK, gin.H{"message": "object created"})
}
