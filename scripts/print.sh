#! /bin/bash

GREEN='\033[1;32m'
NO_COLOUR='\033[0m'

if [ "$2" == "false" ]
then
  printf "${GREEN}[neo-local]${NO_COLOUR} $1"
  exit 0
fi

printf "${GREEN}[neo-local]${NO_COLOUR} $1\n"