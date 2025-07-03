package server

import (
	"net/http"
	"time"

	"erp-opity/pkg/logger"
)

// LoggingMiddleware registra informações sobre as requisições HTTP
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Criar response writer customizado para capturar status code
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Executar handler
		next.ServeHTTP(wrapped, r)

		// Log da requisição
		duration := time.Since(start)
		logger := logger.New()
		logger.WithFields(map[string]interface{}{
			"method":     r.Method,
			"path":       r.URL.Path,
			"status":     wrapped.statusCode,
			"duration":   duration.String(),
			"user_agent": r.UserAgent(),
			"remote_ip":  r.RemoteAddr,
		}).Info("HTTP Request")
	})
}

// CORSMiddleware configura CORS para requisições cross-origin
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Configurar headers CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Responder imediatamente para requisições OPTIONS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// AuthMiddleware verifica autenticação do usuário
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Rotas públicas que não precisam de autenticação
		publicPaths := map[string]bool{
			"/api/v1/health":        true,
			"/api/v1/auth/login":    true,
			"/api/v1/auth/register": true,
		}

		// Verificar se é uma rota pública
		if publicPaths[r.URL.Path] {
			next.ServeHTTP(w, r)
			return
		}

		// Extrair token do header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token de autorização não fornecido", http.StatusUnauthorized)
			return
		}

		// Validar token (implementação simplificada)
		if !isValidToken(authHeader) {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// RateLimitMiddleware limita a taxa de requisições
func RateLimitMiddleware(next http.Handler) http.Handler {
	// Implementação simplificada - em produção usar Redis ou similar
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implementar rate limiting
		next.ServeHTTP(w, r)
	})
}

// SecurityMiddleware adiciona headers de segurança
func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Headers de segurança
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		next.ServeHTTP(w, r)
	})
}

// responseWriter wrapper para capturar status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// isValidToken valida o token de autenticação (implementação simplificada)
func isValidToken(token string) bool {
	// TODO: Implementar validação real do JWT
	return len(token) > 0
}
