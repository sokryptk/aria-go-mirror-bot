package commandHandlers

import (
	"aria-go-mirror-bot/ariaHelper"
	"aria-go-mirror-bot/env"
	"aria-go-mirror-bot/helpers"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"log"
	"os"
	"path/filepath"
)

//We are not using a logger yet, so just simply return a nil
func MirrorHandler(b ext.Bot, u *gotgbot.Update, args []string) error {
	gid, err := ariaHelper.RPC.AddURI(args[0])
	if err != nil {
		log.Fatalf("Could not add URI to RPC : %v\n", err)
	}

	m, err := u.EffectiveMessage.ReplyText("Download Started!")
	if err != nil {
		log.Println(err)
	}

	//Meh will just ignore here
	status, _ := ariaHelper.RPC.TellStatus(gid)

	Loop:
	for {
		switch status.Status {
		case "active":
			status, _ = ariaHelper.RPC.TellStatus(gid)
			text := helpers.GetDownloadFormat(
						getBasePath(status.Files[0].Path,
						len(status.Files) > 1),
						status.CompletedLength,
						status.TotalLength,
						status.DownloadSpeed,
					)

			if text != m.Text {
				m, _ = m.EditText(text)

			}

		case "complete":
			m, err = m.EditText("Uploading!")
			if err != nil {
				log.Println(err)
			}
			break Loop
		}
	}

	s, err := helpers.GetService()
	if err != nil {
		log.Fatalf("Could not retrive service : %v\n", err)
	}

	//Is a Dir
	if len(status.Files) > 1 {

	} else {
		uFile, err := helpers.UploadFile(s, status.Files[0].Path, env.Config.FolderId)
		if err != nil {
			log.Fatal(err)
		}

		_, err = m.EditHTML(fmt.Sprintf("<a href='%s'>%s</a>", fmt.Sprintf("https://drive.google.com/open?id=%s", uFile.Id), filepath.Base(status.Files[0].Path)))
		if err != nil {
			log.Fatal(err)
		}

		go os.RemoveAll(filepath.Base(getBasePath(status.Files[0].Path, false)))

	}

	return nil

}

func getBasePath(path string, isDir bool) string {
	if isDir {
		return filepath.Base(filepath.Dir(path))
	} else {
		return filepath.Base(path)
	}
}
