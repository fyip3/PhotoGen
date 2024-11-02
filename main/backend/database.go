package main

import (
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "time"
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
    return client.Ping(context.TODO(), nil)
}

func saveGeneratedImages(sessionID string, images []GeneratedImage) error {
    collection := client.Database("imagegen").Collection("sessions")
    _, err := collection.InsertOne(context.TODO(), map[string]interface{}{
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

    filter := map[string]interface{}{"session_id": sessionID}
    err := collection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        return nil, err
    }

    return result.GeneratedImages, nil
}
