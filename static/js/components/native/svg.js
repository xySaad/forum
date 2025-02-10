import { importSvg } from "../../utils/index.js";

export const svg = (name) => {
  const svg = document.createElement("div");
  fetch(importSvg(name)).then(async (res) => {
    svg.innerHTML = await res.text();
  });
  return svg;
};
