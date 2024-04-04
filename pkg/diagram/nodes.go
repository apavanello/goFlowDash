package diagram

import (
	"context"
	"time"

	"github.com/apavanello/goflowdash/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NodeDelete struct {
	Id string `bson:"_id,omitempty" json:"id" binding:"required"`
}

type Node struct {
	Id       string `bson:"_id,omitempty" json:"id" binding:"required"`
	NodeType string `bson:"nodeType" json:"type" binding:"required"`
	Label    string `bson:"label" json:"label" binding:"required"`
	Position struct {
		X float64 `bson:"x" json:"x"`
		Y float64 `bson:"y" json:"y"`
	} `bson:"position" json:"position"`
	Data struct {
		Status          string `bson:"status" json:"status"`
		Description     string `bson:"description" json:"description"`
		Repo            string `bson:"repo" json:"repo"`
		StartTime       string `bson:"startTime" json:"startTime"`
		EndTime         string `bson:"endTime" json:"endTime"`
		PlanedStartTime string `bson:"planedStartTime" json:"planedStartTime"`
		PlanedEndTime   string `bson:"planedEndTime" json:"planedEndTime"`
		Squad           string `bson:"squad" json:"squad"`
	} `bson:"data" json:"data"`
}

type NodeStatus struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func (n *Node) List(c *mongo.Client) ([]Node, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(c, "nodes")

	defer cancel()

	cursor, err := col.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var nodes []Node
	for cursor.Next(ctx) {
		err := cursor.Decode(&n)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, *n)
	}

	return nodes, nil
}

func (n *Node) New(c *mongo.Client) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(c, "nodes")

	res, err := col.InsertOne(ctx, n)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (n *Node) SavePos(client *mongo.Client) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(client, "nodes")

	_, err := col.UpdateOne(
		ctx,
		bson.M{"_id": n.Id},
		bson.D{
			{"$set", bson.D{
				{Key: "position.x", Value: n.Position.X},
				{Key: "position.y", Value: n.Position.Y},
			}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (nd *NodeDelete) Delete(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(client, "nodes")

	_, err := col.DeleteOne(ctx, bson.M{"_id": nd.Id})
	if err != nil {
		return err
	}

	return nil

}

func (ns *NodeStatus) UpdateStatus(c *mongo.Client) (*mongo.UpdateResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(c, "nodes")

	res, err := col.UpdateOne(
		ctx,
		bson.M{"_id": ns.Id},
		bson.D{
			{"$set", bson.D{{Key: "data.status", Value: ns.Status}}},
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (n *Node) Update(c *mongo.Client) (*mongo.UpdateResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(c, "nodes")

	res, err := col.ReplaceOne(ctx, bson.D{{"_id", n.Id}}, n)
	if err != nil {
		return nil, err
	}

	return res, nil
}
