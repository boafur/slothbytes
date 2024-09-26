package main

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func setupLogger() *log.Logger {
	// Override the default log level styles
	styles := log.DefaultStyles()
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().SetString("TEST ERR").Padding(0, 3, 0, 2).Foreground(lipgloss.Color("204")).Bold(true)
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().SetString("TEST FAIL").Padding(0, 2, 0, 2).Foreground(lipgloss.Color("204")).Bold(true)
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().SetString("TEST SUCCEED").Padding(0, 0, 0, 1).Foreground(lipgloss.Color("86")).Bold(true)
	styles.Keys["error"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Keys["didAssert"] = lipgloss.NewStyle().Foreground(lipgloss.Color("134"))
	styles.Values["error"] = lipgloss.NewStyle().Bold(true)
	styles.Values["didAssert"] = lipgloss.NewStyle().Bold(true)
	logger := log.New(os.Stderr)
	logger.SetStyles(styles)
	return logger
}

func logAssert(testInput any, testOutput any, shouldEqual any, testErr error) {
	log := setupLogger()
	didAssert := testOutput == shouldEqual
	if testErr != nil {
		log.Warn("", "testInput", testInput, "error", testErr, "shouldEqual", shouldEqual)
	} else if didAssert {
		log.Info("", "testInput", testInput, "testOutput", testOutput, "shouldEqual", shouldEqual, "didAssert", didAssert)
	} else {
		log.Error("", "testInput", testInput, "testOutput", testOutput, "shouldEqual", shouldEqual, "didAssert", didAssert)
	}
}
