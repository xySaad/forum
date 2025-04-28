import { ws } from "../websockets.js";
import div from "./native/div.js";
import img from "./native/img.js";
const USERS_API = "/api/users";
const MESSAGE_API = "/api/msg"

const getActiveUsers = async (parentNode) => {
  const resp = await fetch(USERS_API);
  const users = await resp.json();
  console.log(users);
  if (!users.length || users.length < 2) {
    parentNode.append(div("fallback", "It's lonely right here!\nno users."));
    return;
  }
  const ownUserId = document.querySelector(".profile").id;
  users.forEach((user) => {
    if (user.id === ownUserId) return;
    let userholder = div(`user uid-${user.id}`)
    userholder.onclick = () => {
      ws.send(
        JSON.stringify({
          type : "message fetch",
          id: ownUserId,
          receiver: user.id,
        })
      ),
      message(user.id)
    };


    parentNode.add(
      userholder.add(
        div("publisher").add(
          img(user.profilePicture, "no-profile"),
          div("username", user.username),
          div(`status ${user.status}`, user.status)
        )
      )
    );
  });
};

export function message(user) {
  const oldConversation = document.querySelector(".ConversationArea");
  if (oldConversation) {
    oldConversation.remove();
  }

  const main = document.querySelector("main");

  let creatbtn = document.createElement("button");
  creatbtn.onclick = async () => {

    const ownUserId = document.querySelector(".profile").id;
    console.log(user);
    ws.send(
      JSON.stringify({
        type : "send message",
        id: ownUserId,
        receiver: user,
        msg: inputmsg.value
      })
    );


  }
  creatbtn.textContent = "Send";

  console.log("Button clicked");

  let inputmsg = document.createElement('input');
  inputmsg.type = 'text';
  inputmsg.className = "input";

  const conversation = document.createElement('div');
  conversation.className = "ConversationArea";

  const msgdiv = document.createElement('div');
  msgdiv.className = "msgdiv";

  const inputContainer = document.createElement('div');
  inputContainer.className = "inputmsg";
  inputContainer.appendChild(inputmsg);
  inputContainer.appendChild(creatbtn);

  conversation.appendChild(msgdiv);
  conversation.appendChild(inputContainer);

  main.appendChild(conversation);
}
export const ActiveUsers = () => {
  const usersContainer = div("users");
  getActiveUsers(usersContainer);
  return usersContainer.add(div("title", "Users"));
};
