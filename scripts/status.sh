#! /bin/bash
# Script used within the integration tests to check that all the services are operational.

CURRENT_DIRECTORY="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
EXPECTED_NUMBER_OF_SERVICES=5

NUMBER_OF_RUNNING_SERVICES=$(docker-compose ps | grep "Up" | wc -l)

if [ $NUMBER_OF_RUNNING_SERVICES -ne $EXPECTED_NUMBER_OF_SERVICES ]
then
  docker-compose ps
  TRIMED_NUMBER_OF_RUNNING_SERVICES="$(echo -e "${NUMBER_OF_RUNNING_SERVICES}" | tr -d '[:space:]')"
  $CURRENT_DIRECTORY/print.sh prefix "FAIL! Only $TRIMED_NUMBER_OF_RUNNING_SERVICES/$EXPECTED_NUMBER_OF_SERVICES services running"
  exit 1
fi

$CURRENT_DIRECTORY/print.sh prefix "PASS! All services running as expected"
