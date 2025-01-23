package Adm

// import (
// 	"context"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// // Add this struct for the response
// type UserListResponse struct {
// 	ID   string `json:"id" bson:"_id"`
// 	Name string `json:"name" bson:"name"`
// 	Role string `json:"role" bson:"role"`
// }

// // Add this method to Application struct
// func AdminListUsers(ctx context.Context) ([]UserListResponse, error) {
// 	// Set timeout for operation
// 	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
// 	defer cancel()

// 	// Create pipeline to project only needed fields
// 	pipeline := []bson.M{
// 		{
// 			"$project": bson.M{
// 				"_id":  1,
// 				"name": 1,
// 				"role": 1,
// 			},
// 		},
// 	}

// 	// Execute aggregation
// 	cursor, err := app.userCollection.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	// Decode results
// 	var users []UserListResponse
// 	if err := cursor.All(ctx, &users); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }
