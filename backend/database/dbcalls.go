package database

import (
	"context"
	"fmt"
	"goTodo/constant"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	instance := mgr.connection.Database(constant.Database).Collection(collectionName)
	res, err := instance.InsertOne(context.TODO(), data)
	return res.InsertedID, err
}

func (mgr *manager) Delete(id primitive.ObjectID, collectionName string) error {
	instance := mgr.connection.Database(constant.Database).Collection(collectionName)
	filter := bson.M{"_id": id}
	_, err := instance.DeleteOne(context.TODO(), filter)
	return err
}

func (mgr *manager) DeleteAll(collectionName string) error {
	instance := mgr.connection.Database(constant.Database).Collection(collectionName)
	_, err := instance.DeleteMany(context.TODO(), bson.M{}) // Delete all documents
	return err
}

func (mgr *manager) Fetch(id primitive.ObjectID, collectionName string) (interface{}, error) {
	instance := mgr.connection.Database(constant.Database).Collection(collectionName)
	filter := bson.M{"_id": id}
	result := instance.FindOne(context.TODO(), filter)
	var data interface{}
	err := result.Decode(&data)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("document with ID %s not found in collection %s", id, collectionName)
	}
	return data, err
}

func (mgr *manager) FetchAll(collectionName string) (interface{}, error) {
	instance := mgr.connection.Database(constant.Database).Collection(collectionName)
	//Find returns a matching cursor over the database document
	cur, err := instance.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var results []interface{}
	for cur.Next(context.TODO()) {
		var elem interface{}
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (mgr *manager) Update(id primitive.ObjectID, data map[string]interface{}, collectionName string) (interface{}, error) {
	instance := mgr.connection.Database(constant.Database).Collection(collectionName)
	filter := bson.M{"_id": id}

	updateDoc := bson.M{"$set": data} // Update specific fields based on data

	updateResult := instance.FindOneAndUpdate(context.TODO(), filter, updateDoc, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var updatedData interface{}
	err := updateResult.Decode(&updatedData)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("document with ID %s not found in collection %s", id, collectionName)
	}
	return updatedData, err
}
