import { importSvg } from "../../utils/index.js";
const svgCache = new Map();
const parser = new DOMParser();

export const svg = (name) => {
  const svgElem = document.createElement("svg");
  const replaceSvg = (parsedSvg) => {
    svgElem.onload?.(parsedSvg);
    svgElem.replaceWith(parsedSvg.cloneNode(true));
  };

  svgElem.classList.add(name);
  const url = importSvg(name);
  const cachedSvg = svgCache.get(url);

  if (!cachedSvg) {
    const promise = fetchAndParseSvg(url);
    promise.then(replaceSvg);
    svgCache.set(url, promise);
    return svgElem;
  }
  if (cachedSvg instanceof Promise) {
    cachedSvg.then(replaceSvg);
    return svgElem;
  } else {
    return cachedSvg.cloneNode(true);
  }
};

const fetchAndParseSvg = async (url) => {  
  const resp = await fetch(url);
  const text = await resp.text();
  const svgDoc = parser.parseFromString(text, "image/svg+xml");
  const parsedSvg = svgDoc.firstChild;
  svgCache.set(url, parsedSvg);

  return parsedSvg;
};
