run:
	go run cmd/app/main.go
swagger-install:
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag@latest
swagger-init:
	swag init -g internal/controller/http/v1/router.go