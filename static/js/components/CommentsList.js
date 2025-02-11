import { timePassed } from "../utils/time.js";
import div from "./native/div.js";
import img from "./native/img.js";
import { reaction } from "./reaction.js";
export const CommentContainer = (comment) => {
  const reactionEndpoint = `/api/reactions/comments/${comment.id}/`;
  const [like, onLike] = reaction("like", comment);
  const [dislike, onDislike] = reaction("dislike", comment);

  onLike(dislike, reactionEndpoint);
  onDislike(like, reactionEndpoint);

  return div("comment").add(
    div("publisher").add(
      img(comment.publisher.profilePicture, "no-profile"),
      div("username", comment.publisher.username),
      div("time", ` • ${timePassed(comment.creationTime)}`)
    ),
    div("text", comment.content),
    div("reactionsContainer").add(like, dislike)
  );
};

const fetchComments = async (commentsList, postId) => {
  const resp = await fetch(`/api/posts/${postId}/comments`);
  const json = await resp.json();
  json.forEach((comment) => {
    commentsList.add(CommentContainer(comment));
  });
};

export const CommentsList = (postId) => {
  const commentsList = div("commentsList");
  fetchComments(commentsList, postId);
  return commentsList;
};
