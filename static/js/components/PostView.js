import { timePassed } from "../utils/time.js";
import div from "./native/div.js";
import img from "./native/img.js";
import { CommentInput } from "./CommentInput.js";
import { CommentsList } from "./CommentsList.js";
import { onResize } from "../utils/events.js";

const PostView = (postData) => {
  const postView = div("postView");
  postView.onclick = (e) => {
    if (e.target == postView) {
      postView.remove();
    }
  };

  const commentsWrap = div("commentsWrap");
  const postImg = img(postData.image);

  // const adjustCommentsListSize = () => {
  //   commentsWrap.style.height = getComputedStyle(
  //     postView.querySelector(".post")
  //   ).height;
  // };

  // postImg.onload = adjustCommentsListSize;
  // onResize(adjustCommentsListSize);

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
      div("post").add(
        div("publisher").add(
          img(postData.publisher.profilePicture, "no-profile"),
          div(null, postData.publisher.username),
          div(null, timePassed(postData.creationTime))
        ),
        cts , 
        div("title", postData.content.title),
        div("text", postData.content.text),
        postImg
      ),
      
      commentsWrap.add(CommentsList(postView.id), CommentInput(postData.id))
    )
  );
};

export default PostView;