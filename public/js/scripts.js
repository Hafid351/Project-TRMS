/*!
    * Start Bootstrap - SB Admin v7.0.7 (https://startbootstrap.com/template/sb-admin)
    * Copyright 2013-2023 Start Bootstrap
    * Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-sb-admin/blob/master/LICENSE)
    */
    // 
// Scripts
// 

window.addEventListener('DOMContentLoaded', event => {
    // Toggle the side navigation
    const sidebarToggle = document.body.querySelector('#sidebarToggle');
    if (sidebarToggle) {
        // Uncomment Below to persist sidebar toggle between refreshes
        // if (localStorage.getItem('sb|sidebar-toggle') === 'true') {
        //     document.body.classList.toggle('sb-sidenav-toggled');
        // }
        sidebarToggle.addEventListener('click', event => {
            event.preventDefault();
            document.body.classList.toggle('sb-sidenav-toggled');
            localStorage.setItem('sb|sidebar-toggle', document.body.classList.contains('sb-sidenav-toggled'));
        });
    }
});

function validateEmail() {
    const emailInput = document.getElementById('email');
    const emailError = document.getElementById('emailError');
    // Regular expression for email validation
    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (emailPattern.test(emailInput.value)) {
      // Email is valid
    emailError.textContent = '';
    emailInput.style.borderColor = 'green';
      return true; // Return true if email is valid
    } else {
      // Email is invalid
    emailError.textContent = 'Invalid email address';
    emailInput.style.borderColor = 'red';
      return false; // Return false if email is invalid
    }
}

