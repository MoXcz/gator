# Gator (TypeScript)

This is the Node.js/TypeScript implementation of the `gator` RSS feed aggregator.

## Requirements

- Node.js (v18+)
- PostgreSQL

## Setup

First, ensure you have PostgreSQL installed and running. Set up your configuration file at `$HOME/.gatorconfig.json` with your database URL:

```json
{
  "db_url": "postgres://username:password@host:port/database"
}
```

Install the dependencies:

```sh
npm install
```

Generate and migrate the database schema using Drizzle:

```sh
npm run generate
npm run migrate
```

## Usage

You can run the CLI using `npm start -- <command> [args...]`:

```sh
npm start -- register mocos
npm start -- addfeed "Hacker News" "https://news.ycombinator.com/rss"
npm start -- agg 1m
npm start -- browse 5
```

### Available Commands

- `login <user>`: Login to a registered user
- `register <user>`: Add a new user
- `reset`: Remove all users, including their feeds and posts
- `users`: Print all registered users (plus show current user)
- `agg <time_between_reqs>`: Aggregate all added feeds to posts
- `addfeed <name> <feed_url>`: Add a new feed
- `feeds`: Print all followed feeds
- `follow <feed_url>`: Follow a new feed for current user
- `unfollow <feed_url>`: Unfollow a followed feed for current user
- `following`: Print followed feeds of current user
- `browse <limit>`: Print a defined number of posts from followed feeds
