package log

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

const IpHeader = "X-FORWARDED_FOR"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
}

func WithRequest(logger string, r *http.Request) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"Logger":   logger,
		"ClientIp": getIpFromRequest(r),
	})
}

func getIpFromRequest(r *http.Request) string {
	forwardedIp := r.Header.Get(IpHeader)
	if forwardedIp != "" {
		return forwardedIp
	}

	return r.RemoteAddr
}
