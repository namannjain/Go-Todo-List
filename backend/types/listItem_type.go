package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type ListItemDao struct {
	ID     primitive.ObjectID `json:"_id,omitempty"  bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status, omitempty"`
}
