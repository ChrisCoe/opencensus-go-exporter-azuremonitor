package main
// Package: Runs code for using Azure exporter

import (
	"context"
	"fmt"
	"log"
	
	"github.com/ChrisCoe/opencensus-go-exporter-azuremonitor/azuremonitor"
	"github.com/ChrisCoe/opencensus-go-exporter-azuremonitor/azuremonitor/common"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()

	exporter, err := azuremonitor.NewAzureTraceExporter(common.Options{
		InstrumentationKey: "111a0d2f-ab53-4b62-a54f-4722f09fd136", // add your InstrumentationKey
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done CREATING")
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
  
	_, span := trace.StartSpan(ctx, "/foo") // This calls the function ExportSpan written in azuremonitor.go 
	span.End()
}
