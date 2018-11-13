#!/bin/bash
#
# Use this script to query all nodes for the last block and confirm they are in sync
# Usase: watch -n .5 ./watch
#

function getBlockCount {
	echo " ----------- 1 -----------";
	curl -s -X POST http://localhost:30333 -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getversion", "params": [] }'
	echo
	curl -s -X POST http://localhost:30333 -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getblockcount", "params": [] }'
	echo
	echo " ----------- 2 -----------"
	curl -s -X POST http://localhost:30334 -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getversion", "params": [] }'
	echo
	curl -s -X POST http://localhost:30334 -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getblockcount", "params": [] }'
	echo
	echo " ----------- 3 -----------"
	curl -s -X POST http://localhost:30335 -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getversion", "params": [] }'
	echo
	curl -s -X POST http://localhost:30335 -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getblockcount", "params": [] }'
	echo
	echo " ----------- 4 -----------"
	curl -s -X POST http://localhost:30336 -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getversion", "params": [] }'
	echo
	curl -s -X POST http://localhost:30336 -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getblockcount", "params": [] }'
	echo
	echo " -------------------------"
}

export -f getBlockCount
watch -t -x -n.5 bash -c "getBlockCount"
