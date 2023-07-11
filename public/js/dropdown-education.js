function createFormGroup(labelText, inputElement) {
    var formGroup = document.createElement("div");
    formGroup.classList.add("form-group");

    var label = document.createElement("label");
    label.textContent = labelText;

    formGroup.appendChild(label);
    formGroup.appendChild(inputElement);

    return formGroup;
}

function createInput(labelText, inputType, inputName, inputId) {
    var input = document.createElement("input");
    input.classList.add("form-control");
    input.type = inputType;
    input.name = inputName;
    input.id = inputId;

    return createFormGroup(labelText, input);
}

function createNumericInput(labelText, inputName, inputId) {
    var input = document.createElement("input");
    input.classList.add("form-control");
    input.type = "number";
    input.name = inputName;
    input.id = inputId;

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
        formContainer.appendChild(createInput("School Name", "text", "originschool, origin"));
        formContainer.appendChild(createInput("GPA / Score", "text", "gpa", "gpa"));
        formContainer.appendChild(createNumericInput("Year Start", "yearstart", "yearstrat"));
        formContainer.appendChild(createNumericInput("Year End", "yearend", "yearend"));
    } else if (qualification === "3") {
        formContainer.appendChild(createInput("School Name", "text", "origin", "origin"));
        formContainer.appendChild(createInput("Major", "text", "major", "major"));
        formContainer.appendChild(createInput("GPA / Score", "text", "gpa", "gpa"));
        formContainer.appendChild(createNumericInput("Year Start", "yearstart", "yearstart"));
        formContainer.appendChild(createNumericInput("Year End", "yearend", "yearend"));
    } else if (qualification === "4" || qualification === "5" || qualification === "6" || qualification === "7" || qualification === "8" || qualification === "9" || qualification === "10") {
        fetch(`/profile/profile-wizard/qualification`)
            .then((response) => response.json())
            .then((value) => {
                formContainer.appendChild(createSelect("University", "universityid", "universityid", value.University));
                formContainer.appendChild(createSelect("Departement", "departementid", "departementid", value.Departement));
                formContainer.appendChild(createInput("GPA / Score", "text", "gpa", "gpa"));
                formContainer.appendChild(createNumericInput("Year Start:", "yearstart", "yearstart"));
                formContainer.appendChild(createNumericInput("Year End", "yearend", "yearend"));
            })
    };
}
