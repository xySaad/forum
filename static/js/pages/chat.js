import { toggleIcon } from "../components/Headers.js";
import div from "../components/native/div.js";
import { input } from "../components/native/input.js";

export const Chat = () => {
  const chatBubble = document.querySelector(".chat-bubble");
  toggleIcon(".chat-bubble");
  chatBubble.onload = (svg) => {
    svg.classList.add("active");
  };
  return infinitusers("/api/users")

};
export const infinitusers = (url) => {
  const UserArea = div("users");
  getusers(UserArea, url);
  return UserArea;
}
export async function getusers(UserArea, url) {
  //let nextUrl
  let json
  try {
    const resp = await fetch(url);
    if (resp.status === 401) {
      go("/login");
      return;
    }
    if (!resp.ok) {
      throw new Error("status not ok:", resp.status);
    }
    json = await resp.json();
    if (json) {

      for (const user of json) {
        let img = document.createElement("img");
        img.src = user.profilePicture;
        let jj = div("userholder")
        jj.onclick(conversation)
        jj.add(
          div("profilepic").add(img),
          user.username
        )
        UserArea.append(jj);
      }
      nextUrl = new URL(resp.url);
    }
  } catch (error) {
    console.error("Error fetching users:", error);
  }
}
export function conversation() {
  return callfunction(infinitusers)
}
export function callfunction(func1) {
  func1()
  let creatbtn = document.createElement("button").value("send")
  let inputmsg = input("msj-input", "message", true)
  const conversation = div("ConversationArea").add(
    div(msgdiv),
    div("inputmsg").add(inputmsg, creatbtn),
  )
  return conversation
}