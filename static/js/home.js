import Post from "./components/Post.js";
import { onResize } from "./utils/events.js";
import { Reaction } from "./reactions.js";
import CreatePost from "./components/createPost.js";
import { filterCat } from "./filter.js";
import ensureAuth from "./utils/ensureAuth.js";
import { CreatePostsArea } from "./components/NewPost.js";
export const Home = async () => {
  document.body.querySelector(".homePage").append(CreatePostsArea())
 
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