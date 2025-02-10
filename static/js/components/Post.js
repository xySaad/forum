import div from "./native/div.js";
import img from "./native/img.js";
import { timePassed } from "../utils/time.js";
import Frame from "./Frame.js";
import PostView from "./PostView.js";
import HandleReactions from "../handlReactions.js";
const Post = (postData) => {
  console.log(postData.content.categories);
  
  const showPost = () => {
    document.body.prepend(PostView(postData));
  };
  const readMore = div("readmore", "Read more");
  readMore.onclick = showPost;
  const comment = document.createElement("div");
  comment.className="comntBtn"
  comment.innerHTML = `<svg class="comment" width="17" height="16" viewBox="0 0 17 16" fill="none" xmlns="http://www.w3.org/2000/svg">
<path fill-rule="evenodd" clip-rule="evenodd" d="M10.626 8.27507C10.626 7.90687 10.9244 7.6084 11.2926 7.6084H11.2986C11.6668 7.6084 11.9653 7.90687 11.9653 8.27507C11.9653 8.64327 11.6668 8.94173 11.2986 8.94173H11.2926C10.9244 8.94173 10.626 8.64327 10.626 8.27507Z" fill="white"/>
<path fill-rule="evenodd" clip-rule="evenodd" d="M7.95312 8.27507C7.95312 7.90687 8.25159 7.6084 8.61979 7.6084H8.62579C8.99399 7.6084 9.29246 7.90687 9.29246 8.27507C9.29246 8.64327 8.99399 8.94173 8.62579 8.94173H8.61979C8.25159 8.94173 7.95312 8.64327 7.95312 8.27507Z" fill="white"/>
<path fill-rule="evenodd" clip-rule="evenodd" d="M5.28125 8.27507C5.28125 7.90687 5.57972 7.6084 5.94792 7.6084H5.95392C6.3221 7.6084 6.62058 7.90687 6.62058 8.27507C6.62058 8.64327 6.3221 8.94173 5.95392 8.94173H5.94792C5.57972 8.94173 5.28125 8.64327 5.28125 8.27507Z" fill="white"/>
<path fill-rule="evenodd" clip-rule="evenodd" d="M13.027 3.63789C10.6214 1.23138 6.71257 1.23138 4.307 3.63789C2.42307 5.52145 2.01385 8.31361 3.07122 10.5986C3.07477 10.6063 3.07813 10.614 3.08128 10.6219C3.14273 10.7745 3.14135 10.9475 3.13858 11.0454C3.13511 11.1683 3.12152 11.3098 3.10474 11.4538C3.08888 11.5899 3.06873 11.7399 3.04826 11.8924L3.04401 11.9241C3.02183 12.0893 2.99925 12.2589 2.97953 12.4288C2.93956 12.7733 2.91446 13.095 2.92563 13.3523C2.93119 13.4806 2.94521 13.5756 2.96311 13.6407C2.96915 13.6627 2.97467 13.6778 2.97874 13.6875C2.98848 13.6916 3.00359 13.6971 3.02553 13.7031C3.09069 13.721 3.18573 13.735 3.31397 13.7406C3.57133 13.7517 3.89302 13.7266 4.23748 13.6867C4.40732 13.667 4.57691 13.6444 4.74212 13.6223L4.77361 13.6181C4.92613 13.5976 5.07619 13.5775 5.21237 13.5616C5.35635 13.5449 5.49789 13.5313 5.62073 13.5279C5.71859 13.5251 5.89157 13.5237 6.04419 13.5852C6.05199 13.5883 6.0597 13.5917 6.06733 13.5952C8.35279 14.6521 11.1438 14.2435 13.027 12.3594C15.4329 9.95314 15.4377 6.04823 13.027 3.63789ZM13.7341 2.93083C10.938 0.133703 6.39585 0.133733 3.59975 2.93092C1.42373 5.10661 0.942894 8.32287 2.13912 10.9649C2.13948 10.9766 2.13965 10.9937 2.13898 11.0171C2.13678 11.0951 2.12735 11.2017 2.11146 11.3381C2.09666 11.4651 2.07763 11.6068 2.05677 11.7622L2.05289 11.7911C2.03077 11.9559 2.00707 12.1337 1.9862 12.3135C1.94502 12.6684 1.91179 13.0553 1.92657 13.3957C1.93397 13.5663 1.95399 13.7425 1.99887 13.9057C2.04273 14.0653 2.12033 14.2496 2.26856 14.3979C2.41683 14.546 2.60129 14.6237 2.7608 14.6675C2.92405 14.7123 3.10013 14.7323 3.27071 14.7397C3.61102 14.7544 3.99782 14.7212 4.35265 14.68C4.53245 14.6592 4.71021 14.6355 4.87503 14.6134L4.90364 14.6096C5.0591 14.5887 5.20083 14.5697 5.32793 14.5549C5.46436 14.5391 5.57091 14.5297 5.64883 14.5275C5.67231 14.5268 5.68943 14.5269 5.70111 14.5273C8.34339 15.723 11.5587 15.2429 13.7342 13.0664C16.53 10.2703 16.5359 5.73227 13.7341 2.93083Z" fill="white"/>
</svg>
`
  comment.onclick = showPost;
  let reactionsContainer = document.createElement("div")
  reactionsContainer.className = "reactionsContainer"
  let container1 = document.createElement("div")
  container1.className = "reaction like reacted"
  container1.innerHTML = `<svg class="like " width="16" height="14" viewBox="0 0 16 14" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M4.61819 5.96023H3.38328C3.38458 5.89687 3.36124 5.83549 3.31819 5.78899C3.27514 5.74249 3.21572 5.71451 3.15246 5.71094H0.911162V11.5854C0.910249 11.6948 0.930899 11.8034 0.971933 11.9048C1.01297 12.0063 1.07358 12.0986 1.15032 12.1766C1.22705 12.2547 1.3184 12.3168 1.41915 12.3595C1.5199 12.4022 1.62808 12.4247 1.73751 12.4256H3.15246C3.21552 12.4215 3.2746 12.3933 3.31754 12.3469C3.36048 12.3006 3.38401 12.2395 3.38328 12.1763H4.60895C4.65412 12.176 4.69878 12.1668 4.74039 12.1492C4.782 12.1317 4.81975 12.1061 4.85147 12.0739C4.88319 12.0418 4.90827 12.0037 4.92528 11.9619C4.94228 11.92 4.95087 11.8752 4.95057 11.8301V6.30415C4.95062 6.2149 4.91597 6.12913 4.85395 6.06495C4.79193 6.00077 4.70738 5.96322 4.61819 5.96023ZM2.43229 11.4723C2.43259 11.5096 2.42555 11.5466 2.41156 11.5811C2.39757 11.6157 2.37692 11.6471 2.35077 11.6737C2.32462 11.7003 2.29349 11.7215 2.25916 11.736C2.22483 11.7506 2.18797 11.7582 2.15068 11.7585C2.07558 11.7573 2.00395 11.7266 1.95127 11.6731C1.89859 11.6195 1.86907 11.5474 1.86908 11.4723V10.8491C1.87144 10.7756 1.90201 10.7058 1.95444 10.6542C2.00687 10.6026 2.07715 10.5732 2.15068 10.5721C2.2244 10.5727 2.295 10.6019 2.34755 10.6536C2.4001 10.7053 2.43049 10.7754 2.43229 10.8491V11.4723Z" fill="white"/>
              <path d="M11.7942 2.12819C11.7942 3.23614 11.1248 3.85937 10.9355 4.60031H13.5554C13.9586 4.60445 14.3443 4.76569 14.6305 5.04974C14.9167 5.33379 15.0809 5.71826 15.0881 6.12144C15.0876 6.59711 14.9099 7.05554 14.5895 7.40712C14.7251 7.74562 14.775 8.11237 14.7347 8.47479C14.6944 8.8372 14.5653 9.18406 14.3587 9.48453C14.4565 9.82019 14.4692 10.175 14.3954 10.5167C14.3216 10.8585 14.1638 11.1765 13.9362 11.4419C13.9955 11.6377 14.0114 11.844 13.9827 12.0466C13.954 12.2491 13.8814 12.4429 13.7701 12.6145C13.2415 13.3878 11.9373 13.3878 10.8363 13.3878H10.7693C9.52752 13.3878 8.50959 12.9261 7.69247 12.5568C7.28161 12.3721 6.74379 12.1413 6.33523 12.1344C6.2538 12.132 6.17655 12.0978 6.12003 12.0391C6.06351 11.9804 6.03221 11.9019 6.03285 11.8205V6.23223C6.03158 6.18948 6.03912 6.14692 6.055 6.10721C6.07089 6.0675 6.09478 6.03149 6.12518 6.00141C7.14311 4.97655 7.58168 3.89399 8.41957 3.04225C8.7935 2.66139 8.92276 2.07741 9.06587 1.50728C9.17898 1.02486 9.41903 0 9.93839 0C10.557 0 11.7942 0.205433 11.7942 2.12819Z" fill="white"/>
            </svg>
            
            `            +postData.likes;
  let container2 = document.createElement("div")
  container2.className = "reaction dislike"
  container2.innerHTML =   `
  <svg class="" width="16" height="14" viewBox="0 0 16 14" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M4.61819 7.42845H3.38328C3.38458 7.4918 3.36124 7.55319 3.31819 7.59968C3.27514 7.64618 3.21572 7.67416 3.15246 7.67773H0.911162V1.80328C0.910249 1.69385 0.930899 1.58532 0.971933 1.48387C1.01297 1.38242 1.07358 1.29005 1.15032 1.21203C1.22705 1.13401 1.3184 1.07187 1.41915 1.02915C1.5199 0.986439 1.62808 0.963988 1.73751 0.963082H3.15246C3.21552 0.967213 3.2746 0.995362 3.31754 1.04173C3.36048 1.08811 3.38401 1.14918 3.38328 1.21237H4.60895C4.65412 1.21267 4.69878 1.22187 4.74039 1.23943C4.782 1.257 4.81975 1.28258 4.85147 1.31474C4.88319 1.34689 4.90827 1.38497 4.92528 1.42682C4.94228 1.46866 4.95087 1.51344 4.95057 1.55861V7.08452C4.95062 7.17377 4.91597 7.25954 4.85395 7.32372C4.79193 7.3879 4.70739 7.42545 4.61819 7.42845ZM2.43229 1.91638C2.43259 1.8791 2.42555 1.84212 2.41156 1.80756C2.39757 1.77299 2.37692 1.74153 2.35077 1.71495C2.32462 1.68837 2.29349 1.6672 2.25916 1.65265C2.22483 1.63811 2.18797 1.63046 2.15068 1.63016C2.07558 1.63137 2.00395 1.66206 1.95127 1.71561C1.89859 1.76916 1.86907 1.84127 1.86908 1.91638V2.53961C1.87144 2.61311 1.90201 2.68289 1.95444 2.73446C2.00687 2.78604 2.07715 2.81545 2.15068 2.81659C2.2244 2.81602 2.295 2.78679 2.34755 2.7351C2.4001 2.68341 2.43049 2.6133 2.43229 2.53961V1.91638Z" fill="white"/>
    <path d="M11.7947 11.2605C11.7947 10.1525 11.1253 9.52931 10.936 8.78836H13.5559C13.9591 8.78422 14.3448 8.62299 14.631 8.33893C14.9172 8.05488 15.0814 7.67041 15.0885 7.26724C15.0881 6.79156 14.9104 6.33314 14.59 5.98155C14.7256 5.64306 14.7755 5.2763 14.7352 4.91389C14.6949 4.55147 14.5657 4.20461 14.3591 3.90414C14.457 3.56848 14.4696 3.21371 14.3959 2.87194C14.3221 2.53017 14.1643 2.21218 13.9367 1.94676C13.996 1.75099 14.0119 1.54464 13.9831 1.34211C13.9544 1.13958 13.8819 0.945763 13.7705 0.774175C13.242 0.000916481 11.9378 0.000916481 10.8368 0.000916481H10.7698C9.52801 0.000916481 8.51008 0.462563 7.69296 0.831881C7.2821 1.01654 6.74428 1.24736 6.33572 1.25429C6.25429 1.2567 6.17703 1.29091 6.12051 1.34958C6.06399 1.40826 6.0327 1.48674 6.03334 1.56821V7.15644C6.03207 7.19919 6.03961 7.24175 6.05549 7.28146C6.07138 7.32117 6.09526 7.35719 6.12567 7.38726C7.1436 8.41212 7.58217 9.49468 8.42006 10.3464C8.79399 10.7273 8.92325 11.3113 9.06636 11.8814C9.17946 12.3638 9.41952 13.3887 9.93887 13.3887C10.5575 13.3887 11.7947 13.1832 11.7947 11.2605Z" fill="white"/>
  </svg>

`+ postData.dislikes
container1.onclick = (e) => HandleReactions(e , postData);
container2.onclick = (e) => HandleReactions(e,postData);
reactionsContainer.append(container1,container2)
console.log(postData.categories);


let cts = div("categoriesInPost")
if(!postData.categories) {}
postData.content.categories.forEach(cat => {
  if(!cat) {
    return
  }
  cts.append(div("cat","#"+ cat))
})
  return div("postContainer").add(
    Frame(
      div("post").add(
        div("publisher").add(
          img(postData.publisher.profilePicture, "no-profile"),
          div(null, postData.publisher.username),
          div(null, timePassed(postData.creationTime))
        ),
        cts,
        div("title", postData.content.title),
        div("text", postData.content.text),
        readMore
      )
    ),
    div("leftBar").add(
      reactionsContainer,
      comment,
    )
  );
};
export default Post;