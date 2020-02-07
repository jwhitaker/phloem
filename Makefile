CMDS = apiservice webapi

all: $(CMDS) ## Build all commands

$(CMDS): ## Build
	go build -o ./bin/$@ ./cmd/$@
	
clean:  ## Clean up everything
	rm -rf ./bin

run: 
	(trap 'kill 0' SIGINT; ./bin/apiservice & ./bin/webapi)
	
gogo:
	@docker-compose up -d

