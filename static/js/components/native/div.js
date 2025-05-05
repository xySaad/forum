import { asyncAppend } from "./index.js";

const div = (className = "", textContent) => {
  const divElement = document.createElement("div");
  divElement.className = className;
  divElement.textContent = textContent;
  divElement.on = function (eventName, callback) {
    divElement["on" + eventName] = callback;
    return this;
  };
  divElement.add = asyncAppend;
  return divElement;
};
export default div;
