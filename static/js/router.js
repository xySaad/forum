import div from "./components/native/div.js";

const routes = {
  "/404": () => div("404", "404 - page not found"),
};

export const AddRoute = (route, page, ...args) => {
  routes[route] = () => page(...args);
};

export const go = (route, popup) => {
  const page = routes[route];
  if (!page) {
    go("/404");
    return;
  }

  document.querySelector("popup").innerHTML = "";
  if (popup || history.state?.popup) {
    document.querySelector("popup").append(page());
  } else {
    document.querySelector("main").innerHTML = "";
    document.querySelector("main").append(page());
  }

  if (!history.state) {
    history.replaceState({ prev: null, path: route }, "");
    return;
  }

  if (history.state.prev?.path != route && history.state.path != route) {
    history.pushState({ prev: history.state, path: route, popup }, "", route);
  }
  console.log(history.state);
};

export const back = () => {
  if (!history.state.prev) {
    go("/");
  } else {
    history.back();
  }
};

export const replacePath = (path) => {
  console.log(history.length, history.state);
  history.replaceState({ ...history.state, path }, "", path);
  console.log(history.length, history.state);
};
