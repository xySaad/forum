import { atBottom } from "./CommentsList.js";
import { NewReference } from "../utils/reference.js";
import Post from "./Post.js";
const getPosts = async (PostsArea, isfetch, offset,categories) => {
    isfetch(true)
    try {
      const resp = await fetch(`/api/posts?page=${offset()}&categories=${categories}`);
      if (!resp.ok) {
        throw new Error('Network response was not ok');
      }
      const json = await resp.json();
  
      offset((prev)=>prev+1)
  
      json.forEach(post => {
        console.log(post)
        PostsArea.append(Post(post))
    })
    } catch (error) {
      console.error('Error fetching comments:', error);
    } finally {
        isfetch(false); 
    }
  };
  const getcategories=()=>{
    let selectedCategories = ["0","0","0","0"]; 
    const categories = document.querySelectorAll(".category");
    categories.forEach((child,index)=>{
if (child.classList.contains("Selected")) {
    if (index==0) {
        return selectedCategories.join("")  
    }
selectedCategories[index-1]="1"    
}
    })
return selectedCategories.join("")  
}

export const CreatePostsArea=()=>{
    const PostsArea=document.querySelector(".posts")
    let offset = NewReference(0);
    let isfetch =NewReference(false)
    let categoiers=getcategories() 
    getPosts(PostsArea,isfetch,offset,categoiers)
    window.addEventListener("scroll",()=>{
        console.log("post=",atBottom(PostsArea));
     console.log("document=",atBottom(document));
   console.log("window=",atBottom(window));
if (!atBottom(window)||isfetch()) {
    return
}
getPosts(PostsArea,isfetch,offset,categoiers)
})
return PostsArea
}