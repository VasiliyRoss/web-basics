var form = document.getElementById('form');
var alertFieldError = document.getElementById('alertFieldError');

form.addEventListener('submit', async (e) => {
  e.preventDefault();
  var alertFieldError = document.getElementById('alertFieldError');
  //var alertErrorInvalidUser = document.getElementById('alertErrorInvalidUser');

  var formData = new FormData(form);
  var formValues = {};

  var hasEmptyFields = false;
  var invalidEmailFormat = false;

  for (var pair of formData.entries()) {
    var inputField = form.querySelector(`[name="${pair[0]}"]`);
    var error = document.getElementById('error-' + pair[0]);

    if (!pair[1]) {
      error.classList.remove('block_hidden')
      inputField.classList.add('form__field_error');
      if(pair[0] === "user_email") {
        error.textContent = 'Email is required';
      } else {
        error.textContent = 'Password is required';
      }      
      hasEmptyFields = true;
    }

    if(pair[0] === "user_email") {
      if(!validateEmail(pair[1]) && pair[1]){
        invalidEmailFormat = true
        error.classList.remove('block_hidden')
        inputField.classList.add('form__field_error');  
        error.textContent = 'Incorrect email format. Correct format is ****@**.***';
      }
    }
      formValues[pair[0]] = pair[1];
  }

  if (hasEmptyFields || invalidEmailFormat) {
    alertFieldError.classList.add('show');
  }

  if (!hasEmptyFields && !invalidEmailFormat) {
    alert('Success!');
  }

  console.log(JSON.stringify(formValues));

  function validateEmail(email) {
    var regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return regex.test(email);
  }
});

var inputFields = ['userEmailInput', 'userPasswordInput'];

function attachInputEventToElements(inputFields) {
    for (var i = 0; i < inputFields.length; i++) {
      (function(index) {
        var element = document.getElementById(inputFields[index]);
        element.addEventListener('input', function() {
          var errorMessage = element.nextElementSibling;
          alertFieldError.classList.remove('show');
          errorMessage.classList.add('block_hidden');
          element.classList.remove('form__field_error');  
        });
      })(i);
    }
  }

attachInputEventToElements(inputFields)