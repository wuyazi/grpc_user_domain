package aggregate

import (
	"fmt"
	"github.com/wuyazi/gddd"
	"github.com/wuyazi/grpc_user_domain/user_domain"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

type UserAggregate struct {
	gddd.AbstractAggregate
	Nickname   string    `json:"nickname"`
	Age        int64     `json:"age"`
	Gender     string    `json:"gender"`
	CreateTime time.Time `json:"create_time"`
}

func (b *UserAggregate) Create(ctx gddd.Context, nickname string) (err error) {
	nickname = strings.TrimSpace(nickname)
	if nickname == "" {
		err = fmt.Errorf("nickname is empty")
		return
	}

	ctx.Apply(b, &user_domain.UserCreated{
		Nickname:   nickname,
		CreateTime: timestamppb.Now(),
	})
	return
}

func (b *UserAggregate) UpdateNickname(ctx gddd.Context, nickname string) (err error) {
	nickname = strings.TrimSpace(nickname)
	if nickname == "" {
		err = fmt.Errorf("nickname is empty")
		return
	}

	ctx.Apply(b, &user_domain.UserChangedNickname{
		Nickname: nickname,
	})
	return
}

func (b *UserAggregate) UpdateAge(ctx gddd.Context, age int64) (err error) {
	if age == 0 {
		err = fmt.Errorf("age is 0")
		return
	}

	ctx.Apply(b, &user_domain.UserChangedAge{
		Age: age,
	})
	return
}

func (b *UserAggregate) UpdateGender(ctx gddd.Context, gender string) (err error) {
	gender = strings.TrimSpace(gender)
	if gender == "" {
		err = fmt.Errorf("gender is empty")
		return
	}

	ctx.Apply(b, &user_domain.UserChangedGender{
		Gender: gender,
	})
	return
}
