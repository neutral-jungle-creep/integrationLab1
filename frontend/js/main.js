const form = document.getElementById("form")
const urlGetDoc = "/api/v1/doc/get"
const urlDownloadDoc = "/api/v1/doc/download/"

form.addEventListener("submit", (event)=> {
    event.preventDefault()
    const formData = new FormData(event.target);
    const fData = Object.fromEntries(formData.entries());
    console.log(fData.clientFullName)

    const body = {
        clientFullName: fData.clientFullName,
        clientPhoneNumber: fData.clientPhoneNumber,
        clientEmail: fData.clientEmail
    }

    async function sendFormData() {
        try {
            const response = await fetch(urlGetDoc, {
                method: "POST",
                body: JSON.stringify(body)
            })
            if (response.status === 200) {
                let res = await response.json()
                downloadFile(res.fileName)
                console.log(res.fileName)
            }
        } catch (err) {
            console.log(err)
        }
    }
    sendFormData()
})

async function downloadFile(filename) {
    const response = await fetch(urlDownloadDoc + filename);

    const blob = await response.blob();

    let link = document.createElement('a');
    link.href = window.URL.createObjectURL(blob);
    link.download = filename;
    link.click();
}




