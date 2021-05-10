#!/bin/zsh

curl -X POST http://localhost:8081/v1/todos \
  -H "Content-Type: application/json" \
  -d "{\"title\": \"curl todo\", \"Description\": \"curl description\"}" \
  -v | jq # need jq installed to pretty print json or remove the | jq

