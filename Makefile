gen:
	protoc --go_out=:. --go-grpc_out=:. ./proto/*.proto

clean:
	rm -rf ./pb/*.go

run:
	go run main.go