package services

import (
	"../../../db"
	"../models"
	"../models/response"
	"github.com/op/go-logging"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var log = logging.MustGetLogger("Todo")

func AddTodo(todo models.Todo) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("tranquy2512_todo").C("todos")
	err := collection.Insert(todo)
	res := response.Status{}
	if err != nil {
		log.Error(err)
		res.Error = err
		res.Success = false
	} else {
		res.Error = err
		res.Success = true
	}
	return res
}

func GetTodoById(id string) models.DataResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("tranquy2512_todo").C("todos")
	result := models.Todo{}
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	res := models.DataResponse{}
	if err != nil {
		log.Error(err)
		res.Status.Error = err
		res.Status.Success = false
	} else {
		res.Data = result
		res.Status.Success = true
		res.Status.Error = err
	}
	return res
}

func RemoveTodo(id string) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("tranquy2512_todo").C("todos")
	err := collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	res := response.Status{}
	if err != nil {
		log.Error(err)
		res.Error = err
		res.Success = false
	} else {
		res.Success = true
		res.Error = err
	}
	return res
}

func UpdateTodo(todo models.Todo) response.Status {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("tranquy2512_todo").C("todos")
	err := collection.Update(bson.M{"_id": todo.Id}, bson.M{"$set": todo})
	res := response.Status{}
	if err != nil {
		log.Error(err)
		res.Error = err
		res.Success = false
	} else {
		res.Success = true
		res.Error = err
	}
	return res
}

func GetTodoList() models.ListResponse {
	var session = db.GetMongoSession()
	defer session.Close() // session must close at the end
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("tranquy2512_todo").C("todos")
	result := []models.Todo{}
	err := collection.Find(nil).All(&result)
	res := models.ListResponse{}
	if err != nil {
		log.Error(err)
		res.Status.Error = err
		res.Status.Success = false
	} else {
		res.List = result
		res.Status.Success = true
		res.Status.Error = err
	}
	return res
}
