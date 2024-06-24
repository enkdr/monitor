console.log("M O N I T O R");


function updateDial(e) {
    
    const data = JSON.parse(e.data);
    const statsTypes = Object.keys(data);
    
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
            updateDial(event);
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
