// Code generated by gen-fields; DO NOT EDIT.

package transmission

// TorrentField is a field of Torrent
type TorrentField string

const (
	TorrentFieldID                       TorrentField = "id"
	TorrentFieldHash                     TorrentField = "hashString"
	TorrentFieldName                     TorrentField = "name"
	TorrentFieldStatus                   TorrentField = "status"
	TorrentFieldCreator                  TorrentField = "creator"
	TorrentFieldComment                  TorrentField = "comment"
	TorrentFieldETA                      TorrentField = "eta"
	TorrentFieldIdleETA                  TorrentField = "etaIdle"
	TorrentFieldErrorType                TorrentField = "error"
	TorrentFieldError                    TorrentField = "errorString"
	TorrentFieldFile                     TorrentField = "torrentFile"
	TorrentFieldMagnetLink               TorrentField = "magnetLink"
	TorrentFieldDownloadDirectory        TorrentField = "downloadDir"
	TorrentFieldCreatedAt                TorrentField = "dateCreated"
	TorrentFieldAddedAt                  TorrentField = "addedDate"
	TorrentFieldStartedAt                TorrentField = "startDate"
	TorrentFieldLastActiveAt             TorrentField = "activityDate"
	TorrentFieldDoneAt                   TorrentField = "doneDate"
	TorrentFieldCanManuallyAnnounceAt    TorrentField = "manualAnnounceTime"
	TorrentFieldDownloadRate             TorrentField = "rateDownload"
	TorrentFieldUploadRate               TorrentField = "rateUpload"
	TorrentFieldDownloadRateLimit        TorrentField = "downloadLimit"
	TorrentFieldDownloadRateLimitEnabled TorrentField = "downloadLimited"
	TorrentFieldUploadRateLimit          TorrentField = "uploadLimit"
	TorrentFieldUploadRateLimited        TorrentField = "uploadLimited"
	TorrentFieldHonorSessionLimits       TorrentField = "honorsSessionLimits"
	TorrentFieldDownloadedTotal          TorrentField = "downloadedEver"
	TorrentFieldUploadedTotal            TorrentField = "uploadedEver"
	TorrentFieldCorruptedTotal           TorrentField = "corruptEver"
	TorrentFieldPriority                 TorrentField = "bandwidthPriority"
	TorrentFieldPositionInQueue          TorrentField = "queuePosition"
	TorrentFieldIdleSeedingLimit         TorrentField = "seedIdleLimit"
	TorrentFieldIdleSeedingLimitMode     TorrentField = "seedIdleMode"
	TorrentFieldUploadRatioLimit         TorrentField = "seedRatioLimit"
	TorrentFieldUploadRatioLimitMode     TorrentField = "seedRatioMode"
	TorrentFieldUploadRatio              TorrentField = "uploadRatio"
	TorrentFieldDownloadingFor           TorrentField = "secondsDownloading"
	TorrentFieldSeedingFor               TorrentField = "secondsSeeding"
	TorrentFieldTotalSize                TorrentField = "totalSize"
	TorrentFieldWantedSize               TorrentField = "sizeWhenDone"
	TorrentFieldWantedAvailable          TorrentField = "desiredAvailable"
	TorrentFieldWantedLeft               TorrentField = "leftUntilDone"
	TorrentFieldUncheckedSize            TorrentField = "haveUnchecked"
	TorrentFieldValidSize                TorrentField = "haveValid"
	TorrentFieldDataDone                 TorrentField = "percentDone"
	TorrentFieldDataChecked              TorrentField = "recheckProgress"
	TorrentFieldMetadataDone             TorrentField = "metadataPercentComplete"
	TorrentFieldIsFinished               TorrentField = "isFinished"
	TorrentFieldIsPrivate                TorrentField = "isPrivate"
	TorrentFieldIsStalled                TorrentField = "isStalled"
	TorrentFieldPeerLimit                TorrentField = "peer-limit"
	TorrentFieldConnectedPeers           TorrentField = "peersConnected"
	TorrentFieldPeersGettingFromUs       TorrentField = "peersGettingFromUs"
	TorrentFieldPeersSendingToUs         TorrentField = "peersSendingToUs"
	TorrentFieldPeers                    TorrentField = "peers"
	TorrentFieldPeersFrom                TorrentField = "peersFrom"
	TorrentFieldWebSeedsSendingToUs      TorrentField = "webseedsSendingToUs"
	TorrentFieldWebSeeds                 TorrentField = "webseeds"
	TorrentFieldWanted                   TorrentField = "wanted"
	TorrentFieldFiles                    TorrentField = "files"
	TorrentFieldFileStats                TorrentField = "fileStats"
	TorrentFieldPriorities               TorrentField = "priorities"
	TorrentFieldPieceCount               TorrentField = "pieceCount"
	TorrentFieldPieceSize                TorrentField = "pieceSize"
	TorrentFieldPieces                   TorrentField = "pieces"
	TorrentFieldTrackers                 TorrentField = "trackers"
	TorrentFieldTrackerStats             TorrentField = "trackerStats"
)

