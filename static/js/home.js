import Post from "./components/Post.js";
import { onResize } from "./utils/events.js";
import { Reaction } from "./reactions.js";
import CreatePost from "./components/createPost.js";
import { filterCat } from "./filter.js";
import ensureAuth from "./utils/ensureAuth.js";
let page=0
export const Home = async () => {
  try {
    const resp = await fetch(`/api/posts?page=${page}&categories=0000`);
    if (!resp.ok) {
      console.log("Didn't get posts from api");
      return;
    }else{
      page+=1
    }
    const posts = await resp.json();

    const postsElement = document.querySelector(".posts");
    posts.forEach((post) => {
      postsElement.append(Post(post));
    });
  } catch (error) {
    console.error(error);
  }
  onResize(AdjustPostLines);
  Reaction();
  document.getElementById("create-post-btn").onclick = async () => {
    if (!(await ensureAuth())) {
      return;
    }

    CreatePost();
  };
  filterCat();
};

const AdjustPostLines = () => {
  const post = document.querySelectorAll(".post");

  post.forEach((elm) => {
    const text = elm.querySelector(".text");
    text.style["-webkit-line-clamp"] = "unset";
    const linesToFit = getLineCount(text);
    text.style["-webkit-line-clamp"] = linesToFit;
  });
};

function getLineCount(element) {
  const lineHeight = parseFloat(getComputedStyle(element).lineHeight);
  const height = element.clientHeight;
  let lineCount = Math.floor(height / lineHeight);
  return lineCount;
}
