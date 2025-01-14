import { timePassed } from "./utils.js";

export const Home = async () => {
  try {
    const resp = await fetch("/api/posts");
    if (!resp.ok) {
      console.log("Didn't get posts from api");
      return;
    }
    const posts = await resp.json();
    posts.forEach((post) => {
      let categories = "";

      post.categories?.forEach((category) => {
        console.log(category);

        categories += `<div class="category">${category}</div>`;
      });

      document.querySelector(".posts").innerHTML += `<div class="postContainer">
      <div class="post">
          <div class="publisher">
            <img src="/static/svg/no-profile.svg" alt="no-profile">
            <div>${post.publisher.name}</div>
            <div>${timePassed(post.creationTime)}</div>
          </div>
          <div class="title">${post.title}</div>
          <div class="categories">
          ${categories}
          </div>
          <div class="text">${post.text}</div>
      </div>
      <div class="leftBar">
        <img src="/static/svg/arrow-up.svg" alt="arrow-up">
        <img src="/static/svg/comment-bubble.svg" alt="comment-bubble">
        <img src="/static/svg/arrow-down.svg" alt="arrow-down">
      </div>
    </div>`;
    });
  } catch (error) {
    console.error(error);
  }
};
