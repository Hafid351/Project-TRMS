function saveEducation() {
    const profileid =  localStorage.getItem("profile")
    const qualificationid = document.querySelector('#qualificationid')?.value;
    const universityid = document.querySelector('#universityid')?.value;
    const departementid = document.querySelector('#departementid')?.value;
    const origindasar = document.querySelector('#origindasar')?.value;
    const originlanjut = document.querySelector('#originlanjut')?.value;
    const major = document.querySelector('#major')?.value;
    const gpa = document.querySelector('#gpa')?.value;
    const yearstart = document.querySelector('#yearstart')?.value;
    const yearend = document.querySelector('#yearend')?.value;
    const payload = {
        "profileid": Number(profileid)?? '',
        "qualificationid": Number(qualificationid)?? '',
        "universityid": Number(universityid)?? '',
        "departementid": Number(departementid)?? '',
        "originschool": qualificationid === "1" || qualificationid === "2" ? origindasar : originlanjut,
        "majorsma": major?? '',
        "gpa": gpa?? '',
        "yearstart": Number(yearstart)?? '',
        "yearend": Number(yearend)?? '',
        }
        fetch(`/profile/profile-wizard/education`, {
            method: `POST`,
            headers: {
                'content-Type': 'application/json'
            },
            body: JSON.stringify(payload)
        })
        .then((response) => response.json())
        .then((value) => {
            if (value.Data == "Ok") {
                const profileid =  localStorage.getItem("profile")
                fetch(`/profile/profile-wizard/education?`+
                    new URLSearchParams({
                    profileid: profileid,
                }))
                .then((response) => response.json())
                .then((value) => {
                const data = value.Data
                const container = document.getElementById("tableEducation")
                let Data = ""
                let no = 1;
                if (data.length > 0) {
                    data.forEach(item => {
                        const qualification = item.Qualification
                        console.log(item)
                        console.log(qualification)
                        if (qualification === "SD" || qualification === "SMP") {
                                Data += `
                                <tr>
                                    <th>${no++}</th>
                                    <th>${item.Qualification}</th>
                                    <th>${item.OriginSchool}</th>
                                    <th>-</th>
                                    <th>${item.Gpa}</th>
                                    <th>${item.YearStart}</th>
                                    <th>${item.YearEnd}</th>
                                </tr>
                                `
                        } else if (qualification === "SMA / SMK") {
                                Data += `
                                <tr>
                                    <th>${no++}</th>
                                    <th>${item.Qualification}</th>
                                    <th>${item.OriginSchool}</th>
                                    <th>${item.MajorSma}</th>
                                    <th>${item.Gpa}</th>
                                    <th>${item.YearStart}</th>
                                    <th>${item.YearEnd}</th>
                                </tr>
                                `
                        } else if (qualification === "D1" || qualification === "D2" || qualification === "D3" || qualification === "D4" || qualification === "S1" || qualification === "S2" || qualification === "S3") {
                                Data += `
                                <tr>
                                    <th>${no++}</th>
                                    <th>${item.Qualification}</th>
                                    <th>${item.University}</th>
                                    <th>${item.Departement}</th>
                                    <th>${item.Gpa}</th>
                                    <th>${item.YearStart}</th>
                                    <th>${item.YearEnd}</th>
                                </tr>
                                `
                        } else {
                            Data = "<tr><td colspan='7'>No data available</td></tr>";
                        }
                    })
                }
                container.innerHTML = Data
                console.log(container)
                })
            }
        })
}

