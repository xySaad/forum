import { importSvg } from "../utils/index.js";
import div from "./native/div.js";
import img from "./native/img.js";
import { Fetch } from "../utils/fetch.js";

export const Input = (endpoint, feedbackComponent) => {
  const inputWrap = div("inputwrap");
  const input = document.createElement("input");
  input.placeholder = "Write a comment...";
  input.className = "commInput";
  input.id = "commInput";
  const sendComment = async () => {
    if (input.value.trim().length === 0) {
      return;
    }
    const body = {
      content: input.value,
    };
    const resp = await Fetch(endpoint, {
      method: "post",
      body: JSON.stringify(body),
    });
    if (resp.ok) {
      const json = await resp.json();
      inputWrap.parentNode.children[0].prepend(feedbackComponent(json));
      input.value = "";
    }
  };

  const button = document.createElement("button");
  button.className = "commentBtn";
  button.append(img(importSvg("arrow-up")));
  input.addEventListener("keydown", (event) => {
    if (event.key === "Enter") {
      event.preventDefault();
      sendComment();
    }
  });

  button.onclick = sendComment;

  return inputWrap.add(input, button);
};
