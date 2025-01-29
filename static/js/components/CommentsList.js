import div from "./native/div.js";
let offset=0

const getComments = async (postId, commentsList) => {
  const resp = await fetch(`/api/coments?p_id=${postId}&offset=${offset}`);
  const json = await resp.json()
  offset=offset+json.length
  
json.forEach(comment => {
  commentsList.append(comment.content)
});
};

export const CommentsList = (postId) => {
  const commentsList = div("commentsList");
  getComments(postId, commentsList);
  return commentsList;
};
