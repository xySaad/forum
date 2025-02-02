import CreatePost from "./components/createPost.js";

const routes = (location) => {
  switch (location) {
    case "create-post":
      return CreatePost();
    default:
      window.history.replaceState({}, "", "/");
      break;
  }
};

export const go = (route) => {
  const page = routes(route);
  if (!page) {
    return;
  }
  window.history.pushState({}, "", route);
  window.onpopstate = () => {
    page && page.remove();
  };
};

export const back = () => {
  if (window.history.length == 1) {
    window.history.replaceState({}, "", "/");
  } else {
    window.history.back();
  }
};
