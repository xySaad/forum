import { importSvg } from "../utils/index.js";
import { timePassed } from "../utils/time.js";
import div from "./native/div.js";
import img from "./native/img.js";
const CommentContainer = (comment) => {
  const reactionsContainer = div("reactionsContainer").add(
    div("reaction like reacted incomment").add(
      img(importSvg("like")),
      comment.likes
    )
  );

  return div("container").add(
    div("comment").add(
      div("publisher").add(img(comment.publisher.profilePicture, "no-profile")),
      div("content2").add(
        div("texts2").add(
          div("commentCreator", comment.publisher.username),
          div("creationTime", ` • ${timePassed(comment.creationTime)}`)
        ),
        div("commentText", comment.content)
      )
    ),
    reactionsContainer
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
