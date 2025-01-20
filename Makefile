ssh:
	#@ssh-keygen -f "/Users/uki/.ssh/known_hosts" -R "[localhost]:2222"
	@cat ssh_http_tunel.go | ssh localhost -p 2222
	
build:
	go build -o bin/price

run: build
	./bin/price

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/service.proto

.PHONY: proto
