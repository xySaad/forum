import div from "./components/native/div.js";
import { query } from "./components/native/index.js";
import { UserCard } from "./components/UserCard.js";
import users from "./context/users.js";
import { Message } from "./pages/chat.js";
import { GetParams } from "./router.js";
import { Deferred } from "./utils/Deferred.js";

const WS_API = "/api/ws";
export let ws;
const handleMessage = (e) => {
  const msg = JSON.parse(e.data);
  if (msg.type === "error") {
    query("popup").append(div("error", msg.value));
    return;
  }
  switch (msg.type) {
    case "status":
      console.log(`.user.uid-${msg.id} .status`);

      const userStatus = query(`.user.uid-${msg.id} .status`);
      if (userStatus) {
        userStatus.className = `status ${msg.value}`;
        userStatus.textContent = msg.value;
      } else if (msg.id !== users.myself.id) {
        query(".users").add(UserCard(msg));
      }

      break;
    case "DM":
      if (msg.chat == GetParams().id || msg.chat === users.myself.id) {
        let messages = document.getElementsByClassName("messages")[0];
        messages.prepend(Message(msg));
      }
      break;
    case "logout":
      location.reload();
      break;
    default:
      break;
  }
};
export const InitWS = async () => {
  const deferred = new Deferred();
  const tempWs = new WebSocket(WS_API);

  tempWs.onopen = () => {
    ws = tempWs;
    deferred.resolve();
  };
  tempWs.onerror = deferred.resolve;
  tempWs.onclose = () => {
    deferred.resolve();
    ws = null;
  };
  tempWs.onmessage = handleMessage;
  await deferred.promise;
};
