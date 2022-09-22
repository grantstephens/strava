#!/bin/sh

java -jar swagger-codegen-cli.jar generate -l go -o tmp -i api/swagger.yaml -c api/config.json 

go mod tidy

gofumpt -w .
