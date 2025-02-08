import { onResize } from "./utils/events.js";
import { Reaction } from "./reactions.js";
import { filterCat } from "./filter.js";
import { InfinitePosts } from "./components/InfinitePosts.js";
import div from "./components/native/div.js";
import { PostCreationBar } from "./components/createPost.js";

export const Home = () => {
  onResize(AdjustPostLines);
  Reaction();
  filterCat();
  return div("homePage").add(PostCreationBar(), InfinitePosts());
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
