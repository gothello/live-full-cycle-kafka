package web

import (
	"encoding/json"
	"log"
	"net/http"

	"gituhb.com/gothello/live-full-cycle-kafka/internal/usecase"
)

type ProductHandlers struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductUseCase   *usecase.ListProductUseCase
}

func NewProductHandlers(cpu *usecase.CreateProductUseCase, lpu *usecase.ListProductUseCase) *ProductHandlers {
	return &ProductHandlers{
		cpu,
		lpu,
	}
}

func (p *ProductHandlers) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var input usecase.CreateProductInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	prc, err := p.CreateProductUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(prc)
}

func (p *ProductHandlers) ListAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	all, err := p.ListProductUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(all)
}
