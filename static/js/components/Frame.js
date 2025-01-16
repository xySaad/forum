import div from "./div.js";

const Frame = (HTMLelement) => {
  HTMLelement.append(
    div("frame").add(
        div("top"), 
        div("bottom")
    )
);

return HTMLelement
};

export default Frame;
