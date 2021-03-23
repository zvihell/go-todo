run:
	go run main.go

build:
	go build -o main main.go

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up