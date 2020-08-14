.PHONY: proto
hello:
	echo "hello"

build:
	go build -o ./bin/lion main.go

http:
	echo "starting http server"
	bin/lion http

grpc:
	echo "starting grpc server"
	bin/lion grpc
proto:
	./regenerate.sh
