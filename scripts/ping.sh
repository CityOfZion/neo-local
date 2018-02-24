#! /bin/bash

until $(curl --output /dev/null --silent --head --fail http://localhost:4001); do
    printf '.'
    sleep 3
done

printf "\n"