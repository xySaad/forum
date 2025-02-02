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
  const postImg = img(postData.image);


  postView.id = postData.ID;
  return postView.add(
    div("postCard").add(
      Post(postData).add(postImg),
      commentsWrap.add(CommentsList(postData.ID), CommentInput(postData.ID))
    )
  );
};

export default PostView;
