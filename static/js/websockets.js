import div from "./components/native/div.js";

const WS_API = "/api/ws";

export const InitWS = () => {
  const ws = new WebSocket(WS_API);
  ws.onerror = (e) => console.log("ws error", e);
  ws.onopen = (e) => console.log("ws open");
  ws.onclose = (e) => console.log("ws close");
  ws.onmessage = (e) => {
    const msg = JSON.parse(e.data);
    if (msg.type === "status") {
      const userStatus = document.querySelector(`.user.uid-${msg.id} .status`);
      userStatus?.replaceWith(div(`status ${msg.status}`, msg.status));
    }
  };
};
