package ports

// GameRepository defines the repository interface for game data
type TimeZoneConverter interface {
	Convert(timestamp string, timezone string) error
}
