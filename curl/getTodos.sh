#!/bin/zsh

curl http://localhost:8081/v1/todos -v | jq # need jq installed to pretty print json or remove the | jq
  # -H "Content-Type: application/json" \
  # -d "{\"title\": \"curl todo\", \"Description\": \"curl description\"}" \

