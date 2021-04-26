package tweet

import (
	"github.com/alaanzr/go-api/internals/database"
)

func getTweets() ([]Tweet, error) {
	rows, err := database.DB.Query(`select * from tweet;`)

	if err != nil {
		return nil, err
	}

	tweets := []Tweet{}
	for rows.Next() {
		var r Tweet
		err = rows.Scan(&r.Id, &r.Content, &r.Author, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tweets = append(tweets, r)
	}

	defer rows.Close()

	return tweets, nil
}

func createTweet(tweet Tweet) error {
	rows, err := database.DB.Query(`insert into tweet (content) values ($1);`, tweet.Content)

	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}