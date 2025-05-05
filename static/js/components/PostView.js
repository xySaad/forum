import div from "./native/div.js";
import { Input } from "./Input.js";
import { CommentContainer, CommentsList } from "./CommentsList.js";
import { Post } from "./Post.js";
import { back, GetParams } from "../router.js";
import { Fetch } from "../utils/fetch.js";

const PostView = async (postData) => {
  let pData = postData;
  const postView = div("postView");
  postView.onclick = (e) => {
    if (e.target === postView) back();
  };

  if (!pData) {
    const { id } = GetParams();
    const resp = await fetch(`/api/posts/${id}`);
    pData = await resp.json();
  }

  const sendComment = async (input) => {
    const resp = await Fetch(`/api/posts/${pData.id}/comments/`, {
      method: "POST",
      body: JSON.stringify({ content: input.value }),
    });
    const json = await resp.json();
    const output = input.parentNode.previousSibling;
    output.prepend(CommentContainer(json));
  };

  return postView.add(
    div("postCard").add(
      Post(pData),
      div("commentsWrap").add(
        CommentsList(pData.id),
        Input(sendComment)
      )
    )
  );
};

export default PostView;
