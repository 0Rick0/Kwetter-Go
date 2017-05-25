package EndPoints

import (
	"github.com/emicklei/go-restful"
	"../types"
)

func (sc *ServiceContainer) DefineKwetEndpoints(container *restful.Container)  {
	ws := new(restful.WebService)
	ws.Path("/kwets").
	Doc("Access to the kwets resource").
	Consumes(restful.MIME_JSON, restful.MIME_XML).
	Produces(restful.MIME_JSON, restful.MIME_XML)

	//ws.Route(ws.GET("/{username}").
	//	To(sc.getUser).
	//	Doc("Get a user object").
	//	Operation("getUser").
	//	Param(ws.
	//		PathParameter("username", "The username of the user to get").
	//		DataType("string")).
	//	Writes(types.User{}))

	ws.Route(ws.GET("/ofuser/{username}").
		 To(sc.getKwetsOfUser).
		  Doc("Get all kwets of an user").
		  Operation("getKwetsOfUser").
		  Param(ws.
	 		PathParameter("username", "The username of the user").DataType("string")).
		  Writes([]types.Kwet{}))

	container.Add(ws)
}

func (sc *ServiceContainer) getKwetsOfUser(req *restful.Request, resp *restful.Response)  {
	username := req.PathParameter("username")

	var kwets *[]types.Kwet = sc.Service.GetKwetsOfUser(username,0,100)//todo parameter count offset
	// check if the user is found
	if kwets == nil{
		//if not, report an service error
		resp.WriteErrorString(404, "User not found")
	}else{
		resp.WriteEntity(&kwets)
	}
}
