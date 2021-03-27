package models

type PlaylistSource struct {
	PlaylistID int
	TrackID    int
	Track      *Track
	Playlist   *Playlist
}
