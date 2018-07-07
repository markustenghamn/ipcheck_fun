
function IPCheck() {
    $.ajax({
        method: "GET",
        url: "/json"
    }).done(function (obj) {
        $("#ipresult").text(obj);
    });
}

IPCheck();
