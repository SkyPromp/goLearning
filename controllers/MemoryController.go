package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/SkyPromp/goLearning/services"
	"github.com/SkyPromp/goLearning/models"
)

// GetMemoryManagementExample returns memory usage statistics.
// @Summary Get memory usage
// @Description Returns memory usage and GC statistics. Optional query parameter "goroutine" determines safe (GC-tracked) or unsafe (non-GC) allocation.
// @Tags memory
// @Param goroutine query bool false "Set to false to use unsafe memory allocation outside GC"
// @Produce json
// @Success 200 {object} models.MemStats
// @Router /memory [get]
func GetMemoryManagementExample(context *gin.Context){
	isSafe, err := strconv.ParseBool(context.DefaultQuery("goroutine", "true"))

	if(err != nil){
		isSafe = true
	}

	var data models.MemStats

	if(isSafe){
		data = services.MemoryManagement()
	} else{
		data = services.UnsafeMemoryManagement()
	}

	context.IndentedJSON(http.StatusOK, data)
}

// GetByteAlignment godoc
// @Summary Get byte alignment information
// @Description Returns byte alignment data from the service layer
// @Tags memory
// @Produce json
// @Success 200  {object}  models.StructSizes
// @Router /alignment [get]
func GetByteAlignment(context *gin.Context){
	data := services.GetByteAlignment()

	context.IndentedJSON(http.StatusOK, data)
}


// GetUnsafeExample godoc
// @Summary Byte manipulation
// @Description Returns an integer that has been manipulated in an unsafe way
// @Tags memory
// @Produce json
// @Success 200  {object}  int32
// @Router /unsafe [get]
func GetUnsafeExample(context *gin.Context){
	data := services.GetUnsafeExample()

	context.IndentedJSON(http.StatusOK, data)
}
