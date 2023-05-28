function displayText(inputId, displayClass) {
  var inputText = document.getElementById(inputId).value;
  var displayElements = document.getElementsByClassName(displayClass);
  
  for (var i = 0; i < displayElements.length; i++) {
    displayElements[i].textContent = inputText;
  }
}

function uploadImage(event, id) {
  var input = event.target;
  var reader = new FileReader();

  reader.onload = function () {
    var placeholderImage = document.getElementById('placeholder-' + id);
    placeholderImage.src = reader.result;

    var previewImage = document.getElementById('preview-' + id);
    previewImage.src = reader.result;

    var inputImage = document.getElementById('inputImage-' + id);
    inputImage.value = reader.result;
  };

  reader.readAsDataURL(input.files[0]);
}
var form = document.getElementById('form');

form.addEventListener('submit', (e) => {
  e.preventDefault();

  var formData = new FormData(form);
  var formValues = {};

  for (var pair of formData.entries()) {
    formValues[pair[0]] = pair[1];
  }

  console.log(JSON.stringify(formValues));
})