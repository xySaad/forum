import { importSvg } from "../utils/index.js";
import div from "./native/div.js";
import img from "./native/img.js";
import { input } from "./native/input.js";

export const Input = (sendFunction) => {
  const submit = async (input) => {
    const value = input.value.trim();
    if (value.length === 0) return;
    await sendFunction(input);
    input.value = ""
  };

  const inputElm = input("commInput", "Write something...");
  inputElm.onkeydown = (e) => {
    if (e.key === "Enter") {
      e.preventDefault();
      submit(e.target);
    }
  };

  const button = document.createElement("button");
  button.className = "commentBtn";
  button.onclick = () => submit(inputElm);
  button.append(img(importSvg("arrow-up")));

  return div("inputwrap").add(inputElm, button);
};
