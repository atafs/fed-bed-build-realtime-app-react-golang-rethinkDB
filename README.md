# fed-bed-build-realtime-app-react-golang-rethinkDB
Chat web app that uses WebSockets

## Set Environment Variables [~./bash_profile]
> export PATH=$PATH:/usr/local/go/bin

> export GOPATH=$HOME/Repository/go

## Test the WebSocket
#### connect and send to test the handshake of the sockets:
- commit => Section5: 27

> go run main.go

http://websocket.org/echo.html

#### send a json data to a web socket:
- commit => Section5: 28 

> go run main.go

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

#### test the client for the websocket
- commit => Section5: 30

> go run client.go

#### run the app after building backend go server
-  commit => Section 37

> go build *.go

> go run *.go

https://jsbin.com/?js,console
```javascript
let ws = new WebSocket('ws://localhost:4000');

let message = {
  name: 'channel add',
  data: {
    name: 'Hardware Support'
  }
};

ws.onopen = () => {
  ws.send(JSON.stringify(message));
}

ws.onmessage = (e) => {
  console.log(JSON.parse(e.data));
}
```




