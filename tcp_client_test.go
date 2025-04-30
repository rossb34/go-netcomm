package netcomm

import (
	"bytes"
	"net"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		// want    *TCPClient
		wantErr bool
	}{
		{
			"Reject connection",
			args{address: "localhost:5555"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Connect(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTCPClient_Close(t *testing.T) {
	type fields struct {
		conn *net.TCPConn
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"nil connection",
			fields{nil},
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TCPClient{
				conn: tt.fields.conn,
			}
			if err := c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("TCPClient.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func newMockConn(b []byte) *MockConn {
	m := &MockConn{b, len(b)}
	return m
}

type MockConn struct {
	buf []byte
	n   int
}

func (c *MockConn) Close() error {
	return nil
}

func (c *MockConn) Write(b []byte) (n int, err error) {
	copy(c.buf, b)
	c.n = len(b)
	n = c.n
	err = nil
	return n, err
}

func (c *MockConn) Read(b []byte) (n int, err error) {
	copy(b, c.buf)
	n = c.n
	err = nil
	return n, err
}

func (c *MockConn) LocalAddr() net.Addr {
	return nil
}

func (c *MockConn) RemoteAddr() net.Addr {
	return nil
}

func (c *MockConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *MockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *MockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func TestTCPClient_Read(t *testing.T) {
	type fields struct {
		conn net.Conn
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantN    int
		wantErr  bool
		expected []byte
	}{
		{
			"Read",
			fields{newMockConn([]byte("hello"))},
			args{make([]byte, 10)},
			5,
			false,
			[]byte("hello"),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TCPClient{
				conn: tt.fields.conn,
			}
			gotN, err := c.Read(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TCPClient.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("TCPClient.Read() = %v, want %v", gotN, tt.wantN)
			}
			if !bytes.Equal(tt.expected, tt.args.b[:tt.wantN]) {
				t.Errorf("expected %v, actual %v", tt.expected, tt.args.b[:tt.wantN])
			}
		})
	}
}

func TestTCPClient_Write(t *testing.T) {
	type fields struct {
		conn net.Conn
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		{
			"Test Write",
			fields{newMockConn(make([]byte, 10))},
			args{[]byte("hello")},
			5,
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TCPClient{
				conn: tt.fields.conn,
			}
			gotN, err := c.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TCPClient.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("TCPClient.Write() = %v, want %v", gotN, tt.wantN)
			}

			// Now read the bytes back
			buf := make([]byte, 10)
			n, _ := c.Read(buf)
			if n != tt.wantN {
				t.Errorf("TCPClient.Read() = %v, want %v", n, tt.wantN)
			}
			if !bytes.Equal(tt.args.b, buf[:n]) {
				t.Errorf("expected = %v, actual %v", tt.args.b, buf[:n])
			}
		})
	}
}
