package dial

import (
	"time"

	"github.com/MyTempoESP/serial"
)

type Arduino struct {
	config *serial.Config
	port   *serial.Port
}

func NewArduino() (ino Arduino, err error) {

	// Configure the serial port
	ino.config = &serial.Config{
		Name: "/dev/ttyUSB0", // Replace with your Arduino port (e.g., COM3 on Windows)
		Baud: 115200,         // Match the baud rate of your Arduino sketch
	}

	// Open the serial port
	ino.port, err = serial.OpenPort(ino.config)

	// Allow time for Arduino to reset
	time.Sleep(2 * time.Second)

	return
}

func (ino *Arduino) Close() {

	ino.port.Close()
}

func (ino *Arduino) Send(msg string) (err error) {

	_, err = ino.port.Write([]byte(msg))

	return
}
