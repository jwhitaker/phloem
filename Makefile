# CMDS = apiservice webapi

# build: $(CMDS) ## Build all commands

# $(CMDS): ## Build
# 	go build -o ./bin/$@ ./cmd/$@

# webapi: build_webapi

help:
	@echo showing help

webapi:
	go build -o ./bin/webapi ./cmd/webapi

	
clean:  ## Clean up everything
	rm -rf ./bin

run: 
	(trap 'kill 0' SIGINT; ./bin/apiservice & ./bin/webapi)
	
gogo:
	@docker-compose up -d

# docker_make_base:
# 	docker build -t jwhitaker/recipebook-base:latest -f ./build/package/servicebase.dockerfile .

docker_webapi: 
	docker build -t jwhitaker/recipebook-base:latest \
		--build-arg SERVICE_NAME=webapi \
		-f ./build/package/servicebase.dockerfile .