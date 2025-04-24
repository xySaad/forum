import div from "./native/div.js";
import img from "./native/img.js";
const USERS_API = "/api/users";

const getActiveUsers = async (parentNode) => {
  const resp = await fetch(USERS_API);
  const users = await resp.json();
  console.log(users);

  if (!users.length || users.length < 2) {
    parentNode.append(div("fallback", "It's lonely right here!\nno users."));
    return;
  }
  const ownUserId = document.querySelector(".profile").id;
  users.forEach((user) => {
    if (user.id === ownUserId) return;
    parentNode.add(
      div("user").add(
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
  const usersContainer = div("users");
  getActiveUsers(usersContainer);
  return usersContainer.add(div("title", "Users"));
};
