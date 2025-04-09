import { importSvg } from "../../utils/index.js";
const parser = new DOMParser();

export const svg = (name) => {
  const svg = document.createElement("svg");
  svg.classList.add(name);
  
  fetch(importSvg(name)).then(async (res) => {
    const text = await res.text();
    const svgDoc = parser.parseFromString(text, "image/svg+xml");
    const parsedSvg = svgDoc.firstChild;
    svg.onload?.(parsedSvg);
    svg.replaceWith(parsedSvg);
  });
  return svg;
};
