import div from "./native/div.js";
import img from "./native/img.js";
import { timePassed } from "../utils/time.js";
import Frame from "./Frame.js";
import { importSvg } from "../utils/index.js";
import PostView from "./PostView.js";

const Post = (postData) => {
  const showPost = () => {
    document.body.prepend(PostView(postData));
  };

  const readMore = div("readmore", "Read more");
  readMore.onclick = showPost;
  const comment = img(importSvg("comment-bubble"), "comment-bubble");
  comment.onclick = showPost;

  return div("postContainer").add(
    Frame(
      div("post").add(
        div("publisher").add(
          img(postData.publisher.profilePicture, "no-profile"),
          div(null, postData.publisher.username),
          div(null, timePassed(postData.creationTime))
        ),
        div("title", postData.title),
        div("text", postData.text),
        readMore
      )
    ),
    div("leftBar").add(
      img(importSvg("arrow-up"), "arrow-up", "reaction-arrow", postData.ID),
      comment,
      img(importSvg("arrow-down"), "arrow-down", "reaction-arrow", postData.ID)
    )
  );
};
export default Post;
