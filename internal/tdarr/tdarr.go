package tdarr

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/robertlestak/tdarr_exporter/internal/prom"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Host      string
	VerifySSL bool
	Interval  time.Duration
}

func NewServerFromEnv() Server {
	l := log.WithFields(log.Fields{
		"app": "tdarr_exporter",
		"fn":  "NewServerFromEnv",
	})
	l.Debug("creating new server from env")
	s := Server{
		Host:      os.Getenv("TDARR_HOST"),
		VerifySSL: os.Getenv("TDARR_VERIFY_SSL") != "false", // default to true
	}
	if s.Host == "" {
		s.Host = "http://tdarr:8265"
		l.WithField("host", s.Host).Warnf("TDARR_HOST not set, defaulting to %s", s.Host)
	}
	if os.Getenv("TDARR_INTERVAL") != "" {
		d, err := time.ParseDuration(os.Getenv("TDARR_INTERVAL"))
		if err != nil {
			log.WithError(err).Error("error parsing interval")
			os.Exit(1)
		}
		s.Interval = d
	} else {
		s.Interval = time.Minute
		l.WithField("interval", s.Interval).Warnf("TDARR_INTERVAL not set, defaulting to %s", s.Interval)
	}
	return s
}

type TdarrStatsRequest struct {
	Collection string `json:"collection"`
	Mode       string `json:"mode"`
	DocID      string `json:"docID"`
}

type TdarrStatsRequestData struct {
	Data TdarrStatsRequest `json:"data"`
}

type LanguageMetric struct {
	Count int `json:"count"`
}

type TranscodeInfo struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type CategoryInfo struct {
	Library               string          `json:"library"`
	ID                    string          `json:"_id"`
	TotalFileCount        int             `json:"totalFileCount"`
	TotalTranscodeCount   int             `json:"totalTranscodeCount"`
	SizeDiff              float64         `json:"sizeDiff"`
	TotalHealthCheckCount int             `json:"totalHealthCheckCount"`
	TranscodeStatus       []TranscodeInfo `json:"transcode_status"`
	Health                []TranscodeInfo `json:"health"`
	VideoCodec            []TranscodeInfo `json:"video_codec"`
	Container             []TranscodeInfo `json:"container"`
	Resolution            []TranscodeInfo `json:"resolution"`
	AudioCodec            []TranscodeInfo `json:"audio_codec"`
	AudioContainer        []TranscodeInfo `json:"audio_container"`
}

type TdarrStatsResponse struct {
	ID                    string         `json:"_id"`
	TotalFileCount        int            `json:"totalFileCount"`
	TotalTranscodeCount   int            `json:"totalTranscodeCount"`
	TotalHealthCheckCount int            `json:"totalHealthCheckCount"`
	SizeDiff              float64        `json:"sizeDiff"`
	DBFetchTime           string         `json:"DBFetchTime"`
	DBLoadStatus          string         `json:"DBLoadStatus"`
	DBQueue               int            `json:"DBQueue"`
	Pies                  []any          `json:"pies"`
	ParsedPies            []CategoryInfo `json:"parsedPies"`
	TdarrScore            string         `json:"tdarrScore"`
	HealthCheckScore      string         `json:"healthCheckScore"`
	ProcessWarning        string         `json:"processWarning"`
	ProcessWarningQueues  bool           `json:"processWarningQueues"`
	Table0Count           int            `json:"table0Count"`
	Table1Count           int            `json:"table1Count"`
	Table2Count           int            `json:"table2Count"`
	Table3Count           int            `json:"table3Count"`
	Table4Count           int            `json:"table4Count"`
	Table5Count           int            `json:"table5Count"`
	Table6Count           int            `json:"table6Count"`
	Table0ViewableCount   int            `json:"table0ViewableCount"`
	Table1ViewableCount   int            `json:"table1ViewableCount"`
	Table2ViewableCount   int            `json:"table2ViewableCount"`
	Table3ViewableCount   int            `json:"table3ViewableCount"`
	Table4ViewableCount   int            `json:"table4ViewableCount"`
	Table5ViewableCount   int            `json:"table5ViewableCount"`
	Table6ViewableCount   int            `json:"table6ViewableCount"`
	StreamStats           struct {
		Duration struct {
			Average int `json:"average"`
			Highest int `json:"highest"`
			Total   int `json:"total"`
		} `json:"duration"`
		BitRate struct {
			Average int   `json:"average"`
			Highest int   `json:"highest"`
			Total   int64 `json:"total"`
		} `json:"bit_rate"`
		NbFrames struct {
			Average int `json:"average"`
			Highest int `json:"highest"`
			Total   int `json:"total"`
		} `json:"nb_frames"`
	} `json:"streamStats"`
	AvgNumberOfStreamsInVideo float64                   `json:"avgNumberOfStreamsInVideo"`
	Languages                 map[string]LanguageMetric `json:"languages"`
}

