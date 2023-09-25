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
    let myChannel = socket.subscribe('create-users');
    let myChannel2 = socket2.subscribe('create-users');
    
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
    
    (async () => {
        for await (let data of socket.listener('raw')) {
            console.log('raw',data);
        }
    })()

    // let { response, error } = await socket.invoke('request', {data:9});
    // if (error != null) {
    //     console.log(error);
    // } else {
    //     console.log(response);
    // }

    // socket.transmitPublish('some-data', {data:1});
})();


