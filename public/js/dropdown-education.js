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

function createNumericInput(labelText, inputName, inputId, placeholder) {
    var input = document.createElement("input");
    input.classList.add("form-control");
    input.type = "number";
    input.name = inputName;
    input.id = inputId;
    input.placeholder = placeholder;

    return createFormGroup(labelText, input);
}

function createSelect(labelText, selectId, selectName, options, label) {
    var select = document.createElement("select");
    select.classList.add("form-control");
    select.name = selectName;
    select.id = selectId;

    var defaultOption = document.createElement("option");
    defaultOption.value = "";
    defaultOption.textContent = label;
    defaultOption.selected = true;
    defaultOption.disabled = true;
    select.appendChild(defaultOption);

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

    if (qualification === "1" || qualification === "2") {
        formContainer.appendChild(createInput("School Name", "text", "origindasar", "origindasar", "Enter Your School Name"));
        formContainer.appendChild(createInput("Score", "text", "gpa", "gpa", "Enter Your Score (10 - 100)"));
        formContainer.appendChild(createNumericInput("Year Start", "yearstart", "yearstart", "Enter Your Start Year (2000)"));
        formContainer.appendChild(createNumericInput("Year End", "yearend", "yearend", "Enter Your End Year (2000)"));
    } else if (qualification === "3") {
        formContainer.appendChild(createInput("School Name", "text", "originlanjut", "originlanjut", "Enter Your School Name"));
        formContainer.appendChild(createInput("Major", "text", "major", "major", "Enter Your Major (IPA / IPS)"));
        formContainer.appendChild(createInput("Score", "text", "gpa", "gpa", "Enter Your Score (10 - 100)"));
        formContainer.appendChild(createNumericInput("Year Start", "yearstart", "yearstart", "Enter Your Start Year (2000)"));
        formContainer.appendChild(createNumericInput("Year End", "yearend", "yearend", "Enter Your End Year (2000)"));
    } else if (qualification === "4" || qualification === "5" || qualification === "6" || qualification === "7" || qualification === "8" || qualification === "9" || qualification === "10") {
        fetch(`/profile/profile-wizard/qualification`)
            .then((response) => response.json())
            .then((value) => {
                formContainer.appendChild(createSelect("University", "universityid", "universityid", value.University));
                formContainer.appendChild(createSelect("Departement", "departementid", "departementid", value.Departement));
                formContainer.appendChild(createInput("GPA", "text", "gpa", "gpa", "Enter Your GPA (1.00 - 4.00)"));
                formContainer.appendChild(createNumericInput("Year Start:", "yearstart", "yearstart", "Enter Your Start Year (2000)"));
                formContainer.appendChild(createNumericInput("Year End", "yearend", "yearend", "Enter Your End Year (2000)"));
            })
    };
}