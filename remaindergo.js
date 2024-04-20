var map = L.map('map', {
    maxBounds: [
        [90, -180], // North West
        [-90, 180] // South East
    ]
}).setView([12.861591, 77.664556], 13);

L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);

var marker = L.marker([12.861591, 77.664556], {
    draggable: true
}).addTo(map);

marker.on('dragend', function(event) {
    var marker = event.target;
    var position = marker.getLatLng();
    $('#latitude').val(position.lat.toFixed(6));
    $('#longitude').val(position.lng.toFixed(6));
});

function onMapClick(e) {
    marker.setLatLng(e.latlng);
    $('#latitude').val(e.latlng.lat.toFixed(6));
    $('#longitude').val(e.latlng.lng.toFixed(6));
}

map.on('click', onMapClick);

$('#submitBtn').click(function() {
    var title = $('#title').val();
    var description = $('#description').val();
    var latitude = $('#latitude').val();
    var longitude = $('#longitude').val();

    // Check if either title or description is empty
    if (title.trim() === '' || description.trim() === '') {
        alert('Please enter a reminder first');
        return; // Exit the function early if reminder is empty
    }

    var reminderData = {
        title: title,
        description: description,
        latitude: parseFloat(latitude),
        longitude: parseFloat(longitude)
    };

    $.ajax({
        url: 'http://localhost:8080/reminders',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(reminderData),
        success: function(response) {
            console.log('Reminder set successfully');
            $('#title').val('');
            $('#description').val('');
            sendCurrentLocation(); // Call sendCurrentLocation after successful submission
            alert('Reminder updated successfully!');
        },
        error: function(error) {
            console.log('Error setting reminder:', error);
            alert('Error updating reminder: ' + error);
        }
    });
});

// Add event listener to the current location button
$('#currentLocationBtn').click(goToCurrentLocation);

// Function to go to the user's current location on the map
function goToCurrentLocation() {
    // Check if geolocation is supported by the browser
    if (navigator.geolocation) {
        // Get the user's current position
        navigator.geolocation.getCurrentPosition(function(position) {
            var latitude = position.coords.latitude;
            var longitude = position.coords.longitude;

            // Set the map view to the user's current location
            map.setView([latitude, longitude], 13);

            // Update the marker position
            marker.setLatLng([latitude, longitude]);

            // Update the latitude and longitude inputs
            $('#latitude').val(latitude.toFixed(6));
            $('#longitude').val(longitude.toFixed(6));

            console.log('Moved map to current location - Latitude:', latitude, 'Longitude:', longitude);
        }, function(error) {
            console.log('Error getting current location:', error);
            alert('Error getting current location. Please try again.');
        });
    } else {
        console.log('Geolocation is not supported by this browser.');
        alert('Geolocation is not supported by this browser.');
    }
}


// Get user's current location and send to Go backend server
function sendCurrentLocation() {
    navigator.geolocation.getCurrentPosition(function(position) {
        var latitude = position.coords.latitude;
        var longitude = position.coords.longitude;

        console.log('Sending location - Latitude:', latitude, 'Longitude:', longitude);

        $.ajax({
            url: 'http://localhost:8080/location',
            type: 'POST',
            data: JSON.stringify({
                latitude: latitude,
                longitude: longitude
            }),
            contentType: 'application/json', // Add this line to specify JSON content type
            success: function(response) {
                console.log('Location sent successfully');
            },
            error: function(xhr, status, error) {
                console.log('Error sending location:', error);
            }
        });
    }, function(error) {
        console.log('Error getting location:', error);
    });
}
// Send user's location to the server every 10 seconds
setInterval(sendCurrentLocation, 10000);