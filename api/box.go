package api

import (
	"box-service/db"
	"box-service/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func createBox(c *gin.Context) {
	var box model.Box
	if err := c.BindJSON(&box); err != nil {
		log.Warnf("/createBox can't parse request body Error :  %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"data":    err.Error(),
		})

		return
	}

	if err := db.Client.Create(box); err != nil {
		log.Warnf("/createBox can't same in mongoDB Error :  %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't save in DB",
			"data":    err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Box created successfully",
	})
}

func findBox(c *gin.Context) {
	capacity, err := strconv.ParseInt(c.Query("capacity"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid capacity",
			"data":    err.Error(),
		})

		return
	}

	maxWeight, err := strconv.ParseInt(c.Query("max_weight"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid max weight",
			"data":    err.Error(),
		})

		return
	}

	box, err := db.Client.FindAvailable(int32(capacity), int32(maxWeight))
	if err != nil {
		log.Warnf("/findBox can't find box Error :  %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can't find available box",
			"data":    err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, box)
}
