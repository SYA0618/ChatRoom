document.addEventListener('DOMContentLoaded', function() {
    const token = localStorage.getItem('token');
    if (!token) {
        window.location.href = '/index.html';
    } else {
        document.getElementById('welcome-message').textContent = 'Welcome!';
    }
});

var uname = sessionStorage.getItem('user_name');
var ws = new WebSocket("ws://127.0.0.1:8083/ws");

ws.onopen = function () {
    var data = "System Notification: Connected.";
    listMsg(data);
};
ws.onmessage = function (e) {
    var msg = JSON.parse(e.data);
    var sender, user_name, name_list, change_type;
    switch (msg.type) {
        case 'system':
            sender = 'System Notification: ';
            break;
        case 'user':
            sender = msg.from + ': ';
            break;
        case 'handshake':
            var user_info = {'type': 'login', 'content': uname};
            sendMsg(user_info);
            return;
        case 'login':
        case 'logout':
            user_name = msg.content;
            name_list = msg.user_list;
            change_type = msg.type;
            dealUser(user_name, change_type, name_list);
            return;
    }
    var data = sender + msg.content;
    listMsg(data);
};
ws.onerror = function () {
    var data = "System Notification: Page error, please refresh.";
    listMsg(data);
};
function confirm(event) {
    var key_num = event.keyCode;
    if (13 == key_num) {
        send();
    } else {
        return false;
    }
}
function send() {
    var msg_box = document.getElementById("msg");
    var content = msg_box.value;
    var reg = new RegExp("\r\n", "g");
    content = content.replace(reg, "");
    var msg = {'content': content.trim(), 'type': 'user'};
    sendMsg(msg);
    msg_box.value = '';
}
function listMsg(data) {
    var msg_list = document.getElementById("chat");
    var msg = document.createElement("p");
    msg.innerHTML = data;
    msg_list.appendChild(msg);
    msg_list.scrollTop = msg_list.scrollHeight;
}
function dealUser(user_name, type, name_list) {
    var user_list = document.getElementById("user_list");
    var user_num = document.getElementById("user_num");
    while(user_list.hasChildNodes()) {
        user_list.removeChild(user_list.firstChild);
    }
    for (var index in name_list) {
        var user = document.createElement("p");
        user.innerHTML = name_list[index];
        user_list.appendChild(user);
    }
    user_num.innerHTML = name_list.length;
    user_list.scrollTop = user_list.scrollHeight;
    var change = type == 'login' ? 'online' : 'offline';
    var data = 'System Notification: ' + user_name + ' is ' + change;
    listMsg(data);
}
function sendMsg(msg) {
    var data = JSON.stringify(msg);
    ws.send(data);
}
