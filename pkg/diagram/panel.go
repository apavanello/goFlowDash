package diagram

type Diagram struct {
	Boxes []Box  `json:"boxes" required:"true" binding:"required"`
	Edges []Edge `json:"edges" required:"true" binding:"required"`
}
