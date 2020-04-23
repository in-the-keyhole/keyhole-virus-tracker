'use strict';

/*
 * Hyperledger Fabric Sample Query Program
 */

var hfc = require('fabric-client');
var path = require('path');

var options = {
    wallet_path: path.join(__dirname, '../hfc-key-store'),
    user_id: 'PeerAdmin',
    channel_id: 'mychannel',
    chaincode_id: 'lab',
    network_url: 'grpc://localhost:7051'
};

var channel = {};
var client = null;
var peer = null;
Promise.resolve().then(() => {
    console.log("Create a client and set the wallet location");
    client = new hfc();
    return hfc.newDefaultKeyValueStore({ path: options.wallet_path });
}).then((wallet) => {
    console.log("Set wallet path, and associate user ", options.user_id, " with application");
    client.setStateStore(wallet);
    return client.getUserContext(options.user_id, true);
}).then((user) => {
    console.log("Check user is enrolled, and set a query URL in the network");
    if (user === undefined || user.isEnrolled() === false) {
        console.error("User not defined, or not enrolled - error");
    }
    channel = client.newChannel(options.channel_id);
    peer = client.newPeer(options.network_url);
    channel.addPeer(peer);
    return;
}).then(() => {
    console.log("Query TX");
    let tx =  "fef7952cc9b140e18fcfbe4c108c6ff169e3a4754fa629906db2a1e67f576dbe"
    return channel.queryTransaction(tx,peer,false, false);
}).then((query_responses) => {
    console.log("returned from query");
 // console.log(query_responses.transactionEnvelope.payload.data.actions[0].payload.action);
    console.log( query_responses )
}).catch((err) => {
    console.error("Caught Error", err);
});
