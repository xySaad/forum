import users from "../context/users.js";
import div from "./native/div.js";
import img from "./native/img.js";
const USERS_API = "/api/users";

const getActiveUsers = async (parentNode) => {  
  if (users.size === 0) {
    const resp = await fetch(USERS_API);
    if (resp.ok) {
      const json = await resp.json();
      json.forEach(users.add);
    }
  }

  if (users.size < 2) {
    parentNode.append(div("fallback", "It's lonely right here!\nno users."));
    return;
  }

  const ownUserId = document.querySelector(".profile").id;
  users.list.forEach((user) => {    
    if (user.id === ownUserId) return;
    parentNode.add(
      div(`user uid-${user.id}`).add(
        div("publisher").add(
          img(user.profilePicture, "no-profile"),
          div("username", user.username),
          div(`status ${user.status}`, user.status)
        )
      )
    );
  });
};

export const ActiveUsers = () => {
  const usersContainer = div("users").add(div("title", "Users"));
  getActiveUsers(usersContainer);
  return usersContainer;
};
