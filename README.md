# Gator

> Brought to you thanks to [boot.dev](https://www.boot.dev/)

This is `gator` a RSS feed aggregator to follow posts from websites with integrated storage.

## Quick Start

Install `gator` using the Go toolchain (information on how to install it [here](https://go.dev/doc/install)):

```sh
go install github.com/MoXcz/gator@latest
```

You will also need postgresql for storage (information on how to install it [here](https://www.postgresql.org/download/).

With both the program and postgresql installed set the configuration file at
`$HOME/.gatorconfig.json` with the database URL:
```json
{
    "db_url":"protocol://username:password@host:port/database",
}
```

> `$HOME` is defined: On Unix, including macOS, it returns the $HOME environment variable. On Windows, it returns %USERPROFILE%. On Plan 9, it returns the $home environment variable as specified [here](https://pkg.go.dev/os#UserHomeDir)

Run `gator` (make sure `$GOPATH` is set correctly and is part of your `$PATH`. More information [here](https://go.dev/wiki/SettingGOPATH)):

```sh
gator register <user> # register a user
gator addfeed "Hacker News" "https://news.ycombinator.com/rss" # add feed
gator agg 1m # Refresh posts from feeds each minute
gator browse 5 # Print the 5 most recent posts from feeds
```

## Usage

- `login <user>`. Login to a registered user
- `register <user>`. Add a new user
- `reset`. Remove all users, including their feeds and posts
- `users`. Print all registered users (plus show current user)
- `agg <time_between_reqs>`. Aggregate all added feeds to posts
- `addfeed <name> <feed_url>`. Add a new feed
- `feeds`. Print all followed feeds
- `follow <feed_url>`. Follow a new feed for current user
- `unfollow <feed_url>`. Unfollow a followed feed for current user
- `following`. Print followed feeds of current user
- `browse <limit>`. Print a defined number of posts from followed feeds

## Examples

```sh
gator register mocos
User '69dd6e48-01c5-484b-b49e-457ebebaf8a2: mocos' added successfully

gator addfeed "Hacker News" "https://news.ycombinator.com/rss"
Feed added successfuly:
ID:  065f4671-84dc-4b0a-a08a-4abaa9902b0c
Created:  2025-03-19 11:37:42.940019 +0000 +0000
Updated:  2025-03-19 11:37:42.940022 +0000 +0000
Name:  Hacker News
URL:  https://news.ycombinator.com/rss
Created:  2025-03-19 11:37:42.940019 +0000 +0000
User:  mocos

Feed followed successfuly

gator following
User feed following: mocos
Hacker News

gator users
 * mocos (current)

gator agg 1m
Collecting feeds every 1m
Feed Hacker News with 30 posts found:
Post Memory Safety for Web Fonts added
Post The Lost Art of Research as Leisure added
Post fd: A simple, fast and user-friendly alternative to 'find' added
Post Konva.js - Declarative 2D Canvas for React, Vue, and Svelte added
Post The clustering behavior of sliding windows added
Post Show HN: AGX – Open-Source Data Exploration for ClickHouse (The New Standard?) added
Post Ikemen-GO: open-source reimplementation of MUGEN added
Post I got 100% off my train travel added
Post Two new PebbleOS watches added
Post The Origin of the Pork Taboo added
Post Wolfram: Learning about Innovation from Half a Century of Conway's Game of Life added
Post Make Ubuntu packages 90% faster by rebuilding them added
Post Selective async commits in PostgreSQL – balancing durability and performance added
Post Visualising data structures and algorithms through animation added
Post Show HN: I made a tool to port tweets to Bluesky mantaining their original date added
Post Zest: a programming language for malleable and legible systems added
Post Google to buy Wiz for $32B added
Post CVE-2024-9956 – PassKey Account Takeover in All Mobile Browsers added
Post Some notes on Grafana Loki's new "structured metadata" added
Post Apple Loses Top Court Fight Over German Antitrust Crackdown added
Post HTTrack Website Copier added
Post Designing Electronics That Work added
Post An early look at cryptographic watermarks for AI-generated content added
Post SheepShaver is an open source PowerPC Apple Macintosh emulator added
Post Karatsuba Matrix Multiplication and Its Efficient Hardware Implementations added
Post For Delivery Workers in Latin America, Affordable E-Bikes Are a Superpower added
Post Everything Is Broken: Shipping Rust-Minidump at Mozilla (2022) added
Post US appeals court rules AI generated art cannot be copyrighted added
Post Video game workers in North America now have an industry-wide union added
Post Show HN: "Git who" – A new CLI tool for industrial-scale Git blaming added
^C

gator browse 5
Mon 1 Jan from 065f4671-84dc-4b0a-a08a-4abaa9902b0c
--- Designing Electronics That Work ---

Link: https://www.hscott.net/designing-electronics-that-work/
=====================================

Mon 1 Jan from 065f4671-84dc-4b0a-a08a-4abaa9902b0c
--- An early look at cryptographic watermarks for AI-generated content ---

Link: https://blog.cloudflare.com/an-early-look-at-cryptographic-watermarks-for-ai-generated-content/
=====================================

Mon 1 Jan from 065f4671-84dc-4b0a-a08a-4abaa9902b0c
--- SheepShaver is an open source PowerPC Apple Macintosh emulator ---

Link: https://www.emaculation.com/doku.php/sheepshaver
=====================================

Mon 1 Jan from 065f4671-84dc-4b0a-a08a-4abaa9902b0c
--- Karatsuba Matrix Multiplication and Its Efficient Hardware Implementations ---

Link: https://arxiv.org/abs/2501.08889
=====================================

Mon 1 Jan from 065f4671-84dc-4b0a-a08a-4abaa9902b0c
--- HTTrack Website Copier ---

Link: https://www.httrack.com/
===================================
```

## Build

### Clone the repo:

```sh
git clone https://github.com/MoXcz/gator
cd gator
```

### Build the project:

```sh
go build .
./gator
```

