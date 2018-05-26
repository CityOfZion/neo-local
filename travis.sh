#!/bin/bash

set -e

if [[ $TRAVIS_BRANCH == 'master' ]] 
then
  make integration-tests
else
  make check-version
  make integration-tests
fi