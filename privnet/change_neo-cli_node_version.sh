#!/bin/bash
#
CONTAINER=$1
NEO_CLI_VERSION=$2

docker exec -it $CONTAINER bash -c 'pkill screen'
docker exec -it $CONTAINER bash -c 'rm -fr /var/lib/apt/lists/*'
docker exec -it $CONTAINER bash -c 'apt-get clean'
docker exec -it $CONTAINER bash -c 'apt-get update -o Acquire::CompressionTypes::Order::=gz'
docker exec -it $CONTAINER bash -c 'apt-get install wget unzip -y'
docker exec -it $CONTAINER bash -c 'mv /opt/node/neo-cli/ /opt/node/neo-cli_old'
docker exec -it $CONTAINER bash -c "wget -O /opt/neo-cli.zip https://github.com/neo-project/neo-cli/releases/download/v$NEO_CLI_VERSION/neo-cli-linux-x64.zip || echo \"Error downloading neo-cli\" && exit"
docker exec -it $CONTAINER bash -c 'unzip -q -d /opt/node /opt/neo-cli.zip'
docker exec -it $CONTAINER bash -c "wget -O /opt/SimplePolicy.zip https://github.com/neo-project/neo-plugins/releases/download/v2.9.3/SimplePolicy.zip || echo \"Error downloading neo-plugins\" && exit"
docker exec -it $CONTAINER bash -c 'unzip -q -d /opt/node/neo-cli /opt/SimplePolicy.zip'
docker exec -it $CONTAINER bash -c 'rm -fr /opt/node/neo-cli/Index_0000DDB1'
docker exec -it $CONTAINER bash -c 'rm -fr /opt/node/neo-cli/Chain_0000DDB1'
docker exec -it $CONTAINER bash -c 'cp /opt/node/neo-cli_old/wallet*.json /opt/node/neo-cli/'
docker exec -it $CONTAINER bash -c 'cp /opt/node/neo-cli_old/protocol.json /opt/node/neo-cli/'
for f in configs/v$NEO_CLI_VERSION/*.json; do echo $f && docker cp $f $CONTAINER:/opt/node/neo-cli/; done
docker exec -d $CONTAINER bash -c /opt/run.sh
docker exec -it $CONTAINER bash -c 'rm /opt/neo-cli.zip'
docker exec -it $CONTAINER bash -c 'rm /opt/SimplePolicy.zip'
docker exec -it $CONTAINER bash -c 'rm -fr /opt/node/neo-cli_old'
