package types

type User struct {
	Id int		`json:"id" gorm:"primary_key"`
	Username string	`json:"username" sql:"unique"`
	Password string	`json:"password"`
	FullName string	`json:"full_name"`
	Location string	`json:"location"`
	Website  string	`json:"website"`
	Biography string	`json:"biography"`
}

type UserFollowing struct {
	Follower User `json:"follower" gorm:"ForeignKey:FollowerID"`
	FollowerID int `json:"-"`
	Following User `json:"following" gorm:"ForeignKey:FollowingID"`
	FollowingID int `json:"-"`
}

type Tag struct {
	KwetID int
	Text string
}

type Kwet struct {
	Id int		`json:"id" gorm:"primary_key"`
	Text string	`json:"text"`
	Sender User	`json:"sender" gorm:"ForeignKey:SenderID"`
	SenderID int	`json:"-"`
	Tags []Tag 	`json:"tags2" gorm:"ForeignKey:KwetID"`
	TagsString []string 	`json:"tags" gorm:"-"`
	Mentions []User        `json:"mentions" gorm:"many2many:mentions"`
	LikedBy	[]User        `json:"liked_by" gorm:"many2many:liked_by"`
}

func (User) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "A user object",
		"id":	    "The id of the user",
		"username": "The username of the user",
		"password": "The password of the user",
		"full_name":"The full name of the user",
		"location": "The home location of the user",
		"website":  "The website of the user",
		"biography":"The biography of the user",
		"following":"A list of usernames that the user is following",
		"followers":"A list of usernames that follow the user",
	}
}

type PostKwet struct {
	Content string `json:"content"`
}

func (Kwet) SwaggerDoc() map[string]string {
	return map[string]string{
		"":		"A kwet object",
		"id":		"The id of the kwet",
		"text":		"The text of the kwet",
		"sender":	"The sending user of the kwet",
		"tags":		"A list of tags belonging to the kwet",
		"mentions":	"A list of users mentioned in the kwet",
		"liked_by":	"A list of users that like the kwet",
	}
}