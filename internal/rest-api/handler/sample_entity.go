package handler

import (
	"application/internal/biz/sample_entity"
	"application/internal/rest-api/dto"
	"application/internal/rest-api/response"
	"application/pkg/middlewares"
	"application/pkg/middlewares/httplogger"
	"application/pkg/middlewares/httprecovery"
	"application/pkg/utils"
	"log/slog"
	"net/http"
)

var _ HandlerInterface = (*SampleEntityHandler)(nil)

type SampleEntityHandler struct {
	sampleEntityBiz sample_entity.SampleEntity
	logger          *slog.Logger
}

func NewSampleEntityHandler(logger *slog.Logger, sampleEntityBiz sample_entity.SampleEntity) *SampleEntityHandler {
	return &SampleEntityHandler{
		logger: logger.With("layer", "MuxSampleEntityService"), sampleEntityBiz: sampleEntityBiz,
	}
}

func (s *SampleEntityHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := s.logger.With("method", "Create", "ctx", utils.GetLoggerContext(ctx))
	logger.Debug("Create called with ctx")

	var request dto.SampleEntityRequest
	if err := request.FromRequest(r); err != nil {
		response.ResponseBadRequest(w, "invalid request")
		logger.DebugContext(ctx, "SampleEntityHandler.", "url", r.Host, "status", http.StatusBadRequest)
		return
	}

	if err := request.Validate(); err != nil {
		response.ResponseBadRequest(w, "invalid request")
		logger.DebugContext(ctx, "SampleEntityHandler.", "url", r.Host, "status", http.StatusBadRequest)
		return
	}

	_, err := s.sampleEntityBiz.Create(ctx, request.ToEntity())
	if err != nil {
		response.ResponseInternalError(w)
		logger.DebugContext(ctx, "SampleEntityHandler.", "url", r.Host, "status", http.StatusInternalServerError)
		return
	}

	response.ResponseCreated(w)
	logger.DebugContext(ctx, "SampleEntityHandler.", "url", r.Host, "status", http.StatusOK)
}

func (s *SampleEntityHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := s.logger.With("method", "List", "ctx", utils.GetLoggerContext(ctx))
	logger.Debug("list called with ctx")

	es, err := s.sampleEntityBiz.List(ctx)
	if err != nil {
		response.ResponseInternalError(w)
		return
	}

	response.ResponseOk(w, es, "")
	logger.DebugContext(ctx, "SampleEntityHandler.", "url", r.Host, "status", http.StatusOK)
}

func (s *SampleEntityHandler) RegisterMuxRouter(mux *http.ServeMux) {
	recoverMiddleware, err := httprecovery.NewRecoveryMiddleware()
	if err != nil {
		panic(err)
	}

	loggerMiddlewareDebug, err := httplogger.NewLoggerMiddleware(httplogger.WithLevel(slog.LevelDebug))
	if err != nil {
		panic(err)
	}

	middles := []middlewares.Middleware{
		recoverMiddleware.RecoverMiddleware,
		httplogger.SetRequestContextLogger,
		loggerMiddlewareDebug.LoggerMiddleware,
	}

	mux.HandleFunc("POST /api/sample-entities/", middlewares.MultipleMiddleware(s.Create, middles...))
	mux.HandleFunc("GET /api/sample-entities/", middlewares.MultipleMiddleware(s.List, middles...))
}
