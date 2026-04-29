# `gator`

> Brought to you thanks to [boot.dev](https://www.boot.dev/)

Gator is a command-line RSS feed aggregator designed to follow posts from websites, complete with integrated PostgreSQL storage.

This repository contains two different implementations of the Gator CLI:

- **[Go Implementation (`go/`)](./go/README.md)**: Written in Go, utilizing `sqlc` for database interactions.
- **[TypeScript Implementation (`ts/`)](./ts/README.md)**: A Node.js/TypeScript implementation using the Drizzle ORM.

## Overview

Both versions of `gator` use a PostgreSQL database to store users, feeds, and fetched posts. They allow you to:

- Register users and switch between them
- Follow RSS feeds
- Periodically aggregate feeds to fetch new posts
- Browse collected posts from your terminal

See the respective directories for specific build and usage instructions.

