// Package WasaProject provides primitives to interact with the openapi HTTP API.

package WasaProject

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Error The overall error object
type Error struct {
	// Error Describes the more specific error I am putting out
	Error *string `json:"error,omitempty"`
}

// User the entire user with all its attributes
type User struct {
	// Birthday User's birthday
	Birthday *string `json:"birthday,omitempty"`

	// Email User's email
	Email *string `json:"email,omitempty"`

	// Matricola the student's enrolment code
	Matricola *int `json:"matricola,omitempty"`

	// Password User's password
	Password *string `json:"password,omitempty"`

	// SecurityQuestion User's security question
	SecurityQuestion *string `json:"security_question,omitempty"`

	// UserId A unique user identifier
	UserId *string `json:"userId,omitempty"`

	// Username User's username
	Username *string `json:"username,omitempty"`
}

// PhotoId A unique photo identifier
type PhotoId = string

// UserId A unique user identifier
type UserId = string

// CommentPhotoJSONBody defines parameters for CommentPhoto.
type CommentPhotoJSONBody struct {
	// PhotoId A unique photo identifier
	PhotoId *PhotoId `json:"photoId,omitempty"`
}

// LikePhotoJSONBody defines parameters for LikePhoto.
type LikePhotoJSONBody struct {
	// UserID A unique user identifier
	UserID *UserId `json:"userID,omitempty"`
}

// DoLoginJSONBody defines parameters for DoLogin.
type DoLoginJSONBody struct {
	// Name User's name
	Name *string `json:"name,omitempty"`
}

// BanUserJSONBody defines parameters for BanUser.
type BanUserJSONBody struct {
	// UserId A unique user identifier
	UserId *UserId `json:"userId,omitempty"`
}

// FollowUserJSONBody defines parameters for FollowUser.
type FollowUserJSONBody struct {
	// UserId A unique user identifier
	UserId *UserId `json:"userId,omitempty"`
}

// UploadPhotoJSONBody defines parameters for UploadPhoto.
type UploadPhotoJSONBody struct {
	// Photo Base64-encoded image data
	Photo *string `json:"photo,omitempty"`
}

// UnfollowUserJSONBody defines parameters for UnfollowUser.
type UnfollowUserJSONBody struct {
	// UserId A unique user identifier
	UserId *UserId `json:"userId,omitempty"`
}

// CommentPhotoJSONRequestBody defines body for CommentPhoto for application/json ContentType.
type CommentPhotoJSONRequestBody CommentPhotoJSONBody

// LikePhotoJSONRequestBody defines body for LikePhoto for application/json ContentType.
type LikePhotoJSONRequestBody LikePhotoJSONBody

// DoLoginJSONRequestBody defines body for DoLogin for application/json ContentType.
type DoLoginJSONRequestBody DoLoginJSONBody

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = User

// BanUserJSONRequestBody defines body for BanUser for application/json ContentType.
type BanUserJSONRequestBody BanUserJSONBody

// FollowUserJSONRequestBody defines body for FollowUser for application/json ContentType.
type FollowUserJSONRequestBody FollowUserJSONBody

// SetMyUserNameJSONRequestBody defines body for SetMyUserName for application/json ContentType.
type SetMyUserNameJSONRequestBody = User

// UploadPhotoJSONRequestBody defines body for UploadPhoto for application/json ContentType.
type UploadPhotoJSONRequestBody UploadPhotoJSONBody

// UnbanUserJSONRequestBody defines body for UnbanUser for application/json ContentType.
type UnbanUserJSONRequestBody = User

