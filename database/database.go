package database

import (
	"context"
	"log"
	"time"

	"github.com/damianino/gradient-graphql/graph/model"
	"github.com/damianino/gradient-graphql/utils"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://o111oo11o:@cluster0.l6uvs.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("testDB").Collection("testCol")
	_, err = collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	return &DB{client}
}

func (db *DB) Save(input *model.NewGradient) (*model.Gradient, error) {
	collection := db.client.Database("gradientGql").Collection("gradients")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	g := utils.GenerateMissingValues(input)

	res, err := collection.InsertOne(ctx, g)
	if err != nil {
		return nil, err
	}

	g.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return g, nil
}

func (db *DB) FindById(id string) (*model.Gradient, error) {
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	collection := db.client.Database("gradientGql").Collection("gradients")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})

	g := model.Gradient{}

	err = res.Decode(&g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (db *DB) All() ([]*model.Gradient, error) {
	collection := db.client.Database("gradientGql").Collection("gradients")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var res []*model.Gradient
	for cur.Next(ctx) {
		var g *model.Gradient
		err := cur.Decode(&g)
		if err != nil {
			return nil, err
		}
		res = append(res, g)
	}

	return res, nil
}

func (db *DB) Comment(input *model.NewComment) (*model.Gradient, error) {
	ObjectID, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, err
	}
	collection := db.client.Database("gradientGql").Collection("gradients")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$push": bson.M{"comments": bson.M{"body": input.Comment}}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	g := model.Gradient{}
	res.Decode(&g)
	return &g, nil
}
