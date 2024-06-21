console.log("M O N I T O R");

document.addEventListener("DOMContentLoaded", function () {

    // const dataElement = document.querySelector('.stats');

    let eventSource;

    const statsDial = document.querySelector('stats-dial');
    
    // Accessing individual attributes
    const value = statsDial.getAttribute('value'); // "50"
    
    function connect() {
        eventSource = new EventSource('http://localhost:8080/stats');

        eventSource.onmessage = function (event) {

            
            let d = JSON.parse(event.data);
            let s = JSON.parse(d.stats_json);

            // console.log(d)

            const memoryUsagePercentage = (s.allocated_memory / s.system_memory) * 100;
            
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
