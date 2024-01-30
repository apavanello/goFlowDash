package main

import (
	"context"
	"github.com/apavanello/goflowdash/pkg/diagram"
	"github.com/apavanello/goflowdash/pkg/mongodb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

var (
	mongodbClient *mongo.Client
	ctx           context.Context
	err           error
)

func main() {

	load()
	defer mongodbClient.Disconnect(ctx)
	//mongodb.Ping()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET, POST"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/:object/list", func(c *gin.Context) {

		switch c.Param("object") {
		case "box":
			var box diagram.Box
			boxes, err := box.List(mongodbClient)
			log.Default().Println(boxes)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, boxes)
		case "edge":
			var edge diagram.Edge
			boxes, err := edge.List(mongodbClient)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, boxes)
		default:
			c.JSON(404, gin.H{"error": "not found"})
		}
	})

	r.POST("/api/:object/new", func(c *gin.Context) {

		switch c.Param("object") {
		case "box":
			var box diagram.Box
			err := c.ShouldBindJSON(&box)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			log.Println(box)
			res, err := box.New(mongodbClient)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, gin.H{"result": res.InsertedID})

		case "edge":
			var edge diagram.Edge
			err := c.ShouldBindJSON(&edge)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			log.Println(edge)
			res, err := edge.New(mongodbClient)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, gin.H{"result": res.InsertedID})

		}
	})

	r.POST("/api/:object/update", func(c *gin.Context) {

		if c.GetHeader("Update-Type") == "status" {

			var boxStatus diagram.BoxStatus

			err := c.ShouldBindJSON(&boxStatus)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			status, err := boxStatus.UpdateStatus(mongodbClient)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, gin.H{"result": status.ModifiedCount})

		}
	})

	err := r.Run(":8082")
	if err != nil {
		panic(err)
	}

}

func load() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctx = context.Background()

	mongodbClient, err = mongodb.Connect(ctx)

	if err != nil {
		panic(err)
	}

}
