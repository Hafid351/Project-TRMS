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
        }
        )
}