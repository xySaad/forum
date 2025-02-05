import Post from "./components/Post.js";

export const filterCat = (page = 1) => {
    let selectedCategories = []; 
    const categories = document.querySelectorAll(".category");
    const allCategoryButton = document.querySelector(".category.active");

    document.addEventListener("click", async (event) => {
        if (event.target.classList.contains("category")) {
            let value = event.target.textContent;
            if (value === "All") {
                selectedCategories = [];
                categories.forEach((btn) => btn.classList.remove("Selected", "active"));
                event.target.classList.add("Selected", "active");
            } else {
                
                if (allCategoryButton.classList.contains("active")) {
                    allCategoryButton.classList.remove("Selected", "active");
                }

                if (selectedCategories.includes(value)) {
                    
                    selectedCategories = selectedCategories.filter(cat => cat !== value);
                    event.target.classList.remove("Selected", "active");
                } else {
                    
                    selectedCategories.push(value);
                    event.target.classList.add("Selected", "active");
                }
            }

            console.log("Selected categories:", selectedCategories);

            let url = selectedCategories.length > 0 
                ? `/api/posts/${page}/categories=${selectedCategories.join("&")}` 
                : `/api/posts/${page}`;
            try {
                const resp = await fetch(url);

                if (!resp.ok) {
                    console.log("Didn't get posts from API");
                    return;
                }

                const posts = await resp.json();
                console.log(posts);

                const postsElement = document.querySelector(".posts");
                postsElement.innerHTML = "";  

                posts.forEach((post) => {
                    postsElement.append(Post(post));
                });
            } catch (error) {
                console.error("Error fetching posts:", error);
            }
        }
    });
};