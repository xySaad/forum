export const onResize = (func, ...args) => {
  func(...args);
  window.addEventListener("resize", () => {
    func(...args);
  });
};
