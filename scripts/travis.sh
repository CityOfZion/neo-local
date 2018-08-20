#!/bin/bash
# Script that Travis CI will run to test the project, it is conditional on the branch.

version=$(cat ./VERSION)

set -e

if [[ $TRAVIS_BRANCH == 'master' || $TRAVIS_TAG == $version ]]
then
  make integration-tests
else
  make check-version
  make integration-tests
fi
