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
  return div("boody").add((infinitusers("/api/users")), conversation())
};

export const infinitusers = (url) => {
  const activeUsers = div("users");
  getusers(activeUsers, url);
  return activeUsers
};

export async function getusers(UserArea, url) {
  //let nextUrl
  let json
  try {
    const resp = await fetch(url);
    if (resp.status === 401) {
      go("/login");
      document.getElementsByClassName
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
        //jj.onclick = conversation
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
  //return callfunction(infinitusers)
  //func1()
  let creatbtn = div("btn").add("send")

  let inputmsg = input("msj-input", "message", true)
  const conversation = div("ConversationArea").add(
    div("msgdiv"),
    div("inputmsg").add(inputmsg, creatbtn),
  )
  let test = document.querySelector("inputmsg")
  /*   if (test !== null){
      test.replaceWith(conversation)
    }else {
      document.querySelector("main").append(conversation);
    }
   */
  return conversation
}
