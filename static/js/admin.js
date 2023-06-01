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


async function createPost() {
  const titleInput = document.getElementById('postTitle')
	const subtitleInput = document.getElementById('postDescription')
  const authorInput = document.getElementById('postAuthor')
  const authorPhotoInput = document.getElementById('inputImage-authorPhoto')
  const publishDateInput = document.getElementById('postPublishDate')
  const postImageInput = document.getElementById('inputImage-postImage')
  const cardImageInput = document.getElementById('inputImage-cardImage')
  const contentInput = document.getElementById('postContent')

  const respose = await fetch('/api/post', {
    method: 'POST',
    body: JSON.stringify({
      title: titleInput.value,
      subtitle: subtitleInput.value,
      author: authorInput.value,
      author_photo: authorPhotoInput.value,
      publish_date: publishDateInput.value,
      post_image: postImageInput.value,
      card_image: cardImageInput.value,
      content: contentInput.value,
    })
  })

    console.log(respose.ok)
  }

  /* старая функция, которая работает с формой. Нужно использовать её 
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
  
  */