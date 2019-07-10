package main
// Package: Runs code for using Azure exporter

import (
	"context"
	"log"
	
	"github.com/opencensus-go-exporter-azuremonitor/azuremonitor"
	"go.opencensus.io/trace"
)

func main() {
	ctx := context.Background()

	exporter, err := azuremonitor.NewAzureTraceExporter("111a0d2f-ab53-4b62-a54f-4722f09fd136")
	if err != nil {
		log.Fatal(err)
	}

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
  
	_, span := trace.StartSpan(ctx, "/foo") // This calls the function ExportSpan written in azuremonitor.go 
	span.End()
}
