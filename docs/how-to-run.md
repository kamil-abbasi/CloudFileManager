# How To Run

**Note**: This guide assumes you are in the repository root directory.

## Prerequisities:

- docker and docker compose installed on your system

## Development

To start in development mode run: `docker compose up -d --build`.

**Tip**: Install go and your editor's tooling for it for better developer experience.

## Production

To start in production mode run: `docker compose -f prod.compose.yaml up -d --build`
