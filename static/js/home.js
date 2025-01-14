import { onResize } from "./utils/events.js";
import { timePassed } from "./utils/time.js";

export const Home = async () => {
  try {
    const resp = await fetch("/api/posts");
    if (!resp.ok) {
      console.log("Didn't get posts from api");
      return;
    }
    const posts = await resp.json();
    posts.forEach((post) => {
      document.querySelector(".posts").innerHTML += `<div class="postContainer">
      <div class="post">
          <div class="publisher">
            <img src="/static/svg/no-profile.svg" alt="no-profile">
            <div>${post.publisher.name}</div>
            <div>${timePassed(post.creationTime)}</div>
          </div>
          <div class="title">${post.title}</div>
          <div class="text">${post.text}</div>
          <div class="readmore">Read more</div>
      </div>
      <div class="leftBar">
        <img src="/static/svg/arrow-up.svg" alt="arrow-up">
        <img src="/static/svg/comment-bubble.svg" alt="comment-bubble">
        <img src="/static/svg/arrow-down.svg" alt="arrow-down">
      </div>
    </div>`;
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
