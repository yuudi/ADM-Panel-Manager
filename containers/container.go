package containers

import (
	"context"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var clientInstance *client.Client
var clientOnce sync.Once

func GetClient() *client.Client {
	clientOnce.Do(func() {
		var err error
		clientInstance, err = client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			panic(err)
		}
	})
	return clientInstance
}

func ListContainers() ([]types.Container, error) {
	c := GetClient()
	return c.ContainerList(context.Background(), types.ContainerListOptions{})
}
