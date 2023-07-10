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
    const label = data.map(obj => obj.Skill);
    const values = data.map(obj => obj.Total);
    console.log(label)
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
}
