package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type IngestedMessage struct {
	Schema   *JsonMessageSchema
	Time     *time.Time
	Location *Location
	Fields   map[string]interface{}
}

const (
	FieldNameLatitude  = "latitude"
	FieldNameLongitude = "longitude"
	FieldNameAltitude  = "altitude"
	FieldNameTime      = "time"
)

func parseLocation(pm *ProcessedMessage, ms *JsonMessageSchema, m map[string]interface{}) (l *Location, err error) {
	coordinates := make([]float32, 0)
	for _, key := range []string{FieldNameLatitude, FieldNameLongitude, FieldNameAltitude} {
		f, err := strconv.ParseFloat(m[key].(string), 32)
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, float32(f))
	}
	return &Location{Coordinates: coordinates}, nil
}

func parseTime(pm *ProcessedMessage, ms *JsonMessageSchema, m map[string]interface{}) (t *time.Time, err error) {
	if ms.UseProviderTime {
		t = pm.Time
	} else {
		if !ms.HasTime {
			return nil, fmt.Errorf("%s: no time information.", pm.SchemaId)
		}

		raw := m[FieldNameTime].(string)
		unix, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return nil, err
		}
		parsed := time.Unix(unix, 0)
		t = &parsed
	}

	return
}

func (i *MessageIngester) ApplySchema(pm *ProcessedMessage, ms *JsonMessageSchema) (im *IngestedMessage, err error) {
	if len(pm.ArrayValues) != len(ms.Fields) {
		return nil, fmt.Errorf("%s: fields S=%v != M=%v", pm.SchemaId, len(ms.Fields), len(pm.ArrayValues))
	}

	mapped := make(map[string]interface{})

	for i, field := range ms.Fields {
		mapped[ToSnake(field.Name)] = pm.ArrayValues[i]
	}

	stream, err := i.Streams.LookupStream(pm.SchemaId)
	if err != nil {
		return nil, err
	}

	time, err := parseTime(pm, ms, mapped)
	if err != nil {
		return nil, err
	}

	if ms.HasLocation {
		location, err := parseLocation(pm, ms, mapped)
		if err != nil {
			return nil, fmt.Errorf("%s: unable to parse location.", pm.SchemaId)
		}
		stream.SetLocation(time, location)
	}

	im = &IngestedMessage{
		Schema:   ms,
		Time:     time,
		Location: nil,
		Fields:   mapped,
	}

	return
}

func (i *MessageIngester) ApplySchemas(pm *ProcessedMessage, schemas []interface{}) (im *IngestedMessage, err error) {
	errors := make([]error, 0)

	if len(schemas) == 0 {
		errors = append(errors, fmt.Errorf("No matching schemas"))
	}

	for _, schema := range schemas {
		jsonSchema, ok := schema.(*JsonMessageSchema)
		if ok {
			im, applyErr := i.ApplySchema(pm, jsonSchema)
			if applyErr != nil {
				errors = append(errors, applyErr)
			} else {
				return im, nil
			}
		} else {
			log.Printf("Unexpected MessageSchema type: %v", schema)
		}
	}
	return nil, fmt.Errorf("Schemas failed: %v", errors)
}

func (i *MessageIngester) HandleMessage(raw *RawMessage) error {
	i.Statistics.Processed += 1

	mp, err := IdentifyMessageProvider(raw)
	if err != nil {
		return err
	}
	if mp == nil {
		log.Printf("(%s)[Error]: No message provider: (UserAgent: %v) (ContentType: %s) Body: %s",
			raw.Data.Context.RequestId, raw.Data.Params.Headers.UserAgent, raw.Data.Params.Headers.ContentType, raw.Data)
		return nil
	}

	pm, err := mp.ProcessMessage(raw)
	if err != nil {
		log.Printf("(%s)[Error]: %v", raw.Data.Context.RequestId, err)
	}
	if pm != nil {
		schemas, err := i.Schemas.LookupSchema(pm.SchemaId)
		if err != nil {
			return err
		}
		fm, err := i.ApplySchemas(pm, schemas)
		if err != nil {
			if true {
				log.Printf("(%s)(%s)[Error]: %v %s", pm.MessageId, pm.SchemaId, err, pm.ArrayValues)
			} else {
				log.Printf("(%s)(%s)[Error]: %v", pm.MessageId, pm.SchemaId, err)
			}
		} else {
			if true {
				log.Printf("(%s)(%s)[Success]", pm.MessageId, pm.SchemaId)
			}
			i.Statistics.Successes += 1
		}
		_ = fm
	}

	return nil
}

type IngestionStatistics struct {
	Processed uint64
	Successes uint64
}

type MessageIngester struct {
	Handler
	Schemas    SchemaRepository
	Streams    StreamsRepository
	Statistics IngestionStatistics
}

func NewMessageIngester(sr SchemaRepository, streams StreamsRepository) *MessageIngester {
	return &MessageIngester{
		Schemas: sr,
		Streams: streams,
	}
}
