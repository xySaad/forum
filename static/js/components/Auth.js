import img from "./native/img.js";
import { importSvg } from "../utils/index.js";
import div from "./native/div.js";
import { appendUserHeader } from "./Headers.js";
import { changeAuthState } from "../utils/ensureAuth.js";
import { back, replacePath } from "../router.js";
import { NewReference } from "../utils/reference.js";
import {
  validateAge,
  validateEmail,
  validateFirstname,
  validateGender,
  validateLastname,
  validatePassword,
  validateUsername,
} from "../utils/auth_validation.js";
import { ActiveUsers } from "./ActiveUsers.js";
import { InitWS } from "../websockets.js";
import users from "../context/users.js";

const input = (type, confirm, placeholdername) => {
  const input = document.createElement("input");
  if (type == "select") {
    const select = document.createElement("select");
    select.name = "gender";
    select.className = "input";
    const fruits = ["male", "female", "other", "prefer not to say"];
    fruits.forEach((fruit) => {
      const option = document.createElement("option");
      option.style.color = "blue";
      option.value = fruit.toLowerCase();
      option.text = fruit;
      select.appendChild(option);
    });
    return select;
  }
  input.required = true;
  input.type = type == "name" ? "text" : type;
  input.className = "input";
  input.placeholder =
    (confirm ? "Confirm" : "Enter") + " your " + placeholdername;
  return input;
};

const createRegisterForm = (authElement, context) => {
  const form = document.createElement("form");
  const username = input("name", false, "nickname");
  const age = input("number", false, "age");
  const gender = input("select", false, "gender");
  const firstname = input("name", false, "firstname");
  const lastname = input("name", false, "lastname");
  const password = input("password", false, "password");
  const email = input("email", false, "email");
  const confirmPassword = input("password", true, "password");
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
      validateFirstname(firstname.value, context()),
      validateLastname(lastname.value, context()),
      validateGender(gender.value, context()),
      validateAge(age.value, context()),
      validatePassword(
        password.value,
        confirmPassword.value,
        context() === "register"
      ),

    ].filter((value) => value);

    if (errors.length > 0) {
      errDisplay.textContent = errors[0];
      return;
    }

    const resp = await fetch("/api/auth/" + context(), {
      method: "POST",
      body: JSON.stringify({
        username: username.value,
        age: age.value,
        gender: gender.value,
        firstname: firstname.value,
        lastname: lastname.value,
        email: email.value,
        password: password.value,
      }),
    });
    if (resp.ok) {
      const resp = await fetch("/api/profile/");
      const json = await resp.json();
      users.myself = json;
      await InitWS();
      changeAuthState(true);
      appendUserHeader();
      const main = document.querySelector("main");
      main.insertAdjacentElement("beforebegin", await ActiveUsers());
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
    div("inputContainer").add(username, age),
    div("inputContainer").add(gender, firstname),
    div("inputContainer").add(lastname, email),
    div("inputContainer").add(password, confirmPassword),
    errDisplay,
    div("btns").add(loginButton, cancelButton)
  );

  return {
    form,
    age,
    gender,
    firstname,
    lastname,
    email,
    confirmPassword,
    errDisplay,
    reset: () => {
      form.innerHTML = "";
      form.append(
        div("inputContainer").add(username, age),
        div("inputContainer").add(gender, firstname),
        div("inputContainer").add(lastname, email),
        div("inputContainer").add(password, confirmPassword),
        errDisplay,
        div("btns").add(loginButton, cancelButton)
      );
    },
  };
};

const Auth = (authType) => {
  let authElement = div("auth");
  const context = NewReference(authType);
  const registerForm = createRegisterForm(authElement, context);

  const changeContext = (registerForm) => {
    registerForm.errDisplay.innerHTML = "";
    if (context() == "register") {
      context("login");
      registerForm.age.remove();
      registerForm.firstname.remove();
      registerForm.lastname.remove();
      registerForm.gender.remove();
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
  changeContext(registerForm)
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
