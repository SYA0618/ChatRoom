
document.getElementById('btn_').addEventListener('click',(event)=>{
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
          });
        return
    }
    var path = '/api/v1/register'
    var data = new FormData();
    var registerError = document.getElementById('register-error');
    data.append('user_name', username);
    data.append('password', password);

    axios.post(path, data)
    .then(res => {
        console.log(res);
        window.location.href = '/index.html';
    })
    .catch(err => {
        console.error(err);
        registerError.style.visibility = 'visible';
    });
}

const inputPassword = document.getElementById('pass');

inputPassword.addEventListener("keypress", (event)=> {
  
    if (event.key === 'Enter') {
      event.preventDefault();
      document.getElementById("btn_").click();
    }
  })


