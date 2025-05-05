import { go } from "../router.js";
import div from "./native/div.js";
import img from "./native/img.js";

export const UserCard = (user) => {
  const userDiv = div(`user uid-${user.id}`);
  userDiv.onclick = () => {
    go(`/chat/${user.id}`);
  };
  return userDiv.add(
    div("publisher").add(
      img(user.profilePicture, "no-profile"),
      div("username", user.username),
      div(`status ${user.status}`, user.status)
    )
  );
};
