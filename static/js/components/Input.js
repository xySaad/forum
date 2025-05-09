import { importSvg } from "../utils/index.js";
import div from "./native/div.js";
import img from "./native/img.js";
import { input } from "./native/input.js";
import { ws } from "../websockets.js";
import users from "../context/users.js";
import { GetParams } from "../router.js";
import { Chat } from "../pages/chat.js";


export const Input = (sendFunction) => {
  const { id } = GetParams()

  const submit = async (input) => {
    const value = input.value.trim();
    if (value.length === 0) return;
    await sendFunction(input);
    input.value = ""
  };

  const inputElm = input("commInput", "Write something...");
  let typingTimeout;
  let isTyping = false;

  inputElm.onkeydown = async (e) => {
    if (e.key === "Enter") {
      e.preventDefault();
      sendTypingStatus(false, id);
      clearTimeout(typingTimeout);
      await submit(e.target);
      return
    }

    if (!isTyping) {
      isTyping = true;
      sendTypingStatus(true, id);
    }

    clearTimeout(typingTimeout);
    typingTimeout = setTimeout(() => {
      isTyping = false;
      sendTypingStatus(false, id);
    }, 2000);
  };


  const button = document.createElement("button");
  button.className = "commentBtn";
  button.onclick = () => submit(inputElm);
  button.append(img(importSvg("arrow-up")));

  return div("inputwrap").add(inputElm, button);
};

function sendTypingStatus(isTyping, id) {
  const value = isTyping ? "typing" : "not typing";
  ws.send(JSON.stringify({
    type: "STATUS",
    chat: id,
    value,
  }));
}

