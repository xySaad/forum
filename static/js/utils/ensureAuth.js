import Auth from "../components/Auth.js";

const ensureAuth = async () => {
  const authForm = Auth();
  const homePage = document.querySelector(".homePage");
  const resp = await fetch("/api/auth/session/");
  if (resp.status === 401) {
    homePage.prepend(authForm);
    return false;
  }

  return true;
};

export default ensureAuth;
