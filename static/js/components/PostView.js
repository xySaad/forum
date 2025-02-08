import div from "./native/div.js";
import { CommentInput } from "./CommentInput.js";
import { CommentsList } from "./CommentsList.js";
import { Post } from "./Post.js";
import { back } from "../router.js";

const PostView = (postData) => {
  const postView = div("postView");
  postView.onclick = (e) => {
    if (e.target == postView) {
      back();
    }
  };

  const commentsWrap = div("commentsWrap");

  postView.id = postData.id;

  return postView.add(
    div("postCard").add(
      Post(postData),
      commentsWrap.add(CommentsList(postData.id), CommentInput(postData.id))
    )
  );
};

export default PostView;
