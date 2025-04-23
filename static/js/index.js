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
AddRoute("/", Home);
AddRoute("/create-post", CreatePost);
AddRoute("/login", () => Auth("login"));
AddRoute("/register", () => Auth("register"));
AddRoute("/post/:id", PostView);
AddRoute("/created-posts", CreatedPosts);
AddRoute("/liked-posts", LikedPosts);
AddRoute("/chat", Chat);

window.onpopstate = () => {
  go(window.location.pathname);
};

const main = async () => {
  const resp = await fetch("/api/profile");
  if (resp.ok) {
    changeAuthState(true);
    const userInfo = await resp.json();
    await appendUserHeader(userInfo);
  } else {
    appendGuestHeader();
  }
  go(window.location.pathname);
  const main = document.querySelector("main");
  main.insertAdjacentElement("beforebegin", ActiveUsers());

  InitWS();
};

main();
