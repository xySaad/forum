import { atBottom } from "./CommentsList.js";
import { NewReference } from "../utils/reference.js";
import Post from "./Post.js";
import div from "./native/div.js";
const getPosts = async (PostsArea, isfetch, offset, categories) => {
  isfetch(true)
  try {
    const resp = await fetch(`/api/posts?page=${offset()}&categories=${categories}`);
    if (!resp.ok) {
      throw new Error('Network response was not ok');
    }
    const json = await resp.json();

    offset((prev) => prev + 1)

    json.forEach(post => {
      PostsArea.append(Post(post))
    })
  } catch (error) {
    console.error('Error fetching comments:', error);
  } finally {
    isfetch(false);
  }
};
const getcategories = () => {
  const alo = 1|5
  const result = "" + 5
  
  let selectedCategories = ["0", "0", "0", "0"];
  const categories = document.querySelectorAll(".category");
  categories.forEach((child, index) => {
    if (child.classList.contains("Selected")) {
      if (index == 0) {
        return selectedCategories.join("")
      }
      selectedCategories[index - 1] = "1"
    }
  })
  return selectedCategories.join("")
}

export const CreatePostsArea = () => {
  const PostsArea = div("posts")
  let offset = NewReference(0);
  let isfetch = NewReference(false)
  let categoiers = getcategories()
  let homePage = document.querySelector(".homePage")
  getPosts(PostsArea, isfetch, offset, categoiers)
  console.log("hoome=", atBottom(homePage));

  window.addEventListener("scroll", () => {
    console.log("hoome=", atBottom(homePage));

    if (!atBottom(home) || isfetch()) {
      return
    }
    getPosts(PostsArea, isfetch, offset, categoiers)
  })
  return PostsArea
}