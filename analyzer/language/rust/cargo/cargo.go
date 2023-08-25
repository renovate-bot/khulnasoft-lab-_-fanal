package cargo

import (
	"context"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"

	"github.com/khulnasoft-lab/fanal/analyzer"
	"github.com/khulnasoft-lab/fanal/analyzer/language"
	"github.com/khulnasoft-lab/fanal/types"
	"github.com/khulnasoft-lab/go-dep-parser/pkg/rust/cargo"
)

func init() {
	analyzer.RegisterAnalyzer(&cargoLibraryAnalyzer{})
}

const version = 1

type cargoLibraryAnalyzer struct{}

func (a cargoLibraryAnalyzer) Analyze(_ context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	res, err := language.Analyze(types.Cargo, input.FilePath, input.Content, cargo.NewParser())
	if err != nil {
		return nil, xerrors.Errorf("error with Cargo.lock: %w", err)
	}
	return res, nil
}

func (a cargoLibraryAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	fileName := filepath.Base(filePath)
	return fileName == types.CargoLock
}

func (a cargoLibraryAnalyzer) Type() analyzer.Type {
	return analyzer.TypeCargo
}

func (a cargoLibraryAnalyzer) Version() int {
	return version
}
