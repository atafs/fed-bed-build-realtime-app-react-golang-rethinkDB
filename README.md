# fed-bed-build-realtime-app-react-golang-rethinkDB
Chat web app that uses WebSockets

## Set Environment Variables [~./bash_profile]
> export PATH=$PATH:/usr/local/go/bin

> export GOPATH=$HOME/Repository/go

## Test the WebSocket
- connect and send to test the handshake of the sockets:
http://websocket.org/echo.html

- send a json data to a web socket:
https://jsbin.com/?js,console
```javascript
let msg = {
  name: 'channel add',
  data: {
    name: 'Hardware Support'
  }
}

let subMsg = {
  name: 'channel subscribe'
}

let ws = new WebSocket('ws://localhost:4000');
ws.onopen = () => {
  ws.send(JSON.stringify(subMsg))
  ws.send(JSON.stringify(msg));
}

ws.onmessage = (e) => {
  console.log(e.data);
}
```



