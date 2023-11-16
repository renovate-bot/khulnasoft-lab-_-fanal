package artifact

import (
	"context"
	"sort"

	"github.com/khulnasoft-lab/fanal/analyzer"
	misconf "github.com/khulnasoft-lab/fanal/analyzer/config"
	"github.com/khulnasoft-lab/fanal/analyzer/licensing"
	"github.com/khulnasoft-lab/fanal/analyzer/secret"
	"github.com/khulnasoft-lab/fanal/types"
)

type Option struct {
	AnalyzerGroup     analyzer.Group // It is empty in OSS
	DisabledAnalyzers []analyzer.Type
	DisabledHandlers  []types.HandlerType
	SkipFiles         []string
	SkipDirs          []string
	NoProgress        bool
	Offline           bool
	InsecureSkipTLS   bool

	MisconfScannerOption misconf.ScannerOption
	SecretScannerOption  secret.ScannerOption
	LicenseScannerOption licensing.ScannerOption
}

func (o *Option) Sort() {
	sort.Slice(o.DisabledAnalyzers, func(i, j int) bool {
		return o.DisabledAnalyzers[i] < o.DisabledAnalyzers[j]
	})
	sort.Strings(o.SkipFiles)
	sort.Strings(o.SkipDirs)
}

type Artifact interface {
	Inspect(ctx context.Context) (reference types.ArtifactReference, err error)
	Clean(reference types.ArtifactReference) error
}
