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


function showIndo() {
    var countryid = document.getElementById("countryid").value;
    var formCountry = document.getElementById('formCountry');
    formCountry.innerHTML = ""; 
    // var defaultProvinceOption = document.createElement("option");
    // defaultProvinceOption.value = "";
    // defaultProvinceOption.textContent = "Select City";
    // formCountry.appendChild(defaultProvinceOption);
    console.log(countryid)
    if (countryid === "102") {
        fetch(`/profile/profile-wizard/country?` +
            new URLSearchParams({
            countryid: countryid,
        }))
            .then((response) => response.json())
            .then((value) => {
                formCountry.appendChild(createSelect("Select Province", "provinceid", "provinceid", value.Province, "Select Province"));
                var citySelect = createSelect("Select City", "cityid", "cityid", []);
                formCountry.appendChild(citySelect);
                document.getElementById("provinceid").onchange = function () {
                fetchCities(citySelect);
            }
        });
    }
}

function fetchCities(citySelect) {
    var selectedProvince = document.getElementById("provinceid").value;
    citySelect.innerHTML = "";
    // var defaultCityOption = document.createElement("option");
    // defaultCityOption.value = "";
    // defaultCityOption.textContent = "Select City";
    // citySelect.appendChild(defaultCityOption);
    fetch(`/profile/profile_wizard/country/city?` +
        new URLSearchParams({
            provinceid: selectedProvince,
        }))
        .then((response) => response.json())
        .then((value) => {
            citySelect.appendChild(createSelect("Select City", "cityid", "cityid", value.City, "Select City"));
        });
}