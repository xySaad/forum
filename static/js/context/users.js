const usersMap = new Map();
let usersList = null;

const users = {
  myself: null,
  add: (user) => usersMap.set(user.id, user),
  get: (id) => usersMap.get(id),
  get list() {
    if (!usersList) usersList = [...usersMap.values()];

    return usersList.sort(
      (a, b) => BigInt(a.lastMessage?.id || 0) < BigInt(b.lastMessage?.id || 0)
    );
  },
  get size() {
    return usersMap.size;
  },
};

export default users;
