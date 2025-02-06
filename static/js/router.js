import div from "./components/native/div.js";

const routesByLevel = [
  {
    404: () => div("404", "404 - page not found"),
  },
];

const trimSlash = (str) => {
  if (str[0] === "/") {
    if (str[str.length - 1] === "/") {
      return str.slice(1, str.length - 1);
    }
    return str.slice(1);
  } else if (str[str.length - 1] === "/") {
    return str.slice(0, str.length - 1);
  }
};

export const AddRoute = (route, page) => {
  const splitedRoute = trimSlash(route).split("/");

  for (let i = 0; i < splitedRoute.length; i++) {
    let path = splitedRoute[i];
    if (!routesByLevel[i]) {
      routesByLevel[i] = [];
    }

    const isArg = path[0] == ":";
    const pageToAdd = i === splitedRoute.length - 1 ? page : null;

    let args = routesByLevel[i - 1]
      ? { ...routesByLevel[i - 1][splitedRoute[i - 1]].args }
      : {};

    if (isArg) {
      args[path.slice(1)] = i;
      path = splitedRoute[i - 1] + "/*";
    }

    routesByLevel[i][path] = { page: pageToAdd, args };
  }
};

/** 
  @param {string} route - the path to lookup for in the router
  @returns {{found: bool, page: HTMLElement}}
*/

const routeLookup = (route) => {
  const splitedRoute = trimSlash(route).split("/");

  for (let i = 0; i < splitedRoute.length; i++) {
    const path = splitedRoute[i];

    if (i === splitedRoute.length - 1) {
      if (!routesByLevel[i][path] || !routesByLevel[i][path].page) {
        if (routesByLevel[i][splitedRoute[i - 1] + "/*"]) {
          return {
            found: true,
            page: routesByLevel[i][splitedRoute[i - 1] + "/*"].page,
          };
        }
        return { found: false, page: routesByLevel[0]["404"] };
      }
      return { found: true, page: routesByLevel[i][path].page };
    }

    if (!routesByLevel[i][path]) {
      return { found: false, page: routesByLevel[0]["404"] };
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
