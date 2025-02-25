package delete

import (
	"errors"
	"log/slog"
	"net/http"
	resp "url-shortner/internal/lib/api/response"
	"url-shortner/internal/lib/logger/sl"
	"url-shortner/internal/storage"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type URLDeleter interface {
	DeleteURL(alias string) error
}

type Response struct {
	resp.Response
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log = log.With(
			slog.String("op:", op),
			slog.String("reques_id:", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")

		err := urlDeleter.DeleteURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url with alias %s not found", "alias", alias)

			render.JSON(w, r, resp.Error("url not found"))

			return
		}

		if err != nil {
			log.Error("failed to delete url", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to delete url"))

			return
		}

		log.Info("url deleted:")

		render.JSON(w, r, Response{
			resp.OK(),
		})

	}
}
