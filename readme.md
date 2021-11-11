
## Install
 - go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
 - go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
 - go get -u google.golang.org/grpc
 - go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
 - go get google.golang.org/api
 - go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

## Create a proto file

 - `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/user.proto`

## Python client
 - sudo pip install grpcio
 - sudo  pip install grpcio-tools
 - pip install protobuf
 - python3 -m grpc_tools.protoc -I./proto --python_out=. --grpc_python_out=. ./proto/user.proto

````
import grpc
import user_pb2_grpc as pb2_grpc
import user_pb2 as pb2

channel = grpc.insecure_channel('localhost:50051')
stub = pb2_grpc.UsersStub(channel)

feature = stub.GetUsers(pb2.User())
print(feature)
 ````