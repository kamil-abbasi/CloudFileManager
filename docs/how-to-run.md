# How To Run

**Note**: This guide assumes you are in the repository root directory.

## Docker (Recommended)

### Prerequisities:

- docker and docker compose installed on your system

### Development

To start in development mode run: `docker compose up -d --build`.

**Tip**: Install go and your editor's tooling for it for better developer experience.

### Production

To start in production mode run: `docker compose -f prod.compose.yaml up -d --build`

## Local Machine

### Prerequisities:

- go installed on your system

### Development

**Note**: Air is a tool for hot reloading go apps, so you don't have to constantly rebuild code after making changes. To install it run this command:
`go install github.com/air-verse/air@latest`

To run in development simply execute this command: `air`.

### Production

For production execute the following commands:

- `go build ./cmd/server/main.go`
- `./main`

**Tip**: If you are running in development mode, install go and your editor's tooling (vscode for example) for better developer experience.
