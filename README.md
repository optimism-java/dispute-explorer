# dispute-explorer

This is a dispute explorer for showing dispute games information who use op stack

You can use Docker to run this.

# Prerequisites

Download and install Docker.

# 1. Run dispute-explorer-backend

# Step 1. Config Environment file

```
mv .env.template .evn
```

```
#log_format you can use console or json
LOG_FORMAT=console   

# config your mysql data source
MYSQL_DATA_SOURCE=<data-source>

# config chain name to tag your block chain name
BLOCKCHAIN=<block-chain-name>

# l1 rpc url example: eth json rpc url
L1_RPC_URL=<l1-rpc>

RPC_RATE_LIMIT=15
RPC_RATE_BURST=5

# the block number which before the first game has been created to make sure can not missing any game
FROM_BLOCK_NUMBER=6034337

# FROM_BLOCK_NUMBER block hash
FROM_BLOCK_HASH=0xafc3e42c5899591501d29649ffef0bfdec68f8d77e6d44ee00ef88cfb1a2f163

# the contract address of dispute game factory proxy
DISPUTE_GAME_PROXY_CONTRACT=0x05F9613aDB30026FFd634f38e5C4dFd30a197Fa1

API_PORT=8080
```

# Step 2. Start Dispute Game Explorer backend service

use docker-compose to run this service

```
cd deploy
docker-compose -f docker-compose.yml up -d
```

Now, this project is running now.

Tip: if you just need a backend service to collect all data, Run Step 1 and Step 2.

# Step 3. Run meiliSearch 

Run a meiliSearch service to sync MySql data for front service to search quickly

```
cd deploy
docker-compose -f docker-compose-meiliSearch.yml up -d meiliSearch
```

Now, meiliSearch is running. 

# Step 4. Run meiliSync

first, we need to find the api_key of meiliSearch 

```
curl -H "Authorization: Bearer <Token>" http://localhost:port/keys
```

You should get a result, similar to :

```json
{
	"results": [{
		"name": "Default Search API Key",
		"description": "Use it to search from the frontend",
		"key": "d09536ef1e2742b4792c607465dc169f659f1b2dcb0107bfdce2542b602ed534",
		"uid": "675ff658-9e73-460c-a3be-c6fcee624edf",
		"actions": ["search"],
		"indexes": ["*"],
		"expiresAt": null,
		"createdAt": "2024-08-06T08:47:38.225365511Z",
		"updatedAt": "2024-08-06T08:47:38.225365511Z"
	}, {
		"name": "Default Admin API Key",
		"description": "Use it for anything that is not a search operation. Caution! Do not expose it on a public frontend",
		"key": "abc40e8457b32aa86d20ab0db0b42a86298b253209c4c31d9936b378e686d132",
		"uid": "db1499f6-59a1-42c7-a13a-e18e191f456c",
		"actions": ["*"],
		"indexes": ["*"],
		"expiresAt": null,
		"createdAt": "2024-08-06T08:47:38.225052792Z",
		"updatedAt": "2024-08-06T08:47:38.225052792Z"
	}],
	"offset": 0,
	"limit": 20,
	"total": 2
}
```

And, use `Default Admin API Key`, `key` to update config.yml

```
meilisearch:
  api_url: http://localhost:7701  
  api_key: abc40e8457b32aa86d20ab0db0b42a86298b253209c4c31d9936b378e686d132
```

launch the meiliSync service

```
cd deploy
docker-compose -f docker-compose-meiliSearch up -d meiliSync
```

# Step 5. Validate meiliSync Service

We can visit meiliSearch api to validate meiliSync service. more [meiliSearch docs](https://www.meilisearch.com/docs/reference/api/overview)

```
curl -H "Authorization: Bearer <Token>" http://localhost:port/indexes
```

You should get a result, similar to :

```json
{
    "results": [
        {
            "uid": "disputegame",
            "createdAt": "2024-08-06T09:24:24.640693956Z",
            "updatedAt": "2024-08-07T07:02:32.402360903Z",
            "primaryKey": "id"
        },
        {
            "uid": "gameclaim",
            "createdAt": "2024-08-06T09:24:24.670117944Z",
            "updatedAt": "2024-08-07T07:02:28.94487306Z",
            "primaryKey": "id"
        },
        {
            "uid": "gamecredit",
            "createdAt": "2024-08-06T10:37:42.013472322Z",
            "updatedAt": "2024-08-07T07:02:32.379350451Z",
            "primaryKey": "id"
        },
        {
            "uid": "syncevents",
            "createdAt": "2024-08-06T09:24:24.696318772Z",
            "updatedAt": "2024-08-07T07:02:30.382386632Z",
            "primaryKey": "id"
        }
    ],
    "offset": 0,
    "limit": 20,
    "total": 4
}
```

If you get information like this, it means our deploy it`s success.

