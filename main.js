let form = document.querySelector("#lform")
form.addEventListener("submit",async(e)=>{
    e.preventDefault()
let password = document.querySelector('input[name="password"]')
let email =  document.querySelector('input[name="email"]')
let user =   document.querySelector('input[name="username"]')
await fetch("/api/register",{

    method:"post",
    body: JSON.stringify({
        password:password.value,
        email:email.value,
        user:user.value,
    })
})
})
location.reload()
