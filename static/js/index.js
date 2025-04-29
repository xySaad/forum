"use strict";
import { appendUserHeader } from "./components/Headers.js";
import { appendGuestHeader } from "./components/Headers.js";
import ensureAuth, { changeAuthState } from "./utils/ensureAuth.js";
import { AddRoute, go } from "./router.js";
import { Home } from "./pages/home.js";
import CreatePost from "./components/createPost.js";
import Auth from "./components/Auth.js";
import PostView from "./components/PostView.js";
import { CreatedPosts, LikedPosts } from "./pages/user-posts.js";
import { Chat } from "./pages/chat.js";
import { ActiveUsers } from "./components/ActiveUsers.js";
import { InitWS } from "./websockets.js";
import users from "./context/users.js";
AddRoute("/", Home);
AddRoute("/create-post", CreatePost, true);
AddRoute("/login", () => Auth("login"), true);
AddRoute("/register", () => Auth("register"), true);
AddRoute("/post/:id", PostView, true);
AddRoute("/created-posts", CreatedPosts);
AddRoute("/liked-posts", LikedPosts);
AddRoute("/chat", Chat);
// AddRoute("/chat/:id", conversation);

window.onpopstate = () => {
  go(window.location.pathname);
};

const main = async () => {
  const resp = await fetch("/api/profile");
  if (resp.ok) {
    changeAuthState(true);
    users.myself = await resp.json();
    await appendUserHeader();
  } else {
    appendGuestHeader();
  }
  go(window.location.pathname);
  const main = document.querySelector("main");
  main.insertAdjacentElement("beforebegin", ActiveUsers());

  InitWS();
};

main();
