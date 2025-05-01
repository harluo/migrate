package argument

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newId,
		newPattern,
	).Build().Apply()
}
