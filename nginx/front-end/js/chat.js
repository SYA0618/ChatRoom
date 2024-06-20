window.onload = function() {
    var conn;
    var msg = document.getElementById("msg")
    var chat = document.getElementById("chat")

    function appendChat(item) {
        var doScroll = chat.scrollTop > chat.scrollHeight - chat.clientHeight - 1;
        chat.appendChild(item);
        if (doScroll) {
            chat.scrollTop = chat.scrollHeight - chat.clientHeight;
        }
    }
    document.getElementById("form").onsubmit = function () {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        conn.send(msg.value);
        msg.value = "";
        return false;
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendChat(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerText = messages[i];
                appendChat(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendChat(item);
    }
}