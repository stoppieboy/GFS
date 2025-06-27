async function uploadHandler(event) {
    event.preventDefault()
    const data = new FormData(event.target)

    var token = window.sessionStorage.getItem("token") ?? ""
    console.log(`token passed: ${token}`)
    const authHeader = `Bearer ${token}`
    console.log(authHeader)

    const result = await fetch("file", {
        method: "POST",
        body: data,
        headers: {
            "Authorization": authHeader
        }
    })

    const file = await result.body.getReader().read()

    console.log(new TextDecoder().decode(file.value))

    if(result.status == 200) {
        event.target.reset()
        setTimeout(() => {
            alert("Upload successful")
        }, 1000);
    }else {
        alert("Upload failed")
    }
}

document.getElementById('file-form').addEventListener('submit', uploadHandler)