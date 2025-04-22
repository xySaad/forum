import { go } from "../router.js";
import ensureAuth, { changeAuthState } from "./ensureAuth.js";

export const Fetch = async (...args) => {
  if (!ensureAuth(true)) {
    return;
  }

  const resp = await fetch(...args);
  if (resp.status == 401) {
    changeAuthState(false);
    go("/login", true);
  } else if (!resp.ok) {
    throw new Error("status not ok:", resp.status);
  }
  return resp;
};
