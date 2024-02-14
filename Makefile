genShare:
	protoc -I ./api_1 \
			--go_out ./api_1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_1 \
			--go-grpc_opt paths=source_relative \
			./api_1/share.proto

genOzon:
	protoc -I ./api_1 \
			--proto_path=./api_1 \
			--go_out ./api_1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_1 \
			--go-grpc_opt paths=source_relative \
			./api_1/ozon.proto

genSettings:
	protoc -I ./api_1 \
			--proto_path=./api_1 \
			--go_out ./api_1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_1 \
			--go-grpc_opt paths=source_relative \
			./api_1/settings.proto

genAuth:
	protoc -I ./api_1 \
			--proto_path=./api_1 \
			--go_out ./api_1 \
			--go_opt paths=source_relative \
			--go-grpc_out ./api_1 \
			--go-grpc_opt paths=source_relative \
			./api_1/auth.proto