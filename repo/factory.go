package content

// import (
// 	"project-art-museum/business/content"
// 	"project-art-museum/util"
// )

// //RepositoryFactory Will return business.content.Repository based on active database connection
// func RepositoryFactory(dbCon *util.DatabaseConnection) content.Repository {
// 	var contentRepo content.Repository
// 	if dbCon.Driver == util.MySQL {
// 		contentRepo = NewMySQLRepository(dbCon.PostgreSQL)
// 	} else if dbCon.Driver == util.MongoDB {
// 		contentRepo = NewMongoDBRepository(dbCon.MongoDB)
// 	}
// 	return contentRepo
// }
// func RepositoryFactory(dbCon *util.DatabaseConnection) content.Repository {
// 	var contentRepo content.Repository
// 	if dbCon.Driver == util.MySQL {
// 		contentRepo = NewMySQLRepository(dbCon.PostgreSQL)
// 	// } else if dbCon.Driver == util.MongoDB {
// 	// 	contentRepo = NewMongoDBRepository(dbCon.MongoDB)
// 	}

// 	return contentRepo
// }
