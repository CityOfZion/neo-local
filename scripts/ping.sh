#! /bin/bash
# Script to check and wait until the NeoScan service is operational. 

while [ "$(docker inspect --format '{{json .State.Health.Status }}' neo-local_autoheal)" != "\"healthy\"" ]
  do
    printf '.'
    sleep 3
  done

while [ "$(docker inspect --format '{{json .State.Health.Status }}' neo-scan-api)" != "\"healthy\"" ]
  do
    printf '.'
    sleep 3
  done

while [ "$(docker inspect --format '{{json .State.Health.Status }}' notifications-server)" != "\"healthy\"" ]
  do
    printf '.'
    sleep 3
  done

while [ "$(docker inspect --format '{{json .State.Health.Status }}' neo-faucet)" != "\"healthy\"" ]
  do
    printf '.'
    sleep 3
  done

printf "\n"
