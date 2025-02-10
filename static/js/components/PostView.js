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
  if (!postData.categories) {
    postData.categories = ["sport" , "art"]
  }
  let cts = div("categoriesInPost")
  postData.categories.forEach(cat => {
    cts.append(div("cat","#"+ cat))
  })
  return postView.add(
    div("postCard").add(
      Post(postData),
      commentsWrap.add(CommentsList(postData.id), CommentInput(postData.id))

    )
  );
};

export default PostView;