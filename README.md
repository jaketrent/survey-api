## Install dependencies

```
brew install dep
dep ensure
```

## Deploy

```
up
```

## DB Setup

- Create RDS instance, survey_api, via console, make publicly available
- `aws rds describe-db-instances --db-instance-identifier licketybit`
- `psql postgres://user:pass@thatEndpoint:5432/survey_api`
- `create user survey_api_app with password 'thepassword';`
- `grant all privileges on database survey_api to survey_api_app;`
- `up env add DATABASE_URL=postgres://survey_api_app:password@endpoint:5432/survey_api`

## DB Migration

```
go get -u github.com/pressly/goose/cmd/goose
cd ./db/migrations
goose postgres "postgres://postgres:postgres@localhost:5432/survey_api?sslmode=disable" up
```
