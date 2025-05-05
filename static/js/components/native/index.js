export const asyncAppend = function (...children) {
  (async () => {
    const settled = await Promise.allSettled(children);
    const results = settled
      .filter((p) => p.status === "fulfilled")
      .map((p) => p.value);
    this.append(...results);
  })();
  return this;
};

export const query = (selector) => {
  const element = document.querySelector(selector);
  if (!element) return element;
  element.add = asyncAppend;
  element.on = function (eventName, callback) {
    element["on" + eventName] = callback;
    return this;
  };
  return element;
};
