
// Port is currently hard-coded.
let conn = new WebSocket('ws://' + location.hostname + ':8082');
let input = document.getElementById("input");
let messages = document.getElementById("messages")

// onopen - Do nothing for now ...
conn.onopen = function () {

};

// Log errors
conn.onerror = function (error) {
    console.log('WebSocket Error ' + error);
};

// Handle incoming messages.
conn.onmessage = function (e) {
    let o = JSON.parse(e.data)
    if (o["msgType"] === "broadcast-message") {
        messages.appendChild(newBroadcastMessage(o));
        messages.scrollTop = messages.clientHeight;
    }
};

function newBroadcastMessage(o) {
    let msg = document.createElement("div");
    msg.className = "message";

    let idSpan = document.createElement("span");
    idSpan.className = "idSpan";
    idSpan.textContent = o["id"];
    msg.appendChild(idSpan);

    let msgSpan = document.createElement("span");
    msgSpan.className = "msgSpan";
    msgSpan.textContent = o["message"];
    msg.appendChild(msgSpan);

    return msg
}


// Execute a function when the user releases a key on the keyboard
input.addEventListener("keydown", function(event) {
    // Act when Enter key is pressed.
    if (event.code === "Enter") {

        // Cancel the default action, if needed
        event.preventDefault();

        // Trigger the button element with a click
        if (input.value.length > 0) {
            // Send just the raw string. It would probably make sense
            // to wrap this in JSON encoding.
            conn.send(input.value);

            // Reset the input value to an empty string.
            input.value = "";
        }
    }
});

// On launch, clear input and place focus on that field.
input.value = "";
input.focus();
