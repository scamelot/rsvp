package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RSVP struct {
	Name      string `json:"name"`
	Attending bool   `json:"attending"`
}

func init() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
}

func connectDB() (*mongo.Client, context.Context, context.CancelFunc) {
	connectionString := os.Getenv("MONGODB_URI")
	if connectionString == "" {
		log.Fatal("MONGODB_CONNECTION_STRING must be set in the .env file")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, cancel
}

func handleRSVP(rsvpCollection *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rsvp RSVP

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		err := json.NewDecoder(r.Body).Decode(&rsvp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Save the RSVP data to MongoDB
		_, err = rsvpCollection.InsertOne(ctx, rsvp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"status":  "success",
			"message": "RSVP submitted",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func main() {

	client, ctx, cancel := connectDB()
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	rsvpCollection := client.Database("rsvp").Collection("responses")

	r := chi.NewRouter()

	// Enable CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsMiddleware.Handler)
	r.Route("/api", func(r chi.Router) {
		r.Post("/rsvp", handleRSVP(rsvpCollection))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s", port)
	http.ListenAndServe(":"+port, r)

}
