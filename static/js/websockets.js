import div from "./components/native/div.js";
import { query } from "./components/native/index.js";
import { Typing } from "./components/Typing.js";
import { UserCard } from "./components/UserCard.js";
import users from "./context/users.js";
import { Message } from "./pages/chat.js";
import { GetParams } from "./router.js";
import { Deferred } from "./utils/Deferred.js";
export const MESSAGE_TYPE = {
  DM: "DM",
  STATUS: "STATUS",
};

const WS_API = "/api/ws";

export let ws;
const handleMessage = (e) => {
  const msg = JSON.parse(e.data);
  if (msg.type === "error") {
    query("popup").append(div("error", msg.value));
    return;
  }

  const { id } = GetParams();
  const openChat = users.get(msg.id);
  const messages = query(".messages");
  switch (msg.type) {
    case MESSAGE_TYPE.STATUS:
      if (msg.value === "afk") {
        openChat.isTyping = false;
        msg.value = openChat.status;
        query(".indicator.typing")?.remove();
      } else if (msg.value === "typing") {
        openChat.isTyping = true;
        console.log(msg.id, id);

        if (msg.id === id) messages.add(Typing());
      } else {
        // msg.value === "online" || "offline"
        openChat.status = msg.value;
      }
      const userStatus = document.querySelectorAll(
        `.user.uid-${msg.id} .status`
      );
      if (!userStatus) return;
      userStatus.forEach((user) => {
        user.className = `status ${msg.value}`;
        user.textContent = msg.value;
      });
      break;
    case MESSAGE_TYPE.DM:
      if (msg.sender !== users.myself.id && msg.sender !== id) {
        query(".notification.message")?.remove();
        const notification = div("notification message").add(Message(msg));
        query("popup").append(notification);
        setTimeout(() => {
          notification.remove();
        }, 2000);
      }
      if (msg.sender === id || msg.chat === id) {
        messages.prepend(Message(msg));
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
