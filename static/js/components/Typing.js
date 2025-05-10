import div from "./native/div.js";

export const Typing = () => {
  return div("indicator typing message").add(
    div("dot"),
    div("dot"),
    div("dot")
  );
};
