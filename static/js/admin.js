const form = document.getElementById('form')
var authorPhoto
var cardImage
var postImage

function loadImage(container, file) {
  const reader = new FileReader()

  reader.addEventListener(
    "load",
    () => {
      authorPhoto = reader.result
    },
    false
  )

  if (file) {
    reader.readAsDataURL(file);
  }
}

function loadAndPreviewAuthorPhoto() {
  const placeholder = document.getElementById("authorPhotoPlaceholder")
  const preview = document.getElementById("authorPhotoDemo")
  const file = document.getElementById("authorPhoto").files[0]

  loadImage(authorPhoto, file)
}


form.addEventListener('submit', (e) => {
  e.preventDefault();

  const formData = new FormData(form)
  const title = formData.get('post_title')
  const subtitle = formData.get('post_short_description')
  const author = formData.get('post_author_name')
  const publishDate = formData.get('post_publish_date')
  const content = formData.get('content')

  let post = JSON.stringify({
    title: title,
    subtitle: subtitle,
    author: author,
    authorPhoto: authorPhoto,
    publishDate: publishDate,
    postImage: postImage,
    cardImage: cardImage,
    content: content
})
  
  console.log(post)
})