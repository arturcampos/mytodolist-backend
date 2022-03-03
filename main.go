package main

import (
	"net/http"
	"strings"

	"todolist/service"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.POST("/tasks", postTask)
	//router.PUT("/tasks/:id", putTask)
	router.Run("localhost:8080")
}

func getTasks(ginCtx *gin.Context) {
	tasks, err := service.GetTask()
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, err)
	}

	ginCtx.JSON(http.StatusOK, tasks)

}

func postTask(ginCtx *gin.Context) {
	var task service.Task

	if contentType := ginCtx.Request.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
		ginCtx.JSON(http.StatusUnsupportedMediaType, nil)
	}

	if ginCtx.ShouldBind(&task) == nil {
		err := service.CreateTask(&task)
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, err)
		}

		ginCtx.JSON(http.StatusOK, nil)
	}

}
