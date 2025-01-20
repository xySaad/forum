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
    if (!alt) {
      imgElement.remove();
      return;
    }
    e.target.src = importSvg(alt);
    imgElement.onerror = null;
  };
  return imgElement;
};
export default img;
