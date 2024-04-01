package main

import (
	"context"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/apavanello/goflowdash/assets"
	"github.com/apavanello/goflowdash/pkg/diagram"
	"github.com/apavanello/goflowdash/pkg/mongodb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongodbClient *mongo.Client
	ctx           context.Context
	err           error
)

func load() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctx = context.Background()

	mongodbClient, err = mongodb.Connect(ctx)

	if err != nil {
		panic(err)
	}

}

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

		if c.GetHeader("Update-Type") == "edge-flow" {

			var edgeFlow diagram.EdgeFlow
			var status *mongo.UpdateResult
			var err error

			err = c.ShouldBindJSON(&edgeFlow)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			if c.GetHeader("Flow-Type") == "serial" || c.GetHeader("Flow-Type") == "parallel" {
				status, err = edgeFlow.Update(mongodbClient, c.GetHeader("Flow-Type"))
				if err != nil {
					log.Default().Println(err)
					c.JSON(400, err)
				}
			} else {
				c.JSON(400, "Flow-Type not supported")
			}

			c.JSON(200, gin.H{"result": status.ModifiedCount})
		}
	})

	r.POST("/api/:object/delete", func(c *gin.Context) {

		switch c.Param("object") {
		case "node":

			var nodeDelete diagram.NodeDelete
			var edge diagram.Edge

			err := c.ShouldBindJSON(&nodeDelete)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			err = nodeDelete.Delete(mongodbClient)
			if err != nil {
				c.JSON(500, err)
			}
			edges, err := edge.GetEdgesBySourceOrTarget(mongodbClient, nodeDelete.Id)
			if err != nil {
				c.JSON(500, err)
			}

			for _, edge := range edges {

				var edgeDelete diagram.EdgeDelete
				edgeDelete.Id = edge.Id
				err = edgeDelete.Delete(mongodbClient)
				if err != nil {
					c.JSON(500, err)
				}

			}
			c.JSON(200, gin.H{"result": "ok"})

		case "edge":
			var edgeDelete diagram.EdgeDelete
			err := c.ShouldBindJSON(&edgeDelete)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			err = edgeDelete.Delete(mongodbClient)
			if err != nil {
				c.JSON(500, err)
			}
			c.JSON(200, gin.H{"result": "ok"})
		default:
			c.JSON(404, gin.H{"error": "not found"})
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

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"result": "ok"})
	})

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}

}
