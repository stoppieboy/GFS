async function loginHandler(event) {
    event.preventDefault()
    const data = new FormData(event.target)

    const value = Object.fromEntries(data.entries())
    console.log({ value })

    var result = await fetch("login", {
        method: "POST",
        body: JSON.stringify(value),
        headers: {
            "Content-type": "application/json; charset:=UTF-8"
        }
    });

    var tokenData = await result.json()
    console.log(tokenData.token)
}

var form = document.getElementById('login-form')
form.addEventListener('submit', loginHandler)