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
    let tx =  "1769ed122b0abb58e9a0661597f6e1c471a4937c08daca257d7ea939e415504c"
    return channel.queryTransaction(tx,peer,false, false);
}).then((query_responses) => {
    console.log("returned from query");
 // console.log(query_responses.transactionEnvelope.payload.data.actions[0].payload.action);
    console.log(JSON.stringify(query_responses))
}).catch((err) => {
    console.error("Caught Error", err);
});
