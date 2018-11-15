$(document).ready(function() {
    $.ajax({
        url: "/api/show"
    }).then(function(data) {
        if(data.nickname !== "") {
            $('#nickname').append(data.nickname);
            $('#message').append(data.message);
        }
    });
});