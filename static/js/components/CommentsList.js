import div from "./div.js";

const getComments = async (postId, commentsList) => {
  const resp = await fetch(`/api/coments?p_id=${postId}&offset=10`);
};

export const CommentsList = (postId) => {
  const commentsList = div("commentsList");
  getComments(postId, commentsList);
  return commentsList;
};
