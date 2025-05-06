import { importSvg } from "../../utils/index.js";

const img = (src, alt, className, id) => {
  const imgElement = document.createElement("img");
  imgElement.src = src ?? "";
  imgElement.alt = alt ?? "";
  imgElement.className = className ?? "";
  imgElement.id = id ?? "";

  imgElement.onerror = () => {
    imgElement.onerror = null;    
    imgElement.src = importSvg(alt) ?? imgElement.remove();
  };
  return imgElement;
};
export default img;
