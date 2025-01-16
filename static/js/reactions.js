export function Reaction() {
  document.querySelectorAll(".reaction-arrow").forEach((button) => {
    button.addEventListener("click", async (event) => {
      console.log("clicked");
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
          console.error(
            `Unexpected reaction type: ${button.alt}!! Only "like" and "dislike" for now. More coming later!!`
          );
          return;
      }

      const postID = button.id;
      const isLiked = button.classList.contains("liked");
      const isDisliked = button.classList.contains("disliked");

      const method = (reactionType === "like" && isLiked) || (reactionType === "dislike" && isDisliked)
        ? "DELETE"
        : "POST";

      try {
        const resp = await fetch("/api/reactions", {
          method,
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            item_id: postID,
            reaction_type: reactionType,
          }),
          credentials: "include",
        });

        if (resp.ok) {
          const data = await resp.json();
          if (method === "DELETE") {
            if (reactionType === "like") {
              button.classList.remove("liked");
            } else {
              button.classList.remove("disliked");
            }
          } else {
            if (reactionType === "like") {
              button.classList.add("liked");
              button.classList.remove("disliked");
            } else {
              button.classList.add("disliked");
              button.classList.remove("liked"); 
            }
          }
        } else {
          console.error(`Failed to process reaction: ${resp.statusText}`);
        }
      } catch (error) {
        console.error("Error processing reaction:", error);
      }
    });

    // button.addEventListener("mouseover", async () => {
    //   let reactionType;
    //   switch (button.alt) {
    //     case "arrow-up":
    //       reactionType = "like";
    //       break;
    //     case "arrow-down":
    //       reactionType = "dislike";
    //       break;
    //     default:
    //       console.error(
    //         `Unexpected reaction type: ${button.alt}!! Only "like" and "dislike" for now. More coming later!!`
    //       );
    //       return;
    //   }

    //   const postID = button.id;
    //   try {
    //     const resp = await fetch(`/api/reactions/${postID}`, {
    //       method: "GET",
    //       headers: {
    //         "Content-Type": "application/json",
    //       },
    //       credentials: "include",
    //     });

    //     if (resp.ok) {
    //       const data = await resp.json();
    //       console.log(data);
    //     } else {
    //       console.error(`Failed to fetch reactions: ${resp.statusText}`);
    //     }
    //   } catch (error) {
    //     console.error("Error fetching reactions:", error);
    //   }
    // });
  });
}
