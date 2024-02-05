package handlers

import (
	"context"
	"ez-note/cmd/api/rpc"
	"ez-note/kitex_gen/notedemo"
	"ez-note/pkg/constants"
	"ez-note/pkg/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

// CreateNote create note info
func CreateNote(ctx context.Context, c *app.RequestContext) {
	var noteVar NoteParam
	if err := c.Bind(&noteVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(noteVar.Title) == 0 || len(noteVar.Content) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))
	err := rpc.CreateNote(context.Background(), &notedemo.CreateNoteRequest{
		UserId:  userID,
		Content: noteVar.Content, Title: noteVar.Title,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// DeleteNote delete note info
func DeleteNote(ctx context.Context, c *app.RequestContext) {
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))
	noteIDStr := c.Param(constants.NoteID)
	noteID, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if noteID <= 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.DeleteNote(context.Background(), &notedemo.DeleteNoteRequest{
		NoteId: noteID, UserId: userID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// QueryNote query list of note info
func QueryNote(ctx context.Context, c *app.RequestContext) {
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))
	var queryVar struct {
		Limit         int64  `json:"limit" form:"limit" query:"limit"`
		Offset        int64  `json:"offset" form:"offset" query:"offset"`
		SearchKeyword string `json:"search_keyword" form:"search_keyword" query:"search_keyword"`
	}
	if err := c.Bind(&queryVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}

	if queryVar.Limit < 0 || queryVar.Offset < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	req := &notedemo.QueryNoteRequest{UserId: userID, Offset: queryVar.Offset, Limit: queryVar.Limit}
	if len(queryVar.SearchKeyword) != 0 {
		req.SearchKey = &queryVar.SearchKeyword
	}
	notes, total, err := rpc.QueryNotes(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{constants.Total: total, constants.Notes: notes})
}

// UpdateNote update user info
func UpdateNote(ctx context.Context, c *app.RequestContext) {
	var noteVar NoteParam
	if err := c.Bind(&noteVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))
	noteIDStr := c.Param(constants.NoteID)
	noteID, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if noteID <= 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	req := &notedemo.UpdateNoteRequest{NoteId: noteID, UserId: userID}
	if len(noteVar.Title) != 0 {
		req.Title = &noteVar.Title
	}
	if len(noteVar.Content) != 0 {
		req.Content = &noteVar.Content
	}
	if err = rpc.UpdateNote(context.Background(), req); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
