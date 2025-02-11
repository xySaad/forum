import { timePassed } from "../utils/time.js";
import div from "./native/div.js";
import img from "./native/img.js";
import { reaction } from "./reaction.js";
const CommentContainer = (comment, postId) => {
  const reactionEndpoint = `/api/reactions/posts/${postId}/comments/`;
  const [like, onLike] = reaction("like", reactionEndpoint);
  const [dislike, onDislike] = reaction("dislike", reactionEndpoint);

  onLike(dislike);
  onDislike(like);

  return div("container").add(
    div("comment").add(
      div("publisher").add(img(comment.publisher.profilePicture, "no-profile")),
      div("commentCreator", comment.publisher.username),
      div("creationTime", ` • ${timePassed(comment.creationTime)}`)
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
