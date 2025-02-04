import { NewReference } from "../utils/reference.js";
import { PostCard } from "./Post.js";
import div from "./native/div.js";
const getPosts = async (PostsArea, isfetch, offset, categories, lastPostId) => {
  isfetch(true);
  try {
    const resp = await fetch(
      `/api/posts?page=${offset()}&categories=${categories}&from=${lastPostId()}`
    );
    if (!resp.ok) {
      throw new Error("Network response was not ok");
    }
    const json = await resp.json();
    if (offset() === 0) {
      lastPostId(json[0].id);
    }
    offset((prev) => prev + 1);

    json.forEach((post) => {
      PostsArea.append(PostCard(post));
    });
  } catch (error) {
    console.error("Error fetching comments:", error);
  } finally {
    isfetch(false);
  }
};
const getcategories = () => {
  let selectedCategories = ["0", "0", "0", "0"];
  const categories = document.querySelectorAll(".category");
  categories.forEach((child, index) => {
    if (child.classList.contains("Selected")) {
      if (index == 0) {
        return selectedCategories.join("");
      }
      selectedCategories[index - 1] = "1";
    }
  });
  return selectedCategories.join("");
};

export const CreatePostsArea = () => {
  const PostsArea = div("posts");
  let offset = NewReference(0);
  let isfetch = NewReference(false);
  let categoiers = getcategories();
  let lastPostId = NewReference(0);
  getPosts(PostsArea, isfetch, offset, categoiers, lastPostId);

  window.onscroll = (e) => {
    const sh =
      document.documentElement.scrollHeight || document.body.scrollHeight; // Total height of the document
    const st = window.scrollY || window.pageYOffset; // Current vertical scroll position
    const ht = window.innerHeight; // Height of the viewport
    if ((ht === 0 || st + ht >= sh) && !isfetch()) {
      getPosts(PostsArea, isfetch, offset, categoiers, lastPostId);
    }
  };
  return PostsArea;
};
