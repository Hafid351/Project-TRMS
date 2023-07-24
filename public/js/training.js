function createFormGroup(labelText, inputElement) {
    var formGroup = document.createElement("div");
    formGroup.classList.add("form-group");

    var label = document.createElement("label");
    label.textContent = labelText;

    formGroup.appendChild(label);
    formGroup.appendChild(inputElement);

    return formGroup;
}

function createInput(labelText, inputType, inputName, inputId, placeholder) {
    var input = document.createElement("input");
    input.classList.add("form-control");
    input.type = inputType;
    input.name = inputName;
    input.id = inputId;
    input.placeholder = placeholder;

    return createFormGroup(labelText, input);
}


function showCompany() {
    var company = document.getElementById('financedbyCompany').checked;
    var formCompany = document.getElementById('formCompany');
    formCompany.innerHTML = "";

    if (company === true) {
        formCompany.appendChild(createInput("Company Name", "text", "company", "company", "Enter Company Name"));
    }
}