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
  
        try {
          const resp = await fetch("/api/reactions", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              item_id: postID, 
              reaction_type: reactionType,
            }),
            credentials: 'include'
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