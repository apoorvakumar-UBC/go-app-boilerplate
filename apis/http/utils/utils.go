package utils

import (
	log "go-boilerplate-api/pkg/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/ralstan-vaz/go-errors"
	"github.com/ralstan-vaz/go-errors/http"
)

// HandleError formats, logs and sets a http response for the error
func HandleError(c *gin.Context, errObj *error) {

	if *errObj == nil {
		return
	}

	err := errors.Get(*errObj)

	if err.Message == "" {
		err.Message = "Something Went Wrong"
	}

	log.Error(err.Code, err.Description, log.Priority1, err.Source)

	statusCode := http.StatusCode(err)
	c.JSON(statusCode, gin.H{
		"code":        err.Code,
		"message":     err.Message,
		"description": err.Description,
	})

}
