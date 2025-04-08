import { go } from "../router.js";
import { PostCard } from "./Post.js";
import div from "./native/div.js";

const getPosts = async (PostsArea, url) => {
  let nextUrl;
  let json;
  try {
    const resp = await fetch(url);
    if (resp.status === 401) {
      go("/login");
      return;
    }
    if (!resp.ok) {
      throw new Error("status not ok:", resp.status);
    }
    json = await resp.json();
    if (json) {
      for (const post of json) {
        PostsArea.append(PostCard(post));
      }
      nextUrl = new URL(resp.url);
    }
  } catch (error) {
    console.error("Error fetching comments:", error);
  }

  const options = {
    root: document.querySelector("main"),
    rootMargin: "0px",
    threshold: 0.3,
  };

  const observer = new IntersectionObserver((entries) => {
    if (entries[0].isIntersecting) {
      observer.unobserve(entries[0].target);
      if (nextUrl) {
        nextUrl.searchParams.set("lastId", json[json.length - 1].id);
        getPosts(PostsArea, nextUrl);
      }
    }
  }, options);

  observer.observe(PostsArea.children[PostsArea.children.length - 1]);
};

export const InfinitePosts = (url) => {
  const PostsArea = div("posts");
  getPosts(PostsArea, url);
  return PostsArea;
};
