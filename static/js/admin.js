const form = document.getElementById('form')
var authorPhoto
var cardImage
var postImage

function loadAndPreviewAuthorPhoto() {
  const placeholder = document.getElementById("authorPhotoPlaceholder")
  const preview = document.getElementById("authorPhotoDemo")
  const file = document.getElementById("authorPhoto").files[0]
  const reader = new FileReader()

  reader.addEventListener(
    "load",
    () => {
      preview.src = reader.result
      authorPhoto = reader.result
      placeholder.classList.toggle("post-description__author-photo_hidden")
      preview.classList.toggle("post-description__author-photo_hidden")
    },
    false
  )

  if (file) {
    reader.readAsDataURL(file);
  }

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