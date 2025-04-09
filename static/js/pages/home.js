import { InfinitePosts } from "../components/InfinitePosts.js";
import div from "../components/native/div.js";
import { PostCreationBar } from "../components/createPost.js";
import { FilterSearch } from "../components/filter.js";
import { toggleIcon } from "../components/Headers.js";

export const Home = () => {
  toggleIcon(".home");
  return div("homePage").add(
    PostCreationBar(),
    FilterSearch(),
    InfinitePosts("/api/posts")
  );
};
