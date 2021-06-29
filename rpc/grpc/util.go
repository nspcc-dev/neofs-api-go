package neofsgrpc

import (
	"fmt"
)

const methodNameFmt = "/%s/%s"

func toMethodName(svc, mtd string) string {
	return fmt.Sprintf(methodNameFmt, svc, mtd)
}
