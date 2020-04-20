#!/bin/bash
# Copyright London Stock Exchange Group All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
set -e


 docker-compose -f docker-compose-go.yaml up

 cd ..

 docker exec -it chaincode bash
