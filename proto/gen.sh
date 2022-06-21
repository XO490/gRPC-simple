#!/bin/bash

go get github.com/googleapis/googleapis

googleapis_ver=latest
googleapis_pkg=github.com/googleapis/googleapis
googleapis_path=$(go list -m -f '{{.Dir}}' ${googleapis_pkg}@${googleapis_ver})

protoc -I. \
    -I "${googleapis_path}" \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=. --grpc-gateway_opt paths=source_relative \
    --descriptor_set_out=proto \
    *.proto




