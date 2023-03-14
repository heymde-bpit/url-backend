package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage interface {
	inserttoDB(*link) error
	getfromDB(string, *link) error
}

type MongoStore struct {
	collection *mongo.Collection
}

func newDBStore() (*MongoStore, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://mdoffical8a26e:zendaya@cluster1.ikmoqw4.mongodb.net/?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return &MongoStore{
		collection: client.Database("URL-Shortener").Collection("id-urls"),
	}, nil

}

func (m *MongoStore) inserttoDB(l *link) error {
	err := m.collection.FindOne(context.TODO(), bson.D{{Key: "url", Value: l.URL}}).Decode(&l)

	if err != nil {
		_, err := m.collection.InsertOne(context.TODO(), l)

		if err != nil {
			return err
		}

		return nil

	}

	return nil
}

func (m *MongoStore) getfromDB(s string, resp *link) error {
	err := m.collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: s}}).Decode(&resp)

	if err != nil {
		return err
	}

	return nil
}
