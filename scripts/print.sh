#! /bin/bash

GREEN='\033[1;32m'
GREY='\033[0;37m'
NO_COLOUR='\033[0m'

if [ "$1" == "prefix" ]
then
  if [ "$3" == "false" ]
  then
    printf "${GREEN}[neo-local]${NO_COLOUR} $2"
    exit 0
  fi

  printf "${GREEN}[neo-local]${NO_COLOUR} $2\n"
  exit 0
fi

if [ "$1" == "grey" ]
then
  printf "${GREY}$2${NO_COLOUR}\n"
fi
