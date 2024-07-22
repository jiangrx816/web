package download

type AlbumAudioRequest struct {
	AppType        string        `json:"app_type"`
	AppVersion     string        `json:"app_version"`
	DeviceID       string        `json:"device_id"`
	MessageID      string        `json:"message_id"`
	DeviceType     string        `json:"device_type"`
	Lon            string        `json:"lon"`
	Version        string        `json:"version"`
	DeviceUUID     string        `json:"device_uuid"`
	Token          string        `json:"token"`
	SysVersion     string        `json:"sys_version"`
	UserID         string        `json:"user_id"`
	SystemVersion  string        `json:"system_version"`
	AlbumID        string        `json:"album_id"`
	Model          string        `json:"model"`
	ChannelType    string        `json:"channel_type"`
	DeviceUniqueID string        `json:"device_unique_id"`
	Brand          string        `json:"brand"`
	AudioIDList    []interface{} `json:"audio_id_list"`
	Lat            string        `json:"lat"`
}

type AlbumAudioInfoRequest struct {
	AppType        string `json:"app_type"`
	AppVersion     string `json:"app_version"`
	DeviceID       string `json:"device_id"`
	MessageID      string `json:"message_id"`
	DeviceType     string `json:"device_type"`
	Lon            string `json:"lon"`
	Version        string `json:"version"`
	DeviceUUID     string `json:"device_uuid"`
	Token          string `json:"token"`
	SysVersion     string `json:"sys_version"`
	UserID         string `json:"user_id"`
	SystemVersion  string `json:"system_version"`
	AlbumID        string `json:"album_id"`
	Model          string `json:"model"`
	ChannelType    string `json:"channel_type"`
	DeviceUniqueID string `json:"device_unique_id"`
	Brand          string `json:"brand"`
	AudioIDList    []struct {
		AudioID string `json:"audio_id"`
	} `json:"audio_id_list"`
	Lat string `json:"lat"`
}

type AudioInfoMp3Request struct {
	AppType        string `json:"app_type"`
	AppVersion     string `json:"app_version"`
	DeviceID       string `json:"device_id"`
	MessageID      string `json:"message_id"`
	DeviceType     string `json:"device_type"`
	Lon            string `json:"lon"`
	Version        string `json:"version"`
	DeviceUUID     string `json:"device_uuid"`
	Token          string `json:"token"`
	SysVersion     string `json:"sys_version"`
	UserID         string `json:"user_id"`
	SystemVersion  string `json:"system_version"`
	Model          string `json:"model"`
	ChannelType    string `json:"channel_type"`
	DeviceUniqueID string `json:"device_unique_id"`
	Brand          string `json:"brand"`
	AudioIDList    []struct {
		AudioID string `json:"audio_id"`
	} `json:"audio_id_list"`
	Lat string `json:"lat"`
}
