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
const MESSAGE_TYPE_DM = "DM";
const CONVERSATION_API = "http://localhost:8080/api/chat/";
let observer;

export const Message = (msg) => {
  console.log(msg);
  
  console.log("jat");
  
  const publisher = users.get(msg.sender);
  const time = new Date(msg.creationTime);
  const minutes = time.getMinutes().toString().padStart(2, "0");
  const hours = time.getHours().toString().padStart(2, "0")
  
  const formatedDate = `${hours}:${minutes}`;
  return div("message").add(
    div("publisher").add(
      img(publisher.profilePicture, "no-profile"),
      div("username", publisher.username),
      div("time", ` • ${formatedDate}`)
    ),
    div("text", msg.value)
  );
};

const fetchNext = async (parentNode, url) => {
  try {
    const resp = await Fetch(url);
    const json = await resp.json();
    if (!json) {
      if (parentNode.children.length < 0) {
        parentNode.add(div("fallback", "it's empty here!"));
      }
      return;
    }
    json.forEach((msg) => {
      parentNode.append(Message(msg));
    });
    const topMessage = parentNode.lastChild;
    topMessage.id = json[json.length - 1].id;
    observer.observe(topMessage);
  } catch (error) {
    console.error(error);
    parentNode.append(div("fallback", "error loading messages"));
  }
};

const observerArgs = (parentNode, url) => {
  const callBack = (e) => {
    if (e[0].isIntersecting) {
      observer.unobserve(e[0].target);
      const topMessage = parentNode.lastChild;
      console.log(topMessage.id, topMessage.textContent);

      url.searchParams.set("lastId", topMessage.id);
      fetchNext(parentNode, url);
    }
  };
  const options = {
    root: parentNode,
    rootMargin: "0px",
    threshold: 0.1,
  };
  return [callBack, options];
};

export const Chat = () => {
  const { id } = GetParams();
  const url = new URL(CONVERSATION_API + id);
  const chatBubble = query(".chat-bubble");
  toggleIcon(".chat-bubble");
  chatBubble?.on("load", (svg) => svg.classList.add("active"));

  const user = users.get(id);
  const messages = div("messages");
  observer = new IntersectionObserver(...observerArgs(messages, url));
  fetchNext(messages, url);

  const sendMessage = ({ value }) => {
    const msg = {
      type: MESSAGE_TYPE_DM,
      chat: id,
      value,
    };
    ws.send(JSON.stringify(msg));
  };

  return div("chat").add(UserCard(user), messages, Input(sendMessage));
};
