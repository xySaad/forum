import div from "./components/native/div.js";
import { query } from "./components/native/index.js";
import { Typing } from "./components/Typing.js";
import { UserCard } from "./components/UserCard.js";
import users from "./context/users.js";
import { Message } from "./pages/chat.js";
import { GetParams } from "./router.js";
import { Deferred } from "./utils/Deferred.js";
const WS_API = "/api/ws";
export const MESSAGE = {
  TYPE: {
    DM: "DM",
    STATUS: "status",
    Action: "action",
    NEW_USER: "new_user",
  },
  STATUS: {
    ONLINE: "online",
    OFFLINE: "offline",
  },
  ACTION: {
    LOGOUT: "logout",
  },
  LOGOUT_REASON: {
    NEW_LOGIN: "new_login",
  },
};

export const updateStatus = (id, status) => {
  const userStatus = document.querySelectorAll(`.user.uid-${id} .status`);
  if (!userStatus) return;
  userStatus.forEach((user) => {
    user.className = `status ${status}`;
    user.textContent = status;
  });
};

export let ws;
const handleMessage = (e) => {
  const json = JSON.parse(e.data);
  const { type, data: msg } = json;
  const { id: chatId } = GetParams();
  const IncomingChat = users.get(msg.id);
  const messages = query(".messages");
  switch (type) {
    case MESSAGE.TYPE.STATUS:
      const id = msg.id;
      let status = msg.status;
      switch (status) {
        case "afk":
          IncomingChat.isTyping = false;
          status = IncomingChat.status;
          query(".indicator.typing")?.remove();
          break;
        case "typing":
          if (IncomingChat.isTyping) {
            break;
          }
          IncomingChat.isTyping = true;
          if (id === chatId) messages.add(Typing());
          break;
        case MESSAGE.STATUS.ONLINE:
        case MESSAGE.STATUS.OFFLINE:
          IncomingChat.status = status;
          console.log("updated status");
          break;
        default:
          console.error("Invalid status", status);
          return;
      }
      updateStatus(id, status);
      break;
    case MESSAGE.TYPE.DM:
      const { sender, chat } = msg;
      if (sender !== users.myself.id && sender !== chatId) {
        query(".notification.message")?.remove();
        const notification = div("notification message").add(Message(msg));
        query("popup").append(notification);
        setTimeout(() => {
          notification.remove();
        }, 2000);
      }
      if (sender === chatId || chat === chatId) {
        messages.prepend(Message(msg));
      }
      users.get(chatId).lastMessage = msg; // useless
      const userElem =
        query(`.users .user.uid-${chat}`) ||
        query(`.users .user.uid-${sender}`);
      query(".users .head").insertAdjacentElement("afterend", userElem);
      break;
    case MESSAGE.TYPE.NEW_USER:
      users.add(msg);
      query(".users").add(UserCard(msg));
      break;
    case MESSAGE.TYPE.Action:
      const { action } = msg;
      switch (action) {
        case MESSAGE.ACTION.LOGOUT:
          location.reload();
          break;
        default:
          console.error("invalid action");
          break;
      }
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
