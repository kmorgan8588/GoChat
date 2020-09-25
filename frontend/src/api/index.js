const socket = new WebSocket("ws://localhost:8080/ws");

const connect = () => {
    console.log("Attempting connection...");

    socket.onopen = () => {
        console.log("Successfully connected");
    };

    socket.onmessage = message => {
        console.log(message);
    };

    socket.onclose = event => {
        console.log("Socket closed connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket error: ", error);
    };

};

const sendMessage = message => {
    console.log("sending message: ", message);
    socket.send(message);
};

module.exports = { connect, sendMessage };
