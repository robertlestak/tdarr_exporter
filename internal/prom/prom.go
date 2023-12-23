package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalFileCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_total_file_count",
		Help: "Total number of files in tdarr",
	})
	TotalTranscodeCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_total_transcode_count",
		Help: "Total number of transcodes in tdarr",
	})
	TotalHealthCheckCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_total_health_check_count",
		Help: "Total number of health checks in tdarr",
	})
	SizeDiff = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_size_diff",
		Help: "Size difference in tdarr",
	})
	DBFetchTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_db_fetch_time",
		Help: "DB fetch time in tdarr",
	})
	DBLoadStatus = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_db_load_status",
		Help: "DB load status in tdarr",
	})
	DBQueue = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_db_queue",
		Help: "DB queue in tdarr",
	})
	TdarrScore = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_score",
		Help: "Tdarr score",
	})
	HealthCheckScore = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_health_check_score",
		Help: "Health check score",
	})
	AverageNumberOfStreamsInVideo = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_average_number_of_streams_in_video",
		Help: "Average number of streams in video",
	})
	Languages = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_languages",
		Help: "Languages",
	}, []string{"language"})
	StreamStatsDurationAverage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_stream_stats_duration_average",
		Help: "Average duration of streams",
	})
	StreamStatsDurationHighest = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_stream_stats_duration_highest",
		Help: "Highest duration of streams",
	})
	StreamStatsDurationTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_stream_stats_duration_total",
		Help: "Total duration of streams",
	})
	StreamStatsBitrateAverage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_stream_stats_bitrate_average",
		Help: "Average bitrate of streams",
	})
	StreamStatsBitrateHighest = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_stream_stats_bitrate_highest",
		Help: "Highest bitrate of streams",
	})
	StreamStatsBitrateTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_stream_stats_bitrate_total",
		Help: "Total bitrate of streams",
	})
	StreamStatsNbFramesAverage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_stream_stats_nb_frames_average",
		Help: "Average number of frames in streams",
	})
	StreamStatsNbFramesHighest = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_stream_stats_nb_frames_highest",
		Help: "Highest number of frames in streams",
	})
	StreamStatsNbFramesTotal = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_stream_stats_nb_frames_total",
		Help: "Total number of frames in streams",
	})
	Table0Count = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_0_count",
		Help: "Table 0 count",
	})
	Table1Count = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_1_count",
		Help: "Table 1 count",
	})
	Table2Count = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_2_count",
		Help: "Table 2 count",
	})
	Table3Count = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_3_count",
		Help: "Table 3 count",
	})
	Table4Count = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_4_count",
		Help: "Table 4 count",
	})
	Table5Count = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_5_count",
		Help: "Table 5 count",
	})
	Table6Count = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_6_count",
		Help: "Table 6 count",
	})
	Table0ViewableCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_0_viewable_count",
		Help: "Table 0 viewable count",
	})
	Table1ViewableCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_1_viewable_count",
		Help: "Table 1 viewable count",
	})
	Table2ViewableCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_2_viewable_count",
		Help: "Table 2 viewable count",
	})
	Table3ViewableCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_3_viewable_count",
		Help: "Table 3 viewable count",
	})
	Table4ViewableCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_4_viewable_count",
		Help: "Table 4 viewable count",
	})
	Table5ViewableCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_5_viewable_count",
		Help: "Table 5 viewable count",
	})
	Table6ViewableCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tdarr_table_6_viewable_count",
		Help: "Table 6 viewable count",
	})
	LibraryTotalFileCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_total_file_count",
		Help: "Total number of files in tdarr library",
	}, []string{"library_name", "library_id"})
	LibraryTotalTranscodeCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_total_transcode_count",
		Help: "Total number of transcodes in tdarr library",
	}, []string{"library_name", "library_id"})
	LibraryTotalHealthCheckCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_total_health_check_count",
		Help: "Total number of health checks in tdarr library",
	}, []string{"library_name", "library_id"})
	LibrarySizeDiff = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_size_diff",
		Help: "Size difference in tdarr library",
	}, []string{"library_name", "library_id"})
	LibraryTranscodeStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_transcode_status",
		Help: "Transcode status in tdarr library",
	}, []string{"library_name", "library_id", "status"})
	LibraryHealth = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_health",
		Help: "Health in tdarr library",
	}, []string{"library_name", "library_id", "health"})
	LibraryVideoCodec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_video_codec",
		Help: "Video codec in tdarr library",
	}, []string{"library_name", "library_id", "codec"})
	LibraryVideoContainer = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_video_container",
		Help: "Video container in tdarr library",
	}, []string{"library_name", "library_id", "container"})
	LibraryVideoResolution = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_video_resolution",
		Help: "Video resolution in tdarr library",
	}, []string{"library_name", "library_id", "resolution"})
	LibraryAudioCodec = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_audio_codec",
		Help: "Audio codec in tdarr library",
	}, []string{"library_name", "library_id", "codec"})
	LibraryAudioContainer = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tdarr_library_audio_container",
		Help: "Audio container in tdarr library",
	}, []string{"library_name", "library_id", "container"})
)

