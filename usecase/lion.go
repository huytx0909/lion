package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"lion/model"
)

type LionUseCaseInterface interface {
	eat()
}

type LionUseCase struct {
	lionUC LionUseCaseInterface
}

func NewLionUseCase() *LionUseCase {
	return &LionUseCase{}
}

func (LionUseCase) eat() {
	fmt.Println("lion looking for prey!")
}

func RegisterLionHandler(
	router *message.Router,
	subscriber message.Subscriber,
	publisher message.Publisher,
	handlerName string,
	subscribeTopic string,
	publishTopic string,
) {
	router.AddHandler(
		handlerName,
		subscribeTopic,
		subscriber,
		publishTopic,
		publisher,
		Handler,
	)

}

func Handler(msg *message.Message) ([]*message.Message, error) {
	data := &model.Prey{}
	err := json.Unmarshal(msg.Payload, data)

	res := fmt.Sprintf("A/An %v has been spotted", data.Name)
	resData, err := json.Marshal(res)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	fmt.Println("Prey has been killed at A site!")

	resMsg := message.NewMessage(watermill.NewUUID(), resData)
	return message.Messages{resMsg}, nil
}
