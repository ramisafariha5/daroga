package inspector

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// ContainerSummary is a small, presentation-ready representation of a Docker container.
type ContainerSummary struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	State   string `json:"state"`
	Status  string `json:"status"`
	Running bool   `json:"running"`
}

// DockerAPI contains only the Docker operations used by the inspector.
type DockerAPI interface {
	ContainerList(context.Context, types.ContainerListOptions) ([]types.Container, error)
	ContainerInspect(context.Context, string) (types.ContainerJSON, error)
}

// DockerInspector discovers and inspects Docker containers.
type DockerInspector struct {
	Client DockerAPI
}

func NewDockerInspector(client DockerAPI) *DockerInspector {
	return &DockerInspector{Client: client}
}

// Inspect lists containers and optionally filters by a full or partial ID or a container name.
func (d *DockerInspector) Inspect(ctx context.Context, nameOrID string) ([]ContainerSummary, error) {
	containers, err := d.Client.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("list Docker containers: %w", err)
	}

	var summaries []ContainerSummary
	selector := strings.TrimSpace(nameOrID)
	for _, container := range containers {
		containerName := containerName(container)
		matches := selector == ""
		if selector != "" {
			matches = strings.HasPrefix(container.ID, selector) || container.ID == selector || containerName == selector
		}
		if !matches {
			continue
		}

		shortID := container.ID
		if len(shortID) > 12 {
			shortID = shortID[:12]
		}

		running := strings.EqualFold(container.State, "running")
		summaries = append(summaries, ContainerSummary{
			ID:      shortID,
			Name:    containerName,
			Image:   container.Image,
			State:   container.State,
			Status:  container.Status,
			Running: running,
		})
	}

	if selector != "" && len(summaries) == 0 {
		return nil, fmt.Errorf("container %q was not found", selector)
	}
	return summaries, nil
}

func containerName(container types.Container) string {
	if len(container.Names) == 0 {
		return ""
	}
	return strings.TrimPrefix(container.Names[0], "/")
}

func InitDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("create Docker client: %w", err)
	}
	return cli, nil
}
