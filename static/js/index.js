import { Home } from "./home.js";
import { appendUserHeader } from "./components/Headers.js";
import { appendGuestHeader } from "./components/Headers.js";
import ensureAuth from "./utils/ensureAuth.js";

if (await ensureAuth()) {
 await appendUserHeader()
}else{
  appendGuestHeader()
}
Home()