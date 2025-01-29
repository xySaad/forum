import Auth from "../components/Auth.js";

export const ensureAuth = async () => {
  console.log("imported!!");
  
  const homePage = document.querySelector(".homePage");
  const resp = await fetch("/api/auth/session/");
  if (resp.status === 401) {
    const authForm = Auth();
    homePage.prepend(authForm);
    return false;
  }
  return true;
};

