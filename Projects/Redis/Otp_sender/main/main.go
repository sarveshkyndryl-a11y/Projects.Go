package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

/* ---------- REQUEST BODY ---------- */
type OTPRequest struct {
	ToEmail string `json:"to_email"`
}

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	http.HandleFunc("/send-otp", sendOTPHandler)

	log.Println("OTP service running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

/* ---------- HANDLER ---------- */
func sendOTPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.ToEmail == "" {
		http.Error(w, "to_email is required", http.StatusBadRequest)
		return
	}

	otp := generateOTP()

	key := "otp:" + req.ToEmail
	err := rdb.Set(ctx, key, otp, 5*time.Minute).Err()
	if err != nil {
		http.Error(w, "Redis error", http.StatusInternalServerError)
		return
	}

	if err := sendOTPEmail(req.ToEmail, otp); err != nil {
		log.Println("Email error:", err)
		http.Error(w, "Failed to send OTP", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OTP sent successfully"))
}

/* ---------- OTP ---------- */
func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(900000)+100000)
}

/* ---------- TEMPLATE LOADER ---------- */
func loadOTPTemplate(otp string) (string, error) {
	data, err := os.ReadFile("templates/otp.html")
	if err != nil {
		return "", err
	}
	html := strings.ReplaceAll(string(data), "{{OTP}}", otp)
	return html, nil
}

/* ---------- SEND EMAIL ---------- */
func sendOTPEmail(toEmail, otp string) error {
	from := mail.NewEmail("OTP Service", os.Getenv("FROM_EMAIL"))
	to := mail.NewEmail("User", toEmail)

	htmlContent, err := loadOTPTemplate(otp)
	if err != nil {
		return err
	}

	message := mail.NewSingleEmail(
		from,
		"Your OTP Code",
		to,
		"Your OTP is "+otp,
		htmlContent,
	)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	resp, err := client.Send(message)

	if err != nil {
		return err
	}

	log.Println("SendGrid status:", resp.StatusCode)
	return nil
}
