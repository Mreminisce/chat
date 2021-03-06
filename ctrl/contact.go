package ctrl

import (
	"chat/args"
	"chat/model"
	"chat/service"
	"chat/util"
	"net/http"
)

var contactService service.ContactService

func CreateCommunity(w http.ResponseWriter, req *http.Request) {
	var arg model.Community
	util.Bind(req, &arg)
	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, com, "")
	}
}

func JoinCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	util.Bind(req, &arg)
	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)
	//todo 刷新用户的群组信息
	AddGroupId(arg.Userid, arg.Dstid)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, nil, "")
	}
}

//加载他的群
func LoadCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	util.Bind(req, &arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespOkList(w, comunitys, len(comunitys))
}

//添加好友
func Addfriend(w http.ResponseWriter, req *http.Request) {
	// request.ParseForm()
	// mobile := request.PostForm.Get("mobile")
	// passwd := request.PostForm.Get("passwd")
	var arg args.ContactArg
	util.Bind(req, &arg)
	err := contactService.AddFriend(arg.Userid, arg.Dstid)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, nil, "好友添加成功")
	}
}

//加载个人的好友
func LoadFriend(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	util.Bind(req, &arg)
	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(w, users, len(users))
}
