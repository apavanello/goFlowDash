package main

import (
	"context"
	"github.com/apavanello/goflowdash/assets"
	"github.com/apavanello/goflowdash/pkg/diagram"
	"github.com/apavanello/goflowdash/pkg/mongodb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"io/fs"
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

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET, POST"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	distFS, err := fs.Sub(assets.Dist, "dist")
	if err != nil {
		panic(err)
	}

	assetsFS, err := fs.Sub(assets.Dist, "dist/assets")
	if err != nil {
		panic(err)
	}

	r.StaticFS("/assets", http.FS(assetsFS))

	r.Any("/", func(c *gin.Context) {
		c.FileFromFS("./", http.FS(distFS))
	})

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/:object/list", func(c *gin.Context) {

		switch c.Param("object") {
		case "node":
			var node diagram.Node
			nodes, err := node.List(mongodbClient)
			log.Default().Println(nodes)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, nodes)
		case "edge":
			var edge diagram.Edge
			nodes, err := edge.List(mongodbClient)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, nodes)
		default:
			c.JSON(404, gin.H{"error": "not found"})
		}
	})

	r.POST("/api/:object/new", func(c *gin.Context) {

		switch c.Param("object") {
		case "node":
			var node diagram.Node
			err := c.ShouldBindJSON(&node)

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			log.Println(node)
			res, err := node.New(mongodbClient)
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

			var nodeStatus diagram.NodeStatus

			err := c.ShouldBindJSON(&nodeStatus)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			status, err := nodeStatus.UpdateStatus(mongodbClient)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, gin.H{"result": status.ModifiedCount})

		}
	})

	r.POST("/api/save", func(c *gin.Context) {

		var panel diagram.Panel

		err := c.ShouldBindJSON(&panel)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		err = panel.Save(mongodbClient)
		if err != nil {
			c.JSON(500, err)
		}
		c.JSON(200, gin.H{"result": "ok"})
	})

	err = r.Run(":8082")
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
