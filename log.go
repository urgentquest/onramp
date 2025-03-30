package onramp

import (
	"github.com/go-i2p/logger"
)

var (
	i2pLogger *logger.Logger
)

func init() {
	i2pLogger = logger.GetGoI2PLogger()
}