function saveWork() {
    const profileid =  localStorage.getItem("profile")
    const companyid = document.querySelector('#companyid')?.value;
    const countryid = document.querySelector('#countrycompany')?.value;
    const provinceid = document.querySelector('#provincecompany')?.value;
    const positionlevelid = document.querySelector('#positionlevelid')?.value;
    const salary = document.querySelector('#salary')?.value;
    const experiencedesc = document.querySelector('#experiencedesc')?.value;
    const startdate = document.querySelector('#startdate')?.value;
    const startpositionjobtittle = document.querySelector('#startpositionjobtittle')?.value;
    const enddate = document.querySelector('#enddate')?.value;
    const lastpositionjobtittle = document.querySelector('#lastpositionjobtittle')?.value;
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
        "jobtittle": Number(startpositionjobtittle)?? '',
        "enddate": new Date(enddate)?? '',
        "lastpositionjobtittle": Number(lastpositionjobtittle)?? '',
        "reasonleaving": reasonleaving?? '',
        "skillid": Number(skillid)?? '',
    }
    fetch(`/profile/profile-wizard/work`, {
        method: `POST`,
        headers: {
            'content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    })
    .then((response) => response.json())
    .then((value) => {
        if (value.Data == "Ok") {
            const profileid =  localStorage.getItem("profile")
            fetch(`/profile/profile-wizard/work?`+
                new URLSearchParams({
                profileid: profileid,
            }))
            .then((response) => response.json())
            .then((value) => {
                const data = value.Data
                const container = document.getElementById("tableWorkExperience")
                console.log(data)
                let Data = ""
                let no = 1;
                data.forEach(item => {
                    Data += `
                    <tr>
                        <th>${no++}</th>
                        <th>${item.Company}</th>
                        <th>${item.Positionlevel}</th>
                        <th>${item.StartDate}}</th>
                        <th>${item.Startpositionjobtittle}</th>
                    </tr>
                    `
                })
                container.innerHTML = Data
                console.log(container)
            })  
        }
    })       
}

function saveLanguage() {
    const profileid =  localStorage.getItem("profile")
    const languagecode = document.querySelector('#languagecode')?.value;
    const spokenlevel = document.querySelector('#spokenlevel')?.value;
    const writtenlevel = document.querySelector('#writtenlevel')?.value;
    const listeninglevel = document.querySelector('#listeninglevel')?.value;
    const payload = {
        "profileid": Number(profileid)?? '',
        "languagecode": Number(languagecode)?? '',
        "spokenlevel": Number(spokenlevel)?? '',
        "writtenlevel": Number(writtenlevel)?? '',
        "listeninglevel": Number(listeninglevel)?? '',
    }
    fetch(`/profile/profile-wizard/language`, {
        method: `POST`,
        headers: {
            'content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    })
    .then((response) => response.json())
    .then((value) => {
        if (value.Data == "Ok") {
            const profileid =  localStorage.getItem("profile")
            fetch(`/profile/profile-wizard/language?`+
                new URLSearchParams({
                profileid: profileid,
            }))
            .then((response) => response.json())
            .then((value) => {
                const data = value.Data
                const container = document.getElementById("tableLanguage")
                console.log(data)
                let Data = ""
                let no = 1;
                data.forEach(item => {
                    Data += `
                    <tr>
                        <th>${no++}</th>
                        <th>${item.Language}</th>
                        <th>${item.Spokenlevel}</th>
                        <th>${item.Writtenlevel}</th>
                        <th>${item.Listeninglevel}</th>
                    </tr>
                    `
                })
                container.innerHTML = Data
                console.log(container)
            })  
        }
    })       
}

function saveTraining() {
    const profileid = localStorage.getItem("profile");
    const trainingtittle = document.querySelector('#trainingtittle')?.value;
    const vendor = document.querySelector('#vendor')?.value;
    const trainingyear = document.querySelector('#trainingyear')?.value;
    const durationday = document.querySelector('#durationday')?.value;
    const financedByCompany = document.getElementById('financedbyCompany').checked;
    const financedByPersonal = document.getElementById('financedbyPersonal').checked;
    const financedBy = financedByCompany ? 'Company' : (financedByPersonal ? 'Personal' : '');
    const companyname = document.querySelector('#company')?.value;
    const payload = {
        "profileid": Number(profileid) || '',
        "trainingtittle": trainingtittle || '',
        "vendor": vendor || '',
        "trainingyear": Number(trainingyear) || '',
        "durationday": Number(durationday) || '',
        "financedby": financedBy || '',
        "companyname": companyname || ''
    };
    fetch(`/profile/profile-wizard/training`, {
        method: `POST`,
        headers: {
            'content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
    })
    .then((response) => response.json())
    .then((value) => {
        if (value.Data === "Ok") {
            fetch(`/profile/profile-wizard/training?` +
                new URLSearchParams({
                    profileid: profileid,
                }))
            .then((response) => response.json())
            .then((value) => {
                const data = value.Data;
                const container = document.getElementById("tableTraining");
                let Data = "";
                let no = 1;
                data.forEach(item => {
                    Data += `
                    <tr>
                        <th>${no++}</th>
                        <th>${item.TrainingTittle}</th>
                        <th>${item.TrainingYear}</th>
                        <th>${item.DurationDay}</th>
                        <th>${item.Vendor}</th>
                        <th>${item.FinancedBy}</th>
                    </tr>
                    `;
                });
                container.innerHTML = Data;
            });
        }
    });
}