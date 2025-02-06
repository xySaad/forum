"use strict";
import { appendUserHeader } from "./components/Headers.js";
import { appendGuestHeader } from "./components/Headers.js";
import ensureAuth from "./utils/ensureAuth.js";
import { AddRoute, go } from "./router.js";
import { Home } from "./home.js";
import { Fetch } from "./utils/fetch.js";
import CreatePost from "./components/createPost.js";
import Auth from "./components/Auth.js";

AddRoute("/", Home);
AddRoute("/create-post", CreatePost);
AddRoute("/login", Auth, "login");
AddRoute("/register", Auth, "register");

window.onpopstate = () => {
  go(window.location.pathname);
};

window.addEventListener("DOMContentLoaded", () => {
  go(window.location.pathname);
});

const main = async () => {
  await Fetch("/api/auth/session/");
  if (ensureAuth()) {
    appendUserHeader();
  } else {
    appendGuestHeader();
  }
};

main();
