package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type Logger interface {
	Error(format string, args ...any)
}
type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

var supportedContentTypes = []string{"application/json", "text/html"}

func CustomCompression(next http.Handler, log Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the request is compressed, then we decompress it.
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				log.Error("cannot create gzip reader: %v", err)
				return
			}
			defer func(gz *gzip.Reader) {
				err := gz.Close()
				if err != nil {
					log.Error("cannot close gzip reader: %v", err)
				}
			}(gz)
			r.Body = gz
		}

		contentType := r.Header.Get("Content-Type")
		isSupportedContentType := checkSupportedContentType(contentType, supportedContentTypes)
		// If the content type is not supported, then we do not compress the response.
		if !isSupportedContentType {
			next.ServeHTTP(w, r)
			return
		}

		// If the client supports gzip, then we compress the response.
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
			if err != nil {
				log.Error("cannot create gzip writer: %v", err)
				return
			}
			defer func(gz *gzip.Writer) {
				err := gz.Close()
				if err != nil {
					log.Error("cannot close gzip writer: %v", err)
				}
			}(gz)
			next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
		}

		next.ServeHTTP(w, r)
	})
}

func checkSupportedContentType(contentType string, supportedContentTypes []string) bool {
	for _, supportedContentType := range supportedContentTypes {
		if contentType == supportedContentType {
			return true
		}
	}
	return false
}
