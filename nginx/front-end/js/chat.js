var uname = prompt('请输入用户名', 'user' + uuid(8, 16));
var ws = new WebSocket("ws://127.0.0.1:8083/ws");
ws.onopen = function () {
    var data = "系統消息:已連線";
    listMsg(data);
};
ws.onmessage = function (e) {
    var msg = JSON.parse(e.data);
    var sender, user_name, name_list, change_type;
    switch (msg.type) {
        case 'system':
            sender = '系统消息: ';
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
    var data = "系统消息 : 出错了,请退出重试.";
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
    var change = type == 'login' ? '上线' : '下线';
    var data = '系统消息: ' + user_name + ' 已' + change;
    listMsg(data);
}
function sendMsg(msg) {
    var data = JSON.stringify(msg);
    ws.send(data);
}
function uuid(len, radix) {
    var chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');
    var uuid = [], i;
    radix = radix || chars.length;
    if (len) {
        for (i = 0; i < len; i++) uuid[i] = chars[0 | Math.random() * radix];
    } else {
        var r;
        uuid[8] = uuid[13] = uuid[18] = uuid[23] = '-';
        uuid[14] = '4';
        for (i = 0; i < 36; i++) {
            if (!uuid[i]) {
                r = 0 | Math.random() * 16;
                uuid[i] = chars[(i == 19) ? (r & 0x3) | 0x8 : r];
            }
        }
    }
    return uuid.join('');
}