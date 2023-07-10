// Set new default font family and font color to mimic Bootstrap's default styling
Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
Chart.defaults.global.defaultFontColor = '#292b2c';

//pie chart example
var ctx = document.getElementById("mypolarAreaRecruiter");
var myPieChart1 = new Chart(ctx, {
    type: 'doughnut',
    data: {
        labels: ["Laki - laki","Perempuan"],
        datasets: [{
            data: [30, 45],
            backgroundColor: ['#007bff', '#dc3545'],
        }],
    },
});