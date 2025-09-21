// this file is for scraping content off of the rss.go
package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/itsemadbattal/rss-aggregator/internal/database"
)

// takes 3 inputs, connection to the db, number of goroutines we want to do scraping on,
// and how much time we want inbetween each rquest to go scrape a new RSSFeed
func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {

	log.Printf("Scraping on %v goroutines every %s duration.", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	//this for loop will execute everytime a new value comes across the ticker's channel,
	// this will run every timeBetweenRequest
	//using this `for ; ; <-ticker.C` will make the for loop execute immediatly when the code reaches its line
	for ; ; <-ticker.C {
		//context.Background() is a global context, and is used when we dont have access to scoped context like we have with http requests
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error fetching feeds: ", err)
			//we continue cause we want this function to run as long as the server is running
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)

		}
		//here we want to wait until the scraping is done then we continue to next iteration
		wg.Wait()
	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched: ", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetching feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		//cause we specified the description to be allowed to be null, we have to workaround cases where we check first if its there
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		//parsing the pubDate
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldnt parse date %v with err %v", item.PubDate, err)
			continue
		}

		_, err = db.CreatePost(context.Background(),
			database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: description,
				PublishedAt: pubAt,
				Url:         item.Link,
				FeedID:      feed.ID,
			})
		if err != nil {
			//we check if the error has the string "duplicate key", we dont want to log it like this so we continue
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("failed to create post: ", err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
