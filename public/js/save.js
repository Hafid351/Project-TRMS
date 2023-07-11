function saveEducation() {
    const profileid =  localStorage.getItem("profile")
    const qualificationid = document.querySelector('#qualificationid')?.value;
    const universityid = document.querySelector('#universityid')?.value;
    const departementid = document.querySelector('#departementid')?.value;
    const origin = document.querySelector('#origin')?.value;
    const major = document.querySelector('#major')?.value;
    const gpa = document.querySelector('#gpa')?.value;
    const yearstart = document.querySelector('#yearstart')?.value;
    const yearend = document.querySelector('#yearend')?.value;
    const payload = {
        "profileid": Number(profileid)?? '',
        "qualificationid": Number(qualificationid)?? '',
        "universityid": Number(universityid)?? '',
        "departementid": Number(departementid)?? '',
        "originschool": origin?? '',
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
                console.log(data)
                let Data = ""
                let no = 1;
                data.forEach(item => {
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
                })
                container.innerHTML = Data
                console.log(container)
                })
            }
        }
        )
}