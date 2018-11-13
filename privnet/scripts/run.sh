#!/bin/bash
#
# This script starts consensus node and waits forever
#
#$NODE_NUMBER is passed as a evironment variable

# copy the correct config to be loaded
cp /opt/node/neo-cli/config$NODE_NUMBER.json /opt/node/neo-cli/config.json

case "$NODE_NUMBER" in
1)
  echo "case $NODE_NUMBER"
  screen -dmS node1 expect /opt/start_consensus_node.sh /opt/node/neo-cli/ wallet1.json one
  ;;
2)
  echo "case $NODE_NUMBER"
  screen -dmS node2 expect /opt/start_consensus_node.sh /opt/node/neo-cli/ wallet2.json two
  ;;
3)
  echo "case $NODE_NUMBER"
  screen -dmS node3 expect /opt/start_consensus_node.sh /opt/node/neo-cli/ wallet3.json three
  ;;
4)
  echo "case $NODE_NUMBER"
  screen -dmS node4 expect /opt/start_consensus_node.sh /opt/node/neo-cli/ wallet4.json four
  ;;
*)
  Message="No NODE_NUMBER environment variable or out of range !!"
  echo $Message
  ;;
esac

sleep infinity
