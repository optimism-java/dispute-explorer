# dispute-explorer

This is a dispute explorer for displaying information about dispute games that utilize the OP Stack.

https://www.superproof.wtf/

Superproof is an convenient tool for blockchain developers working with the OP stack, as well as for individuals eager to engage in dispute games.

For developers, Superproof provides real-time data on dispute games, allowing them to monitor key metrics such as the total games played and win numbers for both "Attackers" and "Defenders." This visibility is easy for understanding game dynamics on the blockchain.

For users, Superproof simplifies participation in the fault-proof module, making it easy to invoke game methods like "attack" and "defend." It also facilitates obtaining honest claims, ensuring transparency in the gaming process. Additionally, the platform offers clear visualizations that detail each step in the dispute process, enhancing user comprehension.

In summary, Superproof not only helps developers easily access vital information but also streamlines user participation in dispute games, benefiting all involved in fault-proof interactions.

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
L1_RPC_URLS=<l1-rpc>
# l2 rpc url example: op json rpc url
L2_RPC_URLS=<l2-rpc>

NODE_RPCURLS=<op-node-rpc>

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

# Step 3. Run the deployment script

Run the script to launch the service

```
cd deploy
./star.sh
```

Now, this project is running.

# Step 4. Validate meiliSync Service

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

If you get information like this, it means our deployment it`s success.
