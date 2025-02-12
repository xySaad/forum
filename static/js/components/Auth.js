import img from "./native/img.js";
import { importSvg } from "../utils/index.js";
import div from "./native/div.js";
import { appendUserHeader } from "./Headers.js";
import { changeAuthState } from "../utils/ensureAuth.js";
import { back, replacePath } from "../router.js";
import { NewReference } from "../utils/reference.js";

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
  password.minLength = "6"
  confirmPassword.minLength = "6"
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
    if (confirmPassword.value && password.value !== confirmPassword.value) {
      document.querySelector(".errors").innerText = "the Passwords aren't identical"
      return
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
      notification.innerText="Authenticated Successfully ✓"
      document.body.appendChild(notification);
      setTimeout(() => {
        notification.remove();
      }, 3000);
    } else {
      let nn = await resp.json()
      document.querySelector(".errors").innerText = nn.message
    }
  };

  form.append(
    username,
    email,
    password,
    confirmPassword,
    div("errors"),
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

const Auth = (authType) => {
  let authElement = div("auth");
  const context = NewReference("register");
  const registerForm = createRegisterForm(authElement, context);

  const changeContext = (registerForm) => {
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
