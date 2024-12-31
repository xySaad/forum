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

const addPostIcon = document.querySelector(".addPost");
addPostIcon.addEventListener("click", () => {
  if (!addPostIcon.classList.contains("active")) {
    addPostIcon.classList.add("active");
    setTimeout(() => {
      addPostIcon.classList.remove("active");
    }, 1000);
  }
});
