# RSS Scraper

Welcome to your RSS Scraper project! This tool allows users to collect, follow, and manage RSS feeds efficiently.

## Features

- **Add RSS Feeds:** Users can add RSS feeds to be collected.
- **Follow/Unfollow Feeds:** Users can follow and unfollow RSS feeds that other users have added.
- **Latest Posts:** Fetch the latest posts from the RSS feeds that users follow.

## Getting Started

### Prerequisites

Before you begin, ensure you have Go installed on your machine. You will also need an external client like [Thunder Client](https://www.thunderclient.io/) or [Postman](https://www.postman.com/) for interacting with the API.

- [Go](https://golang.org/doc/install)
- [Thunder Client (VS Code Extension)](https://www.thunderclient.io/) or [Postman](https://www.postman.com/)

### Installation

Follow these steps to clone and set up the repository:

```bash
# Clone the repository
git clone https://github.com/your-username/rss-scraper.git

# Navigate into the project directory
cd rss-scraper
```


RUNNING THE PROJECT
To run the project, use the following commands:

# Install dependencies
go mod tidy

# Run the application
go run main.go
Copy icon
USING THE API
You can use Thunder Client or Postman to interact with the API. Here are the available endpoints and their descriptions:

AUTHORIZATION
Your API has middleware for authentication that requires an API key to be passed in the headers.

API Key Header:
Header: "Authorization: Bearer YOUR_API_KEY"
Copy icon
USERS
Create a User

Method: POST
Endpoint: /v1/users
Body:
{
  "username": "exampleUser",
  "password": "examplePass"
}
Copy icon
Get User

Method: GET

Endpoint: /v1/users

Authorization: Requires authentication

Example using curl:

curl -H "Authorization: Bearer YOUR_API_KEY" \
     -X GET http://localhost:8080/v1/users
