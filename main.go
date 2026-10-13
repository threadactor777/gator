package main

import (
	"fmt"

	"github.com/threadactor777/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Errorf("Error:", err)
	}

	err = cfg.SetUser("Colby")
	if err != nil {
		fmt.Errorf("Error:", err)
	}

	updatedCfg, err := config.Read()
	if err != nil {
		fmt.Errorf("Error:", err)
	}

	fmt.Printf("Config Contents:", updatedCfg)

}
