generate:
	@git clone https://github.com/Chiyoku/PapierSchema
	@mkdir -p schema
	@cp PapierSchema/proto/auth.proto ./schema/
	@mkdir -p internal/gen
	@protoc -I=./ \
        --go-grpc_out=require_unimplemented_servers=false:./internal/gen \
        --go_out=./internal/gen \
        ./schema/*.proto
	@rm -rf PapierSchema
	@rm -rf schema
