package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
	ID        string               `bson:"_id,omitempty"`
	Name      string               `bson:"name"`
	Team      []primitive.ObjectID `bson:"team"`
	OwnerID   primitive.ObjectID   `bson:"owner_id"`
	ProjectID string               `bson:"project_id"`
	CreatedAt primitive.DateTime   `bson:"created_at"`
}

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password"`
	IsConfirmed  bool               `bson:"is_confirmed"`
	AuthSource   string             `bson:"auth_source"`
	FullName     string             `bson:"name"`
	Tokens       Tokens             `bson:"tokens"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
}

type Tokens struct {
	ConfirmToken string `bson:"confirm_token"`
	OAuth2Token  string `bson:"oauth2_token"`
	ResetToken   string `bson:"reset_token"`
	RefreshToken string `bson:"refresh_token"`
}

//type ProjectRole string
//
//const (
//	RoleAdmin    ProjectRole = "admin"
//	RoleMember   ProjectRole = "member"
//	RoleObserver ProjectRole = "observer"
//)
//
//type Project struct {
//	ID        primitive.ObjectID                 `bson:"_id,omitempty"`
//	Name      string                             `bson:"name"`
//	OwnerID   primitive.ObjectID                 `bson:"owner_id"`
//	ProjectID string                             `bson:"project_id"`
//	CreatedAt primitive.DateTime                 `bson:"createdAt"`
//	Team      map[primitive.ObjectID]ProjectRole `bson:"team"`
//}
