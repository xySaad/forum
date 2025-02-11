import div from "./components/native/div.js";

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

const routesByLevel = [
  {
    404: () => div("404", "404 - page not found"),
  },
];

let Params = {};

export const GetParams = () => {
  const result = {};

  Object.keys(Params).forEach((key) => {
    result[key] = trimSlash(location.pathname).split("/")[Params[key]];
  });

  return result;
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

    let params = routesByLevel[i - 1]
      ? { ...routesByLevel[i - 1][splitedRoute[i - 1]].params }
      : {};

    if (isArg) {
      params[path.slice(1)] = i;
      path = splitedRoute[i - 1] + "/*";
    }

    routesByLevel[i][path] = { page: pageToAdd, params };
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
      if (!routesByLevel[i][path] || !routesByLevel[i][path]) {
        if (routesByLevel[i][splitedRoute[i - 1] + "/*"]) {
          return {
            found: true,
            ...routesByLevel[i][splitedRoute[i - 1] + "/*"],
          };
        }
        return { found: false, ...routesByLevel[0]["404"] };
      }
      return { found: true, ...routesByLevel[i][path] };
    }

    if (!routesByLevel[i][path]) {
      return { found: false, ...routesByLevel[0]["404"] };
    }
  }
};

export const go = (route, popup, ...args) => {
  Params = {};
  console.log(routesByLevel);

  const { found, page, params } = routeLookup(route);
  Params = params;
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
