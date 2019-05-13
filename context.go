package context

//information required for the endpoint
type Context struct {
	R            gin.Router
	MongoSession *mgo.Session
}

func New() *Context {
	// create a session to mongodb
	// create a gin router.
	//
}
