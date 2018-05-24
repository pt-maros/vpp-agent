// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package session

import "reflect"

var Types = map[string]reflect.Type{
	"AcceptSession":              reflect.TypeOf((*AcceptSession)(nil)).Elem(),
	"AcceptSessionReply":         reflect.TypeOf((*AcceptSessionReply)(nil)).Elem(),
	"AppNamespaceAddDel":         reflect.TypeOf((*AppNamespaceAddDel)(nil)).Elem(),
	"AppNamespaceAddDelReply":    reflect.TypeOf((*AppNamespaceAddDelReply)(nil)).Elem(),
	"ApplicationAttach":          reflect.TypeOf((*ApplicationAttach)(nil)).Elem(),
	"ApplicationAttachReply":     reflect.TypeOf((*ApplicationAttachReply)(nil)).Elem(),
	"ApplicationDetach":          reflect.TypeOf((*ApplicationDetach)(nil)).Elem(),
	"ApplicationDetachReply":     reflect.TypeOf((*ApplicationDetachReply)(nil)).Elem(),
	"ApplicationTLSCertAdd":      reflect.TypeOf((*ApplicationTLSCertAdd)(nil)).Elem(),
	"ApplicationTLSCertAddReply": reflect.TypeOf((*ApplicationTLSCertAddReply)(nil)).Elem(),
	"ApplicationTLSKeyAdd":       reflect.TypeOf((*ApplicationTLSKeyAdd)(nil)).Elem(),
	"ApplicationTLSKeyAddReply":  reflect.TypeOf((*ApplicationTLSKeyAddReply)(nil)).Elem(),
	"BindSock":                   reflect.TypeOf((*BindSock)(nil)).Elem(),
	"BindSockReply":              reflect.TypeOf((*BindSockReply)(nil)).Elem(),
	"BindURI":                    reflect.TypeOf((*BindURI)(nil)).Elem(),
	"BindURIReply":               reflect.TypeOf((*BindURIReply)(nil)).Elem(),
	"ConnectSession":             reflect.TypeOf((*ConnectSession)(nil)).Elem(),
	"ConnectSessionReply":        reflect.TypeOf((*ConnectSessionReply)(nil)).Elem(),
	"ConnectSock":                reflect.TypeOf((*ConnectSock)(nil)).Elem(),
	"ConnectSockReply":           reflect.TypeOf((*ConnectSockReply)(nil)).Elem(),
	"ConnectURI":                 reflect.TypeOf((*ConnectURI)(nil)).Elem(),
	"ConnectURIReply":            reflect.TypeOf((*ConnectURIReply)(nil)).Elem(),
	"DisconnectSession":          reflect.TypeOf((*DisconnectSession)(nil)).Elem(),
	"DisconnectSessionReply":     reflect.TypeOf((*DisconnectSessionReply)(nil)).Elem(),
	"MapAnotherSegment":          reflect.TypeOf((*MapAnotherSegment)(nil)).Elem(),
	"MapAnotherSegmentReply":     reflect.TypeOf((*MapAnotherSegmentReply)(nil)).Elem(),
	"ResetSession":               reflect.TypeOf((*ResetSession)(nil)).Elem(),
	"ResetSessionReply":          reflect.TypeOf((*ResetSessionReply)(nil)).Elem(),
	"SessionEnableDisable":       reflect.TypeOf((*SessionEnableDisable)(nil)).Elem(),
	"SessionEnableDisableReply":  reflect.TypeOf((*SessionEnableDisableReply)(nil)).Elem(),
	"SessionRuleAddDel":          reflect.TypeOf((*SessionRuleAddDel)(nil)).Elem(),
	"SessionRuleAddDelReply":     reflect.TypeOf((*SessionRuleAddDelReply)(nil)).Elem(),
	"SessionRulesDetails":        reflect.TypeOf((*SessionRulesDetails)(nil)).Elem(),
	"SessionRulesDump":           reflect.TypeOf((*SessionRulesDump)(nil)).Elem(),
	"UnbindSock":                 reflect.TypeOf((*UnbindSock)(nil)).Elem(),
	"UnbindSockReply":            reflect.TypeOf((*UnbindSockReply)(nil)).Elem(),
	"UnbindURI":                  reflect.TypeOf((*UnbindURI)(nil)).Elem(),
	"UnbindURIReply":             reflect.TypeOf((*UnbindURIReply)(nil)).Elem(),
	"UnmapSegment":               reflect.TypeOf((*UnmapSegment)(nil)).Elem(),
	"UnmapSegmentReply":          reflect.TypeOf((*UnmapSegmentReply)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewAcceptSession":              reflect.ValueOf(NewAcceptSession),
	"NewAcceptSessionReply":         reflect.ValueOf(NewAcceptSessionReply),
	"NewAppNamespaceAddDel":         reflect.ValueOf(NewAppNamespaceAddDel),
	"NewAppNamespaceAddDelReply":    reflect.ValueOf(NewAppNamespaceAddDelReply),
	"NewApplicationAttach":          reflect.ValueOf(NewApplicationAttach),
	"NewApplicationAttachReply":     reflect.ValueOf(NewApplicationAttachReply),
	"NewApplicationDetach":          reflect.ValueOf(NewApplicationDetach),
	"NewApplicationDetachReply":     reflect.ValueOf(NewApplicationDetachReply),
	"NewApplicationTLSCertAdd":      reflect.ValueOf(NewApplicationTLSCertAdd),
	"NewApplicationTLSCertAddReply": reflect.ValueOf(NewApplicationTLSCertAddReply),
	"NewApplicationTLSKeyAdd":       reflect.ValueOf(NewApplicationTLSKeyAdd),
	"NewApplicationTLSKeyAddReply":  reflect.ValueOf(NewApplicationTLSKeyAddReply),
	"NewBindSock":                   reflect.ValueOf(NewBindSock),
	"NewBindSockReply":              reflect.ValueOf(NewBindSockReply),
	"NewBindURI":                    reflect.ValueOf(NewBindURI),
	"NewBindURIReply":               reflect.ValueOf(NewBindURIReply),
	"NewConnectSession":             reflect.ValueOf(NewConnectSession),
	"NewConnectSessionReply":        reflect.ValueOf(NewConnectSessionReply),
	"NewConnectSock":                reflect.ValueOf(NewConnectSock),
	"NewConnectSockReply":           reflect.ValueOf(NewConnectSockReply),
	"NewConnectURI":                 reflect.ValueOf(NewConnectURI),
	"NewConnectURIReply":            reflect.ValueOf(NewConnectURIReply),
	"NewDisconnectSession":          reflect.ValueOf(NewDisconnectSession),
	"NewDisconnectSessionReply":     reflect.ValueOf(NewDisconnectSessionReply),
	"NewMapAnotherSegment":          reflect.ValueOf(NewMapAnotherSegment),
	"NewMapAnotherSegmentReply":     reflect.ValueOf(NewMapAnotherSegmentReply),
	"NewResetSession":               reflect.ValueOf(NewResetSession),
	"NewResetSessionReply":          reflect.ValueOf(NewResetSessionReply),
	"NewSessionEnableDisable":       reflect.ValueOf(NewSessionEnableDisable),
	"NewSessionEnableDisableReply":  reflect.ValueOf(NewSessionEnableDisableReply),
	"NewSessionRuleAddDel":          reflect.ValueOf(NewSessionRuleAddDel),
	"NewSessionRuleAddDelReply":     reflect.ValueOf(NewSessionRuleAddDelReply),
	"NewSessionRulesDetails":        reflect.ValueOf(NewSessionRulesDetails),
	"NewSessionRulesDump":           reflect.ValueOf(NewSessionRulesDump),
	"NewUnbindSock":                 reflect.ValueOf(NewUnbindSock),
	"NewUnbindSockReply":            reflect.ValueOf(NewUnbindSockReply),
	"NewUnbindURI":                  reflect.ValueOf(NewUnbindURI),
	"NewUnbindURIReply":             reflect.ValueOf(NewUnbindURIReply),
	"NewUnmapSegment":               reflect.ValueOf(NewUnmapSegment),
	"NewUnmapSegmentReply":          reflect.ValueOf(NewUnmapSegmentReply),
}

var Variables = map[string]reflect.Value{}

var Consts = map[string]reflect.Value{}
