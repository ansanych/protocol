genStatic:
	protoc -I ./api \
 			--go_out ./api \
 			--go_opt paths=source_relative \
 			--go-grpc_out ./api \
 			--go-grpc_opt paths=source_relative \
 			./api/static/v1/static.proto

genAuth:
	protoc -I ./api \
			--proto_path=./ \
 			--go_out ./api \
 			--go_opt paths=source_relative \
 			--go-grpc_out ./api \
 			--go-grpc_opt paths=source_relative \
 			./api/auth/v1/auth.proto

genSettings:
	protoc -I ./api \
			--proto_path=./ \
 			--go_out ./api \
 			--go_opt paths=source_relative \
 			--go-grpc_out ./api \
 			--go-grpc_opt paths=source_relative \
 			./api/settings/v1/settings.proto

genOzon:
	protoc -I ./api \
			--proto_path=./ \
 			--go_out ./api \
 			--go_opt paths=source_relative \
 			--go-grpc_out ./api \
 			--go-grpc_opt paths=source_relative \
 			./api/ozon/v1/ozon.proto
