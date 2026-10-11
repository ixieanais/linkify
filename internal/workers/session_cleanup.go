// Package workers
package workers

import (
	"context"
	"log"
	"time"

	"linkify/internal/services"
)

type SessionCleanupWorker struct {
	service  *services.SessionService
	interval time.Duration
}

func NewSessionCleanupWorker(
	service *services.SessionService,
	interval time.Duration,
) *SessionCleanupWorker {
	return &SessionCleanupWorker{
		service:  service,
		interval: interval,
	}
}

func (w *SessionCleanupWorker) Run(c context.Context) {
	ticker := time.NewTicker(w.interval)
	defer ticker.Stop()

	for {
		select {
		case <-c.Done():
			return
		case <-ticker.C:
			if err := w.Cleanup(c); err != nil {
				log.Printf("token cleanup error: %v", err)
			}
		}
	}
}

func (w *SessionCleanupWorker) Cleanup(c context.Context) error {
	return w.service.DeleteByExpiresAt(c)
}
