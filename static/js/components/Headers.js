import div from "./native/div.js";
import img from "./native/img.js";
import { changeAuthState } from "../utils/ensureAuth.js";
import { go } from "../router.js";
import { svg } from "./native/svg.js";
import users from "../context/users.js";

export async function toggleIcon(type) {
  let icns = document.querySelectorAll(".icon");
  icns.forEach((icn) => icn.classList.remove("active"));
  let clicked = document.querySelector(type);
  clicked?.classList.add("active");
}
function toggleIt() {
  let ul = document.querySelector(".icons");
  ul.classList.toggle("active");
}
function ToggleDisplay() {
  let item = document.querySelector(".profileCard");
  if (item.style.display === "none" || item.style.display === "") {
    item.style.display = "flex";
  } else {
    item.style.display = "none";
  }
}
async function Logout() {
  const resp = await fetch(`/api/auth/logout`, { method: "POST" });
  if (!resp.ok) {
    console.error("haven't logged out!");
    return;
  } else {
    changeAuthState(false);
    appendGuestHeader();
    location.reload();
  }
}

export async function appendUserHeader() {
  const head = document.querySelector("header");
  head.innerHTML = "";

  const icn1 = div("contain").add(svg("home"));
  const icn2 = div("contain").add(svg("heart"));
  const icn3 = div("contain").add(svg("pen"));
  const chatBubble = div("contain").add(svg("chat-bubble"));
  const logout = div("logoutBtn").add(svg("logout"), div(null, "Logout"));

  const h2 = document.createElement("h2");
  h2.innerText = users.myself.username;

  const h4 = document.createElement("h4");
  h4.innerText = users.myself.email;

  const profileContainer = div("profileContainer").add(
    img(users.myself.profilePicture, "avatar", "profile", users.myself.id)
  );
  profileContainer.onclick = ToggleDisplay;

  chatBubble.onclick = () => {
    toggleIcon(".chat-bubble");
    go(`/chat/${users.list[1]?.id ?? ""}`);
  };

  icn1.onclick = () => {
    toggleIcon(".home");
    go("/");
  };
  icn2.onclick = () => {
    toggleIcon(".liked");
    go("/liked-posts");
  };
  icn3.onclick = () => {
    toggleIcon(".created");
    go("/created-posts");
  };
  logout.onclick = Logout;

  const profileCard = div("profileCard").add(
    div("textContainer").add(h2, h4),
    div("line"),
    logout
  );

  const header = div("header").add(
    img("../../static/svg/logo.svg", "logo", "logo"),
    div("close", "☰").on("click", toggleIt),
    div("icons").add(icn1, icn2, icn3, chatBubble),
    profileContainer,
    profileCard
  );

  head.append(await header);
}
export function appendGuestHeader() {
  let head = div("header"); // replaces document.querySelector("header")
  head.innerHTML = "";

  let Content = div("content");
  let Buttons = div("buttons");

  Content.innerHTML =
    `<svg width="44" height="40" viewBox="0 0 44 40" fill="none" xmlns="http://www.w3.org/2000/svg">
      <path fill-rule="evenodd" clip-rule="evenodd" d="M28.1087 4.32368...Z" fill="white"/>
      <path fill-rule="evenodd" clip-rule="evenodd" d="M9.73134 15.498...Z" fill="#FD5F49"/>
    </svg>` +
    `<div class="texts">
      <p>👋 welcome Guest</p>
      <h2>Join Speak to connect with others</h2>
    </div>`;

  Buttons.innerHTML = `
    <button class="secondary">Register</button>
    <button class="primary">login</button>
  `;

  Content.children[0].onclick = () => go("/");

  let header2 = div("toCenter").add(
    div("headerContainer").add(Content, Buttons)
  );

  head.append(header2);

  let loginBtn = Buttons.querySelector(".primary");
  let registerBtn = Buttons.querySelector(".secondary");

  loginBtn.addEventListener("click", () => go("/login"));
  registerBtn.addEventListener("click", () => go("/register"));
}
