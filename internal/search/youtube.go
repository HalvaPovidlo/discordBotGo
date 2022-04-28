package search

import (
	"regexp"
	"sort"

	iso "github.com/channelmeter/iso8601duration"
	ytdl "github.com/kkdai/youtube/v2"
	"github.com/pkg/errors"
	"google.golang.org/api/youtube/v3"

	"github.com/HalvaPovidlo/discordBotGo/internal/pkg"
	"github.com/HalvaPovidlo/discordBotGo/pkg/contexts"
)

const (
	videoPrefix     = "https://youtube.com/watch?v="
	channelPrefix   = "https://youtube.com/channel/"
	maxSearchResult = 50
)

var (
	ErrSongNotFound = errors.New("song not found")
)

// TODO: work on this
// TODO: race conditions?

// YouTube exports the methods required to access the YouTube service
type YouTube struct {
	ctx     contexts.Context
	ytdl    *ytdl.Client
	youtube *youtube.Service
}

func NewYouTubeClient(ctx contexts.Context, ytdl *ytdl.Client, yt *youtube.Service) *YouTube {
	return &YouTube{
		ytdl:    ytdl,
		youtube: yt,
		ctx:     ctx,
	}
}

func (y *YouTube) getQuery(query string) (*pkg.Song, error) {
	call := y.youtube.Search.List([]string{"id"}).
		Q(query).
		MaxResults(maxSearchResult)

	response, err := call.Do()
	if err != nil {
		// TODO: NOT FOUND?
		return nil, errors.Wrap(err, "could not find any results for the specified query")
	}

	for _, item := range response.Items {
		if item.Id.Kind == "youtube#video" {
			return &pkg.Song{
				Title:        item.Snippet.Title,
				URL:          videoPrefix + item.Id.VideoId,
				Service:      pkg.ServiceYouTube,
				ArtistName:   item.Snippet.ChannelTitle,
				ArtistURL:    channelPrefix + item.Snippet.ChannelId,
				ArtworkURL:   item.Snippet.Thumbnails.Maxres.Url,
				ThumbnailURL: item.Snippet.Thumbnails.Standard.Url,
				ID: pkg.SongID{
					ID:      item.Id.VideoId,
					Service: pkg.ServiceYouTube,
				},
			}, nil
		}
	}

	return nil, ErrSongNotFound
}

// GetMetadata returns the metadata for a given YouTube video URL
func (y *YouTube) GetMetadata(url string) (*pkg.Song, error) {
	videoInfo, err := y.ytdl.GetVideo(url)
	if err != nil {
		return nil, errors.Wrapf(err, "loag video metadata by url %s", url)
	}
	formats := videoInfo.Formats
	if len(formats) == 0 {
		return nil, errors.New("unable to get list of formats")
	}

	sort.SliceStable(formats, func(i, j int) bool {
		return formats[i].ItagNo < formats[j].ItagNo
	})

	format := formats[0]

	videoURL, err := y.ytdl.GetStreamURL(videoInfo, &format)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get streamURL %s", videoInfo.Title)
	}

	ytCall := youtube.NewVideosService(y.youtube).
		List([]string{"snippet", "contentDetails"}).
		Id(videoInfo.ID)

	ytResponse, err := ytCall.Do()
	if err != nil {
		return nil, err
	}

	duration, err := iso.FromString(ytResponse.Items[0].ContentDetails.Duration)
	if err != nil {
		return nil, err
	}

	metadata := &pkg.Metadata{
		Title:      videoInfo.Title,
		DisplayURL: url,
		StreamURL:  videoURL,
		Duration:   duration.ToDuration().Seconds(),
	}

	videoAuthor := &pkg.MetadataArtist{
		Name: ytResponse.Items[0].Snippet.ChannelTitle,
		URL:  channelPrefix + ytResponse.Items[0].Snippet.ChannelId,
	}
	metadata.Artists = append(metadata.Artists, *videoAuthor)

	return metadata, nil
}

func (y *YouTube) FindSong(query string) (*pkg.SongRequest, error) {
	logger := y.ctx.LoggerFromContext()
	logger.Debug("get query")
	url, err := y.getQuery(query)
	if err != nil {
		if err == ErrSongNotFound {
			return nil, ErrSongNotFound
		}
		return nil, errors.Wrap(err, "search in youtube")
	}

	test, err := testURL(url)
	if err != nil {
		return nil, err
	}
	if test {
		logger.Debug("GetMetadata")
		metadata, err := y.GetMetadata(url)
		if err != nil {
			return nil, errors.Wrap(err, "get metadata")
		}
		req := &pkg.SongRequest{
			Metadata:     metadata,
			ServiceName:  pkg.ServiceYouTube,
			ServiceColor: y.GetColor(),
		}
		return req, nil
	}

	return nil, errors.New("wrong youtube url")
}

func testURL(url string) (bool, error) {
	test, err := regexp.MatchString("(?:https?://)?(?:www\\.)?youtu\\.?be(?:\\.com)?/?.*(?:watch|embed)?(?:.*v=|v/|/)[\\w-_]+", url)
	return test, err
}