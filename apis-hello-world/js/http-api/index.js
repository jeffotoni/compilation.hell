const server = require("./server");
const host = "127.0.0.1";
const port = "8080";
server.listen(port, host, function() {
    console.log(`Server running on http://${host}:${port}/ !`);
});