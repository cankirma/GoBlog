package data

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strings"
	"time"
)

func GetAllPosts(conn *pgx.Conn) ([]Post, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, title, body, author_id FROM post")
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Error getting posts")
	}

	var posts []Post
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.AuthorID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (i *Post) Create(conn *pgx.Conn, userID string) error {
	i.Title = strings.Trim(i.Title, " ")
	if len(i.Title) < 1 {
		return fmt.Errorf("Title must not be empty.")
	}
	if len(i.Body) < 1 {
		return fmt.Errorf("Body must not be empty")
	}
	now := time.Now()

	row := conn.QueryRow(context.Background(),
		`INSERT INTO post(
title, 
body, 
author_id, 
created_at, 
updated_at) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING id,author_id`, i.Title, i.Body, userID, now, now)

	err := row.Scan(&i.ID, &i.AuthorID)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("There was an error creating the post")
	}

	return nil
}

func GetPostByID(postID string, conn *pgx.Conn) (Post, error) {
	row := conn.QueryRow(context.Background(), "SELECT id, title, body, author_id FROM post WHERE id=$1", postID)
	post := Post{}
	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.AuthorID)
	if err != nil {
		return post, fmt.Errorf("The post not found")
	}

	return post, nil
}

func GetPostsByCurrentUser(userID string, conn *pgx.Conn) ([]Post, error) {
	rows, err := conn.Query(context.Background(), "SELECT id, title, body, author_id FROM post WHERE author_id = $1", userID)
	if err != nil {
		fmt.Printf("Error getting posts %v", err)
		return nil, fmt.Errorf("There was an error getting the posts")
	}

	var posts []Post
	for rows.Next() {
		i := Post{}
		err = rows.Scan(&i.ID, &i.Title, &i.Body, &i.AuthorID)
		if err != nil {
			fmt.Printf("Error scaning post: %v", err)
			continue
		}
		posts = append(posts, i)
	}

	return posts, nil
}

func (i *Post) Update(conn *pgx.Conn) error {
	i.Title = strings.Trim(i.Title, " ")
	if len(i.Title) < 1 {
		return fmt.Errorf("Title must not be empty")
	}
	if len(i.Body) < 1 {
		return fmt.Errorf("Body must not be empty")
	}
	now := time.Now()
	_, err := conn.Exec(context.Background(), "UPDATE post SET title=$1, body=$2, updated_at=$3 WHERE id=$4", i.Title, i.Body, now, i.ID)

	if err != nil {
		fmt.Printf("Error updating post: (%v)", err)
		return fmt.Errorf("Error updating post")
	}

	return nil
}

func Delete(postID string, conn *pgx.Conn) error {
	_, err := conn.Query(context.Background(), "DELETE FROM post WHERE id=$1", postID)
	if err != nil {
		fmt.Printf("Error getting posts %v", err)
		return fmt.Errorf("There was an error getting the posts")
	}

	return nil
}
