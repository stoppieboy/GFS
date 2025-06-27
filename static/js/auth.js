async function loginHandler(event) {
    event.preventDefault()
    const data = new FormData(event.target)

    // form-data to json
    const value = Object.fromEntries(data.entries())
    console.log({ value })

    var result = await fetch("login", {
        method: "POST",
        body: JSON.stringify(value),
        headers: {
            "Content-type": "application/json; charset:=UTF-8"
        }
    });

    if(result.status != 200) {
        alert('login failed')
        return
    }

    var tokenData = await result.json()
    console.log(tokenData.token)
    window.sessionStorage.setItem("token", tokenData.token)
    window.location.replace("http://localhost:3000/somepath")
}

var form = document.getElementById('login-form')
form.addEventListener('submit', loginHandler)