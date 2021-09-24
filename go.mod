module github.com/catbugdemo/project_order

go 1.12

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/gomodule/redigo v1.8.5
	github.com/google/btree v1.0.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.9.0
	github.com/stretchr/testify v1.7.0
)

replace github.com/stretchr/testify v1.7.0 => github.com/stretchr/testify v1.4.0
