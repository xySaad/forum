import div from "./div.js";
import img from "./img.js";
import { timePassed } from "../utils/time.js";
import Frame from "./Frame.js";
import { importSvg } from "../utils/index.js";
import PostView from "./PostView.js";

const Post = (postData) => {
  const readMore = div("readmore", "Read more");

  readMore.onclick = () => {
    document.body.prepend(PostView(postData));
  };

  return div("postContainer").add(
    Frame(
      div("post").add(
        div("publisher").add(
          img(postData.publisher.profilePicture, "no-profile"),
          div(null, postData.publisher.name),
          div(null, timePassed(postData.creationTime))
        ),
        div("title", postData.title),
        div("text", postData.text),
        readMore
      )
    ),
    div("leftBar").add(
      img(importSvg("arrow-up"), "arrow-up", "reaction-arrow", postData.ID),
      img(importSvg("comment-bubble"), "comment-bubble"),
      img(importSvg("arrow-down"), "arrow-down", "reaction-arrow", postData.ID)
    )
  );
};
export default Post;
