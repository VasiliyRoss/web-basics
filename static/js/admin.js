function displayText(inputId, displayClass) {
  var inputText = document.getElementById(inputId).value;
  var displayElements = document.getElementsByClassName(displayClass);
  
  for (var i = 0; i < displayElements.length; i++) {
    displayElements[i].textContent = inputText;
  }
}

function resetFileInput(id, buttonId) {
  document.getElementById(id).value = null;
  document.getElementById('placeholder-' + id).style.backgroundImage = 'url(/static/images/decorations/placeholder-' + id + '.png)';
  document.getElementById('preview-' + id).style.backgroundImage = 'url(/static/images/decorations/preview-' + id + '.png)';
  document.getElementById('inputImage-' + id).value = null;
  document.getElementById(buttonId).classList.add('block_hidden');
  
  var subdescription = document.getElementById('subdescription-' + id);
  if (subdescription) {
    subdescription.classList.remove('block_hidden')
  }

  var upload = document.getElementById('upload-' + id)
  if (upload) {
    upload.classList.remove('block_hidden')
  }

  var uploadNew = document.getElementById('uploadNew-' + id)
  uploadNew.classList.add('block_hidden')
}

function uploadImage(event, id) {
  var input = event.target;
  var reader = new FileReader();

  reader.onload = function () {
    var placeholderImage = document.getElementById('placeholder-' + id);
    placeholderImage.style.backgroundImage = 'url(' + reader.result + ')';

    var previewImage = document.getElementById('preview-' + id);
    previewImage.style.backgroundImage = 'url(' + reader.result + ')';

    var inputImage = document.getElementById('inputImage-' + id);
    inputImage.value = reader.result;

    var removeButton = document.getElementById('removeButton-' + id);
    removeButton.classList.remove('block_hidden')

    var subdescription = document.getElementById('subdescription-' + id);
    if (subdescription) {
      subdescription.classList.add('block_hidden')
    }

    var upload = document.getElementById('upload-' + id)
    if (upload) {
      upload.classList.add('block_hidden')
    }

    var uploadNew = document.getElementById('uploadNew-' + id)
    uploadNew.classList.remove('block_hidden')
  };

  reader.readAsDataURL(input.files[0]);
}

var form = document.getElementById('form');

form.addEventListener('submit', async (e) => {
  e.preventDefault();
  var alertError = document.getElementById('alertError');
  var alertSuccess = document.getElementById('alertSuccess');
  var errorMessages = document.getElementsByClassName('field__error');

  var formData = new FormData(form);
  var formValues = {};

  function convertDateFormat(inputDate) {
    var parts = inputDate.split('-');
    var year = parts[0];
    var month = parts[1];
    var day = parts[2];
  
    var newDate = month + '/' + day + '/' + year;
  
    return newDate;
  }

  var hasEmptyFields = false;
  for (var i = 0; i < errorMessages.length; i++) {
    errorMessages[i].classList.add('block_hidden');
  }

  function cutBase64Prefix(item) {
    
  }

  for (var pair of formData.entries()) {
    var inputField = form.querySelector(`[name="${pair[0]}"]`);
    var error = document.getElementById('error-' + pair[0]);
    if (!pair[1]) {
      if (error){
        error.classList.remove('block_hidden')
        inputField.classList.add('post-description__field_error');
      }
      hasEmptyFields = true;
    }
    if (pair[0] === "publish_date") {
      formValues[pair[0]] = convertDateFormat(pair[1]);
    } else if (pair[0] === "author_photo" || pair[0] === "post_image" || pair[0] === "card_image") {

    } else {
      formValues[pair[0]] = pair[1];
    }    
  }

  if (hasEmptyFields) {
    alertError.classList.add('show');
    alertSuccess.classList.remove('show');
  } else {
    const response = await fetch('/api/post', {
      method: 'POST',
      body: JSON.stringify(formValues)
    });

    if (response.ok) {
      alertError.classList.remove('show');
      alertSuccess.classList.add('show');
    } else {
      alertError.classList.add('show');
      alertSuccess.classList.remove('show');
    }
  }

  console.log(JSON.stringify(formValues));
});


/*

#реализовать отсекание зашифрованной картинки в JSON
#поправить форматирование mysql в admin.go
#вызов функций сделать в admin.js*/