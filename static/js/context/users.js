const usersMap = new Map();
let usersList = null;

const users = {
  myself: null,
  add: (user) => usersMap.set(user.id, user),
  get: (id) => usersMap.get(id),
  get list() {
    if (!usersList) usersList = [...usersMap.values()];

    return usersList.sort((a, b) => {
      const ida = a.lastMessage?.id;
      const idb = b.lastMessage?.id;
      const idOrder = BigInt(ida || 0) < BigInt(idb || 0);
      if (ida || idb) return idOrder ? 1 : -1;

      const nameA = a.username;
      const nameB = b.username;
      const nameOrder = nameA.localeCompare(nameB);
      return nameOrder;
    });
  },
  get size() {
    return usersMap.size;
  },
  whoIsTyping: null
};

export default users;
