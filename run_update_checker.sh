#!/bin/bash
docker run --rm --env-file update_checker.env -v $(pwd):/root/src -w /root/src golang:1.17-alpine sh -c "go run update_checker.go"