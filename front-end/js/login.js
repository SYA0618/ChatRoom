var path = '/api/v1/login'



document.getElementById('btn_').addEventListener('click', function(event) {
    event.preventDefault();

    var username = document.getElementById('user').value;
    var password = document.getElementById('pass').value;
    if (username == '' || password == '') {
        alert('欄位勿空白')
        return
    }
    var data = new FormData();
    data.append('user_name', username);
    data.append('password', password);

    axios.post(path, data)
    .then(res => {
        console.log(res);
        window.location.href = '/welcome.html';
    })
    .catch(err => {
        console.error(err); 
    });
});


// document.getElementById('btn_').addEventListener('click', function(event) {
//     event.preventDefault(); // 防止表單默認提交行為

//     // 獲取使用者輸入的資料
//     var username = document.getElementById('user').value;
//     var password = document.getElementById('pass').value;

//     // 建立要發送的資料物件
//     var data = new FormData();
//     data.append('user_name', username);
//     data.append('password', password);

//     // 使用 Fetch API 發送 POST 請求
//     fetch('/api/v1/login', {
//         method: 'POST',
//         body: data
//     })
//     .then(function(response) {
//         // 處理後端的回應
//         if (response.ok) {
//             // 登入成功，執行相應的操作，例如重新導向至其他頁面
//             window.location.href = '/api/v1/login';
//         } else {
//             // 登入失敗，顯示錯誤訊息等
//             console.error('登入失敗');
//         }
//     })
//     .catch(function(error) {
//         // 發送請求或處理回應時出現錯誤，例如網路連接問題等
//         console.error('發送請求時出現錯誤:', error);
//     });
// });
