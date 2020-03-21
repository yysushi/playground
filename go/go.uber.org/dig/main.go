package main

import (
	"encoding/json"
	"log"
	"os"

	"go.uber.org/dig"
)

// Config ...
type Config struct {
	Prefix string
}

func main() {
	c := dig.New()

	// Provide a Config object. This can fail to decode.
	err := c.Provide(func() (*Config, error) {
		// In a real program, the configuration will probably be read from a
		// file.
		var cfg Config
		err := json.Unmarshal([]byte(`{"prefix": "[foo] "}`), &cfg)
		return &cfg, err
	})
	if err != nil {
		panic(err)
	}

	// Provide a way to build the logger based on the configuration.
	err = c.Provide(func(cfg *Config) *log.Logger {
		return log.New(os.Stdout, cfg.Prefix, 0)
	})
	if err != nil {
		panic(err)
	}

	// Invoke a function that requires the logger, which in turn builds the
	// Config first.
	err = c.Invoke(func(l *log.Logger) {
		l.Print("You've been invoked")
	})
	if err != nil {
		panic(err)
	}

	// // graph
	// // f, err := os.OpenFile("graph.dot", os.O_WRONLY, 0755)
	// f, err := os.Create("graph.dot")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// err = dig.Visualize(c, f)
	// if err != nil {
	// 	panic(err)
	// }
}
