import users from "../context/users.js";
import { Fetch } from "../utils/fetch.js";
import div from "./native/div.js";
import { svg } from "./native/svg.js";
import { UserCard } from "./UserCard.js";
const USERS_API = "/api/users";

export const ActiveUsers = async () => {
  const chevronLeft = svg("chevron-left");
  const icon = svg("groups");

  const usersContainer = div("users").add(
    div("head").add(icon, div("title", "Users"), chevronLeft)
  );
  const hideUsers = () => usersContainer.classList.toggle("hide");
  chevronLeft.onclick = hideUsers;
  icon.onclick = hideUsers;

  if (users.size === 0) {
    const resp = await Fetch(USERS_API);
    if (resp.ok) {
      const json = await resp.json();
      json.forEach((user) => users.add(user));
    }
  }

  if (users.size < 2) {
    usersContainer.append(
      div("fallback", "It's lonely right here!\nno users.")
    );
    return usersContainer;
  }

  users.list.forEach((user) => {
    if (user.id === users.myself.id) return;
    usersContainer.add(UserCard(user));
  });
  return usersContainer;
};
