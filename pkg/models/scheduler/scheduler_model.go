package schedulermodel

import (
	"github.com/mynaparrot/plugnmeet-server/pkg/config"
	roommodel "github.com/mynaparrot/plugnmeet-server/pkg/models/room"
	"github.com/mynaparrot/plugnmeet-server/pkg/models/roomduration"
	"github.com/mynaparrot/plugnmeet-server/pkg/services/db"
	"github.com/mynaparrot/plugnmeet-server/pkg/services/livekit"
	natsservice "github.com/mynaparrot/plugnmeet-server/pkg/services/nats"
	"github.com/mynaparrot/plugnmeet-server/pkg/services/redis"
	"time"
)

type SchedulerModel struct {
	app         *config.AppConfig
	ds          *dbservice.DatabaseService
	rs          *redisservice.RedisService
	lk          *livekitservice.LivekitService
	natsService *natsservice.NatsService
	rm          *roommodel.RoomModel

	rmDuration  *roomdurationmodel.RoomDurationModel
	closeTicker chan bool
}

func New(app *config.AppConfig, ds *dbservice.DatabaseService, rs *redisservice.RedisService, lk *livekitservice.LivekitService) *SchedulerModel {
	if app == nil {
		app = config.GetConfig()
	}
	if ds == nil {
		ds = dbservice.New(app.DB)
	}
	if rs == nil {
		rs = redisservice.New(app.RDS)
	}
	if lk == nil {
		lk = livekitservice.New(app, rs)
	}

	return &SchedulerModel{
		app:         app,
		ds:          ds,
		rs:          rs,
		lk:          lk,
		rm:          roommodel.New(app, ds, rs, lk),
		rmDuration:  roomdurationmodel.New(app, rs),
		natsService: natsservice.New(app),
	}
}

func (m *SchedulerModel) StartScheduler() {
	m.closeTicker = make(chan bool)
	checkRoomDuration := time.NewTicker(5 * time.Second)
	defer checkRoomDuration.Stop()

	roomChecker := time.NewTicker(5 * time.Minute)
	defer roomChecker.Stop()

	for {
		select {
		case <-m.closeTicker:
			return
		case <-checkRoomDuration.C:
			m.checkRoomWithDuration()
			m.checkOnlineUsersStatus()
		case <-roomChecker.C:
			m.activeRoomChecker()
		}
	}
}
