package inspector

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/ramisafariha5/daroga/config"
)

type DockerInspector struct {
	Config *config.ContainerInfo
	Client *client.Client
}

func NewDockerInspector(config *config.ContainerInfo, client *client.Client) *DockerInspector {
	return &DockerInspector{
		Config: config,
		Client: client,
	}
}

func (d *DockerInspector) Inspect() {
	containers, err := d.Client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		fmt.Println(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
		stat, _ := d.Client.ContainerInspect(context.Background(), d.Config.ContainerId)
		//fmt.Println(cli.ContainerInspect(context.Background(), container.ID))
		//var data map[string]interface{}

		fmt.Println(stat.State.Status)
	}
}

func InitDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println(err)
	}
	return cli
}
