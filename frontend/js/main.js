const form = document.getElementById("form");
const addButton = document.getElementById("addButton");
const urlGetDoc = "/api/v1/doc/get";
const urlDownloadDoc = "/api/v1/doc/download/";

addButton.addEventListener("click", () => {
    const parent = document.getElementById("collapseItem");
    const addItem = document.getElementById("add");

    let num = parent.getElementsByClassName("row").length;
    document.getElementById("inp" + (num-1)).innerHTML = createItemCols(num);
    parent.insertBefore(createItemRow(num), addItem)
})

function createItemCols(number) {
     return  '<div class="col-md-1 col-sm-12 col-xs-12 mt-1" id="del">' +
        '<button class="btn btn-outline-secondary icon-minus" id="delButton" type="button"></button>' +
        '</div>' +
        '<div class="col-md-2 col-sm-12 col-xs-12 mt-1">' +
        '<input id="art' + (number) + '" class="form-control" type="text" name="art1" placeholder="Артикул">' +
        '</div>' +
        '<div class="col-md-4 col-sm-12 col-xs-12 mt-1">' +
        '<input id="itemName1" class="form-control" type="text" name="itemName1" placeholder="Наименование">' +
        '</div>' +
        '<div class="col-md-2 col-sm-12 col-xs-12 mt-1">' +
        '<input id="num1" class="form-control" type="text" name="num1" placeholder="Кол-во">' +
        '</div>' +
        '<div class="col-md-3 col-sm-12 col-xs-12 mt-1">' +
        '<input id="cost1" class="form-control" type="text" name="cost1" placeholder="Цена 1шт.">' +
        '</div>'

}

function createItemRow(number) {
    const inpRow = document.createElement("div")
    inpRow.classList.add("row")
    inpRow.id = "inp" + number
    return inpRow
}

form.addEventListener("submit", (event) => {
    event.preventDefault();

    if (validation(form)) {
        document.getElementsByClassName("error-label").textContent = "";
        newDocumentRequest(event);
    } else {
        document.getElementsByClassName("error-label").textContent = "Для получения договора необходимо заполнить все поля!";
    }
})

function validation(form) {
    let validRes = true;

    form.querySelectorAll("input").forEach(input => {
        if (input.value === "") {
            input.classList.add("error");
            validRes = false;
        } else {
            input.classList.remove("error");
        }
    })

    return validRes;
}

function newDocumentRequest(event) {
    event.preventDefault();
    const formData = new FormData(event.target);
    const fData = Object.fromEntries(formData.entries());
    console.log(fData.clientFullName);

    const body = {
        clientFullName: fData.clientFullName,
        clientPhoneNumber: fData.clientPhoneNumber,
        clientEmail: fData.clientEmail,
        clientCompany: fData.clientCompany,
        clientCompanyFullName: fData.clientCompanyFullName,
        clientCompanyInnKpp: fData.clientCompanyInnKpp,
        clientCompanyAddress: fData.clientCompanyAddress,

        providerFullName: fData.providerFullName,
        providerPhoneNumber: fData.providerPhoneNumber,
        providerEmail: fData.providerEmail,
        providerCompany: fData.providerCompany,
        providerCompanyInnKpp: fData.providerCompanyInnKpp,
        providerCompanyAddress: fData.providerCompanyAddress,

        deliveryAddress: fData.deliveryAddress,
    };

    async function sendFormData() {
        try {
            const response = await fetch(urlGetDoc, {
                method: "POST", body: JSON.stringify(body)
            })
            if (response.status === 200) {
                let res = await response.json()
                downloadFile(res.fileName);
                console.log(res.fileName);
            }
        } catch (err) {
            console.log(err);
        }
    }

    sendFormData()
}

async function downloadFile(filename) {
    const response = await fetch(urlDownloadDoc + filename);

    const blob = await response.blob();

    let link = document.createElement('a');
    link.href = window.URL.createObjectURL(blob);
    link.download = filename;
    link.click();
}




