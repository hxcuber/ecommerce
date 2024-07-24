package model

type Status string

// Enum values for Status
const (
	StatusProcessed Status = "processed"
	StatusShipped   Status = "shipped"
	StatusDelivered Status = "delivered"
)
