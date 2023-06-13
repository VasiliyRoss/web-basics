function displayText(inputId, displayClass) {
  var inputText = document.getElementById(inputId).value;
  var displayElements = document.getElementsByClassName(displayClass);
  
  for (var i = 0; i < displayElements.length; i++) {
    displayElements[i].textContent = inputText;
  }
}

function convertDateFormat(inputDate) {
  var parts = inputDate.split('-');
  var year = parts[0];
  var month = parts[1];
  var day = parts[2];

  var newDate = month + '/' + day + '/' + year;

  return newDate;
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

form.addEventListener('submit', async (e) => {
  e.preventDefault();

  var formData = new FormData(form);
  var formValues = {};

  for (var pair of formData.entries()) {
    console.log(pair[0]);
    if (pair[0] === "publish_date") {
      formValues[pair[0]] = convertDateFormat(pair[1]);
    } else {
      formValues[pair[0]] = pair[1];
    }
  }

  const respose = await fetch('/api/post', {
    method: 'POST',
    body: JSON.stringify(formValues)
  })

  console.log(JSON.stringify(formValues));
  console.log(respose.ok);
})
