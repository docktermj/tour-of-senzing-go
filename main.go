/*
 */
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/senzing/g2-sdk-go-mock/g2config"
	"github.com/senzing/g2-sdk-go-mock/g2configmgr"
	"github.com/senzing/g2-sdk-go-mock/g2engine"
	"github.com/senzing/g2-sdk-go/g2api"
	"github.com/senzing/go-common/truthset"
	"github.com/senzing/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Output the error and exit program.
func errorExit(message string, err error) {
	fmt.Printf("Exit with error: %s   Error: %v\n", message, err)
	os.Exit(1)
}

// Add the CUSTOMERS Datasource.
func loadDatasources(ctx context.Context, g2Config g2api.G2config, g2Configmgr g2api.G2configmgr, g2Engine g2api.G2engine) {
	now := time.Now()

	// Using G2Config: Create a default configuration in memory.

	configHandle, err := g2Config.Create(ctx)
	if err != nil {
		errorExit("", err)
	}

	// Using G2Config: Add the CUSTOMERS data source to in-memory configuration.

	datasourceNames := []string{"CUSTOMERS"}
	for _, datasourceName := range datasourceNames {
		datasource := truthset.TruthsetDataSources[datasourceName]
		_, err := g2Config.AddDataSource(ctx, configHandle, datasource.Json)
		if err != nil {
			errorExit("", err)
		}
	}

	// Using G2Config: Persist configuration to a string.

	configStr, err := g2Config.Save(ctx, configHandle)
	if err != nil {
		errorExit("", err)
	}

	// Using G2Configmgr: Persist configuration string to database.

	configComments := fmt.Sprintf("Created by g2diagnostic_test at %s", now.UTC())
	configID, err := g2Configmgr.AddConfig(ctx, configStr, configComments)
	if err != nil {
		errorExit("", err)
	}

	// Using G2Configmgr: Set new configuration as the default.

	err = g2Configmgr.SetDefaultConfigID(ctx, configID)
	if err != nil {
		errorExit("", err)
	}

	// Using G2Engine: Update the G2Engine to use the new configuration.

	err = g2Engine.Reinit(ctx, configID)
	if err != nil {
		errorExit("", err)
	}
}

// Load the Senzing CUSTOMERS Truth Set.
func loadRecords(ctx context.Context, g2Engine g2api.G2engine) {
	loadId := "Test load"
	for _, record := range truthset.CustomerRecords {
		err := g2Engine.AddRecord(ctx, record.DataSource, record.Id, record.Json, loadId)
		if err != nil {
			errorExit("", err)
		}
	}
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	ctx := context.TODO()

	// Create an observer.

	myObserver := &observer.ObserverNull{
		Id: "My Observer",
	}

	// Create a Senzing G2Config, G2Configmgr, and G2Engine objects and register the observer.
	g2Config := &g2config.G2config{}
	g2Config.RegisterObserver(ctx, myObserver)

	g2Configmgr := &g2configmgr.G2configmgr{}
	g2Configmgr.RegisterObserver(ctx, myObserver)

	g2Engine := &g2engine.G2engine{}
	g2Engine.RegisterObserver(ctx, myObserver)

	// Load DataSources.

	loadDatasources(ctx, g2Config, g2Configmgr, g2Engine)

	// Load records.

	loadRecords(ctx, g2Engine)

	// Let observer finish.

	sleepTime := time.Duration(2)
	fmt.Printf("-------------------  Sleeping %d seconds for Observers -------------------\n", sleepTime)
	time.Sleep(sleepTime * time.Second)
	fmt.Printf("-------------------  Completed running mock Implementation --------------\n")
}
