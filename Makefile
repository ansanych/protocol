genAuth:
	protoc -I ./api \
 			--go_out ./api \
 			--go_opt paths=source_relative \
 			--go-grpc_out ./api \
 			--go-grpc_opt paths=source_relative \
 			./api/auth/v1/auth.proto

genSettings:
	protoc -I ./api \
 			--go_out ./api \
 			--go_opt paths=source_relative \
 			--go-grpc_out ./api \
 			--go-grpc_opt paths=source_relative \
 			./api/settings/v1/settings.proto
