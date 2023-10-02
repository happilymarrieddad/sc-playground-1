const { uuid } = require('uuidv4');

let request =  function(conn, action, reqData = {}, token) {
    return new Promise((resolve, reject) => {
        const id = uuid();

        const handler = (raw) => {
            const res = JSON.parse(raw);
            
            if (res && res.id == id) {
                if (parseInt(res.status / 100) == 2) {
                    resolve(res.data || res.pagedData);
                } else if (parseInt(res.status / 100) == 4) {
                    console.error(action, res.data);
                    process.exit(1)
                }
                conn.off('message', handler); // there's probably a better way...
            } else {
                // if (res.action == 'HEARTBEAT') {
                //     return
                // }
            }
        }

        conn.on('message', handler)

        conn.send(JSON.stringify({
            id: id,
            action: action,
            data: reqData,
            token,
        }))
    })
}

module.exports = request;