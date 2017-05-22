var baseurl = "https://cdn.hackerrank.com/hackerrank/static/contests/capture-the-flag/secret/";

function GetNotes() {
    var keys = Array();
    var notes = Array();



    $.ajax({
        url: baseurl + "key.json",
        dataType: 'json',
        async: false,
        success: function (data) {
            keys = Object.keys(data);
        }
    });
    keys.forEach(function (element, index, array) {
        
        $.ajax({
            url: baseurl + "secret_json/" + element + ".json",
            dataType: 'json',
            async: false,
            success: function (json) {
                notes.push(json["news_title"]);
            }
        });

    });
    notes.sort();
    notes.forEach(function(element){
       $("#result").append(element + "<br>");
    });
    
}

GetNotes();