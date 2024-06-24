console.log("M O N I T O R");

document.addEventListener("DOMContentLoaded", function () {

    // const dataElement = document.querySelector('.stats');

    let eventSource;

    const statsDial = document.querySelector('stats-dial');
    
    // Accessing individual attributes
    const value = statsDial.getAttribute('value');
    
    function connect() {
        eventSource = new EventSource('http://localhost:8080/stats');

        eventSource.onmessage = function (event) {
            
            let data = JSON.parse(event.data);
            let cpuStats = JSON.parse(data.cpu_stats);
            let fsStats = JSON.parse(data.fs_stats);
            let psStats = JSON.parse(data.process_stats);

            // console.log(cpuStats);
            // console.log(fsStats);
            // console.log(psStats);

            const memoryUsagePercentage = (cpuStats.stats_json.allocated_memory / cpuStats.stats_json.system_memory) * 100;
            
            statsDial.setAttribute('value', parseInt(memoryUsagePercentage));

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
