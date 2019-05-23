/* eslint-env browser */

let pc = new RTCPeerConnection({
  iceServers: [
    {
      urls: 'stun:stun.l.google.com:19302'
    }
  ]
})
let log = msg => {
  document.getElementById('logs').innerHTML += msg + '<br>'
}

let sendChannel = pc.createDataChannel('foo')
sendChannel.onclose = () => console.log('sendChannel has closed')
sendChannel.onopen = () => console.log('sendChannel has opened')
sendChannel.onmessage = e => log(`Message from DataChannel '${sendChannel.label}' payload '${e.data}'`)

pc.oniceconnectionstatechange = e => log(pc.iceConnectionState)
pc.onicecandidate = event => {
  if (event.candidate === null) {
    document.getElementById('localSessionDescription').value = btoa(JSON.stringify(pc.localDescription))
    socket.send("Server :"+ btoa(JSON.stringify(pc.localDescription)));
  }
}

pc.onnegotiationneeded = e =>
  pc.createOffer().then(d => pc.setLocalDescription(d)).catch(log)

window.sendMessage = () => {
  let message = document.getElementById('message').value
  if (message === '') {
    return alert('Message must not be empty')
  }

  sendChannel.send(message)
}

	//var input = document.getElementById("input");
	//var output = document.getElementById("output");
	var socket = new WebSocket("ws://localhost:8080/echo");

	socket.onopen = function () {
		document.getElementById("output").innerHTML += "Status: Connected\n";
	};

	socket.onmessage = function (e) {
    document.getElementById("output").innerHTML += "Server: " + e.data + "\n";
    document.getElementById("remoteSessionDescription").value= e.data
	};

	function send() {
		socket.send(document.getElementById("input").value);
		document.getElementById("input").value = "";
	}

window.startSession = () => {
  let sd = document.getElementById('remoteSessionDescription').value
  if (sd === '') {
    return alert('Session Description must not be empty')
  }

  try {
    pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(sd))))
  } catch (e) {
    alert(e)
  }
}
