import { importSvg } from "../utils/index.js";

const img = (src, alt, className, id) => {
  const imgElement = document.createElement("img");
  imgElement.src = src;
  imgElement.alt = alt;
  if (className) {
    imgElement.className = className;
  }
  if (id) {
    imgElement.id = id;
  }

  imgElement.onerror = (e) => {
    e.target.src = importSvg("no-profile");
  };
  return imgElement;
};
export default img;
