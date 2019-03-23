package solr

// ConnectionOption - This information needs to be passed during connection creation
type ConnectionOption struct {
	Host string
	Port string
	Root string
	Core string
}

// Connection structure
type Connection struct {
	host       string
	port       string
	root       string
	core       string
	protocol   string
	logLevel   string
	searchPath string
}

// NewConnection - This method will return a new connection
func NewConnection(connOpt ConnectionOption) *Connection {
	newConn := Connection{connOpt.Host, connOpt.Port, connOpt.Root, connOpt.Core, "http", "Error", "select"}
	return &newConn
}

// UseHTTPS - This function enables HTTPS
func (c *Connection) UseHTTPS() *Connection {
	c.protocol = "https"
	return c
}

// SetLogLevel - Sets the log level
// Valid value - Error, Debug, Info
func (c *Connection) SetLogLevel(loglevel string) *Connection {
	c.logLevel = loglevel
	return c
}

// SetSearchPath - Sets the search handler
// if required any custom handler then set it
func (c *Connection) SetSearchPath(searchPath string) *Connection {
	c.searchPath = searchPath
	return c
}

// MakeHostURL - This builds the URL
func (c *Connection) makeHostURL() string {
	if c.port != "" {
		return c.protocol + "://" + c.host + ":" + c.port
	} else {
		return c.protocol + "://" + c.host
	}
}

// MakeRequestURL - This builds the request URL
func (c *Connection) MakeRequestURL() string {
	return c.makeHostURL() + "/" + c.root + "/" + c.core + "/" + c.searchPath
}
