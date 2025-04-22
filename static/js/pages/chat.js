import { toggleIcon } from "../components/Headers.js";
import div from "../components/native/div.js";
import img from "../components/native/img.js";
import { input } from "../components/native/input.js";
import { Fetch } from "../utils/fetch.js";
import { resolve } from "../utils/promise.js";

export const Chat = () => {
  const chatBubble = document.querySelector(".chat-bubble");
  toggleIcon(".chat-bubble");
  chatBubble.onload = (svg) => svg.classList.add("active");
  return infinitusers("/api/users");
};

export const infinitusers = (url) => {
  const activeUsers = div("users");
  getusers(activeUsers, url);
  return activeUsers;
};

export async function getusers(activeUsers, url) {
  const [resp, err] = await resolve(Fetch(url));
  if (err != null) {
    console.error(err);
    return;
  }

  const json = await resp.json();
  json.forEach((user) => {
    activeUsers.add(
      div("userholder").add(
        div("profilepic").add(img(user.profilePicture, "no-profile")),
        div("username", user.username)
      )
    );
  });
}
export function conversation() {
  return callfunction(infinitusers);
}
export function callfunction(func1) {
  func1();
  let creatbtn = document.createElement("button");
  creatbtn.value = "send";
  let inputmsg = input("msj-input", "message", true);
  const conversation = div("ConversationArea").add(
    div(msgdiv),
    div("inputmsg").add(inputmsg, creatbtn)
  );
  return conversation;
}
