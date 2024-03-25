package middleware

import (
	"context"
	"healthstats/pkg/service"
	"net/http"
)

func TransactionMiddleware(svc *service.Service, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Start a new DB transaction
		tx, err := svc.DB.BeginTx(r.Context(), nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Add the transaction to the request context
		ctx := context.WithValue(r.Context(), "tx", tx)

		// Call the next handler with the updated request
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
