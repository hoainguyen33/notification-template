package model

import (
	"fmt"
	"regexp"
)

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)
)

func init() {

}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
}

type PhoneValidate string

func (p PhoneValidate) IsValid() (bool, error) {
	if p == "" {
		return false, fmt.Errorf("phone is empty")
	}
	match, err := regexp.MatchString("((\\+84|0)+([235789]))+([0-9]{8})", string(p))
	if err != nil || !match {
		return false, fmt.Errorf("Phone number is invalid! ONLY Mobile phone number accepted")
	}
	return true, nil
}
