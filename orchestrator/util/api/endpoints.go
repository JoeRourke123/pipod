package api

import "orchestrator/util"

var (
	AlbumsList = builder("/list/albums")
	Album      = builderParam("/albums")

	PlaylistList = builder("/list/playlists")
	Playlist     = builderParam("/playlists")

	PodcastList = builder("/list/podcasts")
	Podcast     = builderParam("/podcasts")

	QueueList    = builder("/queue")
	QueueTrack   = builderParam("/queue/track")
	QueueAlbum   = builderParam("/queue/album")
	UnqueueTrack = builderParam("/remove/queue/")

	LoginRedirect = builder("/go")
	CompleteAuth  = builder("/auth")
	IsAuth        = builder("/isAuth")

	Websocket = builder("/ws")

	Collections = builder("/db/collections")
	Collection  = builderParam("/db/collections")

	DownloadTrack       = builderParam("/download/track")
	DownloadAlbum       = builderParam("/download/album")
	RemoveDownloadTrack = builderParam("/remove/download/track")
	RemoveDownloadAlbum = builderParam("/remove/download/album")

	Player = builder("/player")

	HomeView  = builder("/views/home")
	MusicView = builder("/views/music")
)

func builder(path string) func() string {
	return func() string {
		return path
	}
}

func builderParam(path string) func(string) string {
	return func(parameter string) string {
		return path + "/" + parameter
	}
}

func Full(path string) string {
	apiUrl := util.GetApiUrl()
	return apiUrl + path
}
