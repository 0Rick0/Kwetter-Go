package EndPoints

import (
	"log"
	"github.com/emicklei/go-restful"
	"../types"
)

func (sc *ServiceContainer) DefineUserEndpoints(container *restful.Container)  {
	ws := new(restful.WebService)
	ws.Path("/users").
	Doc("Access to the user resource").
	Consumes(restful.MIME_JSON, restful.MIME_XML).
	Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{username}").
		To(sc.getUser).
		Doc("Get a user object").
		Operation("getUser").
		Param(ws.
			PathParameter("username", "The username of the user to get").
			DataType("string")).
		Writes(types.User{}))

	ws.Route(ws.GET("/{username}/following").
		To(sc.getFollowing).
		Doc("Get the users the user is following").
		Operation("getFollowing").
		Param(ws.
			 PathParameter("username", "The username of the user of whom to get the followings").
			 DataType("string")).
		Writes([]string{}))

	ws.Route(ws.GET("/{username}/followers").
		To(sc.getFollowers).
		Doc("Get the users that follow the user").
		Operation("getFollowers").
		Param(ws.
			PathParameter("username", "The username of the user of whom to get the followers").
			DataType("string")).
		Writes([]string{}))

	ws.Route(ws.POST("/").To(sc.createUser).
		Doc("Create a new User").
		Operation("createUser").
		Reads(types.User{}).Writes(types.User{}))

	ws.Route(ws.DELETE("/{username}").To(sc.deleteUser).
		Doc("Delete a user from the database, returning the deleted user").
		Operation("deleteUser").
		Param(ws.PathParameter("username", "The username of the user to delete").
		DataType("string")).
		Writes(types.User{}))

	u := new(types.User)
	u.Username = "Test"
	u.Biography = "aapje"

	print(sc.Service.AddUser(*u))

	container.Add(ws)
}

func (sc *ServiceContainer) getUser(req *restful.Request, resp *restful.Response) {
	username := req.PathParameter("username")

	log.Printf("Get user %s", username)

	user := sc.Service.GetUserByUsername(username)

	// check if the user is found
	if user == nil{
		//if not, report an service error
		resp.WriteErrorString(404, "User not found")
	}else{
		// else, write the entity
		resp.WriteEntity(&user)
	}
}

func (sc *ServiceContainer) createUser(req *restful.Request, resp *restful.Response)  {
	user := new(types.User)
	err := req.ReadEntity(user)
	if err != nil {
		resp.WriteError(400, err)
		return
	}
	if !sc.Service.AddUser(*user){
		resp.WriteErrorString(400, "Username already exists")
		return
	}
	//get a username with all automatic values(like id)
	user = sc.Service.GetUserByUsername(user.Username)
	resp.WriteEntity(user)
}

func (sc *ServiceContainer) deleteUser(req *restful.Request, resp *restful.Response)  {
	username := req.PathParameter("username")
	log.Printf("Deleting user %sc", username)

	user := sc.Service.GetUserByUsername(username)
	if user == nil {
		resp.WriteErrorString(404, "User not found")
		return
	}
	if !sc.Service.RemoveUser(*user){
		resp.WriteErrorString(500, "Failed to remove user")
		return
	}
	resp.WriteEntity(user)
}

func (sc *ServiceContainer) getFollowing(req *restful.Request, resp *restful.Response)  {
	username := req.PathParameter("username")

	log.Printf("Get user %s", username)

	var followers *[]string = sc.Service.GetFollowingByUsername(username)
	// check if the user is found
	if followers == nil{
		//if not, report an service error
		resp.WriteErrorString(404, "User not found")
	}else{
		resp.WriteEntity(&followers)
	}
}

func (sc *ServiceContainer) getFollowers(req *restful.Request, resp *restful.Response)  {
	username := req.PathParameter("username")

	log.Printf("Get user %s", username)

	var followers *[]string = sc.Service.GetFollowersByUsername(username)
	// check if the user is found
	if followers == nil{
		//if not, report an service error
		resp.WriteErrorString(404, "User not found")
	}else{
		resp.WriteEntity(&followers)
	}
}
