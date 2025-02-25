package redirect

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

type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.redirect.New"

		log = log.With(
			slog.String("op:%s", op),
			slog.String("reques_id:%s", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, resp.Error("alias is empty:"+alias))

			return
		}

		URL, err := urlGetter.GetURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)

			render.JSON(w, r, resp.Error("url not found"))

			return
		}
		if err != nil {
			log.Error("failed to get url from DB", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get url from DB"))

			return
		}

		log.Info("url redirected", slog.String("url", URL))

		http.Redirect(w, r, URL, http.StatusFound)

	}
}
