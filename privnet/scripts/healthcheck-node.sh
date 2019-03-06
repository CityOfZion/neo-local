#!/bin/bash
#
# This script tests if node is syncing correctly
#
#$NODE_NUMBER is passed as a evironment variable

maxDelayBlocks=10
actualRPCPort=`jq '.ApplicationConfiguration.RPC.Port' < /opt/node/neo-cli/config.json`
actualNode="neo-cli-privatenet-$NODE_NUMBER:$actualRPCPort"
actualNodeBlock=`curl -s -X POST http://$actualNode -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getblockcount", "params": [] }'| jq '.result'`

if [ "$actualNodeBlock" == "" ]
then
    echo "NODE IS DOWN"
    exit 503
fi

nodes=('neo-cli-privatenet-1:30333' 'neo-cli-privatenet-2:30334' 'neo-cli-privatenet-3:30335' 'neo-cli-privatenet-4:30336')

for node in "${nodes[@]}"  
do  
    block=`curl -s -X POST http://$node -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getblockcount", "params": [] }'| jq '.result'`

    if [ "$block" == "" ]
    then
        block=0
    fi

    syncDelay=`expr $block - $actualNodeBlock`

    if [ "$syncDelay" -gt "$maxDelayBlocks" ]
    then 
        echo "NODE OUT OF SYNC"
        exit 408
    fi
done

echo "ALL OK - NODE SYNCED"