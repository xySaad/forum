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

      post.categories.forEach((category) => {
        console.log(category);

        categories += `<div class="category">${category}</div>`;
      });

      document.querySelector(".posts").innerHTML += `<div class="post">
          <div class="title">${post.title}</div>
          <div class="categories">
          ${categories}
          </div>
          <div class="text">${post.text}</div>
        </div>`;
    });
  } catch (error) {
    console.error(error);
  }
};
