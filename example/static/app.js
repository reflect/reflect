$(document).ready(function() {
    $.ajax({
        type: 'GET',
        url:  '/users',
        dataType: 'json',
        async: true,
        username: 'brad',
        password: 'beefsticks',
        success: function(user) {
            var apiToken = user.apiToken;


            window.reflect = {
              tokens: apiToken
            };
        }
    });
});
