import { InfinitePosts } from "./components/InfinitePosts.js";
import div from "./components/native/div.js";
import { PostCreationBar } from "./components/createPost.js";
import { FilterSearch } from "./components/filter.js";

export const Home = () => {
  document.querySelector(".icon.home")?.classList.add("active");
  return div("homePage").add(
    PostCreationBar(),
    FilterSearch(),
    InfinitePosts("/api/posts")
  );
};
