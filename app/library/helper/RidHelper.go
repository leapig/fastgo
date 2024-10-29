package helper

import (
	"errors"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	Permission int64 = iota
	PermissionGroupPermission
	PermissionGroup
	RolePermission
	Role
	RoleGroup
	RoleGroupPermission
	UserPermission
	Enterprise
	UserT
	FilesT
	MinioUrlFileName
	Department
	UserCredentials
	Member
	UserProfession
	Position
	ApplicationRecord
	PageResource
	InterfaceResource
	PageInterface
	MenuResource
	User
	UserClient
	Sms
	Project
	ProjectEnterpriseRelation
	ProjectMember
	ProjectPost
	UserEntry
	EnterpriseUser

	UserRealNameAuthenticationLog

	EnterpriseAreaPermission

	EnterpriseUserAttachment

	TurnoverRecord

	TemporaryWorker
	LogisticsPurchase

	LogisticsWarehousing
	LogisticsDelivery
	WorkSchedule
	WorkSchedulePost

	PatrolPoint

	PatrolRoute

	PatrolRoutePoint
	PatrolRouteRecord
	PatrolRouteRecordPoint
	InspectionPoint
	InspectionPlan
	InspectionPlanUser
	InspectionPlanPoint
)

const epoch int64 = 1633017600 //自定义纪元

var (
	vpc     int64
	ipc     int64
	nodeMap *SMap
	once    sync.Once
)

type SMap struct {
	sync.RWMutex
	Map map[int64]*Business
}

func (l *SMap) readMap(key int64) *Business {
	l.RLock()
	value, ok := l.Map[key]
	l.RUnlock()
	if !ok && value == nil {
		l.writeMap(key, &Business{})
		return l.readMap(key)
	}
	return value
}

func (l *SMap) writeMap(key int64, value *Business) {
	l.Lock()
	l.Map[key] = value
	l.Unlock()
}

func Rid(table int64) string {
	for nodeMap == nil {
		once.Do(initId)
	}
	n := nodeMap.readMap(table)
	n.node = table
	res, err := snowflake(n)
	for err != nil {
		res, err = snowflake(n)
	}
	return strconv.Itoa(int(res))
}

func GetRid(table int64) int64 {
	for nodeMap == nil {
		once.Do(initId)
	}
	n := nodeMap.readMap(table)
	n.node = table
	res, err := snowflake(n)
	for err != nil {
		res, err = snowflake(n)
	}
	return res
}

type Business struct {
	node     int64
	lastTime int64
	lastId   int64
	step     int64
}

func snowflake(n *Business) (int64, error) {
	once.Do(initId)
	now := time.Now().Unix()
	if now == n.lastTime {
		n.step = n.step + 1
		if n.step > 4095 {
			return 0, errors.New("序列号溢出")
		}
	} else {
		n.step = 0
	}
	n.lastTime = now
	res := (now-epoch)<<25 | (vpc << 24) | (ipc << 20) | (n.node << 12) | n.step
	//res := (3501244799-epoch)<<25 | (vpc << 24) | (ipc << 20) | (n.node << 12) | n.step
	if n.lastId >= res {
		res += 1
	}
	n.lastId = res
	nodeMap.writeMap(n.node, n)
	return res, nil
}

func initId() {
	//vpcStr := os.Getenv("VPC")
	//if vpcStr != "" {
	//	vpcInt, err := strconv.Atoi(vpcStr)
	//	if err != nil {
	//		panic(err)
	//	}
	//	vpc = int64(vpcInt)
	//}
	vpc = 1
	ips := getLocalIP()
	if len(ips) == 0 {
		panic(errors.New("获取本地IP失败"))
	}
	ipArr := strings.Split(ips[0], ".")
	if len(ipArr) == 4 {
		ipInt, _ := strconv.Atoi(ipArr[3])
		ipc = int64(ipInt / 16)
	}

	nodeMap = &SMap{
		Map: make(map[int64]*Business),
	}
}

func getLocalIP() []string {
	var ipStr []string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return ipStr
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addr, _ := netInterfaces[i].Addrs()
			for _, address := range addr {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						ipStr = append(ipStr, ipNet.IP.String())
					}
				}
			}
		}
	}
	return ipStr
}
