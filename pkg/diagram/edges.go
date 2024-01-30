package diagram

import (
	"context"
	"github.com/apavanello/goflowdash/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Edge struct {
	Id     string `bson:"_id,omitempty" json:"id" binding:"required"`
	Source string `bson:"source" json:"source" binding:"required"`
	Target string `bson:"target" json:"target" binding:"required"`
}

func (e *Edge) List(c *mongo.Client) ([]Edge, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(c, "edges")

	cursor, err := col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var edges []Edge
	for cursor.Next(ctx) {
		err := cursor.Decode(&e)
		if err != nil {
			return nil, err
		}

		edges = append(edges, *e)
	}

	log.Default().Println(edges)

	return edges, nil
}

func (e *Edge) New(c *mongo.Client) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(c, "edges")

	res, err := col.InsertOne(ctx, e)
	if err != nil {
		return nil, err
	}

	return res, nil
}
