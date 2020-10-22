const url = require('url');

module.exports = {
    "/hello-world": function(request, response) {
        const requestUrl = url.parse(request.url, true);
        const name = requestUrl.query.name || "world";

        response.statusCode = 200;
        response.setHeader("Content-Type", "text/plain");
        response.end(`Hello ${name}`);
    },
    "invalid": function(request, response) {
        response.statusCode = 404;
        response.setHeader("Content-Type", "text/plain");
        response.end("Invalid Request!");
    },
}