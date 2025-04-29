const usersMap = new Map();
const users = {
  myself: null,
  add: (user) => {    
    usersMap.set(user.id, user);
  },
  get list(){
    return usersMap.values()
  },
  get size(){
    return usersMap.size
  }
};

export default users;
