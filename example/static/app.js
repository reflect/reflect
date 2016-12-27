var usernameInput = $('input#username'),
    passwordInput = $('input#password');

function showModalOnClick() {
  $('#login-modal').on('shown.bs.modal', function () {
    $('#myInput').focus();
  });
}

function handleUserDataJson(data) {
  console.log(data);
}

function handleFailure(xhr, status, e) {
  window.location.replace('404.html');
}

function listenForLoginSubmit() {
  $('#login-submit').click(function() {
    var username = usernameInput.val(),
        password = passwordInput.val();

    console.log(username);
    console.log(password);

    if (username.length == 0 || password.length == 0) {
      window.alert("You must supply a username and password!");
    } else {
      $.ajax({
        type: 'GET',
        url: '/users',
        dataType: 'json',
        async: true,
        success: handleUserDataJson,
        error: handleFailure
      })
    }
  });
}

$(function() {
  showModalOnClick();
  listenForLoginSubmit();
});