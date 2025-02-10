import { NewReference } from "../utils/reference.js";
import { timePassed } from "../utils/time.js";
import div from "./native/div.js";
import img from "./native/img.js";

import HandleReactions from "../handlReactions.js";

export const atBottom = (el) => {
  const sh = el.scrollHeight,
    st = el.scrollTop,
    ht = el.offsetHeight;
  return ht == 0 || st == (sh - ht);
}

const getComments = async (postId, commentsList, isfetch, offset, lastPostId) => {
  isfetch(true)
  try {
    const resp = await fetch(`/api/coments?p_id=${postId}&offset=${offset()}&from=${lastPostId()}`);
    if (!resp.ok) {
      throw new Error('Network response was not ok');
    }
    const json = await resp.json();
    console.log(lastPostId()," ",json[0].id)
    if (offset() === 0) {
      lastPostId(json[0].id)
    }
    offset((prev) => prev + json.length)
    // offset += json.length; // Update offset
   
    json.forEach(comment => {
      let reactionsContainer = document.createElement("div")
      reactionsContainer.className = "reactionsContainer"
      let container1 = document.createElement("div")
      container1.className = "reaction like reacted incomment"
      container1.innerHTML = `<svg class="like " width="16" height="14" viewBox="0 0 16 14" fill="none" xmlns="http://www.w3.org/2000/svg">
                  <path d="M4.61819 5.96023H3.38328C3.38458 5.89687 3.36124 5.83549 3.31819 5.78899C3.27514 5.74249 3.21572 5.71451 3.15246 5.71094H0.911162V11.5854C0.910249 11.6948 0.930899 11.8034 0.971933 11.9048C1.01297 12.0063 1.07358 12.0986 1.15032 12.1766C1.22705 12.2547 1.3184 12.3168 1.41915 12.3595C1.5199 12.4022 1.62808 12.4247 1.73751 12.4256H3.15246C3.21552 12.4215 3.2746 12.3933 3.31754 12.3469C3.36048 12.3006 3.38401 12.2395 3.38328 12.1763H4.60895C4.65412 12.176 4.69878 12.1668 4.74039 12.1492C4.782 12.1317 4.81975 12.1061 4.85147 12.0739C4.88319 12.0418 4.90827 12.0037 4.92528 11.9619C4.94228 11.92 4.95087 11.8752 4.95057 11.8301V6.30415C4.95062 6.2149 4.91597 6.12913 4.85395 6.06495C4.79193 6.00077 4.70738 5.96322 4.61819 5.96023ZM2.43229 11.4723C2.43259 11.5096 2.42555 11.5466 2.41156 11.5811C2.39757 11.6157 2.37692 11.6471 2.35077 11.6737C2.32462 11.7003 2.29349 11.7215 2.25916 11.736C2.22483 11.7506 2.18797 11.7582 2.15068 11.7585C2.07558 11.7573 2.00395 11.7266 1.95127 11.6731C1.89859 11.6195 1.86907 11.5474 1.86908 11.4723V10.8491C1.87144 10.7756 1.90201 10.7058 1.95444 10.6542C2.00687 10.6026 2.07715 10.5732 2.15068 10.5721C2.2244 10.5727 2.295 10.6019 2.34755 10.6536C2.4001 10.7053 2.43049 10.7754 2.43229 10.8491V11.4723Z" fill="white"/>
                  <path d="M11.7942 2.12819C11.7942 3.23614 11.1248 3.85937 10.9355 4.60031H13.5554C13.9586 4.60445 14.3443 4.76569 14.6305 5.04974C14.9167 5.33379 15.0809 5.71826 15.0881 6.12144C15.0876 6.59711 14.9099 7.05554 14.5895 7.40712C14.7251 7.74562 14.775 8.11237 14.7347 8.47479C14.6944 8.8372 14.5653 9.18406 14.3587 9.48453C14.4565 9.82019 14.4692 10.175 14.3954 10.5167C14.3216 10.8585 14.1638 11.1765 13.9362 11.4419C13.9955 11.6377 14.0114 11.844 13.9827 12.0466C13.954 12.2491 13.8814 12.4429 13.7701 12.6145C13.2415 13.3878 11.9373 13.3878 10.8363 13.3878H10.7693C9.52752 13.3878 8.50959 12.9261 7.69247 12.5568C7.28161 12.3721 6.74379 12.1413 6.33523 12.1344C6.2538 12.132 6.17655 12.0978 6.12003 12.0391C6.06351 11.9804 6.03221 11.9019 6.03285 11.8205V6.23223C6.03158 6.18948 6.03912 6.14692 6.055 6.10721C6.07089 6.0675 6.09478 6.03149 6.12518 6.00141C7.14311 4.97655 7.58168 3.89399 8.41957 3.04225C8.7935 2.66139 8.92276 2.07741 9.06587 1.50728C9.17898 1.02486 9.41903 0 9.93839 0C10.557 0 11.7942 0.205433 11.7942 2.12819Z" fill="white"/>
                </svg>
                
                `            +comment.likes;
      let container2 = document.createElement("div")
      container2.className = "reaction dislike incomment"
      container2.innerHTML =   `
      <svg class="" width="16" height="14" viewBox="0 0 16 14" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M4.61819 7.42845H3.38328C3.38458 7.4918 3.36124 7.55319 3.31819 7.59968C3.27514 7.64618 3.21572 7.67416 3.15246 7.67773H0.911162V1.80328C0.910249 1.69385 0.930899 1.58532 0.971933 1.48387C1.01297 1.38242 1.07358 1.29005 1.15032 1.21203C1.22705 1.13401 1.3184 1.07187 1.41915 1.02915C1.5199 0.986439 1.62808 0.963988 1.73751 0.963082H3.15246C3.21552 0.967213 3.2746 0.995362 3.31754 1.04173C3.36048 1.08811 3.38401 1.14918 3.38328 1.21237H4.60895C4.65412 1.21267 4.69878 1.22187 4.74039 1.23943C4.782 1.257 4.81975 1.28258 4.85147 1.31474C4.88319 1.34689 4.90827 1.38497 4.92528 1.42682C4.94228 1.46866 4.95087 1.51344 4.95057 1.55861V7.08452C4.95062 7.17377 4.91597 7.25954 4.85395 7.32372C4.79193 7.3879 4.70739 7.42545 4.61819 7.42845ZM2.43229 1.91638C2.43259 1.8791 2.42555 1.84212 2.41156 1.80756C2.39757 1.77299 2.37692 1.74153 2.35077 1.71495C2.32462 1.68837 2.29349 1.6672 2.25916 1.65265C2.22483 1.63811 2.18797 1.63046 2.15068 1.63016C2.07558 1.63137 2.00395 1.66206 1.95127 1.71561C1.89859 1.76916 1.86907 1.84127 1.86908 1.91638V2.53961C1.87144 2.61311 1.90201 2.68289 1.95444 2.73446C2.00687 2.78604 2.07715 2.81545 2.15068 2.81659C2.2244 2.81602 2.295 2.78679 2.34755 2.7351C2.4001 2.68341 2.43049 2.6133 2.43229 2.53961V1.91638Z" fill="white"/>
        <path d="M11.7947 11.2605C11.7947 10.1525 11.1253 9.52931 10.936 8.78836H13.5559C13.9591 8.78422 14.3448 8.62299 14.631 8.33893C14.9172 8.05488 15.0814 7.67041 15.0885 7.26724C15.0881 6.79156 14.9104 6.33314 14.59 5.98155C14.7256 5.64306 14.7755 5.2763 14.7352 4.91389C14.6949 4.55147 14.5657 4.20461 14.3591 3.90414C14.457 3.56848 14.4696 3.21371 14.3959 2.87194C14.3221 2.53017 14.1643 2.21218 13.9367 1.94676C13.996 1.75099 14.0119 1.54464 13.9831 1.34211C13.9544 1.13958 13.8819 0.945763 13.7705 0.774175C13.242 0.000916481 11.9378 0.000916481 10.8368 0.000916481H10.7698C9.52801 0.000916481 8.51008 0.462563 7.69296 0.831881C7.2821 1.01654 6.74428 1.24736 6.33572 1.25429C6.25429 1.2567 6.17703 1.29091 6.12051 1.34958C6.06399 1.40826 6.0327 1.48674 6.03334 1.56821V7.15644C6.03207 7.19919 6.03961 7.24175 6.05549 7.28146C6.07138 7.32117 6.09526 7.35719 6.12567 7.38726C7.1436 8.41212 7.58217 9.49468 8.42006 10.3464C8.79399 10.7273 8.92325 11.3113 9.06636 11.8814C9.17946 12.3638 9.41952 13.3887 9.93887 13.3887C10.5575 13.3887 11.7947 13.1832 11.7947 11.2605Z" fill="white"/>
      </svg>
    
    `+ comment.dislikes
    container1.onclick = (e) => HandleReactions(e , comment);
    container2.onclick = (e) => HandleReactions(e,comment);
    reactionsContainer.append(container1,container2)
      commentsList.add(
        div("container").add(
          div("comment").add(
            div("publisher").add(
              img(comment.publisher.profilePicture, "no-profile"),
            ),
            div("content2").add(
              div("texts2").add(
                div("commentCreator", comment.publisher.username),
                div("creationTime", ` • ${timePassed(comment.creationTime)}`)
              ),
              div("commentText", comment.content),
            )
          ),
          reactionsContainer,
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
  let lastPostId = NewReference(0);

  getComments(postId, commentsList, isfetch, offset, lastPostId)

  commentsList.onscroll = () => {
    if (!atBottom(commentsList) || isfetch()) {
      return
    }
    getComments(postId, commentsList, isfetch, offset, lastPostId)
  };

  return commentsList;
};