export default async function HandleReactions(e,postData) {
    let myDiv = e.target
    if (!e.target.classList.contains("reaction")) {
      myDiv = e.target.closest(".reaction");
    }
    if (myDiv.classList.contains("reacted")) {
      if (myDiv.classList.contains("like")) {
        let res = await fetch("/api/reactions" , {
          method : "DELETE",
          body : {
            itemId : postData.itemId,
            itemType : 1,
            reactionType : 1
          }
        })
      }else {

      }
      myDiv.classList.remove("reacted")
    }else {
      if (myDiv.classList.contains("like")) {

      }else {

      }
      myDiv.classList.add("reacted")
      let otherDiv = myDiv.nextElementSibling
      if (!otherDiv) {
        otherDiv = myDiv.previousElementSibling
      }
      otherDiv.classList.remove("reacted")
    }
}