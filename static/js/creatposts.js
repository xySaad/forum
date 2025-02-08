export let post = {};

export function checkPost() {
    const createPost = document.querySelector(".submitButton");

    if (createPost) {
        console.log("Found submitButton, adding event listener");

        createPost.addEventListener('click', async () => {
            console.log("Submit button clicked");

            const title = document.querySelector('.titleInput');
            const inputElement = document.querySelector('.textInput');
            const selectedCategories = Array.from(
                document.querySelectorAll('.category-checkbox:checked')
            ).map((checkbox) => checkbox.value);

            console.log("Title:", title.value, "Content:", inputElement.value, "Categories:", selectedCategories);

            if (inputElement && title) {
                let input = inputElement.value;
                let titl = title.value;

                if (input.trim() === '' || titl.trim() === '') {
                    alert("Title or Content are Empty!!");
                } else {
                    const data = {
                        title: titl,
                        content: input,
                        categories: selectedCategories,
                    };
                    try {
                        const resp = await fetch('/api/posts/', {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'application/json',
                            },
                            body: JSON.stringify(data),
                            credentials: "include",
                        });

                        if (resp.ok) {
                            const responseData = await resp.json();
                            console.log('Post created successfully:', responseData);
                        } else {
                            console.error('Failed to create post:', resp.statusText);
                        }
                    } catch (error) {
                        console.error('Error occurred while creating post:', error);
                    }
                }
            } else {
                console.error('No input field with class "titleInput" or "textInput" found.');
            }
        });
    } else {
        console.error("submitButton not found");
    }
}
