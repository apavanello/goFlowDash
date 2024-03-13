package diagram

import (
	"context"
	"github.com/apavanello/goflowdash/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Node struct {
	Id       string `bson:"_id,omitempty" json:"id" binding:"required"`
	NodeType string `bson:"nodeType" json:"type" binding:"required"`
	Label    string `bson:"label" json:"label" binding:"required"`
	Position struct {
		X float64 `bson:"x" json:"x"`
		Y float64 `bson:"y" json:"y"`
	} `bson:"position" json:"position"`
	Extras struct {
		Status      string `bson:"status" json:"status"`
		Description string `bson:"description" json:"description"`
		Repo        string `bson:"repo" json:"repo"`
	} `bson:"extras" json:"extras"`
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

	log.Default().Println(cursor)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var nodes []Node

	m := make(map[int]int)

	for cursor.Next(ctx) {
		err := cursor.Decode(&n)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, *n)
	}

	log.Default().Println(m)

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
			{"$set", bson.D{{"position.x", n.Position.X}, {"position.y", n.Position.Y}}},
		},
	)
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
			{"$set", bson.D{{"extras.status", ns.Status}}},
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
