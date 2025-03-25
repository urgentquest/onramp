package onramp

import (
	"github.com/go-i2p/logger"
)

var (
	i2plog *logger.Logger
)

func init() {
	i2plog = logger.GetGoI2PLogger()
}
