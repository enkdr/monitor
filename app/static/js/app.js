import StatsTemplate from './stats-template.js';
console.log("M O N I T O R");

function storeStats(statsString) {
    window.localStorage.setItem("storedStats",statsString);
}

function updateInfo() {

    console.log("calling updateInfo()");
    
    const jsonData = JSON.parse(window.localStorage.getItem("storedStats"));
    const statsTemplate = new StatsTemplate(jsonData);
    const statsInfoTemplate = statsTemplate.generateAllTemplates();

    const container = document.querySelector('.stats-info');
    container.innerHTML = "";
    container.insertAdjacentHTML("afterbegin",statsInfoTemplate);

}


function updateDial(data,statsTypes) {
    
    statsTypes.forEach((d) => {
        
        const dial = document.querySelector(`.${d}_dial`);
        const stats = JSON.parse(data[d]);
        const value = dial.getAttribute('value');

        if (d === "cpu_stats") {

            let cpuStatsInfo = document.querySelector(".cpu_stats_info");
            const memoryUsagePercentage = (stats.stats_json.allocated_memory / stats.stats_json.system_memory) * 100;
            
            dial.setAttribute('value', parseInt(memoryUsagePercentage));
        }        
    });        
}


document.addEventListener("DOMContentLoaded", function () {

    let eventSource;
    
    function connect() {
        eventSource = new EventSource('http://localhost:8080/stats');

        eventSource.onmessage = function (event) {

            const data = JSON.parse(event.data);
            const statsTypes = Object.keys(data);

            updateDial(data,statsTypes);
            // localstorage uses string
            storeStats(event.data);
            updateInfo();
        };

        eventSource.onerror = function () {
            console.log('Error occurred, reconnecting...');
            eventSource.close();
            setTimeout(connect, 2000); // attempt to reconnect after 2 seconds
        };

        eventSource.onopen = function () {
            console.log('Connected');

        };
    }

    connect(); // initial connection

});