// UnfollowUserJSONRequestBody defines body for UnfollowUser for application/json ContentType.
type UnfollowUserJSONRequestBody UnfollowUserJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// UncommentPhoto request
	UncommentPhoto(ctx context.Context, commentId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UnlikePhoto request
	UnlikePhoto(ctx context.Context, likeId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CommentPhotoWithBody request with any body
	CommentPhotoWithBody(ctx context.Context, photoId PhotoId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CommentPhoto(ctx context.Context, photoId PhotoId, body CommentPhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeletePhoto request
	DeletePhoto(ctx context.Context, photoId PhotoId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// LikePhotoWithBody request with any body
	LikePhotoWithBody(ctx context.Context, photoId PhotoId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	LikePhoto(ctx context.Context, photoId PhotoId, body LikePhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DoLoginWithBody request with any body
	DoLoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	DoLogin(ctx context.Context, body DoLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateUserWithBody request with any body
	CreateUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateUser(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetUserProfile request
	GetUserProfile(ctx context.Context, userId UserId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// BanUserWithBody request with any body
	BanUserWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	BanUser(ctx context.Context, userId UserId, body BanUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// FollowUserWithBody request with any body
	FollowUserWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	FollowUser(ctx context.Context, userId UserId, body FollowUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// SetMyUserNameWithBody request with any body
	SetMyUserNameWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	SetMyUserName(ctx context.Context, userId UserId, body SetMyUserNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UploadPhotoWithBody request with any body
	UploadPhotoWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UploadPhoto(ctx context.Context, userId UserId, body UploadPhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetMyStream request
	GetMyStream(ctx context.Context, userId UserId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UnbanUserWithBody request with any body
	UnbanUserWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UnbanUser(ctx context.Context, userId UserId, body UnbanUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UnfollowUserWithBody request with any body
	UnfollowUserWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UnfollowUser(ctx context.Context, userId UserId, body UnfollowUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) UncommentPhoto(ctx context.Context, commentId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUncommentPhotoRequest(c.Server, commentId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UnlikePhoto(ctx context.Context, likeId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUnlikePhotoRequest(c.Server, likeId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CommentPhotoWithBody(ctx context.Context, photoId PhotoId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCommentPhotoRequestWithBody(c.Server, photoId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CommentPhoto(ctx context.Context, photoId PhotoId, body CommentPhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCommentPhotoRequest(c.Server, photoId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeletePhoto(ctx context.Context, photoId PhotoId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeletePhotoRequest(c.Server, photoId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) LikePhotoWithBody(ctx context.Context, photoId PhotoId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLikePhotoRequestWithBody(c.Server, photoId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) LikePhoto(ctx context.Context, photoId PhotoId, body LikePhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLikePhotoRequest(c.Server, photoId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DoLoginWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDoLoginRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DoLogin(ctx context.Context, body DoLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDoLoginRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateUserRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateUser(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateUserRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetUserProfile(ctx context.Context, userId UserId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetUserProfileRequest(c.Server, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BanUserWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBanUserRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) BanUser(ctx context.Context, userId UserId, body BanUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewBanUserRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) FollowUserWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFollowUserRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) FollowUser(ctx context.Context, userId UserId, body FollowUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewFollowUserRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SetMyUserNameWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSetMyUserNameRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) SetMyUserName(ctx context.Context, userId UserId, body SetMyUserNameJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewSetMyUserNameRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UploadPhotoWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUploadPhotoRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UploadPhoto(ctx context.Context, userId UserId, body UploadPhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUploadPhotoRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetMyStream(ctx context.Context, userId UserId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMyStreamRequest(c.Server, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UnbanUserWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUnbanUserRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UnbanUser(ctx context.Context, userId UserId, body UnbanUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUnbanUserRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UnfollowUserWithBody(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUnfollowUserRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UnfollowUser(ctx context.Context, userId UserId, body UnfollowUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUnfollowUserRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewUncommentPhotoRequest generates requests for UncommentPhoto
func NewUncommentPhotoRequest(server string, commentId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "commentId", runtime.ParamLocationPath, commentId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/comments/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUnlikePhotoRequest generates requests for UnlikePhoto
func NewUnlikePhotoRequest(server string, likeId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "likeId", runtime.ParamLocationPath, likeId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/likes/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCommentPhotoRequest calls the generic CommentPhoto builder with application/json body
func NewCommentPhotoRequest(server string, photoId PhotoId, body CommentPhotoJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCommentPhotoRequestWithBody(server, photoId, "application/json", bodyReader)
}

// NewCommentPhotoRequestWithBody generates requests for CommentPhoto with any type of body
func NewCommentPhotoRequestWithBody(server string, photoId PhotoId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "photoId", runtime.ParamLocationPath, photoId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/photos/%s/comments/", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeletePhotoRequest generates requests for DeletePhoto
func NewDeletePhotoRequest(server string, photoId PhotoId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "photoId", runtime.ParamLocationPath, photoId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/photos/%s/delete", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewLikePhotoRequest calls the generic LikePhoto builder with application/json body
func NewLikePhotoRequest(server string, photoId PhotoId, body LikePhotoJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewLikePhotoRequestWithBody(server, photoId, "application/json", bodyReader)
}

// NewLikePhotoRequestWithBody generates requests for LikePhoto with any type of body
func NewLikePhotoRequestWithBody(server string, photoId PhotoId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "photoId", runtime.ParamLocationPath, photoId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/photos/%s/like", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDoLoginRequest calls the generic DoLogin builder with application/json body
func NewDoLoginRequest(server string, body DoLoginJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewDoLoginRequestWithBody(server, "application/json", bodyReader)
}

// NewDoLoginRequestWithBody generates requests for DoLogin with any type of body
func NewDoLoginRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/session")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewCreateUserRequest calls the generic CreateUser builder with application/json body
func NewCreateUserRequest(server string, body CreateUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateUserRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateUserRequestWithBody generates requests for CreateUser with any type of body
func NewCreateUserRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetUserProfileRequest generates requests for GetUserProfile
func NewGetUserProfileRequest(server string, userId UserId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewBanUserRequest calls the generic BanUser builder with application/json body
func NewBanUserRequest(server string, userId UserId, body BanUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewBanUserRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewBanUserRequestWithBody generates requests for BanUser with any type of body
func NewBanUserRequestWithBody(server string, userId UserId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/ban", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewFollowUserRequest calls the generic FollowUser builder with application/json body
func NewFollowUserRequest(server string, userId UserId, body FollowUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewFollowUserRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewFollowUserRequestWithBody generates requests for FollowUser with any type of body
func NewFollowUserRequestWithBody(server string, userId UserId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/follow", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewSetMyUserNameRequest calls the generic SetMyUserName builder with application/json body
func NewSetMyUserNameRequest(server string, userId UserId, body SetMyUserNameJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewSetMyUserNameRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewSetMyUserNameRequestWithBody generates requests for SetMyUserName with any type of body
func NewSetMyUserNameRequestWithBody(server string, userId UserId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/name", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewUploadPhotoRequest calls the generic UploadPhoto builder with application/json body
func NewUploadPhotoRequest(server string, userId UserId, body UploadPhotoJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUploadPhotoRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewUploadPhotoRequestWithBody generates requests for UploadPhoto with any type of body
func NewUploadPhotoRequestWithBody(server string, userId UserId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/photos/", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetMyStreamRequest generates requests for GetMyStream
func NewGetMyStreamRequest(server string, userId UserId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/stream", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUnbanUserRequest calls the generic UnbanUser builder with application/json body
func NewUnbanUserRequest(server string, userId UserId, body UnbanUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUnbanUserRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewUnbanUserRequestWithBody generates requests for UnbanUser with any type of body
func NewUnbanUserRequestWithBody(server string, userId UserId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/unban", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewUnfollowUserRequest calls the generic UnfollowUser builder with application/json body
func NewUnfollowUserRequest(server string, userId UserId, body UnfollowUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUnfollowUserRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewUnfollowUserRequestWithBody generates requests for UnfollowUser with any type of body
func NewUnfollowUserRequestWithBody(server string, userId UserId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s/unfollow", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// UncommentPhotoWithResponse request
	UncommentPhotoWithResponse(ctx context.Context, commentId string, reqEditors ...RequestEditorFn) (*UncommentPhotoResponse, error)

	// UnlikePhotoWithResponse request
	UnlikePhotoWithResponse(ctx context.Context, likeId string, reqEditors ...RequestEditorFn) (*UnlikePhotoResponse, error)

	// CommentPhotoWithBodyWithResponse request with any body
	CommentPhotoWithBodyWithResponse(ctx context.Context, photoId PhotoId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CommentPhotoResponse, error)

	CommentPhotoWithResponse(ctx context.Context, photoId PhotoId, body CommentPhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*CommentPhotoResponse, error)

	// DeletePhotoWithResponse request
	DeletePhotoWithResponse(ctx context.Context, photoId PhotoId, reqEditors ...RequestEditorFn) (*DeletePhotoResponse, error)

	// LikePhotoWithBodyWithResponse request with any body
	LikePhotoWithBodyWithResponse(ctx context.Context, photoId PhotoId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*LikePhotoResponse, error)

	LikePhotoWithResponse(ctx context.Context, photoId PhotoId, body LikePhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*LikePhotoResponse, error)

	// DoLoginWithBodyWithResponse request with any body
	DoLoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DoLoginResponse, error)

	DoLoginWithResponse(ctx context.Context, body DoLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*DoLoginResponse, error)

	// CreateUserWithBodyWithResponse request with any body
	CreateUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateUserResponse, error)

	CreateUserWithResponse(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateUserResponse, error)

	// GetUserProfileWithResponse request
	GetUserProfileWithResponse(ctx context.Context, userId UserId, reqEditors ...RequestEditorFn) (*GetUserProfileResponse, error)

	// BanUserWithBodyWithResponse request with any body
	BanUserWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BanUserResponse, error)

	BanUserWithResponse(ctx context.Context, userId UserId, body BanUserJSONRequestBody, reqEditors ...RequestEditorFn) (*BanUserResponse, error)

	// FollowUserWithBodyWithResponse request with any body
	FollowUserWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*FollowUserResponse, error)

	FollowUserWithResponse(ctx context.Context, userId UserId, body FollowUserJSONRequestBody, reqEditors ...RequestEditorFn) (*FollowUserResponse, error)

	// SetMyUserNameWithBodyWithResponse request with any body
	SetMyUserNameWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SetMyUserNameResponse, error)

	SetMyUserNameWithResponse(ctx context.Context, userId UserId, body SetMyUserNameJSONRequestBody, reqEditors ...RequestEditorFn) (*SetMyUserNameResponse, error)

	// UploadPhotoWithBodyWithResponse request with any body
	UploadPhotoWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UploadPhotoResponse, error)

	UploadPhotoWithResponse(ctx context.Context, userId UserId, body UploadPhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*UploadPhotoResponse, error)

	// GetMyStreamWithResponse request
	GetMyStreamWithResponse(ctx context.Context, userId UserId, reqEditors ...RequestEditorFn) (*GetMyStreamResponse, error)

	// UnbanUserWithBodyWithResponse request with any body
	UnbanUserWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UnbanUserResponse, error)

	UnbanUserWithResponse(ctx context.Context, userId UserId, body UnbanUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UnbanUserResponse, error)

	// UnfollowUserWithBodyWithResponse request with any body
	UnfollowUserWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UnfollowUserResponse, error)

	UnfollowUserWithResponse(ctx context.Context, userId UserId, body UnfollowUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UnfollowUserResponse, error)
}

type UncommentPhotoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Message Telling the user you removed the comment
		Message *string `json:"message,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r UncommentPhotoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UncommentPhotoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UnlikePhotoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Message saying you unliked correctly
		Message *string `json:"message,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r UnlikePhotoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UnlikePhotoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CommentPhotoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Message the message
		Message *string `json:"message,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r CommentPhotoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CommentPhotoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeletePhotoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Message You uploaded succesfully
		Message *string `json:"message,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r DeletePhotoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeletePhotoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type LikePhotoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Message telling the user he liked
		Message *string `json:"message,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r LikePhotoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r LikePhotoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DoLoginResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *User
}

// Status returns HTTPResponse.Status
func (r DoLoginResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DoLoginResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *User
	JSON401      *Error
}

// Status returns HTTPResponse.Status
func (r CreateUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetUserProfileResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *User
	JSON401      *Error
}

// Status returns HTTPResponse.Status
func (r GetUserProfileResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetUserProfileResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type BanUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Message Returing a message about what user is banned
		Message *string `json:"message,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r BanUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r BanUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type FollowUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// Message the success message
		Message *string `json:"message,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r FollowUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r FollowUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type SetMyUserNameResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *User
	JSON401      *Error
}

// Status returns HTTPResponse.Status
func (r SetMyUserNameResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r SetMyUserNameResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UploadPhotoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		// PhotoId A unique photo identifier
		PhotoId *PhotoId `json:"photoId,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r UploadPhotoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UploadPhotoResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMyStreamResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Stream all the stuff user needs for his for you page
		Stream *[]struct {
			// ItemProperty Description of item property
			ItemProperty *string `json:"itemProperty,omitempty"`
		} `json:"stream,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r GetMyStreamResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetMyStreamResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UnbanUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Message telling the user it worked
		Message *string `json:"message,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r UnbanUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UnbanUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UnfollowUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// Message unfollow message
		Message *string `json:"message,omitempty"`
	}
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r UnfollowUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UnfollowUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// UncommentPhotoWithResponse request returning *UncommentPhotoResponse
func (c *ClientWithResponses) UncommentPhotoWithResponse(ctx context.Context, commentId string, reqEditors ...RequestEditorFn) (*UncommentPhotoResponse, error) {
	rsp, err := c.UncommentPhoto(ctx, commentId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUncommentPhotoResponse(rsp)
}

// UnlikePhotoWithResponse request returning *UnlikePhotoResponse
func (c *ClientWithResponses) UnlikePhotoWithResponse(ctx context.Context, likeId string, reqEditors ...RequestEditorFn) (*UnlikePhotoResponse, error) {
	rsp, err := c.UnlikePhoto(ctx, likeId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUnlikePhotoResponse(rsp)
}

// CommentPhotoWithBodyWithResponse request with arbitrary body returning *CommentPhotoResponse
func (c *ClientWithResponses) CommentPhotoWithBodyWithResponse(ctx context.Context, photoId PhotoId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CommentPhotoResponse, error) {
	rsp, err := c.CommentPhotoWithBody(ctx, photoId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCommentPhotoResponse(rsp)
}

func (c *ClientWithResponses) CommentPhotoWithResponse(ctx context.Context, photoId PhotoId, body CommentPhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*CommentPhotoResponse, error) {
	rsp, err := c.CommentPhoto(ctx, photoId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCommentPhotoResponse(rsp)
}

// DeletePhotoWithResponse request returning *DeletePhotoResponse
func (c *ClientWithResponses) DeletePhotoWithResponse(ctx context.Context, photoId PhotoId, reqEditors ...RequestEditorFn) (*DeletePhotoResponse, error) {
	rsp, err := c.DeletePhoto(ctx, photoId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeletePhotoResponse(rsp)
}

// LikePhotoWithBodyWithResponse request with arbitrary body returning *LikePhotoResponse
func (c *ClientWithResponses) LikePhotoWithBodyWithResponse(ctx context.Context, photoId PhotoId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*LikePhotoResponse, error) {
	rsp, err := c.LikePhotoWithBody(ctx, photoId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLikePhotoResponse(rsp)
}

func (c *ClientWithResponses) LikePhotoWithResponse(ctx context.Context, photoId PhotoId, body LikePhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*LikePhotoResponse, error) {
	rsp, err := c.LikePhoto(ctx, photoId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLikePhotoResponse(rsp)
}

// DoLoginWithBodyWithResponse request with arbitrary body returning *DoLoginResponse
func (c *ClientWithResponses) DoLoginWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DoLoginResponse, error) {
	rsp, err := c.DoLoginWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDoLoginResponse(rsp)
}

func (c *ClientWithResponses) DoLoginWithResponse(ctx context.Context, body DoLoginJSONRequestBody, reqEditors ...RequestEditorFn) (*DoLoginResponse, error) {
	rsp, err := c.DoLogin(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDoLoginResponse(rsp)
}

// CreateUserWithBodyWithResponse request with arbitrary body returning *CreateUserResponse
func (c *ClientWithResponses) CreateUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateUserResponse, error) {
	rsp, err := c.CreateUserWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateUserResponse(rsp)
}

func (c *ClientWithResponses) CreateUserWithResponse(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateUserResponse, error) {
	rsp, err := c.CreateUser(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateUserResponse(rsp)
}

// GetUserProfileWithResponse request returning *GetUserProfileResponse
func (c *ClientWithResponses) GetUserProfileWithResponse(ctx context.Context, userId UserId, reqEditors ...RequestEditorFn) (*GetUserProfileResponse, error) {
	rsp, err := c.GetUserProfile(ctx, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetUserProfileResponse(rsp)
}

// BanUserWithBodyWithResponse request with arbitrary body returning *BanUserResponse
func (c *ClientWithResponses) BanUserWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*BanUserResponse, error) {
	rsp, err := c.BanUserWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBanUserResponse(rsp)
}

func (c *ClientWithResponses) BanUserWithResponse(ctx context.Context, userId UserId, body BanUserJSONRequestBody, reqEditors ...RequestEditorFn) (*BanUserResponse, error) {
	rsp, err := c.BanUser(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseBanUserResponse(rsp)
}

// FollowUserWithBodyWithResponse request with arbitrary body returning *FollowUserResponse
func (c *ClientWithResponses) FollowUserWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*FollowUserResponse, error) {
	rsp, err := c.FollowUserWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFollowUserResponse(rsp)
}

func (c *ClientWithResponses) FollowUserWithResponse(ctx context.Context, userId UserId, body FollowUserJSONRequestBody, reqEditors ...RequestEditorFn) (*FollowUserResponse, error) {
	rsp, err := c.FollowUser(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseFollowUserResponse(rsp)
}

// SetMyUserNameWithBodyWithResponse request with arbitrary body returning *SetMyUserNameResponse
func (c *ClientWithResponses) SetMyUserNameWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*SetMyUserNameResponse, error) {
	rsp, err := c.SetMyUserNameWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSetMyUserNameResponse(rsp)
}

func (c *ClientWithResponses) SetMyUserNameWithResponse(ctx context.Context, userId UserId, body SetMyUserNameJSONRequestBody, reqEditors ...RequestEditorFn) (*SetMyUserNameResponse, error) {
	rsp, err := c.SetMyUserName(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseSetMyUserNameResponse(rsp)
}

// UploadPhotoWithBodyWithResponse request with arbitrary body returning *UploadPhotoResponse
func (c *ClientWithResponses) UploadPhotoWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UploadPhotoResponse, error) {
	rsp, err := c.UploadPhotoWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUploadPhotoResponse(rsp)
}

func (c *ClientWithResponses) UploadPhotoWithResponse(ctx context.Context, userId UserId, body UploadPhotoJSONRequestBody, reqEditors ...RequestEditorFn) (*UploadPhotoResponse, error) {
	rsp, err := c.UploadPhoto(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUploadPhotoResponse(rsp)
}

// GetMyStreamWithResponse request returning *GetMyStreamResponse
func (c *ClientWithResponses) GetMyStreamWithResponse(ctx context.Context, userId UserId, reqEditors ...RequestEditorFn) (*GetMyStreamResponse, error) {
	rsp, err := c.GetMyStream(ctx, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMyStreamResponse(rsp)
}

// UnbanUserWithBodyWithResponse request with arbitrary body returning *UnbanUserResponse
func (c *ClientWithResponses) UnbanUserWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UnbanUserResponse, error) {
	rsp, err := c.UnbanUserWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUnbanUserResponse(rsp)
}

func (c *ClientWithResponses) UnbanUserWithResponse(ctx context.Context, userId UserId, body UnbanUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UnbanUserResponse, error) {
	rsp, err := c.UnbanUser(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUnbanUserResponse(rsp)
}

// UnfollowUserWithBodyWithResponse request with arbitrary body returning *UnfollowUserResponse
func (c *ClientWithResponses) UnfollowUserWithBodyWithResponse(ctx context.Context, userId UserId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UnfollowUserResponse, error) {
	rsp, err := c.UnfollowUserWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUnfollowUserResponse(rsp)
}

func (c *ClientWithResponses) UnfollowUserWithResponse(ctx context.Context, userId UserId, body UnfollowUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UnfollowUserResponse, error) {
	rsp, err := c.UnfollowUser(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUnfollowUserResponse(rsp)
}

// ParseUncommentPhotoResponse parses an HTTP response from a UncommentPhotoWithResponse call
func ParseUncommentPhotoResponse(rsp *http.Response) (*UncommentPhotoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UncommentPhotoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Message Telling the user you removed the comment
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseUnlikePhotoResponse parses an HTTP response from a UnlikePhotoWithResponse call
func ParseUnlikePhotoResponse(rsp *http.Response) (*UnlikePhotoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UnlikePhotoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Message saying you unliked correctly
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseCommentPhotoResponse parses an HTTP response from a CommentPhotoWithResponse call
func ParseCommentPhotoResponse(rsp *http.Response) (*CommentPhotoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CommentPhotoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Message the message
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseDeletePhotoResponse parses an HTTP response from a DeletePhotoWithResponse call
func ParseDeletePhotoResponse(rsp *http.Response) (*DeletePhotoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeletePhotoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Message You uploaded succesfully
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseLikePhotoResponse parses an HTTP response from a LikePhotoWithResponse call
func ParseLikePhotoResponse(rsp *http.Response) (*LikePhotoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &LikePhotoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Message telling the user he liked
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseDoLoginResponse parses an HTTP response from a DoLoginWithResponse call
func ParseDoLoginResponse(rsp *http.Response) (*DoLoginResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DoLoginResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateUserResponse parses an HTTP response from a CreateUserWithResponse call
func ParseCreateUserResponse(rsp *http.Response) (*CreateUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseGetUserProfileResponse parses an HTTP response from a GetUserProfileWithResponse call
func ParseGetUserProfileResponse(rsp *http.Response) (*GetUserProfileResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetUserProfileResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseBanUserResponse parses an HTTP response from a BanUserWithResponse call
func ParseBanUserResponse(rsp *http.Response) (*BanUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &BanUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Message Returing a message about what user is banned
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseFollowUserResponse parses an HTTP response from a FollowUserWithResponse call
func ParseFollowUserResponse(rsp *http.Response) (*FollowUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &FollowUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// Message the success message
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseSetMyUserNameResponse parses an HTTP response from a SetMyUserNameWithResponse call
func ParseSetMyUserNameResponse(rsp *http.Response) (*SetMyUserNameResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &SetMyUserNameResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest User
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseUploadPhotoResponse parses an HTTP response from a UploadPhotoWithResponse call
func ParseUploadPhotoResponse(rsp *http.Response) (*UploadPhotoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UploadPhotoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			// PhotoId A unique photo identifier
			PhotoId *PhotoId `json:"photoId,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseGetMyStreamResponse parses an HTTP response from a GetMyStreamWithResponse call
func ParseGetMyStreamResponse(rsp *http.Response) (*GetMyStreamResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetMyStreamResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Stream all the stuff user needs for his for you page
			Stream *[]struct {
				// ItemProperty Description of item property
				ItemProperty *string `json:"itemProperty,omitempty"`
			} `json:"stream,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseUnbanUserResponse parses an HTTP response from a UnbanUserWithResponse call
func ParseUnbanUserResponse(rsp *http.Response) (*UnbanUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UnbanUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Message telling the user it worked
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ParseUnfollowUserResponse parses an HTTP response from a UnfollowUserWithResponse call
func ParseUnfollowUserResponse(rsp *http.Response) (*UnfollowUserResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UnfollowUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// Message unfollow message
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Remove Comment from Photo
	// (DELETE /comments/{commentId})
	UncommentPhoto(ctx echo.Context, commentId string) error
	// Unlike Photo
	// (DELETE /likes/{likeId})
	UnlikePhoto(ctx echo.Context, likeId string) error
	// Comment on Photo
	// (POST /photos/{photoId}/comments/)
	CommentPhoto(ctx echo.Context, photoId PhotoId) error
	// Delete Photo
	// (DELETE /photos/{photoId}/delete)
	DeletePhoto(ctx echo.Context, photoId PhotoId) error
	// Like Photo
	// (POST /photos/{photoId}/like)
	LikePhoto(ctx echo.Context, photoId PhotoId) error
	// Logs in the user
	// (POST /session)
	DoLogin(ctx echo.Context) error
	// Create User
	// (POST /users/)
	CreateUser(ctx echo.Context) error
	// Get User Profile
	// (GET /users/{userId})
	GetUserProfile(ctx echo.Context, userId UserId) error
	// Ban User
	// (POST /users/{userId}/ban)
	BanUser(ctx echo.Context, userId UserId) error
	// Follow User
	// (POST /users/{userId}/follow)
	FollowUser(ctx echo.Context, userId UserId) error
	// Set My User Name
	// (PUT /users/{userId}/name)
	SetMyUserName(ctx echo.Context, userId UserId) error
	// Upload Photo
	// (POST /users/{userId}/photos/)
	UploadPhoto(ctx echo.Context, userId UserId) error
	// Get My Stream
	// (GET /users/{userId}/stream)
	GetMyStream(ctx echo.Context, userId UserId) error
	// Unban User
	// (POST /users/{userId}/unban)
	UnbanUser(ctx echo.Context, userId UserId) error
	// Unfollow User
	// (POST /users/{userId}/unfollow)
	UnfollowUser(ctx echo.Context, userId UserId) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// UncommentPhoto converts echo context to params.
func (w *ServerInterfaceWrapper) UncommentPhoto(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "commentId" -------------
	var commentId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "commentId", runtime.ParamLocationPath, ctx.Param("commentId"), &commentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter commentId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UncommentPhoto(ctx, commentId)
	return err
}

// UnlikePhoto converts echo context to params.
func (w *ServerInterfaceWrapper) UnlikePhoto(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "likeId" -------------
	var likeId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "likeId", runtime.ParamLocationPath, ctx.Param("likeId"), &likeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter likeId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UnlikePhoto(ctx, likeId)
	return err
}

// CommentPhoto converts echo context to params.
func (w *ServerInterfaceWrapper) CommentPhoto(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "photoId" -------------
	var photoId PhotoId

	err = runtime.BindStyledParameterWithLocation("simple", false, "photoId", runtime.ParamLocationPath, ctx.Param("photoId"), &photoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter photoId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CommentPhoto(ctx, photoId)
	return err
}

// DeletePhoto converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePhoto(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "photoId" -------------
	var photoId PhotoId

	err = runtime.BindStyledParameterWithLocation("simple", false, "photoId", runtime.ParamLocationPath, ctx.Param("photoId"), &photoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter photoId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeletePhoto(ctx, photoId)
	return err
}

// LikePhoto converts echo context to params.
func (w *ServerInterfaceWrapper) LikePhoto(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "photoId" -------------
	var photoId PhotoId

	err = runtime.BindStyledParameterWithLocation("simple", false, "photoId", runtime.ParamLocationPath, ctx.Param("photoId"), &photoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter photoId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.LikePhoto(ctx, photoId)
	return err
}

// DoLogin converts echo context to params.
func (w *ServerInterfaceWrapper) DoLogin(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DoLogin(ctx)
	return err
}

// CreateUser converts echo context to params.
func (w *ServerInterfaceWrapper) CreateUser(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateUser(ctx)
	return err
}

// GetUserProfile converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserProfile(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserProfile(ctx, userId)
	return err
}

// BanUser converts echo context to params.
func (w *ServerInterfaceWrapper) BanUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.BanUser(ctx, userId)
	return err
}

// FollowUser converts echo context to params.
func (w *ServerInterfaceWrapper) FollowUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.FollowUser(ctx, userId)
	return err
}

// SetMyUserName converts echo context to params.
func (w *ServerInterfaceWrapper) SetMyUserName(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SetMyUserName(ctx, userId)
	return err
}

// UploadPhoto converts echo context to params.
func (w *ServerInterfaceWrapper) UploadPhoto(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UploadPhoto(ctx, userId)
	return err
}

// GetMyStream converts echo context to params.
func (w *ServerInterfaceWrapper) GetMyStream(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMyStream(ctx, userId)
	return err
}

// UnbanUser converts echo context to params.
func (w *ServerInterfaceWrapper) UnbanUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UnbanUser(ctx, userId)
	return err
}

// UnfollowUser converts echo context to params.
func (w *ServerInterfaceWrapper) UnfollowUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UnfollowUser(ctx, userId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/comments/:commentId", wrapper.UncommentPhoto)
	router.DELETE(baseURL+"/likes/:likeId", wrapper.UnlikePhoto)
	router.POST(baseURL+"/photos/:photoId/comments/", wrapper.CommentPhoto)
	router.DELETE(baseURL+"/photos/:photoId/delete", wrapper.DeletePhoto)
	router.POST(baseURL+"/photos/:photoId/like", wrapper.LikePhoto)
	router.POST(baseURL+"/session", wrapper.DoLogin)
	router.POST(baseURL+"/users/", wrapper.CreateUser)
	router.GET(baseURL+"/users/:userId", wrapper.GetUserProfile)
	router.POST(baseURL+"/users/:userId/ban", wrapper.BanUser)
	router.POST(baseURL+"/users/:userId/follow", wrapper.FollowUser)
	router.PUT(baseURL+"/users/:userId/name", wrapper.SetMyUserName)
	router.POST(baseURL+"/users/:userId/photos/", wrapper.UploadPhoto)
	router.GET(baseURL+"/users/:userId/stream", wrapper.GetMyStream)
	router.POST(baseURL+"/users/:userId/unban", wrapper.UnbanUser)
	router.POST(baseURL+"/users/:userId/unfollow", wrapper.UnfollowUser)

}
