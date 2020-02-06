gogo:
	docker-compose up -d

webapi:
	go build -o bin/webapi ./cmd/webapi

.PHONY: gogo
