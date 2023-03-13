const form = document.getElementById("form");
const templateForm = document.getElementById("form-check");
const downloadMsg = document.getElementById("download-msg");

const addButton = document.getElementById("addButton");
const itemsForm = document.getElementById("collapseItem");

const urlGetDoc = "/api/v1/doc/get";
const urlDownloadDoc = "/api/v1/doc/download/";


form.addEventListener("submit", (event) => {
    event.preventDefault();
    const filePath = getCheckboxValue();

    console.log(filePath);

    if (validation(form)) {
        downloadMsg.classList.remove("error-label");
        downloadMsg.textContent = "";
        newDocumentRequest(filePath);
    } else {
       showDownloadError("Необходимо заполнить все поля");
    }
})

function showDownloadError(msg) {
    downloadMsg.classList.add("error-label");
    downloadMsg.textContent = msg;
}

function getCheckboxValue() {
    let path = "";

    templateForm.querySelectorAll("input").forEach(input => {
        if (input.checked) {
            path = input.value;
        }
    })
    return path;
}

function validation(form) {
    let validRes = true;
    let inputs = form.getElementsByClassName("form-control");

    for (let i = 0; i < inputs.length; i++) {
        let input = inputs[i];
        if (input.value === "") {
            input.classList.add("error");
            validRes = false;
        } else {
            input.classList.remove("error");
        }
    }

    if (mapMsg.classList.contains("error-label") || deliveryAddr.value === "") {
        deliveryAddr.classList.add("error");
        validRes = false;
    } else {
        deliveryAddr.classList.remove("error");
    }
    return validRes;
}

function newDocumentRequest(filePath) {
    const body = makeBody(filePath);
    console.log(body);

    async function sendFormData() {
        try {
            const response = await fetch(urlGetDoc, {
                method: "POST", body: JSON.stringify(body)
            })
            if (response.status === 200) {
                let res = await response.json()
                downloadFile(res.fileName);
                console.log(res.fileName);
                downloadMsg.textContent = "Договор успешно сгенерирован"
            } else {
                showDownloadError("Заполните поля формы корректными данными");
            }
        } catch (err) {
            console.log(err);
        }
    }

    sendFormData()
}

function makeBody(filePath) {
    const formData = new FormData(form);

    return {
        templateFile: filePath,

        clientFullName: formData.get("clientFullName"),
        clientPhoneNumber: formData.get("clientPhoneNumber"),
        clientEmail: formData.get("clientEmail"),
        clientCompany: formData.get("clientCompany"),
        clientCompanyFullName: formData.get("clientCompanyFullName"),
        clientCompanyInnKpp: formData.get("clientCompanyInnKpp"),
        clientCompanyAddress: formData.get("clientCompanyAddress"),

        providerFullName: formData.get("providerFullName"),
        providerPhoneNumber: formData.get("providerPhoneNumber"),
        providerEmail: formData.get("providerEmail"),
        providerCompany: formData.get("providerCompany"),
        providerCompanyFullName: formData.get("providerCompanyFullName"),
        providerCompanyInnKpp: formData.get("providerCompanyInnKpp"),
        providerCompanyAddress: formData.get("providerCompanyAddress"),

        deliveryAddress: formData.get("deliveryAddress"),

        items: getItems(formData),
    };
}

function getItems(formData) {
    const inputForms = itemsForm.getElementsByClassName("input-item");
    const items = []

    for (let i = 0; i < inputForms.length; i++) {
        items[i] = {
            vendorCode: formData.get("vendorCode" + i),
            itemName: formData.get("itemName" + i),
            quantity: parseInt(formData.get("quantity" + i)),
            price: parseInt(formData.get("price" + i)),
        }
    }
    return items
}

async function downloadFile(filename) {
    const response = await fetch(urlDownloadDoc + filename);

    const blob = await response.blob();

    let link = document.createElement('a');
    link.href = window.URL.createObjectURL(blob);
    link.download = filename;
    link.click();
}


addButton.addEventListener("click", () => {
    const addItem = document.getElementById("add");

    let num = itemsForm.getElementsByClassName("row").length - 1;
    itemsForm.insertBefore(createItemRow(num), addItem);

    document.getElementById("inp" + (num)).innerHTML = createItemCols(num);
})

function createItemCols(number) {
    // TODO add button for del item row
    return '<div class="col-md-2 col-sm-12 col-xs-12 mt-1">' +
        '<input name="vendorCode' + (number) + '" class="form-control" type="text" placeholder="Артикул">' +
        '</div>' +
        '<div class="col-md-5 col-sm-12 col-xs-12 mt-1">' +
        '<input name="itemName' + (number) + '" class="form-control" type="text" placeholder="Наименование">' +
        '</div>' +
        '<div class="col-md-2 col-sm-12 col-xs-12 mt-1">' +
        '<input name="quantity' + (number) + '" class="form-control" type="number" placeholder="Кол-во">' +
        '</div>' +
        '<div class="col-md-3 col-sm-12 col-xs-12 mt-1">' +
        '<input name="price' + (number) + '" class="form-control" type="number" placeholder="Цена 1шт.">' +
        '</div>'

}

function createItemRow(number) {
    const inpRow = document.createElement("div")
    inpRow.classList.add("row")
    inpRow.classList.add("input-item")
    inpRow.id = "inp" + number
    return inpRow
}