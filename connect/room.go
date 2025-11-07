/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 15:18
 */
package connect

import (
	"casher-server/proto"
	"context"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

const NoRoom = -1
const Plaza = 0

type Room struct {
	Id          int64
	OnlineCount int // room online user count
	rLock       sync.RWMutex
	drop        bool // make room is live
	next        *Channel
}

func NewRoom(roomId int64) *Room {
	room := new(Room)
	room.Id = roomId
	room.drop = false
	room.next = nil
	room.OnlineCount = 0
	return room
}

func (r *Room) Put(ch *Channel) (err error) {
	//doubly linked list
	r.rLock.Lock()
	defer r.rLock.Unlock()
	if !r.drop {
		if r.next != nil {
			r.next.Prev = ch
		}
		ch.Next = r.next
		ch.Prev = nil
		r.next = ch
		r.OnlineCount++
	} else {
		err = errors.New("room drop")
	}
	return
}

func (r *Room) Push(ctx context.Context, msg *proto.Msg) {
	r.rLock.RLock()
	defer r.rLock.RUnlock()
	// 从第一个用户开始推送
	var firstUserId int64
	ch := r.next
	if ch != nil {
		firstUserId = ch.UserId
		if err := ch.Push(ctx, msg); err != nil {
			fmt.Printf("push msg err:%s", err.Error())
		}
	}
	for ch = ch.Next; ch != nil; ch = ch.Next {
		if r.drop {
			break
		}
		fmt.Println("Push", ch.UserId)
		if firstUserId == ch.UserId {
			// 重复用户，不推送。防止出现重复推送
			fmt.Println("重复用户，不推送。防止出现重复推送")
			break
		}
		if err := ch.Push(ctx, msg); err != nil {
			fmt.Printf("push msg err:%s", err.Error())
		}
	}
}

func (r *Room) DeleteChannel(ch *Channel) bool {
	r.rLock.RLock()
	if ch.Next != nil {
		//if not footer
		ch.Next.Prev = ch.Prev
	}
	if ch.Prev != nil {
		// if not header
		ch.Prev.Next = ch.Next
	} else {
		r.next = ch.Next
	}
	r.OnlineCount--
	r.drop = false
	if r.OnlineCount <= 0 {
		if r.Id != Plaza {
			r.drop = true
		} else {
			r.OnlineCount = 0
		}
	}
	r.rLock.RUnlock()
	return r.drop
}
