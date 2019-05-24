
  //############################# Message Channel

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


  //############################# Web Socket for Server

	//var input = document.getElementById("input");
	//var output = document.getElementById("output");
  var socket = new WebSocket("ws://localhost:8080/socketread");
  var socketW = new WebSocket("ws://localhost:8080/socketwrite");

	socket.onopen = function () {
		document.getElementById("output").innerHTML += "Status: Connected\n";
  };
  
  socketW.onopen = function () {
		document.getElementById("output").innerHTML += "Status: Connected\n";
	};

	/*socket.onmessage = function (e) {
    document.getElementById("output").innerHTML += "Server: " + e.data + "\n";
    document.getElementById("remoteSessionDescription").value= e.data
  };*/
  
  socketW.onmessage = function (e) {
    document.getElementById("output").innerHTML += "Server: " + e.data + "\n";
    document.getElementById("remoteSessionDescription").value= e.data
	};

	function send() {
		socket.send(document.getElementById("input").value);
		document.getElementById("input").value = "";
  }

  //############################# Video Streaming
  
  let pc2 = new RTCPeerConnection({
    iceServers: [
      {
        urls: 'stun:stun.l.google.com:19302'
      }
    ]
  })
  
  
  pc2.ontrack = function (event) {
    var el = document.createElement(event.track.kind)
    el.srcObject = event.streams[0]
    el.autoplay = true
    el.controls = true

    if (event.track.kind==="audio") {
      document.getElementById('remoteAudio').appendChild(el)
    }

    if (event.track.kind==="video") {
      document.getElementById('remoteVideos').appendChild(el)
    }
 
  }
  
  pc2.oniceconnectionstatechange = e => log(pc2.iceConnectionState)
  pc2.onicecandidate = event => {
    if (event.candidate === null) {
      document.getElementById('localSessionDescription2').value = btoa(JSON.stringify(pc2.localDescription))
    }
  }
  
  // Offer to receive 1 audio, and 2 video tracks
  pc2.addTransceiver('audio', {'direction': 'sendrecv'})
  pc2.addTransceiver('video', {'direction': 'sendrecv'})
 
  pc2.createOffer().then(d => pc2.setLocalDescription(d)).catch(log)

//############################# Start Session
window.startSession = () => {
  let sd = document.getElementById('remoteSessionDescription').value
  let sd2 = document.getElementById('remoteSessionDescription2').value
  if (sd === '') {
    return alert('Session Description must not be empty')
  }

  if (sd2 === '') {
    return alert('Session Description must not be empty')
  }

  try {
    pc.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(sd))))
    pc2.setRemoteDescription(new RTCSessionDescription(JSON.parse(atob(sd2))))
  } catch (e) {
    alert(e)
  }
}
