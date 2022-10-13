.PHONY: build clean test docker run


gen-all:
	protoc  -I=. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./thingmodel/thingmodel.proto