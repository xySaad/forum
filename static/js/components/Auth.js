import img from "./native/img.js";
import { importSvg } from "../utils/index.js";
import div from "./native/div.js";
import { appendUserHeader } from "./Headers.js";
import { changeAuthState } from "../utils/ensureAuth.js";
import { back, replacePath } from "../router.js";
import { NewReference } from "../utils/reference.js";
import {
  validateEmail,
  validatePassword,
  validateUsername,
} from "../utils/auth_validation.js";

const input = (type, confirm) => {
  const input = document.createElement("input");
  input.required = true;
  input.type = type == "name" ? "text" : type;
  input.className = "input";
  input.placeholder = (confirm ? "Confirm" : "Enter") + " your " + type;
  return input;
};

const createRegisterForm = (authElement, context) => {
  const form = document.createElement("form");
  const username = input("name");
  const password = input("password");
  const email = input("email");
  const confirmPassword = input("password", true);
  const loginButton = document.createElement("button");
  loginButton.textContent = "Submit";

  const cancelButton = document.createElement("button");
  cancelButton.className = "secondary";
  cancelButton.textContent = "Cancel";
  cancelButton.onclick = () => {
    authElement.cleanup();
  };
  const errDisplay = div("errorPlace");
  form.onsubmit = async (e) => {
    e.preventDefault();

    const errors = [
      validateUsername(username.value, context()),
      validateEmail(email.value, context() === "register"),
      validatePassword(password.value,confirmPassword.value, context() === "register"),
    ].filter((value) => value);

    if (errors.length > 0) {
      errDisplay.textContent = errors[0];
      return;
    }

    const resp = await fetch("/api/auth/" + context(), {
      method: "POST",
      body: JSON.stringify({
        username: username.value,
        email: email.value,
        password: password.value,
      }),
    });
    if (resp.ok) {
      changeAuthState(true);
      appendUserHeader();
      authElement.cleanup();
      cancelButton.onclick = null;
      const notification = document.createElement("div");
      notification.classList.add("notification");
      notification.innerText = "Authenticated Successfully ✓";
      document.body.appendChild(notification);
      setTimeout(() => {
        notification.remove();
      }, 3000);
    } else {
      let nn = await resp.json();
      errDisplay.innerText = nn.details;
    }
  };

  form.append(
    username,
    email,
    password,
    confirmPassword,
    errDisplay,
    div("btns").add(loginButton, cancelButton)
  );

  return {
    form,
    email,
    confirmPassword,
    errDisplay,
    reset: () => {
      form.innerHTML = "";
      form.append(
        username,
        email,
        password,
        confirmPassword,
        errDisplay,
        div("btns").add(loginButton, cancelButton)
      );
    },
  };
};

const Auth = (authType) => {
  let authElement = div("auth");
  const context = NewReference("register");
  const registerForm = createRegisterForm(authElement, context);

  const changeContext = (registerForm) => {
    registerForm.errDisplay.innerHTML = "";
    if (context() == "register") {
      context("login");
      registerForm.email.remove();
      registerForm.confirmPassword.remove();
      registerForm.registerSpan.classList.remove("clicked");
      registerForm.loginSpan.classList.add("clicked");
    } else {
      registerForm.loginSpan.classList.remove("clicked");
      registerForm.registerSpan.classList.add("clicked");
      context("register");
      registerForm.reset();
    }
  };

  const loginSpan = document.createElement("span");
  loginSpan.className = "login";
  loginSpan.textContent = "login";
  loginSpan.onclick = () => {
    if (context() != "login") {
      replacePath("/login");
      changeContext(registerForm);
    }
  };

  registerForm.loginSpan = loginSpan;

  const registerSpan = document.createElement("span");
  registerForm.registerSpan = registerSpan;

  registerSpan.className = "register clicked";
  registerSpan.textContent = "register";
  registerSpan.onclick = () => {
    if (context() != "register") {
      replacePath("/register");
      changeContext(registerForm);
    }
  };

  authElement.cleanup = () => {
    back();
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

  if (authType && authType != context()) {
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
