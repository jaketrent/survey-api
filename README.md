

## DB Migration

```
go get -u github.com/pressly/goose/cmd/goose
cd ./db/migrations
goose postgres "postgres://postgres:postgres@localhost:5432/survey_api?sslmode=disable" up
```
