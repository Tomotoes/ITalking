package dao

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"italking.tomotoes.com/m/v1/config"
	"italking.tomotoes.com/m/v1/ent"
	"italking.tomotoes.com/m/v1/ent/notification"
	"italking.tomotoes.com/m/v1/ent/room"
	"italking.tomotoes.com/m/v1/ent/user"
	"italking.tomotoes.com/m/v1/helper"
	"log"
	"sync"
	"time"
)

func GetAdmin() error {
	Admin, _ = G.User.Query().Where(user.NameEQ(config.AdminID)).First(G_)
	if Admin != nil {
		return nil
	}
	Admin, err = G.User.
		Create().
		SetName(config.AdminID).
		SetDescription("欢迎来到 ITalking").
		SetPassword("password").
		Save(G_)
	if err != nil {
		return fmt.Errorf("failed creating admin: %w", err)
	}
	log.Println("admin was created: ", Admin)
	return nil
}

func SendWelcomeNotification(receiver *ent.User) (*ent.Notification, error) {
	welcomeNotification, err := G.Notification.Create().
		SetContent(config.WelcomeNotification).
		SetSender(Admin).
		SetReceiver(receiver).
		SetType(notification.TypeOfficial).
		Save(G_)
	return welcomeNotification, err
}

func SendFollowNotification(receiver *ent.User, sender *ent.User) (*ent.Notification, error) {
	followNotification, err := G.Notification.Create().
		SetContent(config.FollowNotification).
		SetSender(sender).
		SetReceiver(receiver).
		SetType(notification.TypeFollow).
		Save(G_)
	return followNotification, err
}

func SendFollowNotificationByID(receiverID string, senderID string) (*ent.Notification, error) {
	followNotification, err := G.Notification.Create().
		SetContent(config.FollowNotification).
		SetSenderID(senderID).
		SetReceiverID(receiverID).
		SetType(notification.TypeFollow).
		Save(G_)
	return followNotification, err
}

func DestroyIllegalRooms() {
	ticker := time.NewTicker(config.CheckRoomsPeriodicity)
	defer ticker.Stop()

	for range ticker.C {
		rooms := G.Room.Query().AllX(G_)

		if len(rooms) == 0 {
			continue
		}

		notExistRoomChannel := make(chan string, len(rooms))
		var wg sync.WaitGroup
		wg.Add(len(rooms))

		for i, r := range rooms {
			if (i+1)%config.MaxCallLimit == 0 {
				time.Sleep(time.Second)
			}
			go func(r *ent.Room) {
				defer wg.Done()
				if !helper.IsRoomExist(r.ID) {
					notExistRoomChannel <- r.ID
				}
			}(r)
		}

		wg.Wait()
		close(notExistRoomChannel)

		var notExistRooms []interface{}
		for notExistRoom := range notExistRoomChannel {
			notExistRooms = append(notExistRooms, notExistRoom)
		}

		if len(notExistRooms) == 0 {
			continue
		}

		G.Room.Delete().
			Where(func(s *sql.Selector) {
				s.Where(sql.In(room.FieldID, notExistRooms...))
			}).ExecX(G_)
	}
}

var (
	G     *ent.Client
	G_    context.Context
	err   error
	Admin *ent.User
)

func Init() {
	G, err = ent.Open("mysql", config.GetDBMaster())
	if err != nil {
		log.Fatal(err)
	}
	G_ = context.Background()
	if err := G.Schema.Create(G_); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err = GetAdmin(); err != nil {
		log.Fatal(err)
	}

	go DestroyIllegalRooms()
}
