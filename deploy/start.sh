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

# Check if curl is installed
if ! [ -x "$(command -v curl)" ]; then
  echo 'Error: curl is not installed. Please install curl.' >&2
  exit 1
fi

# Start Docker Compose
docker-compose -f docker-compose.yml up -d

sleep 5

# Bearer Token
bearer_token="123456"

# Send request to get API key
curl_response=$(curl -s -H "Authorization: Bearer $bearer_token" http://localhost:7700/keys)
new_api_key=$(echo "$curl_response" | jq -r '.results[] | select(.actions[] | contains("*")) | .key')

sed -i 's/\(api_key:\s*\).*/\1new_api_key_value/' config.yml

echo "api_key was successfully updated to: $new_api_key"

sortable_attributes='["block_number"]'
sortEventUrl="http://localhost:7700/indexes/syncevents/settings/sortable-attributes"
sortEventRsp=$(curl -X PUT $sortEventUrl -H "Content-Type: application/json" -H "Authorization: Bearer $bearer_token" -d "$sortable_attributes")
echo "$sortEventRsp"

sortGamesUrl="http://localhost:7700/indexes/disputegames/settings/sortable-attributes"
sortGamesRsp=$(curl -X PUT $sortGamesUrl -H "Content-Type: application/json" -H "Authorization: Bearer $bearer_token" -d "$sortable_attributes")
echo "$sortGamesRsp"

maxEventTotalHits='{ "maxTotalHits": 100000 }'
totalEventHitsUrl="http://localhost:7700/indexes/syncevents/settings/sortable-attributes"
maxEventsRsp=$(curl -X PATCH $totalEventHitsUrl -H "Content-Type: application/json" -H "Authorization: Bearer $bearer_token" -d "$maxEventTotalHits")
echo "$maxEventsRsp"

maxGamesTotalHits='{ "maxTotalHits": 100000 }'
totalGamesHitsUrl="http://localhost:7700/indexes/disputegames/settings/sortable-attributes"
maxGamesRsp=$(curl -X PATCH $totalGamesHitsUrl -H "Content-Type: application/json" -H "Authorization: Bearer $bearer_token" -d "$maxGamesTotalHits")
echo "$maxGamesRsp"

setFilter='["id","block_number","get_len_status","claim_data_len"]'
setFilterUrl="http://localhost:7700/indexes/disputegames/settings/filterable-attributes"
setFilterRsp=$(curl -X PUT $setFilterUrl -H "Content-Type: application/json" -H "Authorization: Bearer $bearer_token" -d "$setFilter")
echo "$setFilterRsp"
