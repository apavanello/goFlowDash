package diagram

import (
	"context"
	"github.com/apavanello/goflowdash/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Edges []Edge

type Edge struct {
	Id           string `bson:"_id,omitempty" json:"id" binding:"required"`
	Type         string `bson:"type" json:"type" binding:"required"`
	Source       string `bson:"source" json:"source" binding:"required"`
	Target       string `bson:"target" json:"target" binding:"required"`
	SourceHandle string `bson:"sourceHandle" json:"sourceHandle" binding:"required"`
	TargetHandle string `bson:"targetHandle" json:"targetHandle" binding:"required"`
	Data         struct {
	} `json:"data" optional:"true"`
	Label     string `bson:"label" json:"label" optional:"true"`
	MarkerEnd string `bson:"markerEnd" json:"markerEnd"`
}

type EdgeFlow struct {
	Id           string `bson:"_id,omitempty" json:"id" binding:"required"`
	Source       string `bson:"source" json:"source" binding:"required"`
	Target       string `bson:"target" json:"target" binding:"required"`
	SourceHandle string `bson:"sourceHandle" json:"sourceHandle" binding:"required"`
	TargetHandle string `bson:"targetHandle" json:"targetHandle" binding:"required"`
	MarkerEnd    string `bson:"markerEnd" json:"markerEnd"`
}

type EdgeDelete struct {
	Id string `bson:"_id,omitempty" json:"id" binding:"required"`
}

//markerEnd: MarkerType.ArrowClosed

func (ef *EdgeFlow) Update(client *mongo.Client, flowType string) (*mongo.UpdateResult, error) {

	if flowType == "serial" {
		oldTarget := ef.Target
		oldTargetHandle := ef.TargetHandle

		ef.Target = ef.Source
		ef.TargetHandle = ef.SourceHandle
		ef.Source = oldTarget
		ef.SourceHandle = oldTargetHandle
		ef.MarkerEnd = "arrowclosed"

	} else if flowType == "parallel" {
		ef.MarkerEnd = "null"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(client, "edges")

	result, err := col.UpdateOne(ctx, bson.D{{"_id", ef.Id}}, bson.D{{"$set", ef}})
	if err != nil {
		return nil, err
	}

	return result, nil
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

func (e *Edge) Save(client *mongo.Client) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(client, "edges")
	res, err := col.ReplaceOne(ctx, bson.D{{"_id", e.Id}}, e)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		res, err := col.InsertOne(ctx, e)
		if err != nil {
			return err
		}
		log.Default().Println(res)
	}

	return nil
}

func (ed *EdgeDelete) Delete(client *mongo.Client) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(client, "edges")
	_, err := col.DeleteOne(ctx, bson.D{{"_id", ed.Id}})
	if err != nil {
		return err
	}

	return nil

}

func (e *Edge) GetEdgesBySourceOrTarget(c *mongo.Client, id string) (edges Edges, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(c, "edges")

	cursor, err := col.Find(ctx, bson.D{{"$or", bson.A{bson.D{{"source", id}}, bson.D{{"target", id}}}}})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		err := cursor.Decode(&e)
		if err != nil {
			return nil, err
		}

		edges = append(edges, *e)

	}

	log.Default().Println(edges)

	return
}
