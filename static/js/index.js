import { Home } from "./home.js";
import { appendUserHeader } from "./components/Headers.js";
import { appendGuestHeader } from "./components/Headers.js";
import ensureAuth from "./utils/ensureAuth.js";
import { go } from "./router.js";

const addPostIcon = document.querySelector(".addPost");
addPostIcon?.addEventListener("click", () => {
  if (!addPostIcon.classList.contains("active")) {
    addPostIcon.classList.add("active");
    setTimeout(() => {
      addPostIcon.classList.remove("active");
    }, 1000);
  }
});
if (await ensureAuth()) {
  appendUserHeader();
} else {
  appendGuestHeader();
}

Home();
go(location.pathname.split("/")[1]);
