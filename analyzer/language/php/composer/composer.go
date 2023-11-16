package composer

import (
	"context"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"

	"github.com/aquasecurity/go-dep-parser/pkg/php/composer"
	"github.com/khulnasoft-lab/fanal/analyzer"
	"github.com/khulnasoft-lab/fanal/analyzer/language"
	"github.com/khulnasoft-lab/fanal/types"
	"github.com/khulnasoft-lab/fanal/utils"
)

func init() {
	analyzer.RegisterAnalyzer(&composerLibraryAnalyzer{})
}

const version = 1

var requiredFiles = []string{types.ComposerLock}

type composerLibraryAnalyzer struct{}

func (a composerLibraryAnalyzer) Analyze(_ context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	res, err := language.Analyze(types.Composer, input.FilePath, input.Content, composer.NewParser())

	if err != nil {
		return nil, xerrors.Errorf("error with composer.lock: %w", err)
	}
	return res, nil
}

func (a composerLibraryAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	fileName := filepath.Base(filePath)
	return utils.StringInSlice(fileName, requiredFiles)
}

func (a composerLibraryAnalyzer) Type() analyzer.Type {
	return analyzer.TypeComposer
}

func (a composerLibraryAnalyzer) Version() int {
	return version
}
