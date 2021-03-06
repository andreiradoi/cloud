package logging

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"

	"gopkg.in/alexcesaro/statsd.v2"
)

type MetricsSettings struct {
	Prefix  string
	Address string
}

type Metrics struct {
	SC       *statsd.Client
	Settings *MetricsSettings
}

func newClient(ms *MetricsSettings) (*statsd.Client, error) {
	if ms.Address == "" {
		return statsd.New(statsd.Prefix(ms.Prefix))
	}
	return statsd.New(statsd.Prefix(ms.Prefix), statsd.Address(ms.Address))
}

func NewMetrics(ctx context.Context, ms *MetricsSettings) *Metrics {
	log := Logger(ctx).Sugar()

	log.Infow("statsd", "address", ms.Address)

	sc, err := newClient(ms)
	if err != nil {
		log.Warnw("error", "error", err)
	}

	return &Metrics{
		SC:       sc,
		Settings: ms,
	}
}

func (m *Metrics) AuthTry() {
	m.SC.Increment("api.auth.try")
}

func (m *Metrics) AuthSuccess() {
	m.SC.Increment("api.auth.success")
}

func (m *Metrics) AuthRefreshTry() {
	m.SC.Increment("api.auth.refresh.try")
}

func (m *Metrics) AuthRefreshSuccess() {
	m.SC.Increment("api.auth.refresh.success")
}

func (m *Metrics) Ingested(blocks, bytes int) {
	m.SC.Count("api.ingestion.blocks", blocks)
	m.SC.Count("api.ingestion.bytes", bytes)
}

func (m *Metrics) DataErrorsParsing() {
	m.SC.Increment("api.data.errors.parsing")
}

func (m *Metrics) DataErrorsMissingMeta() {
	m.SC.Increment("api.data.errors.meta")
}

func (m *Metrics) MessagePublished() {
	m.SC.Increment("messages.published")
}

type Timing struct {
	sc         *statsd.Client
	timer      statsd.Timing
	timingKey  string
	counterKey string
}

func (t *Timing) Send() {
	t.timer.Send(t.timingKey)
	t.sc.Increment(t.counterKey)
}

func (m *Metrics) FileUpload() *Timing {
	timer := m.SC.NewTiming()

	return &Timing{
		sc:         m.SC,
		timer:      timer,
		timingKey:  "files.uploading.time",
		counterKey: "files.uploaded",
	}
}

func (m *Metrics) HandleMessage() *Timing {
	timer := m.SC.NewTiming()

	return &Timing{
		sc:         m.SC,
		timer:      timer,
		timingKey:  "messages.handling.time",
		counterKey: "messages.processed",
	}
}

func (m *Metrics) UserValidated() {
	m.SC.Increment("api.users.validated")
}

func (m *Metrics) UserAdded() {
	m.SC.Increment("api.users.added")
}

func (m *Metrics) EmailVerificationSent() {
	m.SC.Increment("emails.verification")
}

func (m *Metrics) EmailRecoverPasswordSent() {
	m.SC.Increment("emails.password.recover")
}

func (m *Metrics) DataErrorsUnknown() {
	m.SC.Increment("api.data.errors.unknown")
}

func (m *Metrics) IngestionDevice(id []byte) {
	m.SC.Unique("api.ingestion.devices", hex.EncodeToString(id))
}

func (m *Metrics) UserID(id int32) {
	m.SC.Unique("api.users", fmt.Sprintf("%d", id))
}

func (m *Metrics) GatherMetrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := m.SC.NewTiming()

		h.ServeHTTP(w, r)

		t.Send("http.req.time")
	})
}

func (m *Metrics) RecordsViewed(records int) {
	m.SC.Count("api.data.records.viewed", records)
}

func (m *Metrics) ReadingsViewed(readings int) {
	m.SC.Count("api.data.readings.viewed", readings)
}
