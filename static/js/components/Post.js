import div from "./native/div.js";
import img from "./native/img.js";
import { timePassed } from "../utils/time.js";
import Frame from "./Frame.js";
import { importSvg } from "../utils/index.js";
import { go } from "../router.js";
import { reaction } from "./reaction.js";

export const Post = (postData) => {
  const cts = div("categoriesInPost");
  postData.content.categories.forEach((cat) => {
    if (cat != "") {
      cts.append(div("cat", "#" + cat));
    }
  });
  const post = div("post");

  return post.add(
    div("publisher").add(
      img(postData.publisher.profilePicture, "no-profile"),
      div("username", postData.publisher.username),
      div("time", timePassed(postData.creationTime))
    ),
    cts,
    div("title", postData.content.title),
    div("text", postData.content.text)
  );
};

export const PostCard = (postData) => {
  const showPost = () => {
    go(`/post/${postData.id}`, postData);
  };

  const readMore = div("readmore", "Read more");
  readMore.onclick = showPost;
  const comment = img(importSvg("comment-bubble"));
  comment.onclick = showPost;
  const reactionEndpoint = `/api/reactions/posts/${postData.id}/`;
  const [like, likeOnClick] = reaction("like", postData);
  const [dislike, dislikeOnClick] = reaction("dislike", postData);
  likeOnClick(dislike, reactionEndpoint);
  dislikeOnClick(like, reactionEndpoint);

  return div("postContainer").add(
    Frame(Post(postData).add(readMore)),
    div("leftBar").add(
      div("reactionsContainer").add(like, dislike),
      div("comntBtn").add(comment)
    )
  );
};
export default Post;
