package main

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "time"
    "log"
)

var client *mongo.Client

type GeneratedImage struct {
    URL   string `json:"url"`
    Style string `json:"style"`
}

func connectDatabase() error {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    var err error
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return err
    }
    if err = client.Ping(context.TODO(), nil); err != nil {
        return err
    }

    // Set up TTL index to expire documents after 1 hour
    collection := client.Database("imagegen").Collection("sessions")
    indexModel := mongo.IndexModel{
        Keys: bson.M{"created_at": 1}, 
        Options: options.Index().SetExpireAfterSeconds(3600), 
    }
    _, err = collection.Indexes().CreateOne(context.TODO(), indexModel)
    if err != nil {
        log.Printf("Failed to create TTL index: %v\n", err)
        return err
    }

    log.Println("Connected to MongoDB and TTL index created.")
    return nil
}

func saveGeneratedImages(sessionID string, images []GeneratedImage) error {
    collection := client.Database("imagegen").Collection("sessions")
    _, err := collection.InsertOne(context.TODO(), bson.M{
        "session_id":       sessionID,
        "generated_images": images,
        "created_at":       time.Now(),
    })
    return err
}

func getImagesBySessionID(sessionID string) ([]GeneratedImage, error) {
    collection := client.Database("imagegen").Collection("sessions")
    var result struct {
        GeneratedImages []GeneratedImage `bson:"generated_images"`
    }

    filter := bson.M{"session_id": sessionID}
    err := collection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        return nil, err
    }

    return result.GeneratedImages, nil
}
