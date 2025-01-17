export let post = {}


export function checkPost() {
    let createPost = document.getElementsByClassName('create')
    Array.from(createPost).forEach((element) => {
        element.addEventListener('click', () => {
            const message = document.getElementById('message')
            const title = document.querySelector('.title')
            const inputElement = document.querySelector('.description')
            if (inputElement && title) {
                let input = inputElement.value
                let titl= title.value 
                if ((input.trim() === '' || titl.trim() === "") ) {
                    message.textContent = 'Input cannot be empty!'
                    message.style.color = 'red'
                } else {
                    post.title = titl
                    post.content = input
                    post = JSON.stringify(post)
                    console.log(post);
                    post = {}
                }
            } else {
                console.error('No input field with class "title" found.')
            }
        });

    });
}
checkPost();