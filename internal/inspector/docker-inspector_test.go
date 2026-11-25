package inspector

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types"
)

type mockDockerAPI struct {
	containers []types.Container
	err        error
}

func (m *mockDockerAPI) ContainerList(context.Context, types.ContainerListOptions) ([]types.Container, error) {
	return m.containers, m.err
}

func (m *mockDockerAPI) ContainerInspect(context.Context, string) (types.ContainerJSON, error) {
	return types.ContainerJSON{}, nil
}

func TestInspectFiltersByContainerName(t *testing.T) {
	api := &mockDockerAPI{containers: []types.Container{{
		ID: "1234567890abcdef", Names: []string{"/api"}, Image: "example/api:1", State: "running", Status: "Up 5 minutes",
	}}}
	inspector := NewDockerInspector(api)

	summaries, err := inspector.Inspect(context.Background(), "api")
	if err != nil {
		t.Fatal(err)
	}
	if len(summaries) != 1 || summaries[0].Name != "api" {
		t.Fatalf("unexpected summaries: %+v", summaries)
	}
	if summaries[0].ID != "1234567890ab" {
		t.Fatalf("unexpected short ID: %q", summaries[0].ID)
	}
}

func TestInspectFiltersByIDPrefix(t *testing.T) {
	api := &mockDockerAPI{containers: []types.Container{{
		ID: "1234567890abcdef", Names: []string{"/api"}, State: "exited", Status: "Exited (0)",
	}}}
	inspector := NewDockerInspector(api)

	summaries, err := inspector.Inspect(context.Background(), "1234")
	if err != nil {
		t.Fatal(err)
	}
	if len(summaries) != 1 || summaries[0].Running {
		t.Fatalf("unexpected summaries: %+v", summaries)
	}
}

func TestInspectReturnsNotFound(t *testing.T) {
	api := &mockDockerAPI{containers: nil}
	inspector := NewDockerInspector(api)

	if _, err := inspector.Inspect(context.Background(), "missing"); err == nil {
		t.Fatal("expected a not-found error")
	}
}
