export const input = (className, placeholder, required, type) => {
  const input = document.createElement("input");
  input.className = className;
  input.placeholder = placeholder;
  input.required = !!required;
  input.type = type ?? "text";
  return input;
};
