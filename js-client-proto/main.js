/// <reference path="./build/gen/MyMessages.js"/>

window.addEventListener("load", function (evt) {
    var output = document.getElementById("output");
    var inputName = document.getElementById("inputName");
    var inputCharacteristics = document.getElementById("inputCharacteristics");
    var inputHeight = document.getElementById("inputHeight");
    var inputAddresses = document.getElementById("inputAddresses");
    var inputPolRecord = document.getElementById("inputPolRecord");
    var inputGender = document.getElementById("inputGender");

    var inputYears = document.getElementById("inputYears");
    var queryByName = document.getElementById("queryByName");
    var ws;
    var print = function (message) {
        var d = document.createElement("div");
        d.innerText = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function (evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("ws://127.0.0.1:8081/ws");
        ws.binaryType = 'arraybuffer';
        ws.onopen = function (evt) {
            print("OPEN");
        }
        ws.onclose = function (evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function (evt) {
            let arrayBuffer = event.data;
            let message = proto.QueryPersonResult.deserializeBinary(arrayBuffer);

            console.log(message.toObject());
            let persons_retrieved = message.getPersonsList()
            print("QUERY RESULT TOTAL:" + persons_retrieved.length);
            for (let i = 0; i < persons_retrieved.length; i++) {
                print("QUERY RESULT:");
                print(prettyStringPerson(persons_retrieved[i]));
            }

        }
        ws.onerror = function (evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function (evt) {
        if (!ws) {
            return false;
        }

        let message = new proto.ClientEvent();
        let upsertEvent = new proto.UpsertPerson();
        let newPerson = new proto.Person();
        newPerson.setName(inputName.value);
        newPerson.setCharacteristicsList(inputCharacteristics.value.split(";"));
        newPerson.setYears(inputYears.value);
        newPerson.setHeight(inputHeight.value);
        newPerson.setKnownAddressesList(getAddressesList(inputAddresses.value));
        newPerson.setGender(proto.Gender[inputGender.value]);
        newPerson.setHasPoliceRecord((inputPolRecord.value === 'true'));

        upsertEvent.setPerson(newPerson);
        message.setUpsertPerson(upsertEvent);
        print("SENT: " + message.toString());
        ws.send(message.serializeBinary(), { binary: true });
        return false;
    };
    document.getElementById("close").onclick = function (evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

    document.getElementById("query").onclick = function (evt) {
        if (!ws) {
            return false;
        }
        let message = new proto.ClientEvent();
        let queryEvent = new proto.QueryPerson();
        queryEvent.setNameRegexp(queryByName.value);
        message.setQueryPerson(queryEvent);
        print("SENT QUERY: " + queryEvent.toString());
        ws.send(message.serializeBinary(), { binary: true });
        return false;
    };
});

//receives: string ; separated in the form street|number|city|country
//resturn list of Addresses
function getAddressesList(addressesString){
    let result = [];

    let addressesList = addressesString.split(";");
    for (let i = 0; i < addressesList.length; i++) {
        result.push(toAddressProto(addressesList[i]));
    }
    return result;
}

function toAddressProto(addressString){
    let addressFields = addressString.split("|");
    let addressProto = new proto.Address();
    
    addressProto.setStreet(addressFields[0]);
    addressProto.setNumber(addressFields[1]);
    addressProto.setCity(addressFields[2]);
    addressProto.setCountry(addressFields[3]);

    return addressProto;
}

/**
 * @param {proto.Person} person - The person
 */
function prettyStringPerson(person){
    let result = "";
    result += "Name: " + person.getName() + "\n";
    result += "Characteristics: " + person.getCharacteristicsList() + "\n";
    result += "Height: " + person.getHeight() + "\n";
    result += "Years: " + person.getYears() + "\n";
    let addresses = person.getKnownAddressesList();
    for (let i = 0; i < addresses.length; i++) {
        result += "Address: "+ i + " - " + addresses[i] + "\n";
    }
    result += "Police record: " + person.getHasPoliceRecord() + "\n";
    result += "Gender: " + person.getGender() + "\n";
    return result;
}