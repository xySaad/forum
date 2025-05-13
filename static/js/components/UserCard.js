import { GetParams, go } from "../router.js";
import { sendTypingStatus } from "./Input.js";
import div from "./native/div.js";
import img from "./native/img.js";

export const UserCard = (user, clickable = true) => {
  const userDiv = div(`user uid-${user.id}`);
  if (clickable) {
    userDiv.onclick = () => {
      const { id } = GetParams();
      sendTypingStatus(false, id);
      go(`/chat/${user.id}`);
    };
  }
  return userDiv.add(
    div("publisher").add(
      img(user.profilePicture, "no-profile"),
      div("username", user.username),
      div(`status ${user.status}`, user.status)
    )
  );
};
