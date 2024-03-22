install:
	go mod tidy
	go mod download
	go mod vendor

start:
	go run main.go start

admin:
	go run main.go set-admin $(email) $(database-url)

balance:
	go run main.go add-balance $(email) $(amount) $(database-url)