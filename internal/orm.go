package wtfd

import (
	"os"
	"fmt"
	"errors"
    _ "github.com/mattn/go-sqlite3"
    "github.com/go-xorm/xorm"
	"xorm.io/core"
)

const (
	_ORMUserName            = "User"			// Table name for db
	_ORMChallengeByUserName = "ChallengeByUser" // Table name for db
)

var (
	engine *xorm.Engine
	ErrORMGeneric = errors.New("Database error (check log)")
)

////////////////////////////////////////////////////////////////////////////////
// ORM definitions /////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
type _ORMUser struct {
	Name      string `xorm:"unique"`
	Hash      []byte
	Mail	  string
}

type _ORMChallengesByUser struct {
	UserName      string // Foregin keys don't exist
	ChallengeName string
}

func (u _ORMUser) TableName() string {
	return _ORMUserName
}

func (c _ORMChallengesByUser) TableName() string {
	return _ORMChallengeByUserName
}

func ormSync() {
	engine.Sync(_ORMUser{})
	engine.Sync(_ORMChallengesByUser{})
}
////////////////////////////////////////////////////////////////////////////////

func ormStart(logFile string) {
    var err error
    engine, err = xorm.NewEngine("sqlite3", "./state.db")

	if err != nil {
		panic(fmt.Sprintf("Could not start xorm engine: %s\n", err.Error()))
	}

	if logFile != "" {
		f, err := os.Create(logFile)
		if err != nil {
			fmt.Errorf("Could not create DB Logfile: %s\n", err.Error())
		} else {
			engine.SetLogger(xorm.NewSimpleLogger(f))
		}
	}

	engine.SetMapper(core.SameMapper{})

	ormSync()
	ormCreateTestDB()
}

func ormCreateTestDB() {
	u := &_ORMUser {
		Name: "TestUser",
		Hash: []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
		Mail: "test@example.com",
	}

	cu := &_ORMChallengesByUser {
		UserName: u.Name,
		ChallengeName: "TestChallenge",
	}

	engine.Insert(u)
	engine.Insert(cu)
}

func _ORMGenericError(desc string) error {
	return errors.New(fmt.Sprintf("ORM Error %s", desc))
}

////////////////////////////////////////////////////////////////////////////////
// DB Operations ///////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

// Create new User DB record from a User struct
func ormNewUser(user User) error {
	exists, err := ormUserExists(user)
	if err != nil {
		return err
	}

	if exists {
		return ErrUserExisting
	}

	_, err = engine.Insert(_ORMUser {
		Name: user.Name,
		Hash: user.Hash,
		Mail: user.Mail,
	})

	return err
}

// Update existing user record (user.Name) with other values from user
func ormUpdateUser(user User) error {
	// TODO
	return nil
}

// remove user from db
func ormDeleteUser(user User) error {
	// TODO
	return nil
}

// check if user exists in db
func ormUserExists(user User) (bool, error) {
	// TODO
	return false, nil
}

// get Challenges{} solved by user
func ormChallengesSolved(user User) (Challenges, error) {
	// TODO
	return Challenges{}, nil
}

// Write solved state (user solved chall) in db
func ormSolvedChallenge(user User, chall Challenge) (error) {
	// TODO
	return nil
}

// load a single user from db (search by u.Name)
// The remaining fields of u will be filled by this function
func ormLoadUser(u *User) error {
	// TODO
	return nil
}

// Fills u with all users in db
func ormLoadAllUsers(u *Users) error {
	// TODO
	return nil
}
////////////////////////////////////////////////////////////////////////////////
