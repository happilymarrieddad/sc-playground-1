const socketClusterClient = require('socketcluster-client');

let socket = socketClusterClient.create({
    hostname: 'localhost',
    port: 8000
});

let socket2 = socketClusterClient.create({
    hostname: 'localhost',
    port: 8001
});

(async () => {
    let myChannel = socket.subscribe('some-data');
    let myChannel2 = socket2.subscribe('some-data');
    
    (async () => {
        for await (let data of myChannel) {
            console.log('socket1',data);
        }
    })();

    (async () => {
        for await (let data of myChannel2) {
            console.log('socket2',data);
        }
    })();

    let { response, error } = await socket.invoke('request', {data:9});
    if (error != null) {
        console.log(error);
    } else {
        console.log(response);
    }

    socket.transmitPublish('some-data', {data:1});
})();


