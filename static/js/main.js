var output = null;

$(document).ready(function () {
  $("#input").on("keyup", debounce(1000, processInput));
  $("#font").change(debounce(1000, processInput));
});

function processInput() {
  $.ajax({
    type: "POST",
    url: "/process",
    dataType: "json",
    data: {
      text: $("#input").val(),
      font: $("#font").val() || "standard",
    },
    traditional: true,

    success: function (data) {
      output = data;
      document.getElementById("output").innerHTML = data;
    },
    error: function (_, _, errorThrown) {
      console.log(errorThrown);
    },
  });
}

function exportFile() {
  var format = ".txt";
  var fileName = $("#file-name").val();
  $.when(processInput()).done(function (data) {
    if (output) {
      window.location = `/export?output=${encodeURIComponent(
        output
      )}&input=${encodeURIComponent(
        `${fileName}` || "export"
      )}&format=${encodeURIComponent(format)}`;
    }
  });
}

function debounce(wait, func) {
  let timeout;
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout);
      func(...args);
    };
    clearTimeout(timeout);
    timeout = setTimeout(later, wait);
  };
}
