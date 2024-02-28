package module

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"io/ioutil"
	"net/http"
	"strings"
)

type GithubModule struct{}

type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	Name              string `json:"name"`
	Company           string `json:"company"`
	Blog              string `json:"blog"`
	Location          string `json:"location"`
	Email             string `json:"email"`
	Hireable          bool   `json:"hireable"`
	Bio               string `json:"bio"`
	TwitterUsername   string `json:"twitter_username"`
	PublicRepos       int    `json:"public_repos"`
	PublicGists       int    `json:"public_gists"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

func getUserData(username string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *GithubModule) Handle(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := update.Message.Text
	if strings.HasPrefix(message, "/github") {

		parts := strings.Fields(message)
		if len(parts) > 1 {
			// Extract the arguments
			args := parts[1:]
			// Process the arguments
			username := strings.Join(args, " ")
			user, err := getUserData(username)
			if err != nil {
				fmt.Println("Error getting user data:", err)
				return
			}

			// Format the user data into Markdown
			text := fmt.Sprintf(`<b>GitHub Info for</b> %s

<b>Name:</b> %s
<b>Bio:</b> %s
<b>Location:</b> %s
<b>Email:</b> %s
<b>Public Repos:</b> %d
<b>Followers:</b> %d
<b>Following:</b> %d
<b>Created At:</b> %s
<b>Last Updated At:</b> %s`,
				user.Login, user.Name, user.Bio, user.Location, user.Email, user.PublicRepos, user.Followers, user.Following, user.CreatedAt, user.UpdatedAt)

			params := &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				ParseMode: "HTML",
				Text:      text,
			}
			_, err = b.SendMessage(ctx, params)
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}
		} else {
			// If no arguments provided
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "You need to specify GitHub Username!",
			})
		}

		fmt.Println("[LOG] Gtihub module executed successfully")
	}
}

func init() {
	RegisterModule(&GithubModule{})
}
