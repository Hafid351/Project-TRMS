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

function nextEducation(step) {
    if (step == 1) {
        const fullname = document.querySelector('#fullname')?.value;
        const gender = document.querySelector('#gender')?.value;
        const photo = document.querySelector('#photo')?.value;
        const filename = document.querySelector('#avatar')?.value;
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
            "pob": pob?? ''
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
        fetch( `/profile/profile-wizard?step=${step}`)
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
        const container5 = document.getElementById("startpositionjobtitle")
        const container6 = document.getElementById("lastpositionjobtitle")
        const container7 = document.getElementById("skillid")
        let Company = ""
        company.forEach(item => {
            Company += `<option value="${item.ID}">${item.Name}</option>`
        })
        let Country = ""
        country.forEach(item => {
            Country += `<option value="${item.ID}">${item.Name}</option>`
        })
        let Province = ""
        province.forEach(item => {
            Province += `<option value="${item.ID}">${item.Name}</option>`
        })
        let PositionLevel = ""
        positionlevel.forEach(item => {
            PositionLevel += `<option value="${item.ID}">${item.Name}</option>`
        })
        let StartPosition = ""
        position.forEach(item => {
            StartPosition += `<option value="${item.ID}">${item.Name}</option>`
        })
        let LastPosition = ""
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
        const profileid =  localStorage.getItem("profile")
        const companyid = document.querySelector('#companyid')?.value;
        const countryid = document.querySelector('#countrycompany')?.value;
        const provinceid = document.querySelector('#provincecompany')?.value;
        const positionlevelid = document.querySelector('#positionlevelid')?.value;
        const salary = document.querySelector('#salary')?.value;
        const experiencedesc = document.querySelector('#experiencedesc')?.value;
        const startdate = document.querySelector('#startdate')?.value;
        const jobtitle = document.querySelector('#jobtitle')?.value;
        const enddate = document.querySelector('#enddate')?.value;
        const lastpositionjobtitle = document.querySelector('#lastpositionjobtitle')?.value;
        const reasonleaving = document.querySelector('#reasonleaving')?.value;
        const skillid = document.querySelector('#skillid')?.value;
        const payload = {
            "profileid": Number(profileid)?? '',
            "companyid": Number(companyid)?? '',
            "countryid": Number(countryid)?? '',
            "provinceid": Number(provinceid)?? '',
            "positionlevelid": Number(positionlevelid)?? '',
            "salary": Number(salary)?? '',
            "experiencedesc": experiencedesc?? '',
            "startdate": new Date(startdate)?? '',
            "jobtitle": Number(jobtitle)?? '',
            "enddate": new Date(enddate)?? '',
            "lastpositionjobtitle": Number(lastpositionjobtitle)?? '',
            "reasonleaving": reasonleaving?? '',
            "skillid": Number(skillid)?? '',
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
            fetch(`/profile/profile_wizard/step-4?`)
                .then((response) => response.json())
                .then((value) => {
                const data = value.Data
                const language = value.Language
                const languagelevel = value.LanguageLevel
                const container = document.getElementById("languagecode")
                const container2 = document.getElementById("spokenlevel")
                const container3 = document.getElementById("writenlevel")
                const container4 = document.getElementById("listeninglevel")
                let Language = ""
                language.forEach(item => {
                    Language += `<option value="${item.ID}">${item.Name}</option>`
                })
                let SpokenLevel = ""
                languagelevel.forEach(item => {
                    SpokenLevel += `<option value="${item.ID}">${item.Name}</option>`
                })
                let WritenLevel = ""
                languagelevel.forEach(item => {
                    WritenLevel += `<option value="${item.ID}">${item.Name}</option>`
                })
                let ListeningLevel = ""
                languagelevel.forEach(item => {
                    ListeningLevel += `<option value="${item.ID}">${item.Name}</option>`
                })
                container.innerHTML = Language
                container2.innerHTML = SpokenLevel
                container3.innerHTML = WritenLevel
                container4.innerHTML = ListeningLevel
                })
            } else {
                alert (value.Data)
            }
        })
    }
}

function nextTraining(step) {
    if (step == 4) {
        const profileid =  localStorage.getItem("profile")
        const languagecode = document.querySelector('#languagecode')?.value;
        const spokenlevel = document.querySelector('#spokenlevel')?.value;
        const writenlevel = document.querySelector('#writenlevel')?.value;
        const listeninglevel = document.querySelector('#listeninglevel')?.value;
        const payload = {
            "profileid": Number(profileid)?? '',
            "languagecode": Number(languagecode)?? '',
            "spokenlevel": Number(spokenlevel)?? '',
            "writenlevel": Number(writenlevel)?? '',
            "listeninglevel": Number(listeninglevel)?? '',
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
                stepper.next(step)
            } else {
                alert (value.Data)
            }
        })
    }
}

function nextFiles(step) {
    if (step == 5) {
        const profileid =  localStorage.getItem("profile")
        const trainingtitle = document.querySelector('#trainingtitle')?.value;
        const vendor = document.querySelector('#vendor')?.value;
        const trainingyear = document.querySelector('#trainingyear')?.value;
        const durationday = document.querySelector('#durationday')?.value;
        const financedby = document.querySelector('#financedby')?.value;
        const payload = {
            "profileid": Number(profileid)?? '',
            "trainingtitle": trainingtitle?? '',
            "vendor": vendor?? '',
            "trainingyear": new Date(trainingyear)?? '',
            "durationday": Number(durationday)?? '',
            "financedby": financedby?? '',
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
                stepper.next(step)
            } else {
                alert (value.Data)
            }
        })
    }
}

