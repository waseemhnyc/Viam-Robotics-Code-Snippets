package main

import (
	"context"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/utils/rpc"
)

func main() {
	logger := logging.NewLogger("client")
	robot, err := client.New(
		context.Background(),
		"ADDRESS FROM THE VIAM APP",
		logger,
		client.WithDialOptions(rpc.WithEntityCredentials(
			// Replace "<API-KEY-ID>" (including brackets) with your machine's API key ID
			"<API-KEY-ID>",
			rpc.Credentials{
				Type: rpc.CredentialsTypeAPIKey,
				// Replace "<API-KEY>" (including brackets) with your API key
				Payload: "<API-KEY>",
			})),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer robot.Close(context.Background())
	logger.Info("Resources:")
	logger.Info(robot.ResourceNames())
}
