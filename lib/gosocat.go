package lib

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"

	"golang.org/x/net/websocket"
)

// Gosocat represetns a websocket cat connection
type Gosocat struct {
	io.Reader
	io.Writer
	conn *websocket.Conn
}

// New returns a new Gosocat connection
func New(t string) (*Gosocat, error) {
	u, err := url.Parse(t)
	if err != nil {
		return nil, err
	}

	c, err := websocket.Dial(t, "tcp", fmt.Sprintf("http://%s", u.Host))
	if err != nil {
		return nil, err
	}

	return &Gosocat{
		Reader: os.Stdin,
		Writer: os.Stdout,
		conn:   c,
	}, nil
}

// SetReader specifies the Gosocat reader
func (c *Gosocat) SetReader(r io.Reader) *Gosocat {
	c.Reader = r
	return c
}

// SetWriter specifies the Gosocat writer
func (c *Gosocat) SetWriter(w io.Writer) *Gosocat {
	c.Writer = w
	return c
}

// Start starts processing data on the connection
func (c *Gosocat) Start() <-chan error {
	errC := make(chan error)

	go func() {
		s := bufio.NewScanner(c.Reader)
		for s.Scan() {
			text := s.Text()
			websocket.Message.Send(c.conn, text)
		}
		if err := s.Err(); err != nil {
			errC <- err
		}
	}()

	go func() {
		for {
			var msg string
			if err := websocket.Message.Receive(c.conn, &msg); err != nil {
				errC <- err
			}
			fmt.Fprintln(c.Writer, msg)
		}
	}()

	return errC
}

// Close closes the connection
func (c *Gosocat) Close() error {
	return c.conn.Close()
}
