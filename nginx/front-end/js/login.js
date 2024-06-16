
document.getElementById('btn_').addEventListener('click', (event) => {
  event.preventDefault();
  submitForm();
});

function submitForm() {
  var username = document.getElementById('user').value;
  var password = document.getElementById('pass').value;
  if (username == '' || password == '') {
    Swal.fire({
      icon: "error",
      title: "Oops...",
      text: "Please do not leave this field blank.",
      footer: '<a href="#">Why do I have this issue?</a>'
    });
    return
  }
  var path = 'api/v1/login'
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
      showError();
    });
}

const inputUser = document.getElementById('user');
const inputPassword = document.getElementById('pass');
const label = document.querySelector('.labelUser');
const label1 = document.querySelector('.labelPass');

inputPassword.addEventListener("keypress", (event) => {

  if (event.key === 'Enter') {
    event.preventDefault();
    document.getElementById("btn_").click();
  }
})

inputUser.addEventListener('focus', () => {
  label.classList.add('focus');
});

inputUser.addEventListener('blur', () => {
  if (inputUser.value === '') {
    label.classList.remove('focus');
  }
});

inputPassword.addEventListener('focus', () => {
  label1.classList.add('focus');
});

inputPassword.addEventListener('blur', () => {
  if (inputPassword.value === '') {
    label1.classList.remove('focus');
  }
});

// 获取错误消息元素
var loginError = document.getElementById('login-error');
// 在某些条件下显示错误消息
function showError() {
  loginError.style.visibility = 'visible';
  inputPassword.value = '';
}
// 在某些条件下隐藏错误消息
function hideError() {
  loginError.style.visibility = 'hidden';
}


