package item

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Item struct {
		Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title       string             `json:"title" bson:"title"`
		Price       float64            `json:"price" bson:"price"`
		Damage      int                `json:"damage" bson:"damage"`
		ImageUrl    string             `json:"image_url" bson:"image_url"`
		UsageStatus bool               `json:"usage_status" bson:"usage_status"`
		CreateAt    time.Time          `json:"create_at" bson:"create_at"`
		UpdateAt    time.Time          `json:"update_at" bson:"update_at"`
	}
)
