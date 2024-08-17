package natsservice

import (
	"fmt"
	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nats.go/jetstream"
)

func (s *NatsService) CreatePublicChatConsumer(roomId, userId string) (jwt.StringList, error) {
	_, err := s.js.CreateOrUpdateConsumer(s.ctx, roomId, jetstream.ConsumerConfig{
		Durable: fmt.Sprintf("%s:%s", s.app.NatsInfo.Subjects.ChatPublic, userId),
		FilterSubjects: []string{
			fmt.Sprintf("%s:%s.>", roomId, s.app.NatsInfo.Subjects.ChatPublic),
		},
	})
	if err != nil {
		return nil, err
	}

	permission := jwt.StringList{
		fmt.Sprintf("$JS.API.CONSUMER.INFO.%s.%s:%s", roomId, s.app.NatsInfo.Subjects.ChatPublic, userId),
		fmt.Sprintf("$JS.API.CONSUMER.MSG.NEXT.%s.%s:%s", roomId, s.app.NatsInfo.Subjects.ChatPublic, userId),
		fmt.Sprintf("%s:%s.%s", roomId, s.app.NatsInfo.Subjects.ChatPublic, userId),
		fmt.Sprintf("$JS.ACK.%s.%s:%s.>", roomId, s.app.NatsInfo.Subjects.ChatPublic, userId),
	}

	return permission, nil
}

func (s *NatsService) CreatePrivateChatConsumer(roomId, userId string) (jwt.StringList, error) {
	_, err := s.js.CreateOrUpdateConsumer(s.ctx, roomId, jetstream.ConsumerConfig{
		Durable: fmt.Sprintf("%s:%s", s.app.NatsInfo.Subjects.ChatPrivate, userId),
		FilterSubjects: []string{
			fmt.Sprintf("%s:%s.%s.>", roomId, s.app.NatsInfo.Subjects.ChatPrivate, userId),
		},
	})
	if err != nil {
		return nil, err
	}

	permission := jwt.StringList{
		fmt.Sprintf("$JS.API.CONSUMER.INFO.%s.%s:%s", roomId, s.app.NatsInfo.Subjects.ChatPrivate, userId),
		fmt.Sprintf("$JS.API.CONSUMER.MSG.NEXT.%s.%s:%s", roomId, s.app.NatsInfo.Subjects.ChatPrivate, userId),
		fmt.Sprintf("%s:%s.*.%s", roomId, s.app.NatsInfo.Subjects.ChatPrivate, userId),
		fmt.Sprintf("$JS.ACK.%s.%s:%s.>", roomId, s.app.NatsInfo.Subjects.ChatPrivate, userId),
	}

	return permission, nil
}

func (s *NatsService) CreateSystemPublicConsumer(roomId, userId string) (jwt.StringList, error) {
	_, err := s.js.CreateOrUpdateConsumer(s.ctx, roomId, jetstream.ConsumerConfig{
		Durable:       fmt.Sprintf("%s:%s", s.app.NatsInfo.Subjects.SystemPublic, userId),
		DeliverPolicy: jetstream.DeliverNewPolicy,
		FilterSubjects: []string{
			fmt.Sprintf("%s:%s.>", roomId, s.app.NatsInfo.Subjects.SystemPublic),
		},
	})
	if err != nil {
		return nil, err
	}

	permission := jwt.StringList{
		fmt.Sprintf("$JS.API.CONSUMER.INFO.%s.%s:%s", roomId, s.app.NatsInfo.Subjects.SystemPublic, userId),
		fmt.Sprintf("$JS.API.CONSUMER.MSG.NEXT.%s.%s:%s", roomId, s.app.NatsInfo.Subjects.SystemPublic, userId),
		fmt.Sprintf("$JS.ACK.%s.%s:%s.>", roomId, s.app.NatsInfo.Subjects.SystemPublic, userId),
	}

	return permission, nil
}

func (s *NatsService) CreateSystemPrivateConsumer(roomId, userId string) (jwt.StringList, error) {
	_, err := s.js.CreateOrUpdateConsumer(s.ctx, roomId, jetstream.ConsumerConfig{
		Durable:       fmt.Sprintf("%s:%s", s.app.NatsInfo.Subjects.SystemPrivate, userId),
		DeliverPolicy: jetstream.DeliverNewPolicy,
		FilterSubjects: []string{
			fmt.Sprintf("%s:%s.%s.>", roomId, s.app.NatsInfo.Subjects.SystemPrivate, userId),
		},
	})
	if err != nil {
		return nil, err
	}

	permission := jwt.StringList{
		fmt.Sprintf("$JS.API.CONSUMER.INFO.%s.%s:%s", roomId, s.app.NatsInfo.Subjects.SystemPrivate, userId),
		fmt.Sprintf("$JS.API.CONSUMER.MSG.NEXT.%s.%s:%s", roomId, s.app.NatsInfo.Subjects.SystemPrivate, userId),
		fmt.Sprintf("$JS.ACK.%s.%s:%s.>", roomId, s.app.NatsInfo.Subjects.SystemPrivate, userId),
	}

	return permission, nil
}

func (s *NatsService) CreateWhiteboardConsumer(roomId, userId string) (jwt.StringList, error) {
	_, err := s.js.CreateOrUpdateConsumer(s.ctx, roomId, jetstream.ConsumerConfig{
		Durable: fmt.Sprintf("%s:%s", s.app.NatsInfo.Subjects.Whiteboard, userId),
		FilterSubjects: []string{
			fmt.Sprintf("%s:%s.>", roomId, s.app.NatsInfo.Subjects.Whiteboard),
		},
	})
	if err != nil {
		return nil, err
	}

	permission := jwt.StringList{
		fmt.Sprintf("$JS.API.CONSUMER.INFO.%s.%s:%s", roomId, s.app.NatsInfo.Subjects.Whiteboard, userId),
		fmt.Sprintf("$JS.API.CONSUMER.MSG.NEXT.%s.%s:%s", roomId, s.app.NatsInfo.Subjects.Whiteboard, userId),
		fmt.Sprintf("$JS.ACK.%s.%s:%s.>", roomId, s.app.NatsInfo.Subjects.Whiteboard, userId),
	}

	return permission, nil
}
