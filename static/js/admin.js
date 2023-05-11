const form = document.getElementById('form')
const formData = new FormData(form)
const title = formData.get('post_title')
const subtitle = formData.get('post_short_description')
const author = formData.get('post_author_name')
const authorPhoto = formData.get('post_author_photo')
const publishDate = formData.get('post_publish_date')
const postImage = formData.get('post_image')
const cardImage = formData.get('post_card_image')
const content = formData.get('content')




form.addEventListener('submit', (e) => {
  e.preventDefault();

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