import Post from "./components/Post.js";
import { onResize } from "./utils/events.js";

export const Home = async () => {
  try {
    const resp = await fetch("/api/posts");
    if (!resp.ok) {
      console.log("Didn't get posts from api");
      return;
    }
    const posts = await resp.json();
    const postsElement = document.querySelector(".posts");
    posts.forEach((post) => {
      postsElement.append(Post(post));
    });
    onResize(AdjustPostLines);
  } catch (error) {
    console.error(error);
  }
};

const AdjustPostLines = () => {
  const post = document.querySelectorAll(".post");

  post.forEach((elm) => {
    const title = elm.querySelector(".title");
    const titleLines = getLineCount(title) - 1;
    console.log(titleLines, title.textContent);

    const text = elm.querySelector(".text");
    text.style["-webkit-line-clamp"] = "unset";
    const linesToFit = getLineCount(text);
    text.style["-webkit-line-clamp"] = linesToFit - titleLines;
  });
};

function getLineCount(element) {
  const lineHeight = parseFloat(getComputedStyle(element).lineHeight);
  const height = element.clientHeight;
  let lineCount = Math.floor(height / lineHeight);
  return lineCount;
}
