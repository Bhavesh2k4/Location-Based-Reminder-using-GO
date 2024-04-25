package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"net/smtp"

	"github.com/rs/cors"
)

// "gopkg.in/gomail.v2"
type Reminder struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	EmailSent   bool    `json:"emailSent"`
}

type UserLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var reminders []Reminder

func main() {
	handler := cors.Default().Handler(http.DefaultServeMux)

	http.HandleFunc("/reminders", handleReminders)
	http.HandleFunc("/location", handleUserLocation)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func handleReminders(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var reminder Reminder
		err := json.NewDecoder(r.Body).Decode(&reminder)
		if err != nil {
			log.Println("Error decoding reminder JSON:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		reminder.EmailSent = false // Initialize EmailSent field to false
		reminders = append(reminders, reminder)
		log.Println("Reminder added:", reminder)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println("Invalid request method")
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleUserLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userLocation UserLocation
		err := json.NewDecoder(r.Body).Decode(&userLocation)
		if err != nil {
			log.Println("Error decoding user location JSON:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Received user location:", userLocation)
		checkAndSendReminders(userLocation.Latitude, userLocation.Longitude)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println("Invalid request method")
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

var remindersMap = make(map[string]bool) // Map to keep track of reminders that have been sent

func checkAndSendReminders(latitude, longitude float64) {
	for _, reminder := range reminders {
		distance := calculateDistance(latitude, longitude, reminder.Latitude, reminder.Longitude)
		log.Printf("Distance from reminder '%s': %f meters", reminder.Title, distance)
		if distance <= 100 && !reminder.EmailSent && !remindersMap[reminder.Title] {
			sendReminderEmail(reminder.Title, reminder.Description)
			reminder.EmailSent = true
			remindersMap[reminder.Title] = true // Mark reminder as sent in the map
			log.Printf("Email sent for reminder '%s'", reminder.Title)
		}
	}
}




func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	// Convert degrees to radians
	lat1 = lat1 * math.Pi / 180
	lon1 = lon1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180
	lon2 = lon2 * math.Pi / 180

	// Haversine formula
	dlon := lon2 - lon1
	dlat := lat2 - lat1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	r := 6371.0 // Earth's radius in kilometers

	return c * r * 1000 // Distance in meters
}

func sendReminderEmail(title, description string) error {
	from := "gagansl62004@gmail.com"
	password := ""
	to := "primal.music.6@gmail.com"
	host := "smtp.gmail.com"
	port := "587"
	subject := "Reminder: " + title // Set the subject of the email
	content := "Description: " + description
	body := []byte("Subject: " + subject + "\r\n\r\n" + content) // Include subject in the email body

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, body)
	if err != nil {
		log.Printf("Failed to send reminder email: %v", err)
		return err
	}

	log.Printf("Reminder email sent: %s", title)
	return nil
}

