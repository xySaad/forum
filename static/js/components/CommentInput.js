import { importSvg } from "../utils/index.js";
import div from "./native/div.js";
import img from "./native/img.js";

export const CommentInput = (postID) => {
  const input = document.createElement("input");
  input.placeholder = "Write a comment...";
  const sendComment = async () => {
    if (input.value.length == 0) {
      return;
    }
    const body = {
      post_id: postID,
      content: input.value,
    };
    const resp = await fetch("/api/coments", {
      method: "post",
      body: JSON.stringify(body),
    });
    if (resp.ok) {
      input.value = "";
    }
  };

  const button = document.createElement("button");
  button.className ="commentBtn"
  button.append(img(importSvg("arrow-up")));
  input.addEventListener("keydown", (event) => {
    if (event.key === "Enter") {
      event.preventDefault();
      sendComment();
    }
  });

  button.onclick = sendComment;

  return div("inputwrap").add(input, button);
};
