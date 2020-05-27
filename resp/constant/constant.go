package constant

const (
	BUFFER_SIZE int = 1024 * 256

	SIMPLE_STRING byte = '+'
	INTEGER       byte = ':'
	ERROR         byte = '-'
	BULK_STRING   byte = '$'
	ARRAY         byte = '*'

	SENTINEL_COMMAND string = "sentinel get-master-addr-by-name %s\r\n"
)

var (
	REDIS_QUIT []byte = []byte("*1\r\n$4\r\nquit\r\n")
)
