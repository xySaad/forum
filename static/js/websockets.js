const WS_API = "/api/ws";

export const InitWS = () => {
  const ws = new WebSocket(WS_API);
  ws.onerror = (e) => console.log(e);
  ws.onopen = (e) => console.log(e);
  ws.onclose = (e) => console.log(e);
  ws.onmessage = (e) => console.log(e);
};
