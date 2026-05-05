

genprotos:
	protoc \
		--proto_path=protos/proto \
		--go_out=. \
		--go_opt=paths=import \
		--go-grpc_out=. \
		--go-grpc_opt=paths=import \
		protos/proto/collections.proto \
		protos/proto/translations.proto
	cd protos && go mod tidy
	go work sync


