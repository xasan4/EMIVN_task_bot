package database

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int
	TUsername string
	Nickname  string
	UserType  string
}

type Card struct {
	Id             int
	CardNumber     string
	IssuingBank    string
	DailyLimit     int
	CurrentBalance int
	DaimyoID       int
}

type Session struct {
	Id int
	UserId int
	EntityType string
	StartTime time.Time
	EndTime time.Time
}

type Action struct {
	Id int
	SessionId int
	EntityId int
	ActionTime time.Time
	ActionType string
}

type Relations struct {
	Id int
	EntityId int
	RelatedEntityId int
	CreationType string
	CreationTime time.Time
}

func AddUser(db *sql.DB, u User) error {
	if _, err := db.Exec("insert into Users (Telegram_Username, Nickname, UserType) values ($1, $2, $3)", u.TUsername, u.Nickname, u.UserType); err != nil {
		return err
	}
	return nil
}

func AddCard(db *sql.DB, c Card) error {
	if _, err := db.Exec("insert into Users (Card_Number, Issuing_Bank, Daily_Limit, Daimyo_ID) values ($1, $2, $3, $4)", c.CardNumber, c.IssuingBank, c.DailyLimit, c.DaimyoID); err != nil {
		return err
	}
	return nil
}

func GetUserWithNick(db *sql.DB, nickname string) (User, error) {
	var u User
	err := db.QueryRow("select (ID, Telegram_Username, Nickname, UserType) from Users where Nickname = $1", nickname).Scan(&u.Id, &u.TUsername, &u.Nickname, &u.UserType)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func GetUserWithUser(db *sql.DB, username string) (User, error) {
	var u User
	err := db.QueryRow("select (ID, Telegram_Username, Nickname, UserType) from Users where Telegram_Username = $1", username).Scan(&u.Id, &u.TUsername, &u.Nickname, &u.UserType)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func LinkEntity(db *sql.DB, fnickname, ftitle, tnickname, ttitle string) error {
	var exist bool 
	err := db.QueryRow("select exists (select * from Users where Nickname = $1 and UserType = $2)", fnickname, ftitle).Scan(&exist)
	if err != nil {
		return err
	}
	err = db.QueryRow("select exists (select * from Users where Nickname = $1 and UserType = $2)", tnickname, ttitle).Scan(&exist)
	if err != nil {
		return err
	}
	if _, err := db.Exec("insert into Relations (Entity_Nickname, Related_Entity_Nickname, Relation_Type, Creation_Date_Time, ) values ($1, $2, $3, $4)", tnickname, fnickname, ftitle+ttitle, time.Now()); err != nil {
		return err
	}
	return nil
}

func LinkCard(db *sql.DB, fnickname, ftitle, tnickname, ttitle string) error {
	var exist bool 
	err := db.QueryRow("select exists (select * from Users where Nickname = $1 and UserType = $2)", fnickname, ftitle).Scan(&exist)
	if err != nil {
		return err
	}
	err = db.QueryRow("select exists (select * from Users where Nickname = $1 and UserType = $2)", tnickname, ttitle).Scan(&exist)
	if err != nil {
		return err
	}
	if _, err := db.Exec("insert into Relations (Entity_Nickname, Related_Entity_Nickname, Relation_Type, Creation_Date_Time, ) values ($1, $2, $3, $4)", tnickname, fnickname, ftitle+ttitle, time.Now()); err != nil {
		return err
	}
	return nil
}
