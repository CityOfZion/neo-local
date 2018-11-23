#!/bin/bash
# Script that Travis CI will run to test the project, it is conditional on the branch.

version=$(cat ./VERSION)

set -e

# Check that the VERSION file has been bumped on release branches
if [[ $TRAVIS_BRANCH == release/* ]]
then
  make check-version
fi

# Build and test the CLI
cd ./cli
make build
make test
cd ..

# Test the Docker Compose stack
make integration-tests
