import div from "./native/div.js";
import { Input } from "./Input.js";
import { CommentContainer, CommentsList } from "./CommentsList.js";
import { Post } from "./Post.js";
import { back, GetParams } from "../router.js";

const PostView = (postData) => {
  const postView = div("postView");

  postView.onclick = (e) => {
    if (e.target == postView) {
      back();
    }
  };

  if (!postData) {
    const { id } = GetParams();
    fetch(`/api/posts/${id}`).then(async (res) => {
      const postData = await res.json();
      postView.append(
        div("postCard").add(
          Post(postData),
          div("commentsWrap").add(
            CommentsList(postData.id),
            Input(`/api/posts/${postData.id}/comments/`, CommentContainer)
          )
        )
      );
    });

    return postView;
  }

  return postView.add(
    div("postCard").add(
      Post(postData),
      div("commentsWrap").add(
        CommentsList(postData.id),
        Input(`/api/posts/${postData.id}/comments/`, CommentContainer)
      )
    )
  );
};

export default PostView;
