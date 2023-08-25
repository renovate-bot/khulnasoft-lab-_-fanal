package binary

import (
	"context"
	"errors"
	"os"

	"golang.org/x/xerrors"

	"github.com/khulnasoft-lab/fanal/analyzer"
	"github.com/khulnasoft-lab/fanal/analyzer/language"
	"github.com/khulnasoft-lab/fanal/types"
	"github.com/khulnasoft-lab/go-dep-parser/pkg/golang/binary"
)

func init() {
	analyzer.RegisterAnalyzer(&gobinaryLibraryAnalyzer{})
}

const version = 1

type gobinaryLibraryAnalyzer struct{}

func (a gobinaryLibraryAnalyzer) Analyze(_ context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	p := binary.NewParser()
	libs, deps, err := p.Parse(input.Content)
	if errors.Is(err, binary.ErrUnrecognizedExe) || errors.Is(err, binary.ErrNonGoBinary) {
		return nil, nil
	} else if err != nil {
		return nil, xerrors.Errorf("go binary parse error: %w", err)
	}

	return language.ToAnalysisResult(types.GoBinary, input.FilePath, "", libs, deps), nil
}

func (a gobinaryLibraryAnalyzer) Required(_ string, fileInfo os.FileInfo) bool {
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		return false
	}

	// Check executable file
	if mode.Perm()&0111 != 0 {
		return true
	}
	return false
}

func (a gobinaryLibraryAnalyzer) Type() analyzer.Type {
	return analyzer.TypeGoBinary
}

func (a gobinaryLibraryAnalyzer) Version() int {
	return version
}