function nextEducation(step) {
if (step == 1) {
        const isValidEmail = validateEmail();
        if (!isValidEmail) {
        // Show alert if email is invalid
        alert('Please enter a valid email address.');
        return; // Stop further execution if email is invalid
        }
        const fullname = document.querySelector('#fullname')?.value;
        const gender = document.querySelector('#gender')?.value;
        const photo = document.querySelector('#photo')?.value;
        //const filename = document.querySelector('#avatar')?.value;
        const filename = $("#browse_image").prop("files")[0].name;
        const countryid = document.querySelector('#countryid')?.value;
        const provinceid = document.querySelector('#provinceid')?.value;
        const cityid = document.querySelector('#cityid')?.value;
        const address = document.querySelector('#address')?.value;
        const nationalityid = document.querySelector('#nationalityid')?.value;
        const height = document.querySelector('#height')?.value;
        const weight = document.querySelector('#weight')?.value;
        const religionid = document.querySelector('#religionid')?.value;
        const maritalstatus  = document.querySelector('#maritalstatus')?.value;
        const identificationid = document.querySelector('#identificationid')?.value;
        const idcardno = document.querySelector('#idcardno')?.value;
        const phonenumber = document.querySelector('#phonenumber')?.value;
        const otherphonenumber = document.querySelector('#otherphonenumber')?.value;
        const email = document.querySelector('#email')?.value;
        const salaryexpectation = document.querySelector('#salaryexpectation')?.value;
        const jobtitle = document.querySelector('#jobtitle')?.value;
        const about = document.querySelector('#about')?.value;
        const instagramprofile = document.querySelector('#instagramprofile')?.value;
        const facebookprofile = document.querySelector('#facebookprofile')?.value;
        const linkedinprofile = document.querySelector('#linkedinprofile')?.value;
        const dob = document.querySelector('#dob')?.value;
        const pob = document.querySelector('#pob')?.value;
        const image = document.querySelector("#cropped_image_result img");
        const base64data = image.src;
        const payload = {
            "fullname": fullname?? '',
            "gender": gender?? '',
            "photo": photo?? '',
            "filename": filename?? '',
            "countryid": Number(countryid)?? '',
            "provinceid": Number(provinceid)?? '',
            "cityid": Number(cityid)?? '',
            "address": address?? '',
            "nationalityid": nationalityid?? '',
            "height": Number(height)?? '',
            "weight": Number(weight)?? '',
            "religionid": Number(religionid)?? '',
            "maritalstatus": Number(maritalstatus)?? '',
            "identificationid": Number(identificationid)?? '',
            "idcardno": idcardno?? '',
            "phonenumber": phonenumber?? '',
            "otherphonenumber": otherphonenumber?? '',
            "email": email?? '',
            "salaryexpectation": Number(salaryexpectation)?? '',
            "jobtitle": Number(jobtitle)?? '',
            "about": about?? '',
            "instagramprofile": instagramprofile?? '',
            "facebookprofile": facebookprofile?? '',
            "linkedinprofile": linkedinprofile?? '',
            "dob": new Date(dob)?? '',
            "pob": pob?? '',
            "image": base64data?? ''
        }
        fetch(`/profile/profile-wizard?step=${step}`, {
            method: `POST`,
            headers: {
                'content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        })
        .then((response) => response.json())
        .then((value) => {
        if (value.Data == "Ok") {
            localStorage.setItem("profile", value.Profileid)
            stepper.next()
            fetch(`/profile/profile-wizard?step=${step}`)
            .then((response) => response.json())
            .then((value) => {
            const data = value.Data
            const qualification = value.Qualification
            const container = document.getElementById("qualificationid");
            let Qualification = ""
            Qualification += `<option selected disabled>Select Education Level</option>`
            qualification.forEach(item => {
                Qualification += `<option value="${item.ID}">${item.Name}</option>`
            })
            console.log(Qualification)
            container.innerHTML = Qualification
            console.log(container)
            })
        } else {
            alert (value.Data)
        }
        })
    }  
}

function nextWork(step) {
    if (step == 2) {
        stepper.next()
        fetch(`/profile/profile-wizard?step=${step}`)
        .then((response) => response.json())
        .then((value) => {
        const data = value.Data
        const company = value.Company
        const country = value.Country
        const province = value.Province
        const positionlevel = value.PositionLevel
        const position = value.Position
        const skill = value.Skill
        const container = document.getElementById("companyid")
        const container2 = document.getElementById("countrycompany")
        const container3 = document.getElementById("provincecompany")
        const container4 = document.getElementById("positionlevelid")
        const container5 = document.getElementById("startpositionjobtittle")
        const container6 = document.getElementById("lastpositionjobtittle")
        const container7 = document.getElementById("skillid")
        let Company = ""
        Company +=  `<option selected disabled>Select Company</option>`
        company.forEach(item => {
            Company += `<option value="${item.ID}">${item.Name}</option>`
        })
        let Country = ""
        Country += `<option selected disabled>Select Country</option>`
        country.forEach(item => {
            Country += `<option value="${item.ID}">${item.Name}</option>`
        })
        let Province = ""
        Province += `<option selected disabled>Select Province</option>`
        province.forEach(item => {
            Province += `<option value="${item.ID}">${item.Name}</option>`
        })
        let PositionLevel = ""
        PositionLevel += `<option selected disabled>Select Position Level</option>`
        positionlevel.forEach(item => {
            PositionLevel += `<option value="${item.ID}">${item.Name}</option>`
        })
        let StartPosition = ""
        StartPosition += `<option selected disabled>Select Start Position</option>`
        position.forEach(item => {
            StartPosition += `<option value="${item.ID}">${item.Name}</option>`
        })
        let LastPosition = ""
        LastPosition += `<option selected disabled>Select Last Position</option>`
        position.forEach(item => {
            LastPosition += `<option value="${item.ID}">${item.Name}</option>`
        })
        let Skill = ""
        skill.forEach(item => {
            Skill += `<option value="${item.ID}">${item.Name}</option>`
        })
        container.innerHTML = Company 
        container2.innerHTML = Country
        container3.innerHTML = Province
        container4.innerHTML = PositionLevel
        container5.innerHTML = StartPosition
        container6.innerHTML = LastPosition
        container7.innerHTML = Skill
        })
    } else {
        alert (value.Data)
    }
}

function nextLanguage(step) {
    if (step == 3) {
        stepper.next()
        fetch(`/profile/profile-wizard?step=${step}`)
        .then((response) => response.json())
        .then((value) => {
        const data = value.Data
        const language = value.Language
        const languagelevel = value.LanguageLevel
        const container = document.getElementById("languagecode")
        const container2 = document.getElementById("spokenlevel")
        const container3 = document.getElementById("writtenlevel")
        const container4 = document.getElementById("listeninglevel")
        let Language = ""
        Language += `<option selected disabled>Select Language</option>`
        language.forEach(item => {
            Language += `<option value="${item.ID}">${item.Name}</option>`
        })
        let SpokenLevel = ""
        SpokenLevel += `<option selected disabled>Select Spoken Level</option>`
        languagelevel.forEach(item => {
            SpokenLevel += `<option value="${item.ID}">${item.Name}</option>`
        })
        let WrittenLevel = ""
        WrittenLevel += `<option selected disabled>Select Written Level</option>`
        languagelevel.forEach(item => {
            WrittenLevel += `<option value="${item.ID}">${item.Name}</option>`
        })
        let ListeningLevel = ""
        ListeningLevel += `<option selected disabled>Select Listening Level</option>`
        languagelevel.forEach(item => {
            ListeningLevel += `<option value="${item.ID}">${item.Name}</option>`
        })
        container.innerHTML = Language
        container2.innerHTML = SpokenLevel
        container3.innerHTML = WrittenLevel
        container4.innerHTML = ListeningLevel
        })
    } else {
        alert (value.Data)
    }
}

function nextTraining(step) {
    if (step == 4) {
        stepper.next(step)
    } else {
        alert (value.Data)
    }
}

function nextFiles(step) {
    if (step == 5) {
        stepper.next(step)
    } else {
        alert (value.Data)
    }
}
