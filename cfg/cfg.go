package cfg

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToDB() (*mongo.Client, error) {
	//dburi := os.Getenv("MONGODB_URI")
	//if dburi == "" {
	//	return nil, fmt.Errorf("MONGODB_URI is not set")
	//}
	clientOpt := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(clientOpt)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}
