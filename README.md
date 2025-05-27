# Blog Aggregator

This is a blog aggregator, built with go and uses postgres. It was a guided project and one that I had to rely heavily on given solutions to solve. This is mainly me going through the motions as I try and get better.

## Installation

This requires [Postgres](https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql) and [Go](https://golang.org/dl/) installed on your machine. 

Then install the binary through your prefered terminal.

```bash
go install ...
```

## Config
Create a `.gatorconfig.json` file in your home directory in the following style.

```json
{
    "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

Create a new user:

```bash
Blog_Aggregator register <name>
```

Add a feed:
```bash
Blog_Aggregator addfeed <url>
```

Start the aggregator:
```bash
Blog_Aggregator agg 30s
```

View posts:
```bash
Blog_Aggregator browse [limit]
```

Other commands include
-`Blog_Aggregator login` - Logs in as an already registered user
-`Blog_Aggregator users` - Lists all users
-`Blog_Aggregator feeds` - Lists all feeds
-`Blog_Aggregator follow <url>` - Follow a feed that already exists in the database
-`Blog_Aggregator unfollow <url>` - Unfollow a feed that already exists in the database
-`Blog_Aggregator reset` - Deletes all users in the database