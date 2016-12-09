$(document).ready(function() {
    %.ajax({
        type: 'GET',
        url:  '/users',
        dataType: 'json',
        async: false,
        username: 'brad',
        password: 'beefsticks',
        success: function(res) {
            window.alert("It worked!");
        }
    });
});
