package diagram

import (
	"context"
	"github.com/apavanello/goflowdash/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type Box struct {
	Id       string `bson:"_id,omitempty" json:"id" binding:"required"`
	BoxType  string `bson:"boxType" json:"type" binding:"required"`
	Label    string `bson:"label" json:"label" binding:"required"`
	Position struct {
		X int `bson:"x" json:"x"`
		Y int `bson:"y" json:"y"`
	} `bson:"position" json:"position"`
	Extras struct {
		Status      string `bson:"status" json:"status"`
		Description string `bson:"description" json:"description"`
		Repo        string `bson:"repo" json:"repo"`
	} `bson:"extras" json:"extras"`
}

type BoxStatus struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func (b *Box) List(c *mongo.Client) ([]Box, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(c, "boxes")

	defer cancel()

	cursor, err := col.Find(ctx, bson.D{})

	log.Default().Println(cursor)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var boxes []Box

	m := make(map[int]int)

	for cursor.Next(ctx) {
		err := cursor.Decode(&b)
		if err != nil {
			return nil, err
		}
		boxes = append(boxes, *b)
	}

	log.Default().Println(m)

	return boxes, nil
}

func (b *Box) New(c *mongo.Client) (*mongo.InsertOneResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(c, "boxes")

	res, err := col.InsertOne(ctx, b)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *Box) SavePos(client *mongo.Client) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(client, "boxes")

	_, err := col.UpdateOne(
		ctx,
		bson.M{"_id": b.Id},
		bson.D{
			{"$set", bson.D{{"position.x", b.Position.X}, {"position.y", b.Position.Y}}},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (bs *BoxStatus) UpdateStatus(c *mongo.Client) (*mongo.UpdateResult, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := mongodb.GetCollection(c, "boxes")

	res, err := col.UpdateOne(
		ctx,
		bson.M{"_id": bs.Id},
		bson.D{
			{"$set", bson.D{{"extras.status", bs.Status}}},
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
