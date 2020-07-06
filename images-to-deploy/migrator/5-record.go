package main

import (
	"github.com/jmckee46/deployer/flaw"
	"github.com/jmckee46/deployer/logger"
)

func record(state *state) {
	logger.Debug("migrator-record", state)

	_, err := state.PGClient.Exec(
		`INSERT INTO migrations (id, filename, created_at)
		 VALUES ($1, $2, $3)`,
		state.FileID,
		state.FileName,
		state.CreatedAt,
	)

	if err != nil {
		state.FlawError = flaw.From(err).Wrap("cannot record")

		logger.Panic("migrator-record-err", state)
	}

	logger.Info("migrator-record-success", state)
}
