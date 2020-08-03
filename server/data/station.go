package data

import (
	"context"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"

	pb "github.com/fieldkit/app-protocol"
)

type Station struct {
	ID                 int32     `db:"id,omitempty"`
	Name               string    `db:"name"`
	DeviceID           []byte    `db:"device_id"`
	OwnerID            int32     `db:"owner_id,omitempty"`
	CreatedAt          time.Time `db:"created_at,omitempty"`
	UpdatedAt          time.Time `db:"updated_at,omitempty"`
	Battery            *float32  `db:"battery"`
	Location           *Location `db:"location"`
	LocationName       *string   `db:"location_name"`
	RecordingStartedAt *int64    `db:"recording_started_at"`
	MemoryUsed         *int32    `db:"memory_used"`
	MemoryAvailable    *int32    `db:"memory_available"`
	FirmwareNumber     *int32    `db:"firmware_number"`
	FirmwareTime       *int64    `db:"firmware_time"`
	PlaceOther         *string   `db:"place_other"`
	PlaceNative        *string   `db:"place_native"`
	PhotoID            *int64    `db:"photo_id"`
}

func (s *Station) ParseHttpReply(raw string) (*pb.HttpReply, error) {
	bytes, err := DecodeBinaryString(raw)
	if err != nil {
		return nil, err
	}

	buffer := proto.NewBuffer(bytes)
	_, err = buffer.DecodeVarint()
	if err != nil {
		return nil, err
	}

	record := &pb.HttpReply{}
	if err := buffer.Unmarshal(record); err != nil {
		return nil, err
	}

	return record, nil
}

func (s *Station) UpdateFromStatus(ctx context.Context, raw string) error {
	record, err := s.ParseHttpReply(raw)
	if err != nil {
		return err
	}

	if record.Status != nil {
		log := Logger(ctx).Sugar()

		status := record.Status

		if status.Identity != nil {
			if status.Identity.Name != "" {
				s.Name = status.Identity.Name
			}
			if status.Identity.Device != "" {
				s.Name = status.Identity.Device
			}
		}

		if status.Power != nil && status.Power.Battery != nil {
			log.Infow("battery", "device_id", s.DeviceID, "station_id", s.ID, "battery", status.Power)

			battery := float32(status.Power.Battery.Percentage)
			s.Battery = &battery
		}

		if status.Memory != nil {
			log.Infow("memory", "device_id", s.DeviceID, "station_id", s.ID, "memory", status.Memory)

			memoryUsed := int32(status.Memory.DataMemoryUsed)
			memoryInstalled := int32(status.Memory.DataMemoryInstalled)
			s.MemoryUsed = &memoryUsed
			s.MemoryAvailable = &memoryInstalled
		}

		if status.Gps != nil {
			gps := status.Gps

			log.Infow("gps", "device_id", s.DeviceID, "station_id", s.ID, "gps", gps)

			if gps.Fix > 0 {
				log.Infow("status: has gps fix", "station_id", s.ID, "gps", gps)
				s.Location = NewLocation([]float64{
					float64(gps.Longitude),
					float64(gps.Latitude),
				})
			} else if gps.Time > 0 && s.Location == nil {
				log.Infow("status: has old fix", "station_id", s.ID, "gps", gps)
				s.Location = NewLocation([]float64{
					float64(gps.Longitude),
					float64(gps.Latitude),
				})
			}
		} else {
			log.Infow("status: no gps in status", "station_id", s.ID)
		}

		if status.Recording != nil {
			log.Infow("recording", "device_id", s.DeviceID, "station_id", s.ID, "recording", status.Recording)

			recordingStartedAt := int64(status.Recording.StartedTime)
			if recordingStartedAt > 0 {
				s.RecordingStartedAt = &recordingStartedAt
			}
		}

		if status.Firmware != nil {
			log.Infow("firmware", "device_id", s.DeviceID, "station_id", s.ID, "firmware", status.Firmware)

			firmwareTime := int64(status.Firmware.Timestamp)
			s.FirmwareTime = &firmwareTime

			if number, err := strconv.Atoi(status.Firmware.Number); err == nil {
				firmwareNumber := int32(number)
				s.FirmwareNumber = &firmwareNumber
			}
		}
	}

	return nil
}

func (s *Station) DeviceIDHex() string {
	return hex.EncodeToString(s.DeviceID)
}

type AggregatedDataSummary struct {
	StationID     int32      `db:"station_id"`
	Start         *time.Time `db:"start"`
	End           *time.Time `db:"end"`
	NumberSamples *int64     `db:"number_samples"`
}

type StationFull struct {
	Station        *Station
	Owner          *User
	Ingestions     []*Ingestion
	Configurations []*StationConfiguration
	Modules        []*StationModule
	Sensors        []*ModuleSensor
	DataSummary    *AggregatedDataSummary
}

type EssentialStation struct {
	ID                 int64      `db:"id"`
	Name               string     `db:"name"`
	DeviceID           []byte     `db:"device_id"`
	OwnerID            int32      `db:"owner_id"`
	OwnerName          string     `db:"owner_name"`
	CreatedAt          time.Time  `db:"created_at,omitempty"`
	UpdatedAt          time.Time  `db:"updated_at,omitempty"`
	RecordingStartedAt *int64     `db:"recording_started_at"`
	MemoryUsed         *int32     `db:"memory_used"`
	MemoryAvailable    *int32     `db:"memory_available"`
	FirmwareNumber     *int32     `db:"firmware_number"`
	FirmwareTime       *int64     `db:"firmware_time"`
	Location           *Location  `db:"location"`
	LastIngestionAt    *time.Time `db:"last_ingestion_at,omitempty"`
}
