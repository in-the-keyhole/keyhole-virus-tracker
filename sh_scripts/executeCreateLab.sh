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


# Now launch the CLI container in order to install, instantiate chaincode

docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n lab -c '{"function":"createLab","Args":["F","57","Topeka","Shawnee","KS","Mouth Swab","100","2018:07:03:10:00", "questlabs.com","active" ]}'

#docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp" cli peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n lab -c '{"function":"createLab","Args":["F","07/04/1976","Topeka","KS","Mouth Swab","100","2018:07:04:10:00", "questlabs.com" ]}'

printf "\nTotal execution time : $(($(date +%s) - starttime)) secs ...\n\n"
