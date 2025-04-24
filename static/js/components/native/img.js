import { importSvg } from "../../utils/index.js";

const img = (src, alt, className, id) => {
  const imgElement = document.createElement("img");
  imgElement.src = src ?? "";
  imgElement.alt = alt ?? "";
  imgElement.className = className ?? "";
  imgElement.id = id ?? "";

  imgElement.onerror = () => {
    imgElement.onerror = null;
    console.log(imgElement.dataset.alt = true);
    
    imgElement.src = importSvg(alt) ?? imgElement.remove();
  };
  return imgElement;
};
export default img;
