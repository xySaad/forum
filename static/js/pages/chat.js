import { CommentContainer } from "../components/CommentsList.js";
import { toggleIcon } from "../components/Headers.js";
import { Input } from "../components/Input.js";
import div from "../components/native/div.js";
import img from "../components/native/img.js";
import { query } from "../components/native/index.js";
import { UserCard } from "../components/UserCard.js";
import users from "../context/users.js";
import { GetParams } from "../router.js";
import { Fetch } from "../utils/fetch.js";
import { ws } from "../websockets.js";
const CONVERSATION_API = "/api/chat/";
const MESSAGETYPE_DM = "DM";

const Message = (msg) => {
  const publisher = users.get(msg.sender);
  const creationTime = new Date(msg.creationTime);
  const formatedDate = `${creationTime}`;
  return div("message").add(
    div("publisher").add(
      img(publisher.profilePicture, "no-profile"),
      div("username", publisher.username),
      div("time", ` • ${formatedDate}`)
    ),
    div("text", msg.value)
  );
};
export const Chat = async () => {
  const chatBubble = query(".chat-bubble");
  toggleIcon(".chat-bubble");
  chatBubble?.on("load", (svg) => svg.classList.add("active"));

  const { id } = GetParams();  
  const user = users.get(id);
  const messages = div("messages");

  try {
    const resp = await Fetch(CONVERSATION_API + id);

    const json = await resp.json();
    if (!json || json.length < 1) {
      messages.add(div("fallback", "it's empty here!"));
    }
    json?.forEach((msg) => {
      messages.append(Message(msg));
    });
  } catch (error) {
    console.error(error);
    messages.add(div("fallback", "error loading messages"));
  }

  const sendMessage = ({value}) => {
    const msg = {
      type: MESSAGETYPE_DM,
      chat: id,
      value,
    };
    ws.send(JSON.stringify(msg));
  };

  return div("chat").add(UserCard(user), messages, Input(sendMessage));
};
