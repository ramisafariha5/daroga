# Daroga Interview Guide

## 30-second overview

Daroga is a Go CLI that uses Cobra and the Docker SDK to inspect Docker containers. It supports name or ID filtering and table/JSON output, with an optional Slack notification. The project includes unit tests, race detection, vetting, and GitHub Actions CI.

## Technical decisions

- Cobra provides a consistent CLI command and flag model.
- The Docker SDK communicates with Docker through `DOCKER_HOST` or the Docker environment.
- `context.Context` imposes cancellation and timeout limits on Docker operations.
- Filtering is isolated in `DockerInspector`, making name and ID-prefix behavior testable.
- Interface-based Docker dependencies allow mocks instead of requiring Docker for unit tests.
- Slack credentials come from environment variables to keep secrets outside source control.
- GitHub Actions checks formatting, vetting, and race-enabled tests.

## Questions to expect

### Why did you build it?

It began as an exploration of Go, Docker APIs, Cobra, and Slack integrations. I turned it into a complete CLI by replacing prototype behavior with filtering, structured output, errors, tests, and CI.

### How does container lookup work?

The inspector first asks Docker for all containers, including stopped containers. It matches an exact name, full ID, or ID prefix and converts the first container name into a cleaner display name.

### What happens if Docker is unavailable?

The Docker client creation and container list operations return wrapped errors. The CLI returns an error instead of continuing with invalid state or printing misleading output.

### How would you test Slack safely?

Unit tests verify missing configuration and message formatting. Slack API calls are environment-dependent, so a real integration test should use a mock HTTP server or fake notifier interface rather than sending messages during normal unit tests.

### What are the next improvements?

- Persistent state tracking for state-change notifications
- Scheduled monitoring mode
- Container metrics and log access
- Webhook configuration
- Integration tests against Docker
- Signed release binaries

## Resume-safe language

Use: “Developed a Go CLI for Docker container inspection with name/ID filtering, table/JSON output, optional Slack alerts, unit tests, and GitHub Actions CI.”

Do not claim continuous monitoring or Slack state-change alerts until those features are implemented.
