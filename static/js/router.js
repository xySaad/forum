import div from "./components/native/div.js";

const routesByLevel = [];

const routes = {
  "/404": () => div("404", "404 - page not found"),
};

export const AddRoute = (route, page) => {
  const splitedRoute = route.split("/");
  for (let i = 0; i < splitedRoute.length; i++) {
    let path = splitedRoute[i];
    if (i == 0 && path == "") {
      continue;
    }
    if (!routesByLevel[i - 1]) {
      routesByLevel[i - 1] = [];
    }

    const isArg = path[0] == ":";
    const pageToAdd = i === splitedRoute.length - 1 ? page : null;
    let args = {};

    // const isOptional = path[0] == "?";
    // const isArg = (isOptional && path[1] == ":") || path[0] == ":";
    // if (isOptional) {
    //   // set the page component to the previous path that is not optional
    //   routesByLevel[i - 2][splitedRoute[i - 1]].page = page;
    // }

    if (isArg && routesByLevel[i - 2]) {
      args = { ...routesByLevel[i - 2][splitedRoute[i - 1]].args };
      args[path.slice(1)] = i;
      path = splitedRoute[i - 1] + "*";
    }

    routesByLevel[i - 1][path] = { page: pageToAdd, args };
  }

  routes[route] = page;
};

/** 
  @param {string} route - the path to lookup for in the router
  @returns {{found: bool, page: HTMLElement}}
*/

const routeLookup = (route) => {
  const splitedRoute = route.split("/");
  for (let i = 0; i < splitedRoute.length; i++) {
    const path = splitedRoute[i];
    if (i == 0 && path == "") {
      continue;
    }

    if (i === splitedRoute.length - 1) {
      if (!routesByLevel[i - 1][path]) {
        if (routesByLevel[i - 1][splitedRoute[i - 1] + "*"]) {
          return {
            found: true,
            page: routesByLevel[i - 1][splitedRoute[i - 1] + "*"].page,
          };
        }
        return { found: false, page: routes["/404"] };
      }
      return { found: true, page: routesByLevel[i - 1][path].page };
    }

    if (!routesByLevel[i - 1][path]) {
      return { found: false, page: routes["/404"] };
    }
  }
};

export const go = (route, popup, ...args) => {
  const { found, page } = routeLookup(route);

  if (!found) {
    popup = false;
  }

  document.querySelector("popup").innerHTML = "";
  if (popup || history.state?.popup) {
    document.querySelector("popup").append(page(...args));
  } else {
    document.querySelector("main").innerHTML = "";
    document.querySelector("main").append(page(...args));
  }

  if (!history.state) {
    history.replaceState({ prev: null, path: route }, "");
    return;
  }

  if (history.state.prev?.path != route && history.state.path != route) {
    history.pushState({ prev: history.state, path: route, popup }, "", route);
  }
};

export const back = () => {
  if (!history.state.prev) {
    go("/");
  } else {
    history.back();
  }
};

export const replacePath = (path) => {
  history.replaceState({ ...history.state, path }, "", path);
};
