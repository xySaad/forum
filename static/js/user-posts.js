import { InfinitePosts } from "./components/InfinitePosts.js";

export const CreatedPosts = () => {
  const url = "/api/user/created/";
  document.querySelector(".created")?.classList.add("active");
  return InfinitePosts(url);
};
export const LikedPosts = () => {
  const url = "/api/user/liked/";
  document.querySelector(".liked")?.classList.add("active");
  return InfinitePosts(url);
};
