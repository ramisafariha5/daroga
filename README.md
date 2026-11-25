# Daroga

Daroga is a Go command-line tool for inspecting Docker containers. It supports filtering by container name or ID, table and JSON output, and optional Slack notifications.

## Features

- Inspect running and stopped containers
- Filter by exact name, full ID, or ID prefix
- Print table or JSON output
- Send inspection summaries to Slack
- Read Docker configuration from the local Docker environment
- Use context-based Docker API timeouts
- Unit-test Docker filtering and CLI formatting
- Run formatting, vet, and race-enabled tests in CI

## Requirements

- Go 1.26 or newer
- Docker Desktop or Docker Engine
- Docker socket access
- Optional Slack bot token and channel ID

## Install

```bash
go mod download
go build -o daroga ./main.go
```

## Usage

List all containers:

```bash
./daroga run
```

Filter by name:

```bash
./daroga run --name api
```

Filter by ID or ID prefix:

```bash
./daroga run --id 1234567890ab
```

Output JSON:

```bash
./daroga run --format json
```

Send the result to Slack:

```bash
export DAROGA_SLACK_TOKEN="xoxb-your-token"
export DAROGA_SLACK_CHANNEL="C0123456789"
./daroga run --name api --notify
```

Slack variables are read only from the environment. They must never be committed to source control.

## Architecture

- `main.go` and `cmd/` implement the Cobra CLI
- `config/` owns command configuration
- `internal/inspector/` communicates with the Docker API and builds container summaries
- `internal/notifier/` sends optional Slack messages
- `*_test.go` files cover filtering and formatting behavior
- `.github/workflows/ci.yml` runs formatting, vet, and race tests

## Development

```bash
gofmt -w $(find . -name '*.go' -not -path './.git/*')
go vet ./...
go test -race ./...
```

## Security

The Slack credentials found in an earlier implementation must be revoked and considered compromised. Never store API tokens in source files, README examples, environment-specific files, or commit history.

## Limitations

- Notifications are one-shot summaries; Daroga does not currently track state changes or run continuously
- Docker host information comes from `DOCKER_HOST` and Docker environment configuration
- Slack notifications require a bot token and channel permission

## Future Improvements

- State-change monitoring
- Webhook or Prometheus integration
- Container log and resource metrics
- Docker Compose support
- Structured configuration file
- Docker integration tests
- Homebrew and release binaries

## License

Apache License 2.0. See `LICENSE`.