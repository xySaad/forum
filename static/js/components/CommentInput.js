import { importSvg } from "../utils/index.js";
import div from "./div.js";
import img from "./img.js";

export const CommentInput = () => {
  const input = document.createElement("input"),
    button = document.createElement("button");
  button.append(img(importSvg("arrow-up")));
  button.onclick = async () => {
    const body = {};
    const resp = await fetch("/api/coments", {
      method: "post",
      body: JSON.stringify(body),
    });
  };

  return div("inputwrap").add(input, button);
};
