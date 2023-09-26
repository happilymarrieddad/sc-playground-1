package repos

import (
	"fmt"

	"xorm.io/xorm"
)

func handleRollback(sesh *xorm.Session, err error) error {
	if rollBackErr := sesh.Rollback(); rollBackErr != nil {
		fmt.Printf("unable to rollback with err: %s", rollBackErr.Error())
	}

	return err
}
