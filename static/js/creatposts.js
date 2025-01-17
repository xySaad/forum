function checkPost() {
    let createPost = document.getElementsByClassName('create');
    Array.from(createPost).forEach((element) => {
        element.addEventListener('click', () => {
            const message = document.getElementById('message'); 
            const inputElement = document.querySelector('.title'); 

            if (inputElement) {
                let input = inputElement.value;
                if (input.trim() === '' || !/^[a-zA-Z]+$/.test(input)) {
                    message.textContent = 'Input cannot be empty!';
                    message.style.color = 'red';
                } else {
                    
                }
            } else {
                console.error('No input field with class "title" found.');
            }
        });

    });
}
checkPost();