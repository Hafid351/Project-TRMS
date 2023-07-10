function generateRandomColors(n) {
    const colors = [];
    for (let i = 0; i < n; i++) {
        const color = "#" + Math.floor(Math.random() * 16777215).toString(16);
        colors.push(color);
    }
    return colors;
}
window.onload = () => {
    fetch("/dashboard/skill").then((response) => response.json()).then((value) => {
        const data = value.Data.slice(0,5)
        const container = document.getElementById("kategorikeahliankandidat")
        let category = ""
        data.forEach(item => {
        category += `<div class="col-6">
        <div class="card">
            <div class="card-body">
                <div class="row">
                    <div class="col mt-0">
                        <h5 class="card-title">${item.Skill}</h5>
                    </div>
                    <div class="col-auto">
                        <!--div class="stat text-primary">
                            <i class="align-middle" data-feather="users"></i>
                        </div-->
                    </div>
                </div>
                <h1 class="mt-1 mb-3">${item.Total}</h1>
            </div>
        </div>
    </div>`
      })
      container.innerHTML = category
      const label = data.map(obj => obj.Skill);
      const values = data.map(obj => obj.Total);
      // Set new default font family and font color to mimic Bootstrap's default styling
      Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
      Chart.defaults.global.defaultFontColor = '#292b2c';
  
      // Pie Chart Example
      var ctx = document.getElementById("mypieChartSkill");
      var mypieChartSkill = new Chart(ctx, {
        type: 'pie',
        data: {
          labels: label,
          datasets: [{
            data: values,
            backgroundColor: generateRandomColors(5),
          }],
        },
      });
    })

    async function getSkillTable(currentPage = 1) {
      const response = await fetch("/dashboard/skill_table?page=" + currentPage);
      const data = await response.json();
    
      const container = document.getElementById("datatableSkill");
      let table = "";
      data.Table.forEach((item) => {
        table += `
          <tr>
              <td>${item.Name}</td>
              <td>${item.Total}</td>
              <td></td>
          </tr>
          `;
      });
      container.innerHTML = table;
      const Name = data.Table.map((obj) => obj.Name);
      const Total = data.Table.map((obj) => obj.Total);
    
      document.getElementById("Table-Skills-Page").innerHTML = currentPage;
      document.getElementById("Table-Skills-Total-Page").innerHTML = data.TotalPages;
    
      const prevButton = document.getElementById("Table-Skills-Prev");
      const nextButton = document.getElementById("Table-Skills-Next");
    
      const PrevPage = () => {
        prevButton.removeEventListener("click", PrevPage);
        nextButton.removeEventListener("click", NextPage);
        getSkillTable(currentPage - 1);
      };
    
      const NextPage = () => {
        prevButton.removeEventListener("click", PrevPage);
        nextButton.removeEventListener("click", NextPage);
        getSkillTable(currentPage + 1);
      };
    
      prevButton.removeEventListener("click", PrevPage);
      nextButton.removeEventListener("click", NextPage);
      prevButton.addEventListener("click", PrevPage);
      nextButton.addEventListener("click", NextPage);
    }
    
    getSkillTable();


    // fetch("/dashboard/departement").then((response) => response.json()).then((value) => {
    //     const tbody = document.querySelector('tbody')
    //     const container = document.getElementById("datatablesSimple")
    //     let category = ""
    //     data.forEach(item => {
    //       category += `<div class="col-6">
    //       <div class="card">
    //           <div class="card-body">
    //               <div class="row">
    //                   <div class="col mt-0">
    //                       <h5 class="card-title">${item.ProfileWorkExperienceSalary}</h5>
    //                   </div>
    //                   <div class="col-auto">
    //                       <!--div class="stat text-primary">
    //                           <i class="align-middle" data-feather="users"></i>
    //                       </div-->
    //                   </div>
    //               </div>
    //               <h1 class="mt-1 mb-3">${item.Total}</h1>
    //           </div>
    //       </div>
    //   </div>`
    //     })
    //     container.innerHTML = category
    //     const label = data.map(obj => obj.ProfileWorkExperienceSalary);
    //     const values = data.map(obj => obj.Total);
    //     console.log(label)
    //     console.log(values)
    
    fetch("/dashboard/departement").then((response) => response.json()).then((value) => {
      const data = value.Data.slice(0,5)
      const container = document.getElementById("kategoriprogramstudi")
      let category = ""
      data.forEach(item => {
        category += `<div class="col-6">
        <div class="card">
            <div class="card-body">
                <div class="row">
                    <div class="col mt-0">
                        <h5 class="card-title">${item.Name}</h5>
                    </div>
                    <div class="col-auto">
                        <!--div class="stat text-primary">
                            <i class="align-middle" data-feather="users"></i>
                        </div-->
                    </div>
                </div>
                <h1 class="mt-1 mb-3">${item.Total}</h1>
            </div>
        </div>
    </div>`
      })
      container.innerHTML = category
      const label = data.map(obj => obj.Name);
      const values = data.map(obj => obj.Total);
      // Set new default font family and font color to mimic Bootstrap's default styling
      Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
      Chart.defaults.global.defaultFontColor = '#292b2c';
  
      // Doughnuts Chart Example
      var ctx = document.getElementById("mydoughnutChartDepartement");
      var mydoughnutChartDepartement = new Chart(ctx, {
        type: 'doughnut',
        data: {
          labels: label,
          datasets: [{
            data: values,
            backgroundColor: generateRandomColors(5),
          }],
        },
      });
    })

    async function getDepartementTable(currentPage = 1) {
      const response = await fetch("/dashboard/departement_table?page=" + currentPage);
      const data = await response.json();
    
      const container = document.getElementById("datatableDepartement");
      let table = "";
      data.Table.forEach((item) => {
        table += `
          <tr>
              <td>${item.Name}</td>
              <td>${item.Total}</td>
              <td></td>
          </tr>
          `;
      });
      container.innerHTML = table;
      const Name = data.Table.map((obj) => obj.Name);
      const Total = data.Table.map((obj) => obj.Total);
    
      document.getElementById("Table-Departements-Page").innerHTML = currentPage;
      document.getElementById("Table-Departements-Total-Page").innerHTML = data.TotalPages;
    
      const prevButton = document.getElementById("Table-Departements-Prev");
      const nextButton = document.getElementById("Table-Departements-Next");
    
      const PrevPage = () => {
        prevButton.removeEventListener("click", PrevPage);
        nextButton.removeEventListener("click", NextPage);
        getDepartementTable(currentPage - 1);
      };
    
      const NextPage = () => {
        prevButton.removeEventListener("click", PrevPage);
        nextButton.removeEventListener("click", NextPage);
        getDepartementTable(currentPage + 1);
      };
    
      prevButton.removeEventListener("click", PrevPage);
      nextButton.removeEventListener("click", NextPage);
      prevButton.addEventListener("click", PrevPage);
      nextButton.addEventListener("click", NextPage);
    }
    
    getDepartementTable();
  
    fetch("/dashboard/position").then((response) => response.json()).then((value) => {
      const data = value.Data.slice(0,5)
      const container = document.getElementById("kategoriposisiharapan")
      let category = ""
      data.forEach(item => {
        category += `<div class="col-6">
        <div class="card">
            <div class="card-body">
                <div class="row">
                    <div class="col mt-0">
                        <h5 class="card-title">${item.Name}</h5>
                    </div>
                    <div class="col-auto">
                        <!--div class="stat text-primary">
                            <i class="align-middle" data-feather="users"></i>
                        </div-->
                    </div>
                </div>
                <h1 class="mt-1 mb-3">${item.Total}</h1>
            </div>
        </div>
    </div>`
      })
      container.innerHTML = category
      const label = data.map(obj => obj.Name);
      const values = data.map(obj => obj.Total);
      // Set new default font family and font color to mimic Bootstrap's default styling
      Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
      Chart.defaults.global.defaultFontColor = '#292b2c';
  
      // Bar Chart Example
      var ctx = document.getElementById("myBarChartPosition");
      var myBarChartJobPositionLevel = new Chart(ctx, {
        type: 'bar',
        data: {
          labels: label,
          datasets: [{
            label: "Revenue",
            backgroundColor: generateRandomColors(5),
            borderColor: generateRandomColors(5),
            data: values,
          }],
        },
        options: {
          scales: {
            xAxes: [{
              time: {
                unit: 'position'
              },
              gridLines: {
                display: false
              },
              ticks: {
                maxTicksLimit: 6
              }
            }],
            yAxes: [{
              ticks: {
                min: 0,
                max: 20,
                maxTicksLimit: 5
              },
              gridLines: {
                display: true
              }
            }],
          },
          legend: {
            display: false
          }
        }
      });
    })

    // fetch("/dashboard/jobpositionlevel_table").then((response) => response.json()).then((value) => {
    //   const data = value.Table
    //   const container = document.getElementById("datatableJobPositionLevel")
    //   let table = ""
    //   data.forEach(item => {
    //     table += `
    //     <tr>
    //         <td>${item.Name}</td>
    //         <td>${item.Total}</td>
    //         <td></td>
    //     </tr>
    //     `
    //   })
    //   container.innerHTML = table
    //   const Name = data.map(obj => obj.Name);
    //   const Total = data.map(obj => obj.Total);
    //   console.log(Name)
    //   console.log(Total)
    // })

    async function getPositionTable(currentPage = 1) {
      const response = await fetch("/dashboard/position_table?page=" + currentPage);
      const data = await response.json();
    
      const container = document.getElementById("datatablePosition");
      let table = "";
      data.Table.forEach((item) => {
        table += `
          <tr>
              <td>${item.Name}</td>
              <td>${item.Total}</td>
              <td></td>
          </tr>
          `;
      });
      container.innerHTML = table;
      const Name = data.Table.map((obj) => obj.Name);
      const Total = data.Table.map((obj) => obj.Total);
      console.log(Name);
      console.log(Total);
    
      document.getElementById("Table-Position-Page").innerHTML = currentPage;
      document.getElementById("Table-Position-Total-Page").innerHTML = data.TotalPages;
    
      const prevButton = document.getElementById("Table-Position-Prev");
      const nextButton = document.getElementById("Table-Position-Next");
    
      const PrevPage = () => {
        prevButton.removeEventListener("click", PrevPage);
        nextButton.removeEventListener("click", NextPage);
        getPositionTable(currentPage - 1);
      };
    
      const NextPage = () => {
        prevButton.removeEventListener("click", PrevPage);
        nextButton.removeEventListener("click", NextPage);
        getPositionTable(currentPage + 1);
      };
    
      prevButton.removeEventListener("click", PrevPage);
      nextButton.removeEventListener("click", NextPage);
      prevButton.addEventListener("click", PrevPage);
      nextButton.addEventListener("click", NextPage);
    }
    
    getPositionTable();

  
    fetch("/dashboard/profileworkexperience").then((response) => response.json()).then((value) => {
      const data = value.Data.slice(0,5)
      const container = document.getElementById("kategorigajiharapan")
      let category = ""
      data.forEach(item => {
        category += `<div class="col-6">
        <div class="card">
            <div class="card-body">
                <div class="row">
                    <div class="col mt-0">
                        <h5 class="card-title">${item.ProfileWorkExperienceSalary}</h5>
                    </div>
                    <div class="col-auto">
                        <!--div class="stat text-primary">
                            <i class="align-middle" data-feather="users"></i>
                        </div-->
                    </div>
                </div>
                <h1 class="mt-1 mb-3">${item.Total}</h1>
            </div>
        </div>
    </div>`
      })
      container.innerHTML = category
      const label = data.map(obj => obj.ProfileWorkExperienceSalary);
      const values = data.map(obj => obj.Total);
    //   // Set new default font family and font color to mimic Bootstrap's default styling
    //   Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
    //   Chart.defaults.global.defaultFontColor = '#292b2c';
      
    //   console.log(data.map((item) => ({
    //     label: item.Name,
    //     backgroundColor: "#" + Math.floor(Math.random() * 16777215).toString(16),
    //     borderColor: "#" + Math.floor(Math.random() * 16777215).toString(16),
    //     data: item.ProfileWorkExperienceSalary,
    //   })))

    //   // Bar Chart Example
    //   var ctx = document.getElementById("myBarChartProfileWorkExperience");
    //   var myBarChartProfileWorkExperience = new Chart(ctx, {
    //     type: 'bar',
    //     data: {
    //       labels: label,
    //       datasets: data.map((item) => ({
    //         label: item.Name,
    //         backgroundColor: "#" + Math.floor(Math.random() * 16777215).toString(16),
    //         borderColor: "#" + Math.floor(Math.random() * 16777215).toString(16),
    //         data: item.ProfileWorkExperienceSalary,
    //       })),
    //     },
    //     options: {
    //       plugins: {
    //         legend: {position:"left"}
    //       },
    //       scales: {
    //         xAxes: [{
    //           ticks: {
    //             min: 0,
    //             max: 500,
    //             maxTicksLimit: 5
    //           },
    //           gridLines: {
    //             display: true
    //           }
    //         }],
    //         yAxes: [{
    //           // time: {
    //           //   unit: 'salary'
    //           // },
    //           gridLines: {
    //             display: false
    //           },
    //           ticks: {
    //             maxTicksLimit: 6
    //           }
    //         }],
    //       },
    //       legend: {
    //         display: true
    //       }
    //     }
    //   });
    // })
      console.log(data.map((item) => ({
        label: item.Name,
        backgroundColor: "#" + Math.floor(Math.random() * 16777215).toString(16),
        borderColor: "#" + Math.floor(Math.random() * 16777215).toString(16),
        data: item.ProfileWorkExperienceSalary,
      })))
    // Set new default font family and font color to mimic Bootstrap's default styling
    Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
    Chart.defaults.global.defaultFontColor = '#292b2c';

    // Bar Chart Example
    var ctx = document.getElementById("myBarChartProfileWorkExperience");
    var myBarChartProfileWorkExperience = new Chart(ctx, {
      type: 'bar',
      data: {
        labels: label,
        datasets: [{
          label: "Revenue",
          backgroundColor: generateRandomColors(5),
          borderColor: generateRandomColors(5),
          data: values,
        }],
      },
      options: {
        scales: {
          xAxes: [{
            ticks: {
              min: 0,
              max: 500,
              maxTicksLimit: 5
            },
            gridLines: {
              display: true
            }
          }],
          yAxes: [{
            time: {
              unit: 'salary'
            },
            gridLines: {
              display: false
            },
            ticks: {
              maxTicksLimit: 6
            }
          }],
        },
        legend: {
          display: false
        }
      }
    });
  })
}

    // fetch("/dashboard/profileworkexperience_table").then((response) => response.json()).then((value) => {
    //   const data = value.Table
    //   const container = document.getElementById("datatableProfileWorkExperience")
    //   let table = ""
    //   data.forEach(item => {
    //     table += `
    //     <tr>
    //         <td>${item.Name}</td>
    //         <td>${item.Salary}</td>
    //         <td>${item.Total}</td>
    //     </tr>
    //     `
    //   })
    //   container.innerHTML = table
    //   const Name = data.map(obj => obj.Name);
    //   const Salary = data.map(obj => obj.Salary)
    //   const Total = data.map(obj => obj.Total);
    //   console.log(Name)
    //   console.log(Salary)
    //   console.log(Total)
    // })

    async function getProfileWorkExperienceTable(currentPage = 1) {
      const response = await fetch("/dashboard/profileworkexperience_table?page=" + currentPage);
      const data = await response.json();
    
      const container = document.getElementById("datatableProfileWorkExperience");
      let table = "";
      data.Table.forEach((item) => {
        table += `
          <tr>
              <td>${item.Name}</td>
              <td>${item.Salary}</td>
              <td>${item.Total}</td>
              <td></td>
          </tr>
          `;
      });
      container.innerHTML = table;
      const Name = data.Table.map((obj) => obj.Name);
      const Salary = data.Table.map((obj) => obj.Salary);
      const Total = data.Table.map((obj) => obj.Total);
    
      document.getElementById("Table-ProfileWorkExperience-Page").innerHTML = currentPage;
      document.getElementById("Table-ProfileWorkExperience-Total-Page").innerHTML = data.TotalPages;
    
      const prevButton = document.getElementById("Table-ProfileWorkExperience-Prev");
      const nextButton = document.getElementById("Table-ProfileWorkExperience-Next");
    
      const PrevPage = () => {
        prevButton.removeEventListener("click", PrevPage);
        nextButton.removeEventListener("click", NextPage);
        getProfileWorkExperienceTable(currentPage - 1);
      };
    
      const NextPage = () => {
        prevButton.removeEventListener("click", PrevPage);
        nextButton.removeEventListener("click", NextPage);
        getProfileWorkExperienceTable(currentPage + 1);
      };
    
      prevButton.removeEventListener("click", PrevPage);
      nextButton.removeEventListener("click", NextPage);
      prevButton.addEventListener("click", PrevPage);
      nextButton.addEventListener("click", NextPage);
    }
    
    getProfileWorkExperienceTable();