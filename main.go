package main

import (
	"net"
	"strings"
	"bufio"
	"fmt"
	"log"
)

type server struct {
	users		[]user
	channels	[]channel
}

type user struct {
	nick		string
	name		string
	rname		string
	password	string
	conn 		net.Conn
	channels	[]channel

}

type channel struct {
	name 		string
	users		[]user
}

const (
	ERR_NOSUCHNICK			= 401
	ERR_NOSUCHSERVER		= 402
	ERR_NOSUCHCHANNEL		= 403
	ERR_CANNOTSENDTOCHAN	= 404
	ERR_TOOMANYCHANNELS		= 405
	ERR_WASNOSUCHNICK		= 406
	ERR_TOOMANYTARGETS		= 407
	ERR_NOSUCHSERVICE		= 408
	ERR_NOORIGIN		 	= 409
	ERR_NORECIPIENT			= 411
	ERR_NOTEXTTOSEND		= 412
	ERR_NOTOPLEVEL			= 413
	ERR_WILDTOPLEVEL		= 414
	ERR_BADMASK				= 415
	ERR_UNKNOWNCOMMAND		= 421
	ERR_NOMOTD				= 422
	ERR_NOADMININFO			= 423
	ERR_FILEERROR			= 424
	ERR_NONICKNAMEGIVEN		= 431
	ERR_ERRONEUSNICKNAME	= 432
	ERR_NICKNAMEINUSE	 	= 433
	ERR_NICKCOLLISION	 	= 436
	ERR_UNAVAILRESOURCE	 	= 437
	ERR_USERNOTINCHANNEL	= 441
	ERR_NOTONCHANNEL		= 442
	ERR_USERONCHANNEL		= 443
	ERR_NOLOGIN			 	= 444
	ERR_SUMMONDISABLED		= 445
	ERR_USERSDISABLED		= 446
	ERR_NOTREGISTERED		= 451
	ERR_NEEDMOREPARAMS		= 461
	ERR_ALREADYREGISTRED	= 462
	ERR_NOPERMFORHOST		= 463
	ERR_PASSWDMISMATCH		= 464
	ERR_YOUREBANNEDCREEP	= 465
	ERR_YOUWILLBEBANNED	 	= 466
	ERR_KEYSET				= 467
	ERR_CHANNELISFULL		= 471
	ERR_UNKNOWNMODE			= 472
	ERR_INVITEONLYCHAN		= 473
	ERR_BANNEDFROMCHAN		= 474
	ERR_BADCHANNELKEY		= 475
	ERR_BADCHANMASK			= 476
	ERR_NOCHANMODES			= 477
	ERR_BANLISTFULL			= 478
	ERR_NOPRIVILEGES		= 481
	ERR_CHANOPRIVSNEEDED	= 482
	ERR_CANTKILLSERVER		= 483
	ERR_RESTRICTED			= 484
	ERR_UNIQOPPRIVSNEEDED	= 485
	ERR_NOOPERHOST		 	= 491
	ERR_UMODEUNKNOWNFLAG	= 501
	ERR_USERSDONTMATCH		= 502
)

var ErrMsg = map[int]string{
	401 : "<nickname> :No such nick/channel",
	402 : 	"<server name> :No such server",
	403 : 	"<channel name> :No such channel",
	404 : 	"<channel name> :Cannot send to channel",
	405 : 	"<channel name> :You have joined too many channels",
	406 : 	"<nickname> :There was no such nickname",
	407 : "<target> :<error code> recipients. <abort message>",
	408 : 	"<service name> :No such service",
	409 : 	":No origin specified",
	411 : 	":No recipient given (<command>)",
	412 : 	":No text to send",
	413 : 	"<mask> :No toplevel domain specified",
	414 : 	"<mask> :Wildcard in toplevel domain",
	415 : "<mask> :Bad Server/host mask",
	421 : 	"<command> :Unknown command",
	422 : ":MOTD File is missing",
	423 : "<server> :No administrative info available",
	424 : ":File error doing <file op> on <file>",
	431 : ":No nickname given",
	432 : "<nick> :Erroneous nickname",
	433 : "<nick> :Nickname is already in use",
	436 : "<nick> :Nickname collision KILL from <user>@<host>",
	437 : "<nick/channel> :Nick/channel is temporarily unavailable",
	441 : "<nick> <channel> :They aren't on that channel",
	442 : "<channel> :You're not on that channel",
	443 : "<user> <channel> :is already on channel",
	444 : "<user> :User not logged in",
	445 : ":SUMMON has been disabled",
	446 : ":USERS has been disabled",
	451 : ":You have not registered",
	461 : "<command> :Not enough parameters",
	462 : 	":Unauthorized command (already registered)",
	463 : ":Your host isn't among the privileged",
	464 : ":Password incorrect",
	465 : 	":You are banned from this server",
	466 :  "YOU WILL BE BANNED MUTHAFUCKA",
	467 : "<channel> :Channel key already set",
	471 : "<channel> :Cannot join channel (+l)",
	472 : "<char> :is unknown mode char to me for <channel>",
	473 : "<channel> :Cannot join channel (+i)",
	474 : "<channel> :Cannot join channel (+b)",
	475 : "<channel> :Cannot join channel (+k)",
	476 : "<channel> :Bad Channel Mask",
	477 : "<channel> :Channel doesn't support modes",
	478 : "<channel> <char> :Channel list is full",
	481 : ":Permission Denied- You're not an IRC operator",
	482 : "<channel> :You're not channel operator",
	483 : ":You can't kill a server!",
	484 : ":Your connection is restricted!",
	485 : ":You're not the original channel operator",
	491 : ":No O-lines for your host",
	501 : ":Unknown MODE flag",
	502 : ":Cannot change mode for other users",
}



func PassParse(u_comm []string, conn net.Conn, usr *user){
	if cap(u_comm) > 1 {
		handle_pass(usr, u_comm[1])
		return
	} else {
		PutUserError("PASS", ERR_NEEDMOREPARAMS, conn)
	}
}


func NickParse(u_comm []string, conn net.Conn, srv *server, usr *user){
	if cap(u_comm) > 1 {
		handle_nick(srv, usr, u_comm[1])
		return
	} else {
		PutUserError("", ERR_NONICKNAMEGIVEN, conn)
	}





}


func PutUserError(str string, errorcode int, conn net.Conn){
	conn.Write([]byte(":localhost " + fmt.Sprintf("%d", errorcode) + " * " + str + ErrMsg[errorcode] + "\n"))
	//fmt.Fprint(conn, "localhost " + errorcode + " * " + str + ErrMsg[errorcode] + "\n")
}

func PutUserNames(str string, usernick string,  conn net.Conn){
	conn.Write([]byte(":localhost " + "353 " + str + " = :" + usernick +"\n"))
}

func PutUserChNames(str string, usernick string, channelname string , conn net.Conn){
	conn.Write([]byte(":localhost " + "353 " + str + " = " + channelname + " :" + usernick +"\n"))
}

func PutChan(nick string, channelname string, conn net.Conn){
	conn.Write([]byte(":localhost " + "322 " + nick + " " + channelname + "\n"))
}

func pr_endoflist_names(usernick string, channelname string, conn net.Conn) {
	conn.Write([]byte(":localhost " + "366 " + usernick + " " + channelname + " :End of /NAMES" +"\n"))
}

func pr_endoflist_list(usernick string, channelname string, conn net.Conn) {
	conn.Write([]byte(":localhost " + "366 " + usernick + " " + channelname +" :End of /NAMES" + "\n"))
}

func handle_nick(srv *server, usr *user, nick string) {
	if usr.nick != "" {
		for _, val := range srv.users {
			if val.nick == nick {
				PutUserError(usr.nick, ERR_NICKNAMEINUSE, usr.conn)
				return
			}
		}
		usr.nick = nick
	} else {
		for _, val := range srv.users {
			if val.nick == nick {
				PutUserError(usr.nick, ERR_NICKCOLLISION, usr.conn)
				return
			}
		}
		usr.nick = nick
	}
	if usr.name != "" && usr.password != "" {
		if check_usr_in_base(*srv, *usr) == 0 {
			srv.users = append(srv.users, *usr)
		}
	}
	return
}

func handle_pass(usr *user, pass string) {
	if usr.nick == "" && usr.name == "" {
		usr.password = pass
		return
	}
	PutUserError("", ERR_ALREADYREGISTRED, usr.conn)
	return
}

func handle_user(srv *server, usr *user, name, rname string) {



	if usr.name != "" {
		PutUserError("", ERR_ALREADYREGISTRED, usr.conn)
		return
	} else {
		usr.name = name
		usr.rname = rname
	}
	if usr.nick != "" && usr.password != "" {
		if check_usr_in_base(*srv, *usr) == 0 {
			srv.users = append(srv.users, *usr)
		}
	}
	return


}

func get_slice_from_channels(chann []channel) (slc []string) {
	for _, val := range chann {
		slc = append(slc, val.name)
	}
	return slc
}

func delete_user_from_channel(usr *user, chann string) {
	for i, val := range usr.channels {
		if val.name == chann {
			usr.channels[i] = usr.channels[len(usr.channels) - 1]
			usr.channels = usr.channels[:len(usr.channels) - 1]
		}
	}
}

func delete_user_from_channels(usr *user, chann []string) {
	for _, val := range chann {
		delete_user_from_channel(usr, val)
	}
}

func delete_user_from_server(srv *server, name string) {
	for i, usr := range srv.users {
		if usr.name == name {
			srv.users[i] = srv.users[len(srv.users) - 1]
			srv.users = srv.users[:len(srv.users) - 1]
			return
		}
	}
}

func check_usr_in_base(srv server, usr user) (exist int) {
	for _, val := range srv.users {
		if (usr.nick != "" && val.nick == usr.nick) || (usr.name != "" && val.name == usr.name) {
			return 1
		}
	}
	return 0
}

func server_have_channel(srv server, chann string) (exist int) {
	for _, val := range srv.channels {
		if val.name == chann {
			return 1
		}
	}
	return 0
}

func user_have_channel(usr user, chann string) (exist int) {
	for _, val := range usr.channels {
		if val.name == chann {
			return 1
		}
	}
	return 0
}

func handle_join(srv *server, usr *user, chann []string) {
	for _, val := range chann {
		if server_have_channel(*srv, val) == 0 {
			tmp := channel{name: val}
			tmp.users = append(tmp.users, *usr)
			srv.channels = append(srv.channels, tmp)
			usr.channels = append(usr.channels, tmp)
		} else if user_have_channel(*usr, val) == 0 {
			for _, chan_in_serv := range srv.channels {
				if chan_in_serv.name == val {
					break
				}
				usr.channels = append(usr.channels, chan_in_serv)
				chan_in_serv.users = append(chan_in_serv.users, *usr)
			}
		}
	}
}

func handle_names(srv server, usr user, fields []string) {
	if usr.name == "" || usr.nick == "" {
		PutUserError(usr.nick, ERR_NOTREGISTERED, usr.conn)
		return
	} else {
		var counter int
		counter  = 0
		for _, val := range fields {
			for _, val2 := range srv.channels {
				if val2.name == val {
					for _, val3 := range val2.users {
						PutUserChNames(usr.nick, val2.name, val3.name, usr.conn);
					}
				}
			}
			pr_endoflist_names(usr.nick, val, usr.conn)
			counter = 1
		}
		if counter == 0 {
			for _,val := range srv.users {
				PutUserNames(usr.nick, val.nick, usr.conn);
			}
			pr_endoflist_names(usr.nick, "", usr.conn)
		}
	}
}

func handle_part(usr *user, chann []string) {
	if usr.name == "" || usr.nick == "" {
		PutUserError(usr.nick, ERR_NOTREGISTERED, usr.conn)
		return
	} else if len(chann) > 0 {
		delete_user_from_channels(usr, chann)
	} else {
		PutUserError(usr.nick, ERR_NEEDMOREPARAMS, usr.conn)
	}
}

func handle_list(srv server, usr user, fields []string) {
	if usr.name == "" || usr.nick == "" {
		PutUserError(usr.nick, ERR_NOTREGISTERED, usr.conn)
		return
	} else {
		var counter int
		counter  = 0
		for _, val := range fields{
			PutChan(usr.nick, val, usr.conn)
			pr_endoflist_list(usr.nick, val, usr.conn)
			counter = 1
		}
		if counter == 0 {
			for _, val := range srv.channels {
				PutChan(usr.nick, val.name, usr.conn);
			}
			pr_endoflist_list(usr.nick, "", usr.conn)
		}
	}
}

func handle_privmsg(srv server, usr user, receiver string, messag string) {
	for _, val := range srv.users {
		if val.nick == receiver {
			val.conn.Write([]byte(messag + "\n"))
		}
	}
}

func UserParse(str string, conn net.Conn, srv *server, usr *user){
	var str1 string;
	var usrName string;
	var fullName string;
	if strings.Contains(str, ":"){
		a := strings.SplitN(str, ":", 2)
		str1 = a[0]
		fullName = a[1]
	}
	params := strings.Fields(str1);
	fmt.Println(params)
	if len(params) > 3{
		usrName = params[1];
		if fullName == ""{
			fullName = params[4]
		}
		handle_user(srv, usr, usrName, fullName);
		return
	} else {
		PutUserError("USER", ERR_NEEDMOREPARAMS, conn);
		return
	}

}


func JoinParse(u_comm []string, conn net.Conn, srv *server, usr *user){
	if (cap(u_comm) > 1){
		channels := strings.Split(u_comm[1], ",")
		handle_join(srv, usr, channels);
		return
	} else {
		PutUserError("", ERR_NEEDMOREPARAMS, conn)
		return
	}
}


func PartParse(u_comm []string, conn net.Conn, srv *server, usr *user){
	if (cap(u_comm) > 1){
		chann := strings.Split(u_comm[1], ",")
		handle_part(usr, chann);
		return
	} else {
		PutUserError("PART", ERR_NEEDMOREPARAMS, conn);
		return
	}
}


func NamesParse(u_comm []string, conn net.Conn, srv *server, usr *user){
	var channs []string
	if len(u_comm) == 1 {
		channs = append(channs, u_comm[0])
	} else {
		channs = strings.Split(u_comm[1], ",")
	}

	handle_names(*srv, *usr, channs);
	return
}


func ListParse(u_comm []string, conn net.Conn, srv *server, usr *user){
	var channs []string
	if len(u_comm) == 1 {
		channs = append(channs, u_comm[0])
	} else {
		channs = strings.Split(u_comm[1], ",")
	}
	handle_list(*srv, *usr, channs);
	return
}


func PrivmsgParse(str string, conn net.Conn, srv *server, usr *user){

	var str1 string;
	var reciever string;
	var message string;

	if strings.Contains(str, ":"){
		a := strings.SplitN(str, ":", 2)
		if cap(a) > 1{
			str1 = a[0]
			message = a[1]
		}
	}
	params := strings.Fields(str1);
	if cap(params) > 2 {
		reciever = params[1];
		if message != "" {
			message = params[2]
		}
		handle_privmsg(*srv, *usr, reciever, message);
		return
	} else {
		PutUserError("", ERR_NEEDMOREPARAMS, conn);
		return
	}
}


func StringParse(conn net.Conn, srv server) {
	var usr user
	usr.conn = conn

	for {
		var str string;
		str, _ = bufio.NewReader(conn).ReadString('\n')

		str = strings.TrimRight(str, "\n\r");
		u_comm := strings.SplitN(str, " ", 2)

		if len(u_comm) > 0 {
			u_comm[0] = strings.ToUpper(u_comm[0])
			fmt.Printf("%s\n", u_comm[0]);
			switch u_comm[0] {
			case "PASS":
				PassParse(strings.Fields(str), conn, &usr)
			case "NICK":
				NickParse(strings.Fields(str), conn, &srv, &usr)
			case "JOIN":
				JoinParse(strings.Fields(str), conn, &srv, &usr)
			case "PART":
				PartParse(strings.Fields(str), conn, &srv, &usr)
			case "NAMES":
				NamesParse(strings.Fields(str), conn, &srv, &usr)
			case "LIST":
				ListParse(strings.Fields(str), conn, &srv, &usr)
			case "PRIVMSG":
				PrivmsgParse(str, conn, &srv, &usr)
			case "USER":
				UserParse(str, conn, &srv, &usr)
			case "QUIT":
				if usr.password == "" {
					delete_user_from_channels(&usr, get_slice_from_channels(usr.channels))
					delete_user_from_server(&srv, usr.name)
				}
				break
			default:
				PutUserError(u_comm[0], ERR_UNKNOWNCOMMAND, conn);
			}
		}
	}
}



func main() {
	
	l, err := net.Listen("tcp", ":6667")
	var srv server

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go StringParse(conn, srv)
	}
}

