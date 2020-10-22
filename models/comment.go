package models

import (
	"../utilities"
	"fmt"
	"strconv"
)

var (
	CommentsKeyString = "comments"
)

type Comment struct {
	id int64
}

func NewComment(userId int64, text string) (*Comment, error) {
	var id, err = redisClient.Incr(ctx, "comment:next-id").Result()
	if err != nil {
		return nil, err
	}

	var (
		key = getCommentKey(id)
		pipe = redisClient.Pipeline()
	)
	pipe.HSet(ctx, key, "id", id)
	pipe.HSet(ctx, key, "user_id", userId)
	pipe.HSet(ctx, key, "text", text)
	pipe.LPush(ctx, CommentsKeyString, id)
	//pipe.LPush(ctx, fmt.Sprintf("user:%d:comment", userId), id)
	_, err = pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &Comment{id}, nil
}

func (comment *Comment) GetText() (string, error) {
	return redisClient.HGet(ctx, getCommentKey(comment.id), "text").Result()
}

func (comment *Comment) GetAuthor() (*User, error) {
	var userId, err = redisClient.HGet(ctx, getCommentKey(comment.id), "user_id").Int64()
	if err != nil {
		return nil, err
	}
	return GetUserById(userId)
}

func (comment *Comment) GetId() (int64, error) {
	return redisClient.HGet(ctx, getCommentKey(comment.id), "id").Int64()
}

func GetAllComments(start, stop int64) ([]*Comment, error) {
	return queryComments(CommentsKeyString, start, stop)
}

//func GetSpecificComments(userId, start, stop int64) ([]*Comment, error) {
//	return queryComments(fmt.Sprintf("user:%d:comment", userId), start, stop)
//}

func queryComments(key string, start, stop int64) ([]*Comment, error) {
	var commentIds, err = redisClient.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}

	var comments = make([]*Comment, len(commentIds))
	for i, idStr := range commentIds {
		var id, err = strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		comments[i] = &Comment{int64(id)}
	}
	return comments, nil
}

func AddComment(userId int64, text string) error {
	if len(text) == 0 {
		return utilities.ErrEmptyComment
	}
	var _, err = NewComment(userId, text)
	return err
}

func getCommentKey(id int64) string {
	return fmt.Sprintf("comment:%d", id)
}