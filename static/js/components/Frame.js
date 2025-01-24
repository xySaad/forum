import div from "./native/div.js";

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
