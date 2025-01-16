const img = (src, alt, className) => {
  const imgElement = document.createElement("img");
  imgElement.src = src;
  imgElement.alt = alt;
  imgElement.className = className;
  return imgElement;
};
export default img;
