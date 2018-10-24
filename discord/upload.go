package discord

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/TinyKitten/discordgo"
)

type Discord struct {
	dg *discordgo.Session
}

func NewDiscord(token string) (*Discord, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return &Discord{
		dg: dg,
	}, nil
}

func (d *Discord) TransferEmoji(guildID, name, url string) error {
	// 記号があるとアップロードできない
	name = strings.Replace(name, "-", "", -1)
	name = strings.Replace(name, "+", "plus", -1)

	// Get alias:は邪魔
	if strings.HasPrefix(name, "Get alias:") {
		return nil
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error retrieving the file, ", err)
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	img, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading the response, ", err)
		return err
	}

	contentType := http.DetectContentType(img)
	base64img := base64.StdEncoding.EncodeToString(img)

	emoji := fmt.Sprintf("data:%s;base64,%s", contentType, base64img)
	_, err = d.dg.EmojiCreate(guildID, name, emoji)
	if err != nil {
		return err
	}

	return nil
}
