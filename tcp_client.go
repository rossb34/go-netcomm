package netcomm

import "net"

type TCPClient struct {
	conn net.Conn
}

// Closes the connection
func (c *TCPClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Write data to the tcp stream
func (c *TCPClient) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}

// Read data from the tcp stream
func (c *TCPClient) Read(b []byte) (n int, err error) {
	return c.conn.Read(b)
}

// Connect initiates a connection to the address.
//
// The address must be a valid IPv4 address.
func Connect(address string) (*TCPClient, error) {
	addr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp4", nil, addr)
	if err != nil {
		return nil, err
	}
	return &TCPClient{conn}, nil
}
