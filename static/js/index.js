const theme = document.documentElement.attributes.getNamedItem("data-theme");
if (!!localStorage.getItem("theme")) {
  theme.value = localStorage.getItem("theme");
}
const menuIcon = document.querySelector(".menu");
menuIcon.addEventListener("click", () => {
  menuIcon.classList.toggle("active");
});

const themeSwitcher = document.querySelector(".themeSwitcher");
themeSwitcher.addEventListener("click", () => {
  const theme = document.documentElement.attributes.getNamedItem("data-theme");
  theme.value = theme.value == "dark" ? "light" : "dark";
  localStorage.setItem("theme", theme.value);
});
