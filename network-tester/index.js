const socketClusterClient = require('socketcluster-client');

const conn = socketClusterClient.create({
    hostname: 'localhost',
    port: 8000
});
const socket = conn.transport.socket;

socket.on(conn.OPEN, () => {
    console.log('connection open')
    // Everyone listening gets this
    //conn.transmitPublish('test', {data:1})
})
socket.on(conn.CONNECTING, () => {
    console.log('connection connectiong')
})
socket.on(conn.CLOSED, () => {
    console.log('connection closed')
})
socket.on(conn.PENDING, () => {
    console.log('connection pending')
})
socket.on(conn.AUTHENTICATED, () => {
    console.log('connection authenticated')
})
