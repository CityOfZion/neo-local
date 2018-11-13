#!/bin/bash
#
# Waits until a specific block (range) is reached, and then commits the Docker image.
#
SLEEP_TIME=60
CONTAINER=$1

function usage {
    echo "Usage: $0 <container name/id> $1 [--2k|--10k|--20k|--until-block <block-number-regex>]"
    echo "Default: --2k"
}

while [[ "$#" > 0 ]]; do case $2 in
    -h)
        usage
        exit 0
        ;;
    --2k)
        UNTIL_BLOCK="200[1-9]"
        shift 2
        ;;
    --10k)
        UNTIL_BLOCK="1000[1-9]"
        shift 2
        ;;
    --20k)
        UNTIL_BLOCK="2000[1-9]"
        shift 2
        ;;
    --until-block)
        SLEEP_TIME=10 # to prevent skipped blocks
        if [ -z $3 ]; then usage; exit 1; fi
        UNTIL_BLOCK=$3
        shift 3
        ;;
    *)
        usage
        exit 1
        ;;
  esac;
done

if [ -z $UNTIL_BLOCK ]; then
  UNTIL_BLOCK="200[1-9]"
fi

echo "Using container $CONTAINER and waiting until block $UNTIL_BLOCK"

while true; do
    cnt=`curl -s -X POST http://localhost:30333 -H 'Content-Type: application/json' -d '{ "jsonrpc": "2.0", "id": 5, "method": "getblockcount", "params": [] }'`
    echo $cnt
    is2k=`echo $cnt | grep "$UNTIL_BLOCK"`
    if [ $? -eq 0 ]; then
      break
    fi
    sleep $SLEEP_TIME
done

echo "Reached block target of $UNTIL_BLOCK"

echo "Claiming GAS..."
CLAIM_CMD="python3.6 /neo-python/claim_gas_fixedwallet.py"
DOCKER_CMD="docker exec -it $CONTAINER ${CLAIM_CMD}"
echo $DOCKER_CMD
echo
($DOCKER_CMD)

echo "Cleaning up the container"
CLEAN_CMD="rm -rf /neo-python/*"
DOCKER_CMD="docker exec -it $CONTAINER ${CLEAN_CMD}"
echo $DOCKER_CMD
echo
($DOCKER_CMD)

echo "Cleaning up unnecessary packages"
CLEAN_CMD="apt remove -y unzip wget curl git-core python3.6 python3.6-dev python3.6-venv python3-pip man vim"
DOCKER_CMD="docker exec -it $CONTAINER ${CLEAN_CMD}"
echo $DOCKER_CMD
echo
($DOCKER_CMD)

echo "Cleaning up extra packages"
CLEAN_CMD="apt -y autoremove"
DOCKER_CMD="docker exec -it $CONTAINER ${CLEAN_CMD}"
echo $DOCKER_CMD
echo
($DOCKER_CMD)

echo "Cleaning up apt cache packages"
CLEAN_CMD="rm -rf /var/lib/apt/lists/*"
DOCKER_CMD="docker exec -it $CONTAINER ${CLEAN_CMD}"
echo $DOCKER_CMD
echo
($DOCKER_CMD)

echo "Commit the docker container as $CONTAINER:latest image"
docker commit $CONTAINER $CONTAINER:latest

echo "Next steps:"
echo "- docker tag $CONTAINER:latest cityofzion/neo-local-privatenet:<version>"
echo "- docker push $CONTAINER:latest cityofzion/neo-local-privatenet:<version>"
