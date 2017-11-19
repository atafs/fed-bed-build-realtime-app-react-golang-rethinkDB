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

#### database rethinkDB web admin
http://localhost:8080/#dataexplorer

REQL (Rethink Query Language):
```
r.dbList()
r.dbCreate('rtsupport')
r.db('rtsupport').tableCreate('channel')
r.db('rtsupport').tableList()
r.db('rtsupport').table('channel').insert({ name: 'Hardware Support' })
r.db('rtsupport').table('channel').indexCreate('name')
r.db('rtsupport').table('user').get('3aed4c07-6248-4d2e-bec9-079f045b8bb0')
r.db('rtsupport').table('user').get('3aed4c07-6248-4d2e-bec9-079f045b8bb0').update({ name: 'James Moore'})

r.db('rtsupport').table('message').insert({
  author: 'Americo Tomas',
  createAt: r.now(),
  channelId: '452e5e81-232b-4c17-85fe-7de014a4cb6a'
})

r.db('rtsupport').table('message').indexCreate('createAt')
r.db('rtsupport').table('message').get('14db8630-e351-47cf-a2ea-3551026197f3').delete()

//change feed
//one tab at the web browser
r.db('rtsupport').table('channel').changes({ includeInitial: true }) 
//another tab
r.db('rtsupport').table('channel').insert({ name: 'Software Support' })


```





