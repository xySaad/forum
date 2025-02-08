import { PostCard } from "./Post.js";
import div from "./native/div.js";

const getPosts = async (PostsArea) => {
  try {
    const resp = await fetch(`/api/posts`);
    if (!resp.ok) {
      throw new Error("status not ok:", resp.status);
    }
    const json = await resp.json();

    json.forEach((post) => {
      PostsArea.append(PostCard(post));
    });
  } catch (error) {
    console.error("Error fetching comments:", error);
  }
};

export const InfinitePosts = () => {
  const PostsArea = div("posts");
  getPosts(PostsArea);
  return PostsArea;
};
