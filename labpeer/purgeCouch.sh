#!/bin/bash
#
#
# Exit on first error, print all commands.
set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

docker-compose -f  docker-compose.couch.yml down --remove-orphans

