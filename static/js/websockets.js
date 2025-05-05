import div from "./components/native/div.js";
import { query } from "./components/native/index.js";
import { UserCard } from "./components/UserCard.js";
import users from "./context/users.js";
import { Deferred } from "./utils/Deferred.js";

const WS_API = "/api/ws";
export let ws;

const handleMessage = (e) => {
  const msg = JSON.parse(e.data);
  if (msg.type === "error") {
    query("popup").append(div("error", msg.value));
    return;
  }
  if (msg.type === "status" && msg.id !== users.myself.id) {
    const userStatus = query(`.user.uid-${msg.id} .status`);
    if (!userStatus) {
      query(".users").add(UserCard(msg));
      console.log("should append new user", msg);
      return;
    }
    userStatus.className = `status ${msg.value}`;
    userStatus.textContent = msg.value;
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
