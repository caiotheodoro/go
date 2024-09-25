package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiotheodoro/product-service/internal/config"
	"google.golang.org/grpc"
	"github.com/caiotheodoro/auth-grpc-microservice/pkg/proto"
)

type Handler struct {
	config config.Config
	authClient proto.UserClient
}

func NewHandler(cfg config.Config) *Handler {
	conn, err := grpc.Dial(cfg.AuthGRPCAddress, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to auth service: %v", err))
	}

	return &Handler{
		config: cfg,
		authClient: proto.NewUserClient(conn),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/products":
		h.handleProducts(w, r)
	case "/health":
		h.handleHealth(w, r)
	case "/order":
		h.handleOrder(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) handleProducts(w http.ResponseWriter, r *http.Request) {
	// Here you might want to add authentication using the auth-grpc service
	products := []string{"Product 1", "Product 2", "Product 3"}
	json.NewEncoder(w).Encode(products)
}

func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "healthy"}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) handleOrder(w http.ResponseWriter, r *http.Request) {
	// Here you would implement logic to create an order
	// This would involve calling the order-redis service
	resp, err := http.Post(h.config.OrderServiceAddress+"/orders", "application/json", r.Body)
	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Forward the response from the order service
	w.WriteHeader(resp.StatusCode)
	json.NewDecoder(resp.Body).Decode(json.NewEncoder(w))
}