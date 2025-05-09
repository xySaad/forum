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
      const userStatus = query(`.user.uid-${msg.id} .status`);
      if (!userStatus) return;
      userStatus.className = `status ${msg.value}`;
      userStatus.textContent = msg.value;
      break;
    case "STATUS" :
      let div = query(".publisher");
      let oldTp = div.querySelector(".tp");
      if (oldTp) oldTp.remove();
      let msg = { value: "Your message here" }; 
      let typ = document.createElement("div");
      typ.className = "tp";
      typ.textContent = msg.value;
      div.appendChild(typ);
    case "DM":
      const { id } = GetParams();
      if (msg.sender !== users.myself.id && msg.sender !== id) {
        query(".notification.message")?.remove();
        const notification = div("notification message").add(Message(msg));
        query("popup").append(notification);
        setTimeout(() => {
          notification.remove();
        }, 2000);
      }
      if (msg.sender === id || msg.chat === id) {
        query(".messages").prepend(Message(msg));
      }
      users.get(id).lastMessage = msg; // useless
      const userElem =
        query(`.users .user.uid-${msg.chat}`) ||
        query(`.users .user.uid-${msg.sender}`);
      query(".users .title").insertAdjacentElement("afterend", userElem);
      break;
    case "user":
      users.add(msg.value);
      query(".users").add(UserCard(msg.value));
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
