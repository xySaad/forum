import div from "./native/div.js";
import img from "./native/img.js";
import { timePassed } from "../utils/time.js";
import Frame from "./Frame.js";
import { importSvg } from "../utils/index.js";
import { go } from "../router.js";

export const Post = (postData) => {
  const postImg = img(postData.image);

  return div(`post ${postData.id}`).add(
    div("publisher").add(
      img(postData.publisher.profilePicture, "no-profile"),
      div().add(
        div("username", postData.publisher.username),
        div("time", timePassed(postData.creationTime))
      )
    ),
    div("title", postData.content.title),
    div("text", postData.content.text),
    postImg
  );
};

export const PostCard = (postData) => {
  const showPost = () => {
    go(`/post/${postData.id}`, true, postData);
  };

  const readMore = div("readmore", "Read more");
  readMore.onclick = showPost;
  const comment = img(importSvg("comment-bubble"), "comment-bubble");
  comment.onclick = showPost;

  const like = img(importSvg("like"));
  const dislike = img(importSvg("dislike"));

  return div("postContainer").add(
    Frame(Post(postData).add(readMore)),
    div("leftBar").add(like, comment, dislike)
  );
};

export default PostCard;
