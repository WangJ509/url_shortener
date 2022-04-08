# url_shortener
This project implements URL shorten API using gorm and gin with unit testing.

## How to run the code

### Create the table
`mysql -u {user_name} -p {database_name} < create_table_shorten_urls.sql`

### Run the API server
- `go mod tidy`
- `go run .`

### Run the unit test
`go test -v .`

### Imported Modules
- gin
- gorm


## Examples
- upload url: `curl --request POST \
  --url http://localhost:8088/api/v1/urls \
  --header 'Content-Type: application/json' \
  --data '{
		"url": "www.google.com",
	"expireAt": "2022-04-09T09:20:43Z" 
}'`
- redirect api: `curl --request GET \
  --url http://localhost:8088/{shorten_url}`