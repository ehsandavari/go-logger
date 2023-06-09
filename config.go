package logger

type sConfig struct {
	isDevelopment     bool
	disableStacktrace bool
	disableStdout     bool
	level             string
	serviceId         int
	serviceName       string
	serviceNamespace  string
	serviceInstanceId string
	serviceVersion    string
	serviceMode       string
	serviceCommitId   string
}
