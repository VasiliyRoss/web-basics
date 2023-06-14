function displayText(inputId, displayClass, defaultText) {
  var inputText = document.getElementById(inputId).value;
  var displayElements = document.getElementsByClassName(displayClass);
  var inputElement = document.getElementById(inputId);
  var errorMessage = inputElement.nextElementSibling; 
  
  for (var i = 0; i < displayElements.length; i++) {
    if (inputText !== '') {
      displayElements[i].textContent = inputText;
      inputElement.classList.remove('post-description__field_error');
      errorMessage.classList.add('block_hidden');
    } else if (defaultText) {
      displayElements[i].textContent = defaultText;
    }
  }
}

function attachInputEventToElements(inputFieldIds, placeholderClasses, defaultTexts) {
  for (var i = 0; i < inputFieldIds.length; i++) {
    (function(index) {
      var element = document.getElementById(inputFieldIds[index]);
      element.addEventListener('input', function() {
        displayText(inputFieldIds[index], placeholderClasses[index], defaultTexts[index]);
      });
    })(i);
  }
}

var inputFieldIds = ['postTitle', 'postDescription', 'postAuthor', 'postPublishDate'];
var placeholderClasses = ['post-title', 'post-description-text', 'post-author', 'post-publish-date'];
var defaultTexts = ['New Title', 'Please, enter any description', 'Enter author name', '4/19/2023'];

attachInputEventToElements(inputFieldIds, placeholderClasses, defaultTexts);

var formContent = document.getElementById('postContent');
formContent.addEventListener('input', function() {
  var errorMessage = formContent.nextElementSibling;
  errorMessage.classList.add('block_hidden');
})

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

var removeButtons = ['removeButton-authorPhoto', 'removeButton-postImage', 'removeButton-cardImage'];
var imageInputIds = ['authorPhoto', 'postImage', 'cardImage'];

function attachClickEventToElements(imageInputIds, removeButtons) {
  for (var i = 0; i < removeButtons.length; i++) {
    (function(index) {
      var element = document.getElementById(removeButtons[index]);
      element.addEventListener('click', function() {
        resetFileInput(imageInputIds[index], removeButtons[index]);
      });
    })(i);
  }
}

attachClickEventToElements(imageInputIds, removeButtons)

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

var uploadButtonsIds = ['authorPhoto', 'uploadNew-authorPhoto', 'postImage', 'uploadNew-postImage', 'cardImage', 'uploadNew-cardImage'];

function attachUploadEventToElements(ids) {
  for (var i = 0; i < ids.length; i++) {
    (function(index) {
      var element = document.getElementById(ids[index]);
      element.addEventListener('change', function() {
        uploadImage(event, ids[index]);
      });
    })(i);
  }
}

attachUploadEventToElements(uploadButtonsIds)

var form = document.getElementById('form');

form.addEventListener('submit', async (e) => {
  e.preventDefault();
  var alertError = document.getElementById('alertError');
  var alertSuccess = document.getElementById('alertSuccess');

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

  function cutBase64Prefix(base64string) {
    const commaIndex = base64string.indexOf(',');
    const result = base64string.substring(commaIndex + 1);
    return result    
  }

  for (var pair of formData.entries()) {
    var inputField = form.querySelector(`[name="${pair[0]}"]`);
    var error = document.getElementById('error-' + pair[0]);

    if (!pair[1]) {
      error.classList.remove('block_hidden')
      inputField.classList.add('post-description__field_error');
      hasEmptyFields = true;
    }

    if (pair[0] === "publish_date") {
      formValues[pair[0]] = convertDateFormat(pair[1]);
    } else if (pair[0] === "author_photo" || pair[0] === "post_image" || pair[0] === "card_image") {
      formValues[pair[0]] = cutBase64Prefix(pair[1])
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