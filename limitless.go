package limitless

import "net"

//Client to connect to limitless led bridge
type Client struct {
	hostname  string
	port      string
	udpClient *net.UDPConn
}

//NewClient create new client to connect to bridge
func NewClient(hostname string, port string) (*Client, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", hostname+":"+port)
	if err != nil {
		return &Client{}, err
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return &Client{}, err
	}
	return &Client{
		hostname:  hostname,
		port:      port,
		udpClient: conn,
	}, nil
}
