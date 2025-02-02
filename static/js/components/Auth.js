import img from "./native/img.js";
import { importSvg } from "../utils/index.js";
import div from "./native/div.js";
import { appendUserHeader } from "./Headers.js";
import { changeAuthState } from "../utils/ensureAuth.js";
let context = "register";
let authElement = null;

const input = (type, confirm) => {
  const input = document.createElement("input");
  input.required = true;
  input.type = type == "name" ? "text" : type;
  input.className = "input";
  input.placeholder = (confirm ? "Confirm" : "Enter") + " your " + type;
  return input;
};

const createRegisterForm = () => {
  const form = document.createElement("form");
  const username = input("name");
  const password = input("password");
  const email = input("email");
  const confirmPassword = input("password", true);
  const loginButton = document.createElement("button");
  loginButton.textContent = "Login";

  const cancelButton = document.createElement("button");
  cancelButton.className = "secondary";
  cancelButton.textContent = "Cancel";
  cancelButton.onclick = () => {
    authElement.cleanup();
  };

  form.onsubmit = async (e) => {
    e.preventDefault();
    const resp = await fetch("/api/auth/" + context, {
      method: "POST",
      body: JSON.stringify({
        username: username.value,
        email: email.value,
        password: password.value,
      }),
    });
    if (resp.ok) {
      changeAuthState(true)
      appendUserHeader();
      authElement.cleanup();
      cancelButton.onclick = null;
    }
  };

  form.append(
    username,
    email,
    password,
    confirmPassword,
    div("btns").add(loginButton, cancelButton)
  );

  return {
    form,
    email,
    confirmPassword,
    reset: () => {
      form.innerHTML = "";
      form.append(
        username,
        email,
        password,
        confirmPassword,
        div("btns").add(loginButton, cancelButton)
      );
    },
  };
};
const changeContext = (registerForm) => {
  if (context == "register") {
    context = "login";
    registerForm.email.remove();
    registerForm.confirmPassword.remove();
    registerForm.registerSpan.classList.remove("clicked");
    registerForm.loginSpan.classList.add("clicked");
  } else {
    registerForm.loginSpan.classList.remove("clicked");
    registerForm.registerSpan.classList.add("clicked");
    context = "register";
    registerForm.reset();
  }
};

const Auth = (authType) => {
  const registerForm = createRegisterForm();
  const loginSpan = document.createElement("span");
  loginSpan.className = "login";
  loginSpan.textContent = "login";
  loginSpan.onclick = () =>
    context == "register" ? changeContext(registerForm) : null;

  registerForm.loginSpan = loginSpan;

  const registerSpan = document.createElement("span");
  registerSpan.className = "register clicked";
  registerSpan.textContent = "register";
  registerSpan.onclick = () =>
    context == "login" ? changeContext(registerForm) : null;
  registerForm.registerSpan = registerSpan;

  authElement = div("auth");
  authElement.cleanup = () => {
    context = "register";
    loginSpan.onclick = null;
    registerSpan.onclick = null;
    authElement.remove();
    authElement = null;
  };

  const authentication = div("authentication");
  authentication.onclick = (e) => {
    if (e.target == authentication) {
      authElement.cleanup();
    }
  };

  if (authType && authType != context) {
    changeContext(registerForm);
  }

  return authElement.add(
    div("full-screen-background"),
    div("blur-layer"),
    authentication.add(
      div("card").add(
        div("tocenter").add(img(importSvg("logo"))),
        div("toggle").add(registerForm.loginSpan, registerForm.registerSpan),
        div("register").add(registerForm.form)
      )
    )
  );
};

export default Auth;
