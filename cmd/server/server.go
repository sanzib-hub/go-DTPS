package tcpServer

import (
	"fmt"
	"os"
)

func StartTCPServer() error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	if host == "" || port == "" {
		return fmt.Errorf("missing HOST or PORT environment variable")
	}

	address := fmt.Sprintf("%s:%s", host, port)
	fmt.Println("Server will start on:", address)

	// Placeholder for TCP setup (e.g., net.Listen etc.)
	return nil
}
