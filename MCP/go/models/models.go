package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Tokeninfo represents the Tokeninfo schema from the OpenAPI specification
type Tokeninfo struct {
	Expires_in int `json:"expires_in,omitempty"` // The expiry time of the token, as number of seconds left until expiry.
	Issued_to string `json:"issued_to,omitempty"` // To whom was the token issued to. In general the same as audience.
	Scope string `json:"scope,omitempty"` // The space separated list of scopes granted to this token.
	User_id string `json:"user_id,omitempty"` // The obfuscated user id.
	Verified_email bool `json:"verified_email,omitempty"` // Boolean flag which is true if the email address is verified. Present only if the email scope is present in the request.
	Audience string `json:"audience,omitempty"` // Who is the intended audience for this token. In general the same as issued_to.
	Email string `json:"email,omitempty"` // The email address of the user. Present only if the email scope is present in the request.
}

// Userinfo represents the Userinfo schema from the OpenAPI specification
type Userinfo struct {
	Gender string `json:"gender,omitempty"` // The user's gender.
	Name string `json:"name,omitempty"` // The user's full name.
	Picture string `json:"picture,omitempty"` // URL of the user's picture image.
	Verified_email bool `json:"verified_email,omitempty"` // Boolean flag which is true if the email address is verified. Always verified because we only return the user's primary email address.
	Family_name string `json:"family_name,omitempty"` // The user's last name.
	Id string `json:"id,omitempty"` // The obfuscated ID of the user.
	Given_name string `json:"given_name,omitempty"` // The user's first name.
	Link string `json:"link,omitempty"` // URL of the profile page.
	Locale string `json:"locale,omitempty"` // The user's preferred locale.
	Email string `json:"email,omitempty"` // The user's email address.
	Hd string `json:"hd,omitempty"` // The hosted domain e.g. example.com if the user is Google apps user.
}
