import { timePassed } from "../utils/time.js";
import div from "./native/div.js";
import img from "./native/img.js";
let offset=0

const getComments = async (postId, commentsList) => {
  const resp = await fetch(`/api/coments?p_id=${postId}&offset=${offset}`);
  const json = await resp.json()
  offset=offset+json.length
  
json.forEach(comment => {
  console.log(comment);
  
  commentsList.add(
    div("comment").add(
    div("publisher").add(
      
      img(comment.publisher.profilePicture, "no-profile"),
      div(null, comment.publisher.username),
      div(null, timePassed(comment.creationTime))

    ),
    div("text", comment.content),
    div("reactions").add(
      div("likes",comment.likes),
      div("dislike",comment.dislikes)
          ),
  )  )

});
};

export const CommentsList = (postId) => {
  const commentsList = div("commentsList");
  getComments(postId, commentsList);
  return commentsList;
};
