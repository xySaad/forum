export function Reaction() {
    document.querySelectorAll(".reaction-btn").forEach((button) => {

      button.addEventListener("click", async (event) => {
        console.log('clicked')
        event.preventDefault();
  
        let reactionType;
        switch (button.alt) {
          case "arrow-up":
            reactionType = "like";
            break;
          case "arrow-down":
            reactionType = "dislike";
            break;
          default:
            console.error(`Unexpected reaction type: ${button.alt}!! only Like and Dislike for now. more coming later!!`);
            return; 
        }
        const postID = button.id;
        let userID = getUserID();
  
        if (!userID) {
          alert("User is not authenticated. Please log in.");
          return;
          // userID = "user-4"
        }

  
        try {
          const resp = await fetch("/api/reactions", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              user_id: userID,
              item_id: postID, // Assuming `item_id` can refer to post or comment
              reaction_type: reactionType,
            }),
          });
  
          if (resp.ok) {
            const data = await resp.json();
          } else {
            console.error(`Failed to add reaction: ${resp.statusText}`);
          }
        } catch (error) {
          console.error("Error adding reaction:", error);
        }
      });
      button.addEventListener("hover", async (event) => {
        let reactionType;
        switch (button.alt) {
          case "arrow-up":
            reactionType = "like";
            break;
          case "arrow-down":
            reactionType = "dislike";
            break;
          default:
             console.error(`Unexpected reaction type: ${button.alt}!! only Like and Dislike for now. more coming later!!`);
            return; 
        }
        const postID = button.id;
        
      })
    });
  
} 


function getUserID() {
  const cookies = document.cookie.split("; ").reduce((acc, cookie) => {
    const [key, value] = cookie.split("=");
    acc[key] = value;
    return acc;
  }, {});

  return cookies["token"];
}