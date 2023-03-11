const form = document.getElementById("form");
const addButton = document.getElementById("addButton");
const urlGetDoc = "/api/v1/doc/get";
const urlDownloadDoc = "/api/v1/doc/download/";

addButton.addEventListener("click", () => {
    const parent = document.getElementById("collapseItem");
    const addItem = document.getElementById("add");

    let num = parent.getElementsByClassName("row").length;
    const newItem = createItemRow(num);
    parent.insertBefore(newItem, addItem);
})

function createItemRow(number) {
    const itemRow = document.createElement("div");

    const itemCol1 = document.createElement("div");
    const itemCol2 = document.createElement("div");
    const itemCol3 = document.createElement("div");
    const itemCol4 = document.createElement("div");

    const colArray = [itemCol1, itemCol2, itemCol3, itemCol4]

    const itemInput1 = document.createElement("input");
    const itemInput2 = document.createElement("input");
    const itemInput3 = document.createElement("input");
    const itemInput4 = document.createElement("input");


    itemRow.classList.add("row");

    itemCol1.classList.add("col-md-2");
    itemCol2.classList.add("col-md-5");
    itemCol3.classList.add("col-md-2");
    itemCol4.classList.add("col-md-3");

    colArray.forEach(elem => {
            elem.classList.add("col-sm-12")
            elem.classList.add("col-xs-12")
            elem.classList.add("mt-1")
        }
    )
    itemInput1.classList.add("form-control");
    itemInput2.classList.add("form-control");
    itemInput3.classList.add("form-control");
    itemInput4.classList.add("form-control");

    itemInput1.type = "text";
    itemInput2.type = "text";
    itemInput3.type = "text";
    itemInput4.type = "text";

    itemInput1.id = "art" + number;
    itemInput2.id = "itemName" + number;
    itemInput3.id = "num" + number;
    itemInput4.id = "cost" + number;

    itemInput1.nodeName = "art" + number;
    itemInput2.nodeName = "itemName" + number;
    itemInput3.nodeName = "num" + number;
    itemInput4.nodeName = "cost" + number;

    itemInput1.placeholder = "Артикул";
    itemInput2.placeholder = "Наименование";
    itemInput3.placeholder = "Кол-во";
    itemInput4.placeholder = "Цена 1шт.";

    itemCol1.append(itemInput1)
    itemCol2.append(itemInput2)
    itemCol3.append(itemInput3)
    itemCol4.append(itemInput4)

    itemRow.append(itemCol1, itemCol2, itemCol3, itemCol4)

    return itemRow
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




