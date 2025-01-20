import { timePassed } from "../utils/time.js";
import div from "./div.js";
import img from "./img.js";
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

  postView.id = postData.ID;
  return postView.add(
    div("postCard").add(
      div("post").add(
        div("publisher").add(
          img(postData.publisher.profilePicture, "no-profile"),
          div(null, postData.publisher.name),
          div(null, timePassed(postData.creationTime))
        ),
        div("title", postData.title),
        div("text", postData.text),
        postImg
      ),
      commentsWrap.add(CommentsList(postData.ID), CommentInput())
    )
  );
};

export default PostView;
