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
    }else{
      const notification = document.createElement("div");
      notification.classList.add("notification", type);
      notification.textContent = message;
      notification.style.position = 'fixed';
      notification.style.maxWidth = '100px'
      notification.style.top = '50vh';
      notification.style.right = '35%';
      notification.style.left = '35%';
      notification.style.padding = '15px';
      notification.style.background = '';
      notification.style.color = 'white';
      notification.style.borderRadius = '5px';
      notification.style.boxShadow = '0 4px 6px rgba(0,0,0,0.1)';
      notification.style.zIndex = '1000';
      notification.style.fontSize = '16px';
      notification.style.fontWeight = 'bold';
      document.body.appendChild(notification);

      setTimeout(() => {
        notification.remove();
      }, 3000);
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

const Auth = (authType) => {
  try{
    checkerr(authType)
  } 
  catch {
    
  }
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
