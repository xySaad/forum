import { Reaction } from "./reactions.js";
import { InfinitePosts } from "./components/InfinitePosts.js";
import div from "./components/native/div.js";
import { PostCreationBar } from "./components/createPost.js";

const FilterSearch = () => {
  const filterContainer = div("filter");
  filterContainer.innerHTML = `<div class="categories">
          <button class="category active all">All</button>
          <button class="category">Sport</button>
          <button class="category">Technology</button>
          <button class="category">Finance</button>
          <button class="category">Science</button>
        </div>`;
  return filterContainer;
};

export const Home = () => {
  Reaction();
  return div("homePage").add(
    PostCreationBar(),
    FilterSearch(),
    InfinitePosts()
  );
};
