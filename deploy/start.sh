#!/bin/bash

docker-compose -f docker-compose.yml up -d

sleep 5

# Bearer Token
bearer_token="123456"

curl_response=$(curl -s -H "Authorization: Bearer $bearer_token" http://localhost:7700/keys)
new_api_key=$(echo "$curl_response" | jq -r '.results[] | select(.actions[] | contains("*")) | .key')

yq e ".meilisearch.api_key = \"$new_api_key\"" -i config.yml

echo "api_key was successfully updated to: $new_api_key"
