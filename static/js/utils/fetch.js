import { changeAuthState } from "./ensureAuth.js";

export const Fetch = async (...args) => {
    const resp = await fetch(...args)
    if (resp.status == 401) {
        changeAuthState(false)
    } else {
        changeAuthState(true)
    }
    return resp
};
