import { timePassed } from "../utils/time.js";
import div from "./native/div.js";
import img from "./native/img.js";

const CommentContainer = (comment) => {
  return div("container").add(
    div("comment").add(
      div("publisher").add(img(comment.publisher.profilePicture, "no-profile")),
      div("content").add(
        div("").add(
          div("commentCreator", comment.publisher.username),
          div("creationTime", ` • ${timePassed(comment.creationTime)}`)
        ),
        div("commentText", comment.content)
      )
    ),
    div("reactions").add(
      div("reaction likes").add(
        img("like", "like", "like", "like"),
        div("nOfLikes", comment.likes)
      ),
      div("reaction dislike").add(
        img("dislike", "dislike", "dislike", "dislike"),
        div("nOfLikes", comment.likes)
      )
    )
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
