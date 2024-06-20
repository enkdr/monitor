console.log("M O N I T O R");

document.addEventListener("DOMContentLoaded", function () {
    const dataElement = document.querySelector('.stats');
    let eventSource;

    function connect() {
        eventSource = new EventSource('http://localhost:8080/stats');

        eventSource.onmessage = function (event) {
            // dataElement.innerHTML += JSON.stringify(event.data, null, 2) + '<br/>';
            dataElement.innerHTML += event.data + '<br/>';
        };

        eventSource.onerror = function () {
            console.log('Error occurred, reconnecting...');
            eventSource.close();
            setTimeout(connect, 2000); // Attempt to reconnect after 2 seconds
        };

        eventSource.onopen = function () {
            console.log('Connected');
            
        };
    }

    connect(); // Initial connection

});
