import { toggleIcon } from "../components/Headers.js";
import div from "../components/native/div.js";

export const Chat = () => {
  const chatBubble = document.querySelector(".chat-bubble");
  toggleIcon(".chat-bubble");

  chatBubble.onload = (svg) => {
    svg.classList.add("active");
  };
  return div("chat")
};