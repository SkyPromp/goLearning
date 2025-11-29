package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func RerouteToSwagger(context *gin.Context){
	context.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
}
