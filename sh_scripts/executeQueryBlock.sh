#!/bin/bash
#
#
# Exit on first error
set -e

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

starttime=$(date +%s)

if [ ! -d ~/.hfc-key-store/ ]; then
	mkdir ~/.hfc-key-store/
fi
cp ../creds/* ~/.hfc-key-store/
# launch network; create channel and join peer to channel
#cd ./managechaincode
#./start.sh


docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer channel fetch newest -o orderer.example.com:7050 -c mychannel 


printf "\nTotal execution time : $(($(date +%s) - starttime)) secs ...\n\n"
