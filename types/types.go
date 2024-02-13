package types

import "time"

type TigerResponse struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	DOB          string    `json:"dob"`
	LastSeenAt   time.Time `json:"last_seen_at"`
	LastSeenLat  float64   `json:"last_seen_lat"`
	LastSeenLong float64   `json:"last_seen_long"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetTigerResponse struct {
	Message string          `json:"message"`
	Data    []TigerResponse `json:"data"`
}

type CreateTiger struct {
	TigerId int64 `json:"tiger_id"`
}

type CreateTigerResponse struct {
	Message string      `json:"message"`
	Data    CreateTiger `json:"data"`
}

type CreateTigerSighting struct {
	TigerSightingId int64 `json:"tiger_sighting_id"`
}

type CreateTigerSightingResponse struct {
	Message string              `json:"message"`
	Data    CreateTigerSighting `json:"data"`
}

type TigerSightingsResponse struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	TigerId   int64     `json:"tiger_id"`
	SeenAt    time.Time `json:"seen_at"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetTigerSightingResponse struct {
	Message string                   `json:"message"`
	Data    []TigerSightingsResponse `json:"data"`
}
