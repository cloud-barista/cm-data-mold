package routes

import (
	"github.com/cloud-barista/cm-data-mold/websrc/controllers"
	"github.com/labstack/echo/v4"
)

func MigrationRoutes(g *echo.Group) {
	// Migration From On-premise (Linux, Windows) to Object Storage
	MigrationFromOnpremiseToObjectStorage(g)

	// Migration MySQL to MySQL
	MigrationMySQL(g)

	// Migration From Object Storage to Other Object Storage
	MigrationFromS3Routes(g)
	MigrationFromGCPRoutes(g)
	MigrationFromNCPRoutes(g)

	// Migration No-SQL to the other No-SQL
	MigrationNoSQLRoutes(g)
}

func MigrationFromOnpremiseToObjectStorage(g *echo.Group) {
	g.GET("/linux/s3", controllers.MigrationLinuxToS3GetHandler)
	g.POST("/linux/s3", controllers.MigrationLinuxToS3PostHandler)

	g.GET("/linux/gcp", controllers.MigrationLinuxToGCPGetHandler)
	g.POST("/linux/gcp", controllers.MigrationLinuxToGCPPostHandler)

	g.GET("/linux/ncp", controllers.MigrationLinuxToNCPGetHandler)
	g.POST("/linux/ncp", controllers.MigrationLinuxToNCPPostHandler)

	g.GET("/windows/s3", controllers.MigrationWindowsToS3GetHandler)
	g.POST("/windows/s3", controllers.MigrationWindowsToS3PostHandler)

	g.GET("/windows/gcp", controllers.MigrationWindowsToGCPGetHandler)
	g.POST("/windows/gcp", controllers.MigrationWindowsToGCPPostHandler)

	g.GET("/windows/ncp", controllers.MigrationWindowsToNCPGetHandler)
	g.POST("/windows/ncp", controllers.MigrationWindowsToNCPPostHandler)
}

func MigrationMySQL(g *echo.Group) {
	g.GET("/mysql", controllers.MigrationMySQLGetHandler)
	g.POST("/mysql", controllers.MigrationMySQLPostHandler)
}

func MigrationFromS3Routes(g *echo.Group) {
	g.GET("/s3/linux", controllers.MigrationS3ToLinuxGetHandler)
	g.POST("/s3/linux", controllers.MigrationS3ToLinuxPostHandler)

	g.GET("/s3/windows", controllers.MigrationS3ToWindowsGetHandler)
	g.POST("/s3/windows", controllers.MigrationS3ToWindowsPostHandler)

	g.GET("/s3/gcp", controllers.MigrationS3ToGCPGetHandler)
	g.POST("/s3/gcp", controllers.MigrationS3ToGCPPostHandler)

	g.GET("/s3/ncp", controllers.MigrationS3ToNCPGetHandler)
	g.POST("/s3/ncp", controllers.MigrationS3ToNCPPostHandler)
}

func MigrationFromGCPRoutes(g *echo.Group) {
	g.GET("/gcp/linux", controllers.MigrationGCPToLinuxGetHandler)
	g.POST("/gcp/linux", controllers.MigrationGCPToLinuxPostHandler)

	g.GET("/gcp/windows", controllers.MigrationGCPToWindowsGetHandler)
	g.POST("/gcp/windows", controllers.MigrationGCPToWindowsPostHandler)

	g.GET("/gcp/s3", controllers.MigrationGCPToS3GetHandler)
	g.POST("/gcp/s3", controllers.MigrationGCPToS3PostHandler)

	g.GET("/gcp/ncp", controllers.MigrationGCPToNCPGetHandler)
	g.POST("/gcp/ncp", controllers.MigrationGCPToNCPPostHandler)
}

func MigrationFromNCPRoutes(g *echo.Group) {
	g.GET("/ncp/linux", controllers.MigrationNCPToLinuxGetHandler)
	g.POST("/ncp/linux", controllers.MigrationNCPToLinuxPostHandler)

	g.GET("/ncp/windows", controllers.MigrationNCPToWindowsGetHandler)
	g.POST("/ncp/windows", controllers.MigrationNCPToWindowsPostHandler)

	g.GET("/ncp/s3", controllers.MigrationNCPToS3GetHandler)
	g.POST("/ncp/s3", controllers.MigrationNCPToS3PostHandler)

	g.GET("/ncp/gcp", controllers.MigrationNCPToGCPGetHandler)
	g.POST("/ncp/gcp", controllers.MigrationNCPToGCPPostHandler)
}

func MigrationNoSQLRoutes(g *echo.Group) {
	g.GET("/dynamodb/firestore", controllers.MigrationDynamoDBToFirestoreGetHandler)
	g.POST("/dynamodb/firestore", controllers.MigrationDynamoDBToFirestorePostHandler)

	g.GET("/dynamodb/mongodb", controllers.MigrationDynamoDBToMongoDBGetHandler)
	g.POST("/dynamodb/mongodb", controllers.MigrationDynamoDBToMongoDBPostHandler)

	g.GET("/firestore/dynamodb", controllers.MigrationFirestoreToDynamoDBGetHandler)
	g.POST("/firestore/dynamodb", controllers.MigrationFirestoreToDynamoDBPostHandler)

	g.GET("/firestore/mongodb", controllers.MigrationFirestoreToMongoDBGetHandler)
	g.POST("/firestore/mongodb", controllers.MigrationFirestoreToMongoDBPostHandler)

	g.GET("/mongodb/dynamodb", controllers.MigrationMongoDBToDynamoDBGetHandler)
	g.POST("/mongodb/dynamodb", controllers.MigrationMongoDBToDynamoDBPostHandler)

	g.GET("/mongodb/firestore", controllers.MigrationMongoDBToFirestoreGetHandler)
	g.POST("/mongodb/firestore", controllers.MigrationMongoDBToFirestorePostHandler)
}