func (r *TdarrStatsResponse) ParsePies() error {
	l := log.WithFields(log.Fields{
		"app": "tdarr_exporter",
		"fn":  "ParsePies",
	})
	l.Debug("parsing pies")

	var parsedPies []CategoryInfo

	for _, pie := range r.Pies {
		pieArray, ok := pie.([]interface{})
		if !ok {
			l.Warn("Invalid pie format: not an array")
			continue
		}

		if len(pieArray) < 7 {
			l.Warn("Invalid pie format: not enough elements")
			continue
		}

		c := CategoryInfo{
			Library:               pieArray[0].(string),
			ID:                    pieArray[1].(string),
			TotalFileCount:        int(pieArray[2].(float64)),
			TotalTranscodeCount:   int(pieArray[3].(float64)),
			SizeDiff:              pieArray[4].(float64),
			TotalHealthCheckCount: int(pieArray[5].(float64)),
		}

		// Parse the slices
		for i := 6; i < len(pieArray); i++ {
			subArray, ok := pieArray[i].([]interface{})
			if !ok {
				l.Warn("Invalid sub-array format")
				continue
			}

			for _, subElement := range subArray {
				subMap, ok := subElement.(map[string]interface{})
				if !ok {
					l.Warn("Invalid sub-map format")
					continue
				}

				var transcodeInfo TranscodeInfo
				transcodeInfo.Name = subMap["name"].(string)
				transcodeInfo.Value = int(subMap["value"].(float64))

				switch i {
				case 6:
					c.TranscodeStatus = append(c.TranscodeStatus, transcodeInfo)
				case 7:
					c.Health = append(c.Health, transcodeInfo)
				case 8:
					c.VideoCodec = append(c.VideoCodec, transcodeInfo)
				case 9:
					c.Container = append(c.Container, transcodeInfo)
				case 10:
					c.Resolution = append(c.Resolution, transcodeInfo)
				case 11:
					c.AudioCodec = append(c.AudioCodec, transcodeInfo)
				case 12:
					c.AudioContainer = append(c.AudioContainer, transcodeInfo)
				}
			}
		}

		parsedPies = append(parsedPies, c)
	}

	r.ParsedPies = parsedPies
	l.Debugf("parsed pies: %v", parsedPies)
	return nil
}

