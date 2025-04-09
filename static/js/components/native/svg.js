import { importSvg } from "../../utils/index.js";

export const svg = (name) => {
  const svg = document.createElement("svg");
  fetch(importSvg(name)).then(async (res) => {
    svg.outerHTML = await res.text();
  });
  return svg;
};
