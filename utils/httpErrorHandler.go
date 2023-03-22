package utils

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpErrorHandler(err error, c *gin.Context) {
	switch err {
	case sql.ErrNoRows:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
	}

}
