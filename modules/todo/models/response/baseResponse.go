package response

import (
	"gopkg.in/mgo.v2/bson"
)

type Status struct {
	Success    bool
	Error      error
	CustomData bson.ObjectId
}
