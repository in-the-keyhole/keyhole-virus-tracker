#!/bin/bash

docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker volume prune --force

docker rmi $(docker images -aq *example.com*) --force

docker ps -a