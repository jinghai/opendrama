module opendrama

go 1.21

require (
	github.com/jinghai/opendrama/pkg/ai
	github.com/jinghai/opendrama/pkg/config
	github.com/jinghai/opendrama/pkg/image
	github.com/jinghai/opendrama/pkg/logger
	github.com/jinghai/opendrama/pkg/response
	github.com/jinghai/opendrama/pkg/utils
	github.com/jinghai/opendrama/pkg/video
)

require (
	github.com/gin-gonic/gin@v1.9.1
	gorm.io/gorm@v1.25.5
	github.com/jinghai/opendrama/domain
	github.com/jinghai/opendrama/infrastructure
	github.com/jinghai/opendrama/application
	github.com/jinghai/opendrama/api
	github.com/jinghai/opendrama/cmd
)