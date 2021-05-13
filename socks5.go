package proxyM

import (
	"encoding/binary"
	"errors"
	connect "github.com/demonoid81/proxyM/conn"
	"github.com/demonoid81/proxyM/log"
	"github.com/demonoid81/proxyM/pool"
	"github.com/demonoid81/proxyM/proxy"
	"github.com/demonoid81/proxyM/util"
	"net"
	"strconv"
	"time"
)

const (
	socksVer5 = 0x05

	verIndex     = 0
	nMethodIndex = 1
	methodIndex  = 2

	cmdIndex       = 1
	rsvIndex       = cmdIndex + 1
	atypIndex      = rsvIndex + 1
	addrIndex      = atypIndex + 1
	domianLenIndex = atypIndex + 1
	ipv4PortIndex  = addrIndex + 4
	ipv6PortIndex  = addrIndex + 16
)

func SocksHandle(co net.Conn) {
	log.Logger.Debug("[SOCKS] start proxyM.IConn wrap net.Conn")
	conn, err := connect.NewDefaultConn(co, connect.TCP)
	if err != nil {
		log.Logger.Errorf("proxyM.IConn wrap net.Conn failed: %v", err)
		return
	}
	log.Logger.Debugf("[SOCKS] [ID:%d] proxyM.IConn wrap net.Conn success ", conn.GetID())
	log.Logger.Debugf("[SOCKS] [ID:%d] start handShake", conn.GetID())
	err = handShake(conn)
	if err != nil {
		log.Logger.Errorf("[SOCKS] [ID:%d] handShake failed: %s", conn.GetID(), err.Error())
		return
	}
	req, err := parseRequest(conn)
	if err != nil {
		log.Logger.Errorf("[SOCKS] [ID:%d] parseRequest failed: %s", conn.GetID(), err.Error())
		return
	}
	req.protocol = ProtocolSocks
	req.target = req.Host()
	_, err = conn.Write([]byte{socksVer5, 0x00, 0x00, AddrTypeIPv4, 0x00, 0x00, 0x00, 0x00, 0x08, 0x43})
	if err != nil {
		log.Logger.Errorf("[SOCKS] [ID:%d] send connection confirmation: %s", conn.GetID(), err.Error())
		return
	}

	//inner controller domain
	if req.addr == ControllerDomain {
		port, err := strconv.ParseUint(ControllerPort, 10, 16)
		if err == nil {
			req.ip = []byte{127, 0, 0, 1}
			req.port = uint16(port)
		}
	}

	// Whitelist judgment
	if IsPass(req.addr, req.ip.String(), strconv.Itoa(int(req.port))) {
		s, _ := proxy.GetServer(proxy.ProxyDirect)
		sc, err := s.Conn(req)
		if err != nil {
			log.Logger.Errorf("[SOCKS] [ID:%d] ConnectToServer failed [%s] err: %s", conn.GetID(), req.Host(), err.Error())
		}
		direct := &DirectChannel{}
		direct.Transport(conn, sc)
		return
	}

	//RuleFilter by Rules and DNS
	rule, s, err := FilterByReq(req)
	record := &Record{
		ID:       util.NextID(),
		Protocol: req.protocol,
		Created:  time.Now(),
		Status:   RecordStatusActive,
		URL:      req.target,
		Rule:     rule,
		Proxy:    s,
	}
	if err != nil {
		if err == ErrorReject {
			log.Logger.Errorf("[SOCKS] [ID:%d] ConnectToServer failed [%s] err: %s", conn.GetID(), req.Host(), err)
		}
		record.Status = RecordStatusCompleted
		boxChan <- &Box{Op: RecordAppend, Value: record, ID: record.ID}
		conn.Close()
	} else {
		//connnet to server
		log.Logger.Debugf("[SOCKS] [ID:%d] Start connect to Server [%s]", conn.GetID(), s.Name)
		sc, err := s.Conn(req)
		if err != nil {
			log.Logger.Errorf("[SOCKS] [ID:%d] ConnectToServer failed [%s] err: %s", conn.GetID(), req.Host(), err.Error())
			return
		}
		log.Logger.Debugf("[SOCKS] [ID:%d] Server [%s] Connected success", conn.GetID(), s.Name)
		log.Logger.Debugf("[HTTP] [ClientConnID:%d] Bind to [ServerConnID:%d]", conn.GetID(), sc.GetID())
		sc.SetRecordID(record.ID)
		boxChan <- &Box{Op: RecordAppend, Value: record, ID: record.ID}
		direct := &DirectChannel{}
		direct.Transport(conn, sc)
		boxChan <- &Box{record.ID, RecordStatus, RecordStatusCompleted}
	}
}

//socks
func handShake(conn net.Conn) error {
	buf := pool.GetBuf()
	_, err := conn.Read(buf)
	if err != nil {
		return err
	}
	if buf[verIndex] != socksVer5 {
		return errors.New("socks version not supported")
	}
	//todo get methods

	//return supported methods
	conn.Write([]byte{0x05, 0x00})
	pool.PutBuf(buf)
	return nil
}


func parseRequest(conn connect.IConn) (*SocksRequest, error) {
	buf := pool.GetBuf()
	_, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	request := &SocksRequest{
		ver:    uint8(buf[verIndex]),
		cmd:    uint8(buf[cmdIndex]),
		rsv:    uint8(buf[rsvIndex]),
		atyp:   uint8(buf[atypIndex]),
		connID: conn.GetID(),
	}
	switch request.atyp {
	case AddrTypeIPv4:
		request.ip = buf[atypIndex+1 : ipv4PortIndex]
		request.port = binary.BigEndian.Uint16(buf[ipv4PortIndex : ipv4PortIndex+2])
		if request.cmd == CmdUDP {
			request.data = buf[ipv4PortIndex+2:]
		}
	case AddrTypeDomain:
		end := buf[domianLenIndex] + 1 + domianLenIndex
		request.addr = string(buf[domianLenIndex+1 : end])
		request.port = binary.BigEndian.Uint16(buf[end : end+2])
		if request.cmd == CmdUDP {
			request.data = buf[end+2:]
		}
	case AddrTypeIPv6:
		request.ip = buf[atypIndex+1 : ipv6PortIndex]
		request.port = binary.BigEndian.Uint16(buf[ipv6PortIndex : ipv6PortIndex+2])
		if request.cmd == CmdUDP {
			request.data = buf[ipv6PortIndex+2:]
		}
	}
	if request.cmd != CmdUDP {
		pool.PutBuf(buf)
	}
	return request, nil
}
