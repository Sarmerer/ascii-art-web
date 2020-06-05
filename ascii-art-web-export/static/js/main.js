var output = null

function process() {
    var fontSelector = document.getElementById("font");
    var font = fontSelector.options[fontSelector.selectedIndex].value;

    return $.ajax({
        type: "POST",
        url: '/process',
        dataType: "json",
        data: {
            "text": $("#input").val(),
            "font": font
        },
        traditional: true,

        success: function (data) {
            output = data
            document.getElementById("output").innerHTML = data
        },
        error: function (jqXHR, textStatus, errorThrown) {
            alert('500 Internal server error')
        }
    });
}

function exportFile() {
    var formatSelector = document.getElementById("format");
    var format = formatSelector.options[formatSelector.selectedIndex].value;
    var fileName = document.getElementById("file-name").value;
    var input = ""
    if (fileName) {
        input = fileName
    }else{
        input = "exported"
    }
    console.log(input);

    $.when(process()).done(function (data) {
        if (output) {
            window.location = `/export?output=${encodeURIComponent(output)}&input=${encodeURIComponent(input)}&format=${encodeURIComponent(format)}`
        }
    })
}