import { Fetch } from "../utils/fetch.js";
import { svg } from "./native/svg.js";
import div from "./native/div.js";
export const reaction = (type, postData) => {
  const reactionElement = div(`reaction ${type}`).add(
    svg(type),
    div("", postData[type + "s"])
  );
  if (postData.reaction === type) {
    reactionElement.classList.add("reacted");
  }
  return [
    reactionElement,
    (adverse, url) => {
      reactionElement.onclick = () => {
        const isReacted = reactionElement.classList.contains("reacted");
        const isAdverseReacted = adverse.classList.contains("reacted");
        Fetch(url + type, {
          method: isReacted ? "delete" : "post",
        });
        reactionElement.classList.toggle("reacted");
        if (isAdverseReacted) {
          adverse.children[1].textContent--;
        }
        if (isReacted) {
          reactionElement.children[1].textContent--;
        } else {
          reactionElement.children[1].textContent++;
        }
        adverse.classList.remove("reacted");
      };
    },
  ];
};
