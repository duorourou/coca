package tbs

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/adapter"
	"github.com/phodal/coca/core/adapter/call"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"path/filepath"
	"testing"
)

func TestTbsApp_EmptyTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/code/EmptyTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(result[0].Line).To(Equal(8))
	g.Expect(result[0].Type).To(Equal("EmptyTest"))
}

func TestTbsApp_IgnoreTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/code/IgnoreTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(len(result)).To(Equal(1))
	g.Expect(result[0].Line).To(Equal(0))
	g.Expect(result[0].Type).To(Equal("IgnoreTest"))
}

func TestTbsApp_RedundantPrintTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/code/RedundantPrintTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(result[0].Line).To(Equal(9))
	g.Expect(result[0].Type).To(Equal("RedundantPrintTest"))
}

func TestTbsApp_SleepyTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/code/SleepyTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(result[0].Line).To(Equal(8))
	g.Expect(result[0].Type).To(Equal("SleepyTest"))
}

func TestTbsApp_DuplicateAssertTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/code/DuplicateAssertTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(result[0].Line).To(Equal(23))
	g.Expect(result[0].Type).To(Equal("DuplicateAssertTest"))
}

func TestTbsApp_UnknownTest(t *testing.T) {
	g := NewGomegaWithT(t)
	codePath := "../../../_fixtures/tbs/code/UnknownTest.java"
	codePath = filepath.FromSlash(codePath)

	result := buildTbsResult(codePath)

	g.Expect(result[0].Type).To(Equal("EmptyTest"))
	g.Expect(result[0].Line).To(Equal(7))
	g.Expect(result[1].Type).To(Equal("UnknownTest"))
}

func buildTbsResult(codePath string) []TestBadSmell {
	files := support.GetJavaTestFiles(codePath)
	var identifiers []models.JIdentifier

	identifiers = adapter.LoadTestIdentify(files)
	identifiersMap := adapter.BuildIdentifierMap(identifiers)

	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	analysisApp := call.NewJavaCallApp()
	classNodes := analysisApp.AnalysisFiles(identifiers, files, classes)

	app := NewTbsApp()
	result := app.AnalysisPath(classNodes, identifiersMap)
	return result
}