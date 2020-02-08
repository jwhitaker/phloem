package phloem

// KafkaConfiguration defines required configuration elements to hook into kafka
type KafkaConfiguration struct {
	BootstrapServers string `default:"localhost"`
}
