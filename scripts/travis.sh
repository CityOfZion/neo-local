#!/bin/bash
# Script that Travis CI will run to test the project, it is conditional on the branch.

version=$(cat ./VERSION)

set -e

if [[ $TRAVIS_BRANCH == release/* ]]
then
  make check-version
fi

make integration-tests
