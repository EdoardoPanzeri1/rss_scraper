# RSS Scraper

The RSS Scraper project allows users to collect, follow, and manage RSS feeds efficiently.

## Features

- **Add RSS Feeds:** Users can add RSS feeds to be collected.
- **Follow/Unfollow Feeds:** Users can follow and unfollow RSS feeds that other users have added.
- **Latest Posts:** Fetch the latest posts from the RSS feeds that users follow.

## Getting Started

### Prerequisites

Before you begin, ensure you have Go installed on your machine. You will also need an external client like [Thunder Client](https://www.thunderclient.io/) or [Postman](https://www.postman.com/) for interacting with the API.

Additionally, you'll need `sqlc` and `goose` for handling database interactions and migrations:

- [Go](https://golang.org/doc/install)
- [sqlc](https://sqlc.dev/)
- [goose](https://github.com/pressly/goose)
- [Thunder Client (VS Code Extension)](https://www.thunderclient.io/) or [Postman](https://www.postman.com/)

### Installation

Follow these steps to clone and set up the repository:

```bash
# Clone the repository
git clone https://github.com/your-username/rss-scraper.git

# Navigate into the project directory
cd rss-scraper
```
### Database Setup
Install and configure `goose`:

```bash
# Install goose CLI
go install github.com/pressly/goose/v3/cmd/goose@latest

# Run migrations
goose up
```
Generate Go code from SQL queries using `sqlc`:

```bash
# Install sqlc CLI
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# Generate code
sqlc generate
```
### Running the Project
To run the project, use the following commands:
```bash
# Install dependencies
go mod tidy

# Run the application
go run main.go
```
### Using the API

You can use Thunder Client or Postman to interact with the API. Below are the available endpoints with brief descriptions:

#### Authorization

All endpoints requiring authorization use an API key in the headers.

- **API Key Header:**

  Header: "Authorization: Bearer YOUR_API_KEY"
- **USERS:**

  Create User: `POST /v1/users` - Register a new user.
  Get User: `GET /v1/users` - Retrieve the authenticated user's details.
- **FEEDS:**

  Create Feed: `POST /v1/feeds` - Add a new RSS feed.
  Get Feeds: `GET /v1/feeds` - List all available feeds.

- **FEED FOLLOWS:**

  Get Feed Follows: `GET /v1/feed_follows` - List feeds followed by the user.
  Follow Feed: `POST /v1/feed_follows` - Follow a specific feed.
  Unfollow Feed: `DELETE /v1/feed_follows/{feedFollowID}` - Unfollow a specific feed.
- **POSTS:**

   Get Posts: `GET /v1/posts` - List posts from followed feeds.

## DEPENDENCIES
This project uses the following third-party libraries:

- [sqlc](https://sqlc.dev/): sqlc.dev - Generates type-safe Go code from SQL.
- [goose](https://github.com/pressly/goose): GitHub repository - A database migration tool.
- [godotenv](https://github.com/joho/godotenv): GitHub repository - Loads environment variables from `.env` files.
- [pq](https://github.com/lib/pq): GitHub repository - A pure Go Postgres driver for the database/sql package.

Please ensure to review the licenses of these libraries for compliance.
