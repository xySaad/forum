import { InfinitePosts } from "./InfinitePosts.js"
import div from "./native/div.js"

const filterByCat = (e) => {
    const allDiv = document.querySelector(".category.all")
    let categories = document.querySelectorAll(".categories .active")
    if (e.target === allDiv) {
        categories.forEach((cat) => {
            cat.classList.remove("active")
        })
        allDiv.classList.add("active")
    } else {
        if (e.target.classList.contains("active")) {

            e.target.classList.remove("active")
        } else {
            e.target.classList.add("active")
            allDiv.classList.remove("active")
        }
        if (!categories) {
            categories = document.querySelectorAll(".categories .active")
            allDiv.classList.add("active")
        }
    }


    const arr = []
    categories = document.querySelectorAll(".categories .active")
    categories.forEach((category) => {
        if (category === allDiv) {
            return
        }
        arr.push("category=" + category.textContent)
    })
    const homePage = document.querySelector(".homePage")
    homePage.children[2].remove()
    homePage.append(InfinitePosts("api/posts?" + arr.join("&")))
}

export const FilterSearch = () => {
    const filterContainer = div("filter");
    fetch("/api/categories").then(async (resp) => {
        const categories = div("categories")
        filterContainer.add(categories)
        const json = await resp.json()

        if (!resp.ok || json.length === 0) {
            return;
        }
        const allDiv = div("category active all", "All")
        allDiv.onclick = filterByCat
        categories.add(allDiv)
        json?.forEach((category) => {
            const categoryDiv = div("category", category)
            categoryDiv.onclick = filterByCat
            categories.add(categoryDiv)
        })
    })
    return filterContainer;
};