var allTorrentFields = []TorrentField{
	TorrentFieldID,
	TorrentFieldHash,
	TorrentFieldName,
	TorrentFieldStatus,
	TorrentFieldCreator,
	TorrentFieldComment,
	TorrentFieldETA,
	TorrentFieldIdleETA,
	TorrentFieldErrorType,
	TorrentFieldError,
	TorrentFieldFile,
	TorrentFieldMagnetLink,
	TorrentFieldDownloadDirectory,
	TorrentFieldCreatedAt,
	TorrentFieldAddedAt,
	TorrentFieldStartedAt,
	TorrentFieldLastActiveAt,
	TorrentFieldDoneAt,
	TorrentFieldCanManuallyAnnounceAt,
	TorrentFieldDownloadRate,
	TorrentFieldUploadRate,
	TorrentFieldDownloadRateLimit,
	TorrentFieldDownloadRateLimitEnabled,
	TorrentFieldUploadRateLimit,
	TorrentFieldUploadRateLimited,
	TorrentFieldHonorSessionLimits,
	TorrentFieldDownloadedTotal,
	TorrentFieldUploadedTotal,
	TorrentFieldCorruptedTotal,
	TorrentFieldPriority,
	TorrentFieldPositionInQueue,
	TorrentFieldIdleSeedingLimit,
	TorrentFieldIdleSeedingLimitMode,
	TorrentFieldUploadRatioLimit,
	TorrentFieldUploadRatioLimitMode,
	TorrentFieldUploadRatio,
	TorrentFieldDownloadingFor,
	TorrentFieldSeedingFor,
	TorrentFieldTotalSize,
	TorrentFieldWantedSize,
	TorrentFieldWantedAvailable,
	TorrentFieldWantedLeft,
	TorrentFieldUncheckedSize,
	TorrentFieldValidSize,
	TorrentFieldDataDone,
	TorrentFieldDataChecked,
	TorrentFieldMetadataDone,
	TorrentFieldIsFinished,
	TorrentFieldIsPrivate,
	TorrentFieldIsStalled,
	TorrentFieldPeerLimit,
	TorrentFieldConnectedPeers,
	TorrentFieldPeersGettingFromUs,
	TorrentFieldPeersSendingToUs,
	TorrentFieldPeers,
	TorrentFieldPeersFrom,
	TorrentFieldWebSeedsSendingToUs,
	TorrentFieldWebSeeds,
	TorrentFieldWanted,
	TorrentFieldFiles,
	TorrentFieldFileStats,
	TorrentFieldPriorities,
	TorrentFieldPieceCount,
	TorrentFieldPieceSize,
	TorrentFieldPieces,
	TorrentFieldTrackers,
	TorrentFieldTrackerStats,
}
