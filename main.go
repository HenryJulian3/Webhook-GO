package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/HenryJulian3/MundoGoWebhook/docs" // Import para Swagger
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// HelloWorldResponse define la estructura de respuesta JSON
type HelloWorldResponse struct {
	Message string `json:"message"`
}

// WebhookHandler maneja las solicitudes HTTP POST
// @Summary Webhook Hola Mundo
// @Description Este webhook responde con un mensaje de "Hola Mundo"
// @Accept json
// @Produce json
// @Success 200 {object} HelloWorldResponse
// @Router /webhook [post]
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloWorldResponse{
		Message: "Hola Mundo",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Crear enrutador
	router := mux.NewRouter()

	// Ruta del Webhook
	router.HandleFunc("/webhook", WebhookHandler).Methods("POST")

	// Agregar ruta de Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Iniciar servidor
	http.ListenAndServe(":8080", router)
}
