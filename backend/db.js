// db.js
const postgres = require('postgres');

module.exports = postgres({
    host: 'postgres',
    port: 5432,
    pass: 'postgres',
    username: 'postgres',
    publications: 'alltables',
});