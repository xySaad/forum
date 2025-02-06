import div from "./native/div.js";
import img from "./native/img.js";
import { CommentInput } from "./CommentInput.js";
import { CommentsList } from "./CommentsList.js";
import { Post } from "./Post.js";

const PostView = (postData) => {
  const postView = div("postView");
  postView.onclick = (e) => {
    if (e.target == postView) {
      postView.remove();
    }
  };

  const commentsWrap = div("commentsWrap");

  postView.id = postData.id;

  return postView.add(
    div("postCard").add(
      Post(postData),
      commentsWrap.add(CommentsList(postData.ID), CommentInput(postData.ID))
    )
  );
};

export default PostView;
