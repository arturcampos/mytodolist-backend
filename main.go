package main

import (
	"net/http"
	"os"
	"strings"

	"todolist/service"

	"github.com/gin-gonic/gin"
)

func init() {
	os.Setenv("TZ", "America/Sao_Paulo")
}
func main() {

	router := gin.Default()
	router.GET("/tasks", getTasks)
	router.POST("/tasks", postTask)
	router.PUT("/tasks/:id", putTask)
	router.PUT("/tasks/:id/complete", putTask)
	router.PUT("/tasks/:id/uncomplete", putTask)
	router.DELETE("/tasks/:id", deleteTask)
	router.Run(":8080")
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

	if err := ginCtx.ShouldBindJSON(&task); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service.CreateTask(&task)
}

func putTask(ginCtx *gin.Context) {
	var task *service.Task
	var err error

	if contentType := ginCtx.Request.Header.Get("Content-Type"); !strings.Contains(contentType, "application/json") {
		ginCtx.JSON(http.StatusUnsupportedMediaType, nil)
		return
	}

	if err := ginCtx.ShouldBindJSON(&task); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if ginCtx.FullPath() == "/tasks/:id" {
		task, err = service.UpdateTaskText(ginCtx.Param("id"), task.Text)

	} else if ginCtx.FullPath() == "/tasks/:id/complete" {
		task, err = service.CompleteTask(ginCtx.Param("id"))

	} else {
		task, err = service.UncompleteTask(ginCtx.Param("id"))

	}

	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, task)

}

func deleteTask(ginCtx *gin.Context) {
	err := service.DeleteTask(ginCtx.Param("id"))
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, err)
	}

}
