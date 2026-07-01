package camera

const TypeSupportedAudioRecordingConfiguration = "207"

//goland:noinspection ALL
const (
	AudioRecordingCodecTypeAACLC  = 0
	AudioRecordingCodecTypeAACELD = 1

	AudioRecordingSampleRate8Khz  = 0
	AudioRecordingSampleRate16Khz = 1
	AudioRecordingSampleRate24Khz = 2
	AudioRecordingSampleRate32Khz = 3
	AudioRecordingSampleRate44Khz = 4
	AudioRecordingSampleRate48Khz = 5
)

type SupportedAudioRecordingConfiguration struct {
	CodecConfigs []AudioRecordingCodecConfiguration `tlv8:"1"`
}

type AudioRecordingCodecConfiguration struct {
	CodecType   byte                            `tlv8:"1"`
	CodecParams []AudioRecordingCodecParameters `tlv8:"2"`
}

type AudioRecordingCodecParameters struct {
	Channels    uint8  `tlv8:"1"`
	BitrateMode []byte `tlv8:"2"`
	SampleRate  []byte `tlv8:"3"`
}
