function createFormGroup(labelText, inputElement) {
    var formGroup = document.createElement("div");
    formGroup.classList.add("form-group");

    var label = document.createElement("label");
    label.textContent = labelText;

    formGroup.appendChild(label);
    formGroup.appendChild(inputElement);

    return formGroup;
}

function createInput(labelText, inputType, inputName) {
    var input = document.createElement("input");
    input.classList.add("form-control");
    input.type = inputType;
    input.name = inputName;

    return createFormGroup(labelText, input);
}

function createNumericInput(labelText, inputName) {
    var input = document.createElement("input");
    input.classList.add("form-control");
    input.type = "number";
    input.name = inputName;

    return createFormGroup(labelText, input);
}

function createSelect(labelText, selectId, selectName, options) {
    var select = document.createElement("select");
    select.classList.add("form-control");
    select.name = selectName;
    select.id = selectId;

    for (var i = 0; i < options.length; i++) {
    var option = document.createElement("option");
    option.value = options[i].ID;
    option.textContent = options[i].Name;
    select.appendChild(option);
    }

    return createFormGroup(labelText, select);
}

function showForm() {
    var qualification = document.getElementById("qualificationid").value;
    var formContainer = document.getElementById("formContainer");
    formContainer.innerHTML = "";
    console.log(qualification)

    if (qualification === "1" || qualification === "2") {
        formContainer.appendChild(createInput("School Name", "text", "originschool"));
        formContainer.appendChild(createInput("GPA / Score", "text", "gpa"));
        formContainer.appendChild(createNumericInput("Year Start", "yearstart"));
        formContainer.appendChild(createNumericInput("Year End", "yearend"));
    } else if (qualification === "3") {
        formContainer.appendChild(createInput("School Name", "text", "originschool"));
        formContainer.appendChild(createInput("Departement", "text", "departement"));
        formContainer.appendChild(createInput("GPA / Score", "text", "gpa"));
        formContainer.appendChild(createNumericInput("Year Start", "yearstart"));
        formContainer.appendChild(createNumericInput("Year End", "yearend"));
    } else if (qualification === "4" || qualification === "5" || qualification === "6" || qualification === "7" || qualification === "8" || qualification === "9" || qualification === "10") {
        fetch(`/profile/profile-wizard?step=1`)
            .then((response) => response.json())
            .then((value) => {
                formContainer.appendChild(createSelect("University", "universityid", "universityid", value.University));
                formContainer.appendChild(createSelect("Departement", "departementid", "departementid", value.Departement));
                formContainer.appendChild(createInput("GPA / Score", "text", "gpa"));
                formContainer.appendChild(createNumericInput("Year Start:", "yearstart"));
                formContainer.appendChild(createNumericInput("Year End", "yearend"));
            })
    };
}
