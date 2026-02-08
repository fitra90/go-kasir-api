package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"testgo/services"
	"time"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

// HandleReport -- GET /api/report
func (h *ReportHandler) HandleReports(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByDate(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReportHandler) GetByDate(w http.ResponseWriter, r *http.Request) {
	// Simple approach: expect exact string match or format provided by client
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	// Fallback if client sends only "date"
	if startDate == "" {
		startDate = r.URL.Query().Get("date")
	}
	if endDate == "" {
		endDate = startDate
	}

	report, err := h.service.GetReport(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func (h *ReportHandler) HandleReportCustom(w http.ResponseWriter, r *http.Request) {
	// Check if the path is "/api/report/today"
	path := strings.TrimPrefix(r.URL.Path, "/api/report/")

	if path == "today" || path == "hari-ini" {
		today := time.Now().Format("2006-01-02")
		report, err := h.service.GetReport(today, today)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(report)
		return
	}

	// Fallback for other paths or plain calls
	h.HandleReports(w, r)
}
