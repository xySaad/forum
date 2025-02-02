import { onResize } from "./events.js";
import { timePassed } from "./time.js";
const importSvg = (svgName) => svgName ? "./static/svg/" + svgName + ".svg" : "";

export { onResize, timePassed, importSvg };
