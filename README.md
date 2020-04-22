# Keyhole Virus Tracker Blockchain

This project implements a HyperLedger blockchain network with chaincode that manages a ledger of COVID 19 and Influenza Virus lab tests. The chaincode implements functions to create and retrieve Influenza test results.  

It also references a ReactJS client project and supporting API gateway project that provides a user interface to interact with a deployed blockchain.

The instructions will start a Hyperledger network locally on a Linux/Unix/MacOs/Windows operating system and then invoke and access the blockchain chaincode in the following ways:

##### * Start the HyperLedger Orderer, Certificate Authority (CA), and Peer Nodes; create channel "mychannel"; and install chaincode
##### * Interact with Chaincode from a ReactJS/Node Web Application or CLI and Node JavaScript Commands.
##### * Execute Chaincode and Unit tests from CLI 

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Keyhole Virus Tracker Stack Setup](#byzantine-flu-full-stack-setup)
- [Installing and Running](#installing-and-running)
- [Start the Network](#start-the-network)
- [Execute Chaincode on Network Using Network CLI](#execute-chaincode-on-network-using-network-cli)

## Requirements
* [Node](https://nodejs.org/en/download/) 8.9.x to 10.x - **Note: We have seen an issue with node-gyp when using > 10.x**
* [Python](https://www.python.org/downloads/) 2.7+ (v3+ not supported) - **Note: for Windows OS Only!**
* [XCode](https://apps.apple.com/us/app/xcode/id497799835?mt=12) or type `xcode-select --install` - **Note: for OSX Only!**

----
## Keyhole Virus Stracker Full Stack Setup

Follow these steps to get a ReactJS UI and API Gateway for the blockcchain installed and running locally.

#### Setup Steps
1. **-> (You are here)** Set up and run Byzantine Hyperledger Fabric

2. Set up and run the  API Gateway:  https://github.com/hyperledger-labs/keyhole-fabric-api-gateway

    - The communication gateway to the Byzantine Hyperledger Fabric runtime

3. Set up and run the Reactjs UI:  https://github.com/in-the-keyhole/keyhole-virus-tracker-ui

    - A website containing a map displaying the locations and concentrations of reported flu samples


#### Optional Steps:
4. Hyperledger Brower:  https://github.com/in-the-keyhole/byzantine-browser

    - A website showing the actual blockchain and the associated metadata 
-----

# Installing and Running 

Prerequisite: Docker must be installed.

* Clone repo 

* Open terminal and execute shell script below to install Hyperledger Docker images. 

```
$ ./fabric-preload.sh
```
* Verify Docker image installation(s). 

```sh
| => docker images | grep hyperledger
hyperledger/fabric-ca           1.2.0           66cc132bd09c    2 months ago    252MB
hyperledger/fabric-ca           latest          66cc132bd09c    2 months ago    252MB
hyperledger/fabric-tools        1.2.0           379602873003    2 months ago    1.51GB
hyperledger/fabric-tools        latest          379602873003    2 monthsago     1.51GB
hyperledger/fabric-ccenv        1.2.0           6acf31e2d9a4    2 months ago    1.43GB
hyperledger/fabric-ccenv        latest          6acf31e2d9a4    2 months ago    1.43GB
hyperledger/fabric-orderer      1.2.0           4baf7789a8ec    2 months ago    152MB
hyperledger/fabric-orderer      latest          4baf7789a8ec    2 months ago    152MB
hyperledger/fabric-peer         1.2.0           82c262e65984    2 months ago    159MB
hyperledger/fabric-peer         latest          82c262e65984    2 months ago    159MB
hyperledger/fabric-couchdb      latest          3092eca241fc    2 months ago    1.61GB
```
* Install NPM modules - execute the following command in this directory to download and install all of the required npm modules (per the package.json file):

```
$ npm install
```

# Start the network
```
$ ./start.sh

```
This starts the hyperledger fabric network, creates a channel, installs chaincode, and executes the chaincode to verify the network is up and running.

Verify that that output ends with something similar to:
```
...
...
2019-10-28 20:56:37.490 UTC [chaincodeCmd] chaincodeInvokeOrQuery -> INFO 04f Chaincode invoke successful. result: status:200
```

If the script fails due to firewall blocking the ports: 


> Add a firewall rule to allow incoming connection on ports 139 and 445 from IP 10.0.75.0 Subnet 255.255.255.0 (or whatever your Docker config says)


# Execute Chaincode on Network Using Network CLI

* Execute the `queryAllLabs` chaincode function with the following shell script:

```
    $./executeQueryLabs.sh 
```
This is a canned query. Verify the output ends with something similar to 
```
[chaincodeCmd] chaincodeInvokeOrQuery -> DEBU 0a8 ESCC invoke result: version:1 response:<status:200 payload:"
```

* Execute the `createLab` chaincode function with the following shell script:

```
    $./executeCreateLab.sh 
```
verify the output ends  with something similar to 
```
[chaincodeCmd] chaincodeInvokeOrQuery -> DEBU 04f ESCC invoke result: version:1 response:<status:200 payload:"Lab Created"


If the network is not up, start it
```
$ ./start.sh 
```
To run the UI, follow the setup steps in https://github.com/in-the-keyhole/keyhole-virus-tracker-ui


# Compiling and Unit Testing Go Chaincode with the Development CLI 

Chaincode is implemented using the Go Language, which is what Hyperledger is built with. Here's how a Docker-based Go development environment can be started with chaincode that was developed and tested outside of Hyperledger. 

Chaincode (i.e Go implementation) can be found at this location: `chaincodes/lab/lab.go`.

* Stop and remove running Hyperledger Docker instances with following commands:

```
    $ docker stop $(docker ps -a -q) 
    $ docker rm $(docker ps -a -q) 
```

* Change to the `dev` directory and run the following Docker command:

```
$  docker-compose -f docker-compose-go.yaml up
```

* Open new terminal window, navigate to khs-lab-results-blockchain, and issue the following command:

```
$ docker exec -it chaincode bash
```

* Execute unit tests for `labs.go` chaincode bash with the following commands:

```
root@baf0e90e82a6:/opt/gopath/src/chaincode# cd lab
root@baf0e90e82a6:/opt/gopath/src/chaincode# go test 
```

* To compile and build the `labs.go` chaincode, issue the following commands:

```
root@baf0e90e82a6:/opt/gopath/src/chaincode# go build
```
