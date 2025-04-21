import { toggleIcon } from "../components/Headers.js";
import div from "../components/native/div.js";

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
        UserArea.append(
          div("userholder").add(
            div("profilepic").add(img),
            user.username
          )
        );
      }
      nextUrl = new URL(resp.url);
    }
  } catch (error) {
    console.error("Error fetching users:", error);
  }
}
