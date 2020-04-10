package keybuilder

import uuid "github.com/satori/go.uuid"

// BuildCollectKey return a unique key, to collect keys who may be deleted if a request isn't success
func BuildCollectKey() string {
	return uuid.NewV4().String() + "-COLLECT"
}

// BuildAgentOccupyKey return AO:agent:ptfmCode:ptfmMjrTyp
func BuildAgentOccupyKey(agent, ptfmCode, ptfmMjrTyp string) string {
	return "AO:" + agent + ":" + ptfmCode + ":" + ptfmMjrTyp
}

// BuildMemberKey return Mem:member
func BuildMemberKey(member string) string {
	return "Mem:" + member
}

// BuildRolePermissionKey return PERMISSION:role:status:method
func BuildRolePermissionKey(role, status, method string) string {
	return "PERMISSION:" + role + ":" + status + ":" + method
}

// BuildAccessKey return ACCESS:user
func BuildAccessKey(user string) string {
	return "ACCESS:" + user
}

// BuildRefreshKey return REFRESH:user
func BuildRefreshKey(user string) string {
	return "REFRESH:" + user
}
