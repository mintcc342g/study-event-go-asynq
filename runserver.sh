#!/bin/bash
make build CGO_ENABLED=0 GOOS=linux GOARCH=amd64
docker-compose -f local.yml build && docker-compose -f local.yml up --remove-orphans