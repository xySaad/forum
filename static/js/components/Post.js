import div from "./native/div.js";
import img from "./native/img.js";
import { timePassed } from "../utils/time.js";
import Frame from "./Frame.js";
import { importSvg } from "../utils/index.js";
import { go } from "../router.js";
import { Fetch } from "../utils/fetch.js";
import { svg } from "./native/svg.js";

export const Post = (postData) => {
  const cts = div("categoriesInPost");
  postData.categories?.forEach((cat) => {
    cts.append(div("cat", "#" + cat));
  });

  return div("post").add(
    div("publisher").add(
      img(postData.publisher.profilePicture, "no-profile"),
      div("username", postData.publisher.username),
      div("time", timePassed(postData.creationTime))
    ),
    cts,
    div("title", postData.content.title),
    div("text", postData.content.text),
  );
};

export const PostCard = (postData) => {
  const showPost = () => {
    go(`/post/${postData.id}`, true, postData);
  };

  const readMore = div("readmore", "Read more");
  readMore.onclick = showPost;
  const comment = img(importSvg("comment-bubble"));

  const like = div("reaction like reacted").add(
    svg("like"),
    postData.likes
  );

  like.onclick = () => {
    Fetch(`/api/reactions/posts/${postData.id}/like`, {
      method: "post",
    });
    like.classList.toggle("reacted")
    dislike.classList.remove("reacted")
  };

  const dislike = div("reaction dislike").add(
    svg("dislike"),
    postData.likes
  );

  dislike.onclick = async () => {
    Fetch(`/api/reactions/posts/${postData.id}/dislike`, { method: "post" });
    like.classList.remove("reacted")
    dislike.classList.toggle("reacted")
  };

  return div("postContainer").add(
    Frame(
      Post(postData).add(readMore)
    ),
    div("leftBar").add(
      div("reactionsContainer").add(like, dislike),
      div("comntBtn").add(comment)
    )
  );
};
export default Post;
