gogo:
	docker-compose up -d

webapi:
	go build -o bin/webapi ./cmd/webapi

apiservice:
	go build -o bin/apiservice ./cmd/apiservice

build: webapi apiservice

.PHONY: gogo