func InitMetrics() {
	prometheus.MustRegister(TotalFileCount)
	prometheus.MustRegister(TotalTranscodeCount)
	prometheus.MustRegister(TotalHealthCheckCount)
	prometheus.MustRegister(SizeDiff)
	prometheus.MustRegister(DBFetchTime)
	prometheus.MustRegister(DBLoadStatus)
	prometheus.MustRegister(DBQueue)
	prometheus.MustRegister(TdarrScore)
	prometheus.MustRegister(HealthCheckScore)
	prometheus.MustRegister(AverageNumberOfStreamsInVideo)
	prometheus.MustRegister(Languages)
	prometheus.MustRegister(StreamStatsDurationAverage)
	prometheus.MustRegister(StreamStatsDurationHighest)
	prometheus.MustRegister(StreamStatsDurationTotal)
	prometheus.MustRegister(StreamStatsBitrateAverage)
	prometheus.MustRegister(StreamStatsBitrateHighest)
	prometheus.MustRegister(StreamStatsBitrateTotal)
	prometheus.MustRegister(StreamStatsNbFramesAverage)
	prometheus.MustRegister(StreamStatsNbFramesHighest)
	prometheus.MustRegister(StreamStatsNbFramesTotal)
	prometheus.MustRegister(Table0Count)
	prometheus.MustRegister(Table1Count)
	prometheus.MustRegister(Table2Count)
	prometheus.MustRegister(Table3Count)
	prometheus.MustRegister(Table4Count)
	prometheus.MustRegister(Table5Count)
	prometheus.MustRegister(Table6Count)
	prometheus.MustRegister(Table0ViewableCount)
	prometheus.MustRegister(Table1ViewableCount)
	prometheus.MustRegister(Table2ViewableCount)
	prometheus.MustRegister(Table3ViewableCount)
	prometheus.MustRegister(Table4ViewableCount)
	prometheus.MustRegister(Table5ViewableCount)
	prometheus.MustRegister(Table6ViewableCount)
	prometheus.MustRegister(LibraryTotalFileCount)
	prometheus.MustRegister(LibraryTotalTranscodeCount)
	prometheus.MustRegister(LibraryTotalHealthCheckCount)
	prometheus.MustRegister(LibrarySizeDiff)
	prometheus.MustRegister(LibraryTranscodeStatus)
	prometheus.MustRegister(LibraryHealth)
	prometheus.MustRegister(LibraryVideoCodec)
	prometheus.MustRegister(LibraryVideoContainer)
	prometheus.MustRegister(LibraryVideoResolution)
	prometheus.MustRegister(LibraryAudioCodec)
	prometheus.MustRegister(LibraryAudioContainer)
}
