package jobs

import (
	"os"
	"time"

	md "youbei/models"
)

//ExpireDelete ... num  份数
func ExpireDelete(id string, num int) error {
	logs := []md.Log{}
	if err := md.Localdb().Where("deleted=0 and tid=?", id).Desc("created").Find(&logs); err == nil {
		for _, v := range logs {
			_, err := os.Stat(v.Localfilepath)
			if os.IsNotExist(err) {
				log := new(md.Log)
				log.Deleted = time.Now().Unix()
				if _, err := md.Localdb().ID(v.ID).Cols("deleted").Update(log); err != nil {
					return err
				}

			}
		}
	} else {
		return err
	}

	logs = []md.Log{}
	err := md.Localdb().Where("deleted=0 and tid=?", id).Desc("created").Find(&logs)
	if err != nil {
		return err
	}
	for i := num; i < len(logs); i++ {
		f, err := os.Stat(logs[i].Localfilepath)
		if err == nil {
			if !f.IsDir() {
				log := new(md.Log)
				log.Deleted = time.Now().Unix()
				_, err := md.Localdb().ID(logs[i].ID).Cols("deleted").Update(log)
				if err != nil {
					return err
				}
				os.Remove(logs[i].Localfilepath)
			}
		}
	}
	return nil
}
