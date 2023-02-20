package conv

import (
	"encoding/json"
)

func PassToStruct(from interface{}, target interface{}) error {
	byteData, err := json.Marshal(from)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, &target)
	if err != nil {
		return err
	}
	return nil
}
