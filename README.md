#serialtowebsocket

receives serial and send websockets on port 8000/:something

`./serialtowebsocket` to run

`chmod 770 serialtowebsocket` if you get a permission error


###javascript client

```
var ws = new WebSocket('ws://localhost:8000/cool')

// Received a message
ws.onmessage = function(e){
  var msg = JSON.parse(e.data)
  console.log(msg)
}
