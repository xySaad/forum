import { importSvg } from "../utils/index.js";
import div from "./native/div.js";
import img from "./native/img.js";
import { input } from "./native/input.js";
import { ws } from "../websockets.js";
import { GetParams } from "../router.js";
export const Input = (sendFunction) => {
  const { id } = GetParams();

  const submit = (input) => {
    const value = input.value.trim();
    if (value.length === 0) return;
    sendFunction(input);
    input.value = "";
  };

  const inputElm = input("commInput", "Write something...");
  let typingTimeout;
  let isTyping = false;

  inputElm.onkeydown = (e) => {
    if (e.key === "Enter") {
      e.preventDefault();
      sendTypingStatus(false, id);
      clearTimeout(typingTimeout);
      submit(e.target);
      isTyping = false;
    }
  };

  inputElm.oninput = () => {
    if (!isTyping) {
      isTyping = true;
      sendTypingStatus(true, id);
    }

    clearTimeout(typingTimeout);
    typingTimeout = setTimeout(() => {
      isTyping = false;
      sendTypingStatus(false, id);
    }, 10000);
  };

  const button = document.createElement("button");
  button.className = "commentBtn";
  button.onclick = () => submit(inputElm);
  button.append(img(importSvg("arrow-up")));

  return div("inputwrap").add(inputElm, button);
};

function sendTypingStatus(isTyping, id) {
  const value = isTyping ? "typing" : "afk";
  ws.send(
    JSON.stringify({
      type: "status",
      chat: id,
      value,
    })
  );
}