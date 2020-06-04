function process() {
    var fontSelector = document.getElementById("font");
    var font = fontSelector.options[fontSelector.selectedIndex].value;
    console.log("got");

    console.log();
    return $.ajax({
        type: "POST",
        url: '/process',
        dataType: "json",
        data: { "text": $("#input").val(), "font": font },
        traditional: true,
        
        success: function (data) {
            document.getElementById("output").innerHTML = data
        },
        error: function (jqXHR, textStatus, errorThrown) {
            console.log(jqXHR, textStatus,errorThrown);
            
            alert('Internal Server Error!')
        }
    });

}