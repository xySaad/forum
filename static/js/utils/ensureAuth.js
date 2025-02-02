import Auth from "../components/Auth.js";
import { Fetch } from "./fetch.js";

let isAuthenticated = null;

/**
  @param {null | boolean} state - The authentication state, which can be null, true, or false.
*/

export const changeAuthState = (state = null) => {
  if (state === null || typeof state == "boolean") {
    isAuthenticated = state;
  } else {
    throw new Error("state can only be null or boolean");
  }
};

const ensureAuth = async () => {
  if (isAuthenticated === null) {
    await Fetch("/api/auth/session/");
  }

  if (!isAuthenticated) {
    const homePage = document.querySelector(".homePage");
    if (!homePage.querySelector(".auth")) {
      homePage.prepend(Auth("login"));
    }
  }

  return isAuthenticated;
};

export default ensureAuth;
