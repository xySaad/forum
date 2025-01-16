import div from "./div.js";
import img from "./img.js";
import { timePassed } from "../utils/time.js";
import Frame from "./Frame.js";

const importSvg = (svgName) => "./static/svg/" + svgName + ".svg";

const Post = (postData) =>
  div("postContainer").add(
    Frame(
      div("post").add(
        div("publisher").add(
          img(importSvg("no-profile"), "no-profile"),
          div(null, postData.publisher.name),
          div(null, timePassed(postData.creationTime))
        ),
        div("title", postData.title),
        div("text", postData.text),
        div("readmore", "Read more")
      )
    ),
    div("leftBar").add(
      img(importSvg("arrow-up"), "arrow-up"),
      img(importSvg("comment-bubble"), "comment-bubble"),
      img(importSvg("arrow-down"), "arrow-down")
    )
  );
export default Post;
