import { NewReference } from "../utils/reference.js";
import { timePassed } from "../utils/time.js";
import div from "./native/div.js";
import img from "./native/img.js";

const atBottom = (el) => {
  const sh = el.scrollHeight,
    st = el.scrollTop,
    ht = el.offsetHeight;
  return ht == 0 || st == sh - ht;
}

const getComments = async (postId, commentsList, isfetch, offset) => {
  isfetch(true)
  try {
    const resp = await fetch(`/api/coments?p_id=${postId}&offset=${offset()}`);
    if (!resp.ok) {
      throw new Error('Network response was not ok');
    }
    const json = await resp.json();
    console.log(offset());

    offset((prev)=>prev+json.length)
    // offset += json.length; // Update offset

    json.forEach(comment => {
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
            div("reaction likes").add(
              img("like", "like", "like", "like"),
              div("nOfLikes", comment.likes)
            ),
            div("reaction dislike").add(
              img("dislike", "dislike", "dislike", "dislike"),
              div("nOfLikes", comment.likes)
            )
          ),
        )
      );
    });
  } catch (error) {
    console.error('Error fetching comments:', error);
  } finally {
      isfetch(false); 
  }
};

export const CommentsList = (postId) => {
  const commentsList = div("commentsList");
  let offset = NewReference(0); // Move offset outside the function
  let isfetch = NewReference(false);
  getComments(postId, commentsList, isfetch, offset)

  commentsList.onscroll = () => {
    if (!atBottom(commentsList) || isfetch()) {
      return
    }
    getComments(postId, commentsList, isfetch, offset)
  };

  return commentsList;
};