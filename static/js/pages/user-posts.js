import { toggleIcon } from "../components/Headers.js";
import { InfinitePosts } from "../components/InfinitePosts.js";

export const CreatedPosts = () => {
  const url = "/api/user/created/";
  toggleIcon(".created");
  return InfinitePosts(url);
};
export const LikedPosts = () => {
  const url = "/api/user/liked/";
  toggleIcon(".liked");
  return InfinitePosts(url);
};
