import { CreatePostsArea } from "./components/NewPost.js";
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
            document.body.querySelector(".homePage").append(CreatePostsArea())
        }
    });
};