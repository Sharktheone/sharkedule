package locations

import (
	"github.com/Sharktheone/sharkedule/database/db"
	"log"
)

func GetLocations(uuid string) (map[string][]string, error) {
	t, err := db.DB.GetTask(uuid)
	if err != nil {
		return nil, err
	}
	var (
		locations map[string][]string
	)
	for _, b := range t.Boards {
		br, err := db.DB.GetBoard(b)
		if err != nil {
			log.Printf("error getting board: %v", err)
			continue
		}
		for _, c := range br.Columns {
			column, err := db.DB.GetColumn(c)
			if err != nil {
				log.Printf("error getting column: %v", err)
				continue
			}
			for _, t := range column.Tasks {
				if t == uuid {
					locations[b] = append(locations[b], c)
					break
				}
			}
		}
	}
	return locations, nil
}
