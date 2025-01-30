import { timePassed } from "../utils/time.js";
import div from "./native/div.js";
import img from "./native/img.js";
function atBottom(el) {
  let sh = el.scrollHeight,
    st = el.scrollTop,
    ht = el.offsetHeight;
  return ht == 0 || st == sh - ht;
}
const getComments = async (postId, commentsList) => {
  let offset = 0
  let isfetch = false
  console.log(atBottom(commentsList));
  if (atBottom(commentsList) && !isfetch) {
    isfetch = true
    const resp = await fetch(`/api/coments?p_id=${postId}&offset=${offset}`);
    const json = await resp.json()
    offset = offset + json.length
    json.forEach(comment => {
      console.log(comment);

      commentsList.add(
        div("container").add(
          div("comment").add(
            div("publisher").add(
              img(comment.publisher.profilePicture, "no-profile"),
            ),
            div("content").add(
              div("texts").add(
                div("commentCreator", comment.publisher.username),
                div("creationTime", ` • ${timePassed(comment.creationTime)}`)
              ),
              div("commentText", comment.content),
            )
          ),
          div("reactions").add(
            div("reaction likes",).add(
              img("like", "like", "like", "like"),
              div("nOfLikes", comment.likes)
            ),
            div("reaction dislike",).add(
              img("dislike", "dislike", "dislike", "dislike"),
              div("nOfLikes", comment.likes)
            )
          ),
        ))

    });
    isfetch=false
  }
};

export const CommentsList = (postId) => {
  const commentsList = div("commentsList");
  commentsList.addEventListener("scroll",getComments(postId, commentsList));
  return commentsList;
};
