#!/bin/bash
# Script that Travis CI will run to test the project, it is conditional on the branch.

set -e

if [[ $TRAVIS_BRANCH == 'master' ]] 
then
  make integration-tests
else
  make check-version
  make integration-tests
fi