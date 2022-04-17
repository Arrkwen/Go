gen:
	@protoc -I ./pb \
		-I ${GOPATH}/src/github.com/googleapis/ \
		-I ${GOPATH}/src/github.com/protobuf/src \
		-I ${GOPATH}/src/github.com/grpc-gateway \
		-I ${GOPATH}/src/ \
		--go_out=./api --go_opt=paths=source_relative \
		--go-grpc_out=./api --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=./api --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=./api --openapiv2_opt logtostderr=true \
		./pb/*.proto