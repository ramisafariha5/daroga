package inspector

import "context"

// Inspector provides a context-aware container inspection operation.
type Inspector interface {
	Inspect(context.Context, string) ([]ContainerSummary, error)
}
