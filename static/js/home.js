import { onResize } from "./utils/events.js";
import { Reaction } from "./reactions.js";
import { InfinitePosts } from "./components/InfinitePosts.js";
import div from "./components/native/div.js";
import { PostCreationBar } from "./components/createPost.js";

const FilterSearch = ()=>{
  const filterContainer = div("filter")
  filterContainer.innerHTML = `<div class="categories">
          <button class="category active all">All</button>
          <button class="category">Sport</button>
          <button class="category">Technology</button>
          <button class="category">Finance</button>
          <button class="category">Science</button>
        </div>`
  return filterContainer
}
export const Home = () => {
  onResize(AdjustPostLines);
  Reaction();
  return div("homePage").add(PostCreationBar(),FilterSearch(), InfinitePosts());
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
