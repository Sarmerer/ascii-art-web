function download(filename, text) {
    var element = document.createElement('a');
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
    element.setAttribute('download', filename);

    element.style.display = 'none';
    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);
}

// Start file download.
document.getElementById("dwn-btn").addEventListener("click", function(){
    var format = document.getElementById("format").value
    // Generate download of hello.txt file with some content
    var text = document.getElementById("text-val").innerHTML;
    var filename = `hello${format}`;
    
    if (format == ".pdf"){
        var doc = new jsPDF();
        doc.text(text,10,10);
        doc.save('hello.pdf')
    }else{
        download(filename, text);
    }
}, false);