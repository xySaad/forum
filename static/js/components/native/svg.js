import { importSvg } from "../../utils/index.js";
const svgCache = new Map();
const parser = new DOMParser();

export const svg = (name) => {
  const svgElem = document.createElementNS("http://www.w3.org/2000/svg", "svg");
  const replaceSvg = (parsedSvg) => {
    const clonedSvg = parsedSvg.cloneNode(true);
    for (const attr of clonedSvg.getAttributeNames()) {
      svgElem.setAttribute(attr, clonedSvg.getAttribute(attr));
    }
    svgElem.append(...clonedSvg.children);
  };

  svgElem.classList.add(name);
  const url = importSvg(name);
  const cachedSvg = svgCache.get(url);

  if (!cachedSvg) {
    const promise = fetchAndParseSvg(url);
    promise.then(replaceSvg);
    svgCache.set(url, promise);
  } else if (cachedSvg instanceof Promise) {
    cachedSvg.then(replaceSvg);
  } else {
    replaceSvg(cachedSvg)
  }

  return svgElem;
};

const fetchAndParseSvg = async (url) => {
  const resp = await fetch(url);
  const text = await resp.text();
  const svgDoc = parser.parseFromString(text, "image/svg+xml");
  const parsedSvg = svgDoc.firstChild;
  svgCache.set(url, parsedSvg);
  return parsedSvg;
};
