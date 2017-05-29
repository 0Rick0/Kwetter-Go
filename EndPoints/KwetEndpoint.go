package EndPoints

import (
	"github.com/emicklei/go-restful"
	"../types"
	"regexp"
)

var re_hashtag *regexp.Regexp
var re_mention *regexp.Regexp


func (sc *ServiceContainer) DefineKwetEndpoints(container *restful.Container)  {
	re_hashtag = regexp.MustCompile("#(\\S+)")
	re_mention = regexp.MustCompile("@(\\S+)")

	ws := new(restful.WebService)
	ws.Path("/kwets").
	Doc("Access to the kwets resource").
	Consumes(restful.MIME_JSON, restful.MIME_XML).
	Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/ofuser/{username}").
		 To(sc.getKwetsOfUser).
		  Doc("Get all kwets of an user").
		  Operation("getKwetsOfUser").
		  Param(ws.
	 		PathParameter("username", "The username of the user").DataType("string")).
		  Writes([]types.Kwet{}))

	ws.Route(ws.GET("/offollowed/{username}").
		 To(sc.getKwetsOfFollowed).
		 Doc("Get all kwets of the user the user follows").
		 Operation("getKwetsOfFollowed").
		 Param(ws.
			PathParameter("username", "The username of the user").DataType("string")).
		 Writes([]types.Kwet{}))

	ws.Route(ws.POST("/new/{username}").
		 To(sc.postKwet).
		 Doc("Post a new kwet").
		 Operation("postKwet").
		 Param(ws.PathParameter("username", "The username").DataType("string")).
		 Param(ws.BodyParameter("content", "The kwet").DataType("postkwet")).
		 Reads(&types.PostKwet{}).Writes(&types.Kwet{}))

	container.Add(ws)
}

func (sc *ServiceContainer) getKwetsOfUser(req *restful.Request, resp *restful.Response)  {
	username := req.PathParameter("username")

	var kwets *[]types.Kwet = sc.Service.GetKwetsOfUser(username,100,0)//todo parameter count offset
	// check if the user is found
	if kwets == nil{
		//if not, report an service error
		resp.WriteErrorString(404, "User not found")
	}else{
		resp.WriteEntity(&kwets)
	}
}

func (sc *ServiceContainer) getKwetsOfFollowed(req *restful.Request, resp *restful.Response)  {
	username := req.PathParameter("username")

	var kwets *[]types.Kwet = sc.Service.GetKwetsOfFollowed(username, 100, 0)

	if kwets == nil {
		resp.WriteErrorString(404, "User not found")
	}else {
		resp.WriteEntity(&kwets)
	}
}

func (sc *ServiceContainer) postKwet(req *restful.Request, resp *restful.Response)  {
	var username = req.PathParameter("username")
	var kwet types.PostKwet
	req.ReadEntity(&kwet)

	var tags = re_hashtag.FindAllStringSubmatch(kwet.Content, -1) //search for all tags
	var tags_string []string
	if tags != nil {
		tags_string = make([]string, 0, len(tags))
		for _, tag := range tags {
			tags_string = append(tags_string, tag[1])
		}
	}else {
		tags_string = make([]string, 0)
	}

	var mentions = re_mention.FindAllStringSubmatch(kwet.Content, -1) //search for all mentions
	var mentions_string []string
	if mentions != nil {
		mentions_string = make([]string, 0, len(mentions))
		for _, mention := range mentions{
			mentions_string = append(mentions_string, mention[1])
		}
	}else {
		mentions_string = make([]string, 0)
	}

	postedKwet := sc.Service.PostKwet(username, kwet.Content, tags_string, mentions_string)

	if postedKwet != nil {
		resp.WriteEntity(&postedKwet)
	}else{
		resp.WriteErrorString(503, "Failed to post kwet")
	}
}
