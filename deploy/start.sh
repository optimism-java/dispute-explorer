#!/bin/bash

# Install yq if not already installed
if ! [ -x "$(command -v yq)" ]; then
  echo 'Error: yq is not installed. Installing yq...' >&2
  if [ -x "$(command -v brew)" ]; then
    brew install yq
  elif [ -x "$(command -v apt)" ]; then
    sudo apt update
    sudo apt install yq
  else
    echo 'Error: Package manager not found. Please install yq manually.' >&2
    exit 1
  fi
fi

docker-compose -f docker-compose.yml up -d

sleep 5

# Bearer Token
bearer_token="123456"

curl_response=$(curl -s -H "Authorization: Bearer $bearer_token" http://localhost:7700/keys)
new_api_key=$(echo "$curl_response" | jq -r '.results[] | select(.actions[] | contains("*")) | .key')

yq e ".meilisearch.api_key = \"$new_api_key\"" -i config.yml

echo "api_key was successfully updated to: $new_api_key"
