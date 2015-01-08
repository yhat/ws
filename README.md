### WS

A WebSocket command line tool (mostly for testing).

Type to stdin to send along the WebSocket. WebSocket responses will be printed
to stdout.

### Install

If you have Go installed just run:

    go get github.com/yhat/ws

If you're on OS X use brew to install (no go required).

    brew install https://raw.githubusercontent.com/yhat/ws/master/ws.rb

Otherwise direct downloads are available through the [releases page](https://github.com/yhat/ws/releases).

### Usage

Given an example server.

```node
var WebSocketServer = require('ws').Server
  , wss = new WebSocketServer({port: 5000});
wss.on('connection', function(ws) {
  ws.on('message', function(message) {
    ws.send(message)
    ws.send("OVER")
  });
});
```

Let's make sure it's working with `ws`.

```
$ ws ws://127.0.0.1:5000/
That's a mighty fine websocket cli tool you've got there
That's a mighty fine websocket cli tool you've got there
OVER
```