func (s *Server) GetStats() (TdarrStatsResponse, error) {
	l := log.WithFields(log.Fields{
		"app": "tdarr_exporter",
		"fn":  "GetStats",
	})
	var tdarrStatsResponse TdarrStatsResponse
	l.Debug("getting stats from tdarr")
	u := s.Host + "/api/v2/cruddb"
	l.WithField("url", u).Debug("making request")
	statReq := TdarrStatsRequest{
		Collection: "StatisticsJSONDB",
		Mode:       "getById",
		DocID:      "statistics",
	}
	dataWrapper := TdarrStatsRequestData{
		Data: statReq,
	}
	reqJson, err := json.Marshal(dataWrapper)
	if err != nil {
		l.WithError(err).Error("error marshalling request")
		return tdarrStatsResponse, err
	}
	if log.GetLevel() == log.DebugLevel {
		// log the request body
		l.WithField("body", string(reqJson)).Debug("request body")
	}
	req, err := http.NewRequest("POST", u, bytes.NewBuffer(reqJson))
	if err != nil {
		l.WithError(err).Error("error creating request")
		return tdarrStatsResponse, err
	}
	l.Debug("setting request headers")
	req.Header.Add("content-type", "application/json")
	// use default http client
	c := http.DefaultClient
	// set 10s timeout
	c.Timeout = time.Second * 10
	if !s.VerifySSL {
		l.Warn("disabling SSL verification")
		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	l.Debug("making request")
	res, err := c.Do(req)
	if err != nil {
		l.WithError(err).Error("error making request")
		return tdarrStatsResponse, err
	}
	l.Debug("reading response body")
	defer res.Body.Close()
	bd, err := io.ReadAll(res.Body)
	if err != nil {
		l.WithError(err).Error("error reading response body")
		return tdarrStatsResponse, err
	}
	if log.GetLevel() == log.DebugLevel {
		// log the response body
		l.WithField("body", string(bd)).Debug("response body")
	}
	err = json.Unmarshal(bd, &tdarrStatsResponse)
	if err != nil {
		l.WithError(err).Error("error unmarshalling response body")
		return tdarrStatsResponse, err
	}
	l.Debug("parsing pies")
	err = tdarrStatsResponse.ParsePies()
	if err != nil {
		l.WithError(err).Error("error parsing pies")
		return tdarrStatsResponse, err
	}
	return tdarrStatsResponse, err
}

func (s *TdarrStatsResponse) LoadStatusFloat() float64 {
	// There is no documentation on what the load status means
	// I only see "Stable"
	switch s.DBLoadStatus {
	case "Stable":
		return 0
	default:
		// log it so we can see what it is and add it to the switch
		log.WithField("status", s.DBLoadStatus).Warn("unknown DBLoadStatus")
		return 1
	}
}

func (s *TdarrStatsResponse) ExportProm() error {
	l := log.WithFields(log.Fields{
		"app": "tdarr_exporter",
		"fn":  "ExportProm",
	})
	l.Debug("exporting prometheus metrics")
	prom.TotalFileCount.Set(float64(s.TotalFileCount))
	prom.TotalTranscodeCount.Set(float64(s.TotalTranscodeCount))
	prom.TotalHealthCheckCount.Set(float64(s.TotalHealthCheckCount))
	prom.SizeDiff.Set(s.SizeDiff)
	// parse fetch time as duration
	d, err := time.ParseDuration(s.DBFetchTime)
	if err != nil {
		l.WithError(err).Error("error parsing DBFetchTime")
		return err
	}
	prom.DBFetchTime.Set(d.Seconds())
	// parse load status as float
	prom.DBLoadStatus.Set(s.LoadStatusFloat())
	prom.DBQueue.Set(float64(s.DBQueue))
	// parse tdarr score as float
	tf, err := strconv.ParseFloat(s.TdarrScore, 64)
	if err != nil {
		l.WithError(err).Error("error parsing TdarrScore")
		return err
	}
	prom.TdarrScore.Set(tf)
	hf, err := strconv.ParseFloat(s.HealthCheckScore, 64)
	if err != nil {
		l.WithError(err).Error("error parsing HealthCheckScore")
		return err
	}
	prom.HealthCheckScore.Set(hf)
	prom.AverageNumberOfStreamsInVideo.Set(s.AvgNumberOfStreamsInVideo)
	// set languages
	for k, v := range s.Languages {
		prom.Languages.WithLabelValues(k).Set(float64(v.Count))
	}
	prom.StreamStatsDurationAverage.Set(float64(s.StreamStats.Duration.Average))
	prom.StreamStatsDurationHighest.Set(float64(s.StreamStats.Duration.Highest))
	prom.StreamStatsDurationTotal.Set(float64(s.StreamStats.Duration.Total))
	prom.StreamStatsBitrateAverage.Set(float64(s.StreamStats.BitRate.Average))
	prom.StreamStatsBitrateAverage.Set(float64(s.StreamStats.BitRate.Highest))
	prom.StreamStatsBitrateAverage.Set(float64(s.StreamStats.BitRate.Total))
	prom.StreamStatsNbFramesAverage.Set(float64(s.StreamStats.NbFrames.Average))
	prom.StreamStatsNbFramesHighest.Set(float64(s.StreamStats.NbFrames.Highest))
	prom.StreamStatsNbFramesTotal.Set(float64(s.StreamStats.NbFrames.Total))
	prom.Table0Count.Set(float64(s.Table0Count))
	prom.Table1Count.Set(float64(s.Table1Count))
	prom.Table2Count.Set(float64(s.Table2Count))
	prom.Table3Count.Set(float64(s.Table3Count))
	prom.Table4Count.Set(float64(s.Table4Count))
	prom.Table5Count.Set(float64(s.Table5Count))
	prom.Table6Count.Set(float64(s.Table6Count))
	prom.Table0ViewableCount.Set(float64(s.Table0ViewableCount))
	prom.Table1ViewableCount.Set(float64(s.Table1ViewableCount))
	prom.Table2ViewableCount.Set(float64(s.Table2ViewableCount))
	prom.Table3ViewableCount.Set(float64(s.Table3ViewableCount))
	prom.Table4ViewableCount.Set(float64(s.Table4ViewableCount))
	prom.Table5ViewableCount.Set(float64(s.Table5ViewableCount))
	prom.Table6ViewableCount.Set(float64(s.Table6ViewableCount))
	for _, c := range s.ParsedPies {
		prom.LibraryTotalFileCount.WithLabelValues(c.Library, c.ID).Set(float64(c.TotalFileCount))
		prom.LibraryTotalTranscodeCount.WithLabelValues(c.Library, c.ID).Set(float64(c.TotalTranscodeCount))
		prom.LibraryTotalHealthCheckCount.WithLabelValues(c.Library, c.ID).Set(float64(c.TotalHealthCheckCount))
		prom.LibrarySizeDiff.WithLabelValues(c.Library, c.ID).Set(c.SizeDiff)
		for _, t := range c.TranscodeStatus {
			prom.LibraryTranscodeStatus.WithLabelValues(c.Library, c.ID, t.Name).Set(float64(t.Value))
		}
		for _, h := range c.Health {
			prom.LibraryHealth.WithLabelValues(c.Library, c.ID, h.Name).Set(float64(h.Value))
		}
		for _, v := range c.VideoCodec {
			prom.LibraryVideoCodec.WithLabelValues(c.Library, c.ID, v.Name).Set(float64(v.Value))
		}
		for _, v := range c.Container {
			prom.LibraryVideoContainer.WithLabelValues(c.Library, c.ID, v.Name).Set(float64(v.Value))
		}
		for _, r := range c.Resolution {
			prom.LibraryVideoResolution.WithLabelValues(c.Library, c.ID, r.Name).Set(float64(r.Value))
		}
		for _, a := range c.AudioCodec {
			prom.LibraryAudioCodec.WithLabelValues(c.Library, c.ID, a.Name).Set(float64(a.Value))
		}
		for _, v := range c.AudioContainer {
			prom.LibraryAudioContainer.WithLabelValues(c.Library, c.ID, v.Name).Set(float64(v.Value))
		}
	}
	return nil
}
