package defs

// CreateUser、Login requests 用户请求数据结构
type SignedRequest struct {
	Username string `form:"user_name" json:"user_name" binding:"required"`
	Pwd      string `form:"pwd" json:"pwd"`
}

// UserInfo request
type UserInfoRequest struct {
	Username string `form:"username" binding:"required"`
}

// DeleteUser request
type DeleteUserRequest struct {
	Username string `form:"username" binding:"required"`
}

// CreateUser、Login response
type SignedResponse struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
	Character string `json:"character"`
}

// GetUserInfo response
type UserInfoResponse struct {
	Success   bool   `json:"success"`
	Character string `json:"character"`
}

// DeleteUser response
type DeleteUserResponse struct {
	Success bool `json:"success"`
}
