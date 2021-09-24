package streaming

/**
 * @author : Donald Trieu
 * @created : 9/24/21, Friday
**/
type StreamChannelList struct {
	StreamChannel []StreamChannel `xml:"StreamingChannel" json:"StreamingChannel"`
}

type StreamChannel struct {
	ID                 int        `xml:"id" json:"id"`
	ChannelName        string     `xml:"channelName" json:"channelName"`
	Enabled            bool       `xml:"enabled" json:"enabled"`
	Transport          *Transport `xml:"Transport,omitempty" json:"Transport,omitempty"`
	Video              *Video     `xml:"Video,omitempty" json:"Video,omitempty"`
	Audio              *Audio     `xml:"Audio,omitempty" json:"Audio,omitempty"`
	EnableCABAC        bool       `xml:"enableCABAC,omitempty" json:"enableCABAC,omitempty"`
	SubStreamRecStatus bool       `xml:"subStreamRecStatus,omitempty" json:"subStreamRecStatus,omitempty"`
}

type Transport struct {
	MaxPacketSize            int `xml:"maxPacketSize" json:"maxPacketSize"`
	AudioPacketLength        int `xml:"audioPacketLength,omitempty" json:"audioPacketLength,omitempty"`
	AudioInboundPacketLength int `xml:"audioInboundPacketLength,omitempty" json:"audioInboundPacketLength,omitempty"`
	AudioInboundPortNo       int `xml:"audioInboundPortNo,omitempty" json:"audioInboundPortNo,omitempty"`
	VideoSourcePortNo        int `xml:"videoSourcePortNo,omitempty" json:"videoSourcePortNo,omitempty"`
	AudioSourcePortNo        int `xml:"audioSourcePortNo,omitempty" json:"audioSourcePortNo,omitempty"`
	ControlProtocolList      []struct {
		ControlProtocol struct {
			StreamingTransport string `xml:"streamingTransport" json:"streamingTransport"`
		} `xml:"ControlProtocol" json:"ControlProtocol"`
	} `xml:"ControlProtocolList" json:"ControlProtocolList"`
	Unicast struct {
		Enabled          bool   `xml:"enabled" json:"enabled"`
		InterfaceID      string `xml:"interfaceID,omitempty" json:"interfaceID,omitempty"`
		RtpTransportType string `xml:"rtpTransportType" json:"rtpTransportType"`
	} `xml:"Unicast" json:"Unicast"`
	Multicast struct {
		Enabled              bool   `xml:"enabled" json:"enabled"`
		UserTriggerThreshold int    `xml:"userTriggerThreshold,omitempty" json:"userTriggerThreshold,omitempty"`
		DestIPAddress        string `xml:"destIPAddress" json:"destIPAddress"`
		VideoDestPortNo      int    `xml:"videoDestPortNo" json:"videoDestPortNo"`
		AudioDestPortNo      int    `xml:"audioDestPortNo" json:"audioDestPortNo"`
		DestIPv6Address      string `xml:"destIPv6Address,omitempty" json:"destIPv6Address,omitempty"`
		Ttl                  int    `xml:"ttl,omitempty" json:"ttl,omitempty"`
	} `xml:"Multicast" json:"Multicast"`
	Security struct {
		Enabled         bool   `xml:"enabled" json:"enabled"`
		CertificateType string `xml:"certificateType" json:"certificateType"`
	} `xml:"Security" json:"Security"`
}

type Video struct {
	Enabled                 bool   `xml:"enabled" json:"enabled"`
	VideoInputChannelID     string `xml:"videoInputChannelID" json:"videoInputChannelID"`
	VideoCodecType          string `xml:"videoCodecType" json:"videoCodecType"`
	VideoScanType           string `xml:"videoScanType" json:"videoScanType"`
	VideoResolutionWidth    int    `xml:"videoResolutionWidth" json:"videoResolutionWidth"`
	VideoResolutionHeight   int    `xml:"videoResolutionHeight" json:"videoResolutionHeight"`
	VideoPositionX          int    `xml:"videoPositionX,omitempty" json:"videoPositionX,omitempty"`
	VideoPositionY          int    `xml:"videoPositionY,omitempty" json:"videoPositionY,omitempty"`
	VideoQualityControlType string `xml:"videoQualityControlType" json:"videoQualityControlType"`
	ConstantBitRate         int    `xml:"constantBitRate" json:"constantBitRate"`
	FixedQuality            int    `xml:"fixedQuality" json:"fixedQuality"`
	VbrUpperCap             int    `xml:"vbrUpperCap,omitempty" json:"vbrUpperCap,omitempty"`
	VbrLowerCap             int    `xml:"vbrLowerCap,omitempty" json:"vbrLowerCap,omitempty"`
	MaxFrameRate            int    `xml:"maxFrameRate" json:"maxFrameRate"`
	KeyFrameInterval        int    `xml:"keyFrameInterval" json:"keyFrameInterval"`
	RotationDegree          int    `xml:"rotationDegree,omitempty" json:"rotationDegree,omitempty"`
	MirrorEnabled           bool   `xml:"mirrorEnabled,omitempty" json:"mirrorEnabled,omitempty"`
	SnapShotImageType       string `xml:"snapShotImageType,omitempty" json:"snapShotImageType,omitempty"`
	Mpeg4Profile            string `xml:"Mpeg4Profile,omitempty" json:"Mpeg4Profile,omitempty"`
	H264Profile             string `xml:"H264Profile,omitempty" json:"H264Profile,omitempty"`
	GovLength               int    `xml:"GovLength,omitempty" json:"GovLength,omitempty"`
	SVC                     struct {
		Enabled bool `xml:"enabled,omitempty" json:"enabled,omitempty"`
	} `xml:"SVC,omitempty" json:"SVC,omitempty"`
	PacketType  string `xml:"PacketType,omitempty" json:"PacketType,omitempty"`
	Smoothing   int    `xml:"smoothing,omitempty" json:"smoothing,omitempty"`
	H265Profile string `xml:"H265Profile,omitempty" json:"H265Profile,omitempty"`
	SmartCodec  struct {
		Enabled bool `xml:"enabled,omitempty" json:"enabled,omitempty"`
	} `xml:"SmartCodec,omitempty" json:"SmartCodec,omitempty"`
	VbrAverageCap int `xml:"vbrAverageCap,omitempty" json:"vbrAverageCap,omitempty"`
}

type Audio struct {
	Enabled                     bool    `xml:"enabled,omitempty" json:"enabled,omitempty"`
	AudioInputChannelID         string  `xml:"audioInputChannelID,omitempty" json:"audioInputChannelID,omitempty"`
	AudioCompressionType        string  `xml:"audioCompressionType,omitempty" json:"audioCompressionType,omitempty"`
	AudioInboundCompressionType string  `xml:"audioInboundCompressionType,omitempty" json:"audioInboundCompressionType,omitempty"`
	AudioBitRate                int     `xml:"audioBitRate,omitempty" json:"audioBitRate,omitempty"`
	AudioSamplingRate           float32 `xml:"audioSamplingRate,omitempty" json:"audioSamplingRate,omitempty"`
	AudioResolution             int     `xml:"audioResolution,omitempty" json:"audioResolution,omitempty"`
}
