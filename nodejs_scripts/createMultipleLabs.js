var https = require('https');
var connectionInfo = {
    hostname: 'byzanting-gateway.herokuapps.com',
    port: 80,
    path: '/api/v1/createLab',
    method: 'POST'
};


// load state/county data
var fs = require('fs');
var counties = JSON.parse(fs.readFileSync('./usStatesAndCounties.json'));
var countykeys = Object.keys(counties);

// setup data to choose from
var genders = ["Male", "Female"];
var testTypes = [
    {value:"RIDT", name:"Rapid Influenza Diagnostic Tests"},
    {value:"Rapid Molecular Assay", name:"Rapid Molecular Assay"},
    {value:"Immunofluorescence, Direct/Indirect Florescent Antibody Staining", name:"Immunofluorescence, Direct/Indirect Florescent Antibody Staining"},
    {value:"RT-PCR7 and other molecular assays", name:"RT-PCR7 and Other Molecular Assays"},
    {value:"Rapid cell culture", name:"Rapid Cell Culture"},
    {value:"Viral tissue cell culture", name:"Viral Tissue Cell Culture"}
];
var results = [
    {value:"A", name:"Influenza A"},
    {value:"B", name:"Influenza B"},
    {value:"RSV", name:"Respiratory Syncytial Virus"},
    {value:"Strep A", name:"Strep A"}
];

// utility methods
function randomNumber(max) {
    return parseInt(Math.random() * max);
}

function getRandomArrayIndex(array) {
    if (array && array.length) {
        return array[randomNumber(array.length)];
    }

    return null;
}

function getRandomStateCounty() {
    return counties[getRandomArrayIndex(countykeys)];
}

function createRandomLab () {
    var location = getRandomStateCounty();
    return {
        dateTime: new Date().toISOString(),
        gender: getRandomArrayIndex(genders),
        testType: getRandomArrayIndex(testTypes).value,
        age: randomNumber(100),
        result: getRandomArrayIndex(results).value,
        city: 'Springfield', // TODO - do we care?
        county: location.countyCode,
        state: location.stateAbbreviation
    };
}

// recursive sleep/create loop
var sleepAmount = 1500; // default to 1.5seconds
function doSleepAndSend() {
    setTimeout(function () {
            sleepAmount = Math.random() * 1000 * 5; // random amount of time, not to exceed 5 seconds.

            var data = createRandomLab();
            console.log('Creating Lab: ' + JSON.stringify(data));
            // TODO - send POST request with this payload

            console.log('Sleeping for: ' + (sleepAmount/1000) + ' seconds');
            doSleepAndSend();
    }, sleepAmount);
}

// #letsroll
doSleepAndSend();