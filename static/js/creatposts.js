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
                let titl = title.value
                if ((input.trim() === '' || titl.trim() === "")) {
                    message.textContent = 'Input cannot be empty!'
                    message.style.color = 'red'
                } else {
                    const data = {
                        title: titl,
                        description: input
                    }
                    fetch('/api/creatpost', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: (JSON.stringify(data))
                    })
                        .then(response => {
                            if (!response.ok) {
                                throw new Error('Network response was not ok ' + response.statusText)
                            }
                            return response.json()
                        })
                        .then(data => {
                            message.textContent = 'Post created successfully!'
                            message.style.color = 'green'
                            console.log('Success:', data)
                        })
                        .catch(error => {
                            message.textContent = 'Failed to create post!'
                            message.style.color = 'red'
                            console.error('Error:', error)
                        })
                }
            } else {
                console.error('No input field with class "title" found.')
            }
        })
    })
}
checkPost()
