const div = (className = "", textContent) => {
  const divElement = document.createElement("div");
  divElement.className = className;
  divElement.textContent = textContent;

  divElement.add = (...args) => {
    divElement.append(...args);
    return divElement;
  };
  return divElement;
};
export default div;
