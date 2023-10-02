const Websocket = require("ws");
const request = require('./requester');

const creds = {
    email: 'nick@mail.com',
    password: '1234'
}

const run = async function() {
    const conn = new Websocket(`ws://localhost:8000/ws?email`);

    const exit = function() {
        process.exit();
    }

    conn.on('open', async () => {
        console.log('conn open')

        let res = await request(conn, "POST:Login", {
            email: creds.email,
            password: creds.password
        });
        const { token } = res;
        
        /* DO SOMETHING WITH THE DATA */


        /* CLEANUP */

        exit();
    });

    conn.on('error', (err) => {
        console.log(`err: ${err}`)
    })
}

run();