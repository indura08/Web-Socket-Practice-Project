const express = require('express');
const WebSocket  = require('ws');

const app = express();

const server = app.listen(5000, () => {
    console.log('Server is running on port 5000');
})

app.get('/', (req, res) => {
    res.send(
        `<h1>Web Socket test</h1>
        
            <script>
                const socket = new WebSocket("ws://localhost:5000");

                socket.onopen = () => {
                    let usercount = 0
                    console.log("Active user Count =", usercount++)
                }
            </script>
        `
    )
})

//create the web socket to match with the node js server
const webSocket = new WebSocket.Server({
    server
})

//detect new connection
// webSocket.on("connection", (socket) => {
//     console.log("new client connected")
// }) //this the syntax like way to do this real implementation to the project is below

let onlineUsers = 0

webSocket.on("connection", (socket) => {
    onlineUsers++;
    console.log(`Active Users online: ${onlineUsers}`)
})




