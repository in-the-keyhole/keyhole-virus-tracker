#!/bin/bash

export COMPOSE_CONVERT_WINDOWS_PATHS=1
cd labpeer
docker-compose -f  docker-compose.yml down --remove-orphans
cd ../orderer
docker-compose -f  docker-compose.yml down --remove-orphans
cd ../managechaincode
docker-compose -f  docker-compose.yml down --remove-orphans
cd ..
./startOrderer.sh
./startPeer.sh
./startChaincode.sh

# The API server has been moved to byzantine-gateway repository
#./runApiServer.sh
