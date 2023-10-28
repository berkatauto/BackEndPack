package Signup

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type SignupPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var client *mongo.Client

func initMongoDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	return nil
}

func GCFPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload SignupPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Decode Json Payload Failed", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Fail to encrypt Password: %v", err)
		http.Error(w, "Fail to encrypt Password", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"message": "Pendaftaran berhasil"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	collection := client.Database("berkatauto").Collection("userLogin")
	_, err = collection.InsertOne(context.Background(), bson.M{
		"username": payload.Username,
		"password": string(hashedPassword),
	})
	if err != nil {
		log.Printf("Gagal menyimpan data ke MongoDB: %v", err)
		http.Error(w, "Gagal menyimpan data ke MongoDB", http.StatusInternalServerError)
		return
	}
}

func init() {
	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/Signup", GCFPostHandler); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v", err)
	}

	if err := initMongoDB(); err != nil {
		log.Fatalf("Gagal menghubungkan ke MongoDB: %v", err)
	}
}
