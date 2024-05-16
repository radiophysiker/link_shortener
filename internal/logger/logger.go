package logger

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type (
	Log struct {
		*zap.SugaredLogger
	}

	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

func Init() (*Log, error) {
	z, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("logger don't Run! %w", err)
	}
	logger := z.Sugar()
	defer logger.Sync()
	return &Log{logger}, nil
}

func (l *Log) Fatal(format string, args ...any) {
	l.Fatalf(format, args...)
}

func (l *Log) Error(format string, args ...any) {
	l.Errorf(format, args...)
}

func (l *Log) Info(format string, args ...any) {
	l.Infof(format, args...)
}

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	if err != nil {
		return size, fmt.Errorf("cannot write response: %w", err)
	}
	r.responseData.size += size
	return size, nil
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.responseData.status = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (l *Log) CustomMiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}
		next.ServeHTTP(&lw, r)

		l.Infof(
			"URI: %s, Method: %s, Status: %d, Duration: %s, Size: %d",
			r.RequestURI,
			r.Method,
			responseData.status,
			time.Since(start),
			responseData.size,
		)
	})
}
