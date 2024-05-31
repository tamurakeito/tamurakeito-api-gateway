hello:
	echo "hello, world!"

start:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o docker/main src/main.go
	cd docker && docker-compose up --build