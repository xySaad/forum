import img from "./native/img.js";
import { importSvg } from "../utils/index.js";
import div from "./native/div.js";

const passwordInput = (params) => {
  const { confirm } = params ?? {};
  const input = document.createElement("input");
  input.classList = "input tall";
  input.type = "password";

  input.placeholder = confirm ? "Confirm" : "Enter your";
  input.placeholder += " password";
  return input;
};

const userIdInput = (type = "") => {
  const input = document.createElement("input");
  input.classList = "input";
  input.type = type;
  input.placeholder = "Enter your " + type;
  return input;
};

const createRegisterForm = () => {
  const form = document.createElement("form");
  const password = passwordInput();
  const confirmPassword = passwordInput({ confirm: true });
  const username = userIdInput("name");
  const email = userIdInput("email");
  const loginButton = document.createElement("button");
  loginButton.type = "submit";
  loginButton.textContent = "Login";

  const cancelButton = document.createElement("button");
  cancelButton.classList = "secondary";
  cancelButton.textContent = "Cancel";
  cancelButton.onclick = () => {};

  form.append(
    username,
    email,
    password,
    confirmPassword,
    div("btns").add(loginButton, cancelButton)
  );
  return { form, email, confirmPassword };
};

const Auth = () => {
  const registerForm = createRegisterForm();
  registerForm.form.onsubmit = (e)=>{
    e.preventDefault()
  }
  
  const loginSpan = document.createElement("span");
  loginSpan.className = "login clicked";
  loginSpan.textContent = "login";
  loginSpan.onclick = () => {
    loginSpan.classList.add("clicked");
    registerSpan.classList.remove("clicked");

    registerForm.email.style.display = "none";
    registerForm.confirmPassword.style.display = "none";
  };

  const registerSpan = document.createElement("span");
  registerSpan.className = "register";
  registerSpan.textContent = "register";
  registerSpan.onclick = () => {
    registerSpan.classList.add("clicked");
    loginSpan.classList.remove("clicked");

    registerForm.email.style.display = "";
    registerForm.confirmPassword.style.display = "";
  };

  return div("auth").add(
    div("full-screen-background"),
    div("blur-layer"),
    div("authontication").add(
      div("card").add(
        div("tocenter").add(img(importSvg("logo"))),
        div("toggle").add(loginSpan, registerSpan),
        div("register").add(registerForm.form)
      )
    )
  );
};

export default Auth;
