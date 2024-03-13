package diagram

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Panel struct {
	Boxes []Node `json:"nodes" required:"true" binding:"required"`
	Edges []Edge `json:"edges" required:"true" binding:"required"`
}

func (p Panel) Save(client *mongo.Client) error {
	for _, box := range p.Boxes {
		err := box.SavePos(client)
		if err != nil {
			return err
		}
	}
	for _, edge := range p.Edges {
		err := edge.Save(client)
		if err != nil {
			return err
		}
	}
	return nil
}
