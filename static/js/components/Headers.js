import div from "./native/div.js";
import img from "./native/img.js";
function getPosts(type) {
  console.log(type);
  
  let icns = document.querySelectorAll("svg")
  icns.forEach((icn) => icn.classList.remove("active"))
  let clicked = document.querySelector(type)
  clicked.classList.add("active")
}
function toggleIt() {
  let ul = document.querySelector(".icons")
  ul.classList.toggle("active")
}
function ToggleDisplay() {
  let item = document.querySelector(".profileCard");
  if (item.style.display === "none" || item.style.display === "") {
    item.style.display = "flex";
  } else {
    item.style.display = "none";
  }
}
export function appendUserHeader() {
  let icn1 = div("contain")
  let icn2 = div("contain")
  let icn3 = div("contain")
  let logout = div("logoutBtn")
  let h2 = document.createElement("h2")
  h2.innerText = "chakir Ben"
  let h4 = document.createElement("h4")
  h4.innerText = "chakir.benlafkih@gmail.com"
  logout.innerHTML = `<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M22.4206 11.5898L19.4926 8.67383C19.1986 8.38183 18.7246 8.38183 18.4326 8.67583C18.1406 8.96983 18.1416 9.44383 18.4346 9.73583L20.0746 11.3698H17.1396C17.1496 11.8698 17.1496 12.3698 17.1396 12.8698H20.0766L18.4346 14.5058C18.1416 14.7978 18.1406 15.2728 18.4326 15.5668C18.5786 15.7138 18.7716 15.7868 18.9636 15.7868C19.1546 15.7868 19.3466 15.7138 19.4926 15.5688L22.4206 12.6518C22.5626 12.5118 22.6416 12.3198 22.6416 12.1208C22.6416 11.9218 22.5626 11.7308 22.4206 11.5898Z" fill="#FD5F49"/>
          <path fill-rule="evenodd" clip-rule="evenodd" d="M8.88938 12.12C8.88938 11.71 9.21938 11.37 9.63938 11.37L17.1397 11.3698C17.1297 10.1098 17.0594 8.85 16.9594 7.59V7.58C16.5894 3.55 14.7594 2.25 9.45937 2.25C1.85938 2.25 1.85938 5.1 1.85938 12C1.85938 18.9 1.85938 21.75 9.45937 21.75C14.7594 21.75 16.5894 20.45 16.9594 16.41C17.0594 15.24 17.1197 14.0598 17.1397 12.8698L9.63938 12.87C9.21938 12.87 8.88938 12.54 8.88938 12.12Z" fill="#FD5F49"/>
        </svg>` + "<h3>Logout</h3>  ";
  icn1.innerHTML = `<svg class="icon home active"  width="24" height="24" viewBox="0 0 24 22" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M7.30665 18.7733V15.7156C7.30665 14.9351 7.94405 14.3023 8.73031 14.3023H11.6045C11.9821 14.3023 12.3442 14.4512 12.6112 14.7163C12.8782 14.9813 13.0281 15.3408 13.0281 15.7156V18.7733C13.0258 19.0978 13.1539 19.4099 13.3842 19.6402C13.6146 19.8705 13.9279 20 14.2548 20H16.2157C17.1315 20.0023 18.0106 19.6428 18.659 19.0008C19.3075 18.3588 19.6719 17.487 19.6719 16.5778V7.86686C19.6719 7.13246 19.3439 6.43584 18.7765 5.96467L12.1059 0.675869C10.9455 -0.251438 9.28299 -0.221498 8.15727 0.746979L1.63889 5.96467C1.04462 6.42195 0.689427 7.12064 0.671875 7.86686V16.5689C0.671875 18.4639 2.21926 20 4.12805 20H6.04417C6.7231 20 7.27487 19.4562 7.27979 18.7822L7.30665 18.7733Z"/>
          </svg>`;
  icn2.innerHTML = `<svg class="icon liked  "width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M3.79458 12.4099C2.72158 9.05981 3.97658 4.89481 7.49358 3.76281C9.34358 3.16581 11.6266 3.66381 12.9236 5.45281C14.1466 3.59781 16.4956 3.16981 18.3436 3.76281C21.8596 4.89481 23.1216 9.05981 22.0496 12.4099C20.3796 17.7199 14.5526 20.4859 12.9236 20.4859C11.7126 20.4859 8.20858 18.9909 5.78258 16.0149"  />
        <path d="M16.6621 7.52734C17.8691 7.65134 18.6241 8.60834 18.5791 9.94934" />
      </svg>`;
  icn3.innerHTML = `<svg class="icon created"  width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M13.604 19.3789H19.981" />
        <path d="M5.31708 14.3319C3.57808 16.6499 5.22308 19.5219 5.22308 19.5219C5.22308 19.5219 8.46708 20.2679 10.1811 17.9829C11.8961 15.6989 16.9331 8.98786 16.9331 8.98786C17.9411 7.64486 17.6701 5.73786 16.3271 4.72986C14.9831 3.72186 13.0771 3.99386 12.0691 5.33686C12.0691 5.33686 9.07408 9.32686 6.91008 12.2089" />
        <path d="M12.7771 8.63867L15.6371 10.7326" />
      </svg>`;

  let h = div("header").add(
    img("../../static/svg/logo.svg", "logo"),
    div("close", "☰"),
    div("icons").add(
      icn1, icn2, icn3
    ),
    div("profileContainer").add(
      img("avatar", "avatar", "profile")
    ),
    div("profileCard").add(
      div("textContainer").add(
        h2, h4
      ),
      div("line"),
      logout
    )
  )
  let head = document.querySelector("header")
  head.append(h)
  document.querySelector(".profileContainer").addEventListener("click", ()=>{ToggleDisplay()})
  document.querySelector(".home").addEventListener("click", ()=>{getPosts('.home')})
  document.querySelector(".liked").addEventListener("click", ()=>{getPosts('.liked')})
  document.querySelector(".created").addEventListener("click", ()=>{getPosts('.created')})
  document.querySelector(".close").addEventListener("click", ()=>{toggleIt()})
}

export function appendGuestHeader() {

}