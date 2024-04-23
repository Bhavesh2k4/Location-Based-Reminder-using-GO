# Location-Based Reminder

This is a web application that allows users to set reminders based on their location. When the user gets within 100 meters of the specified location, an email reminder is sent to them.

## Features

- Set reminders with a title, description, and location (latitude and longitude)
- Drag and drop a marker on the map to set the reminder location
- Use the "Current Location" button to set the reminder location to the user's current location
- Email reminders are sent when the user gets within 100 meters of the set location

## Technologies Used

- **Frontend**: HTML, CSS, JavaScript, Leaflet.js (map library), jQuery
- **Backend**: Go

## Getting Started

1. Clone the repository:
2. Start the Go backend server
3. Open the `remindergo.html` file in your web browser to access the application.
Note: Make sure to update the `sendReminderEmail` function in `remindergo.go` with your email credentials and the desired recipient email address.

## Usage

1. On the web page, enter a title and description for your reminder.
2. Set the reminder location by either:
- Dragging and dropping the marker on the map
- Clicking the "Current Location" button to use your device's current location
3. Click the "Set Reminder" button to save the reminder.
4. When you get within 100 meters of the set location, an email reminder will be sent to the specified email address.

## Customization

You can customize the following aspects of the application:

- Email sending configuration (SMTP server, credentials, etc.) in `remindergo.go`
- Map styles and settings in `remindergo.html` and `remaindergo.js`
- Frontend styling in `remaindergo.css`

## License

This project is licensed under the [MIT License](LICENSE).

## Images

![image](https://github.com/Bhavesh2k4/Location-Based-Reminder-using-GO/assets/122691714/908ad68c-d5d8-4a6b-a242-0821cedd5745)
![image](https://github.com/Bhavesh2k4/Location-Based-Reminder-using-GO/assets/122691714/a61c6a7f-291c-4f50-b40b-244264ed33d8)
![image](https://github.com/Bhavesh2k4/Location-Based-Reminder-using-GO/assets/122691714/90cf2a6e-167d-4ac7-98d8-a885f037c638)
![image](https://github.com/Bhavesh2k4/Location-Based-Reminder-using-GO/assets/122691714/9d2d225a-6d41-474f-bcff-8e7e76eb40fe)
![image](https://github.com/Bhavesh2k4/Location-Based-Reminder-using-GO/assets/122691714/da5d53e3-bdaa-4333-864a-3bded8098f44)
![image](https://github.com/Bhavesh2k4/Location-Based-Reminder-using-GO/assets/122691714/e88ffeb5-5f08-42e6-80bb-ac2b5b632b3c)






