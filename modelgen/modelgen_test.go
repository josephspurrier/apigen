package modelgen_test

import (
	"go/token"
	"io/ioutil"
	"log"
	"modelgen"
	"testing"
)

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.Lshortfile)
}

// TestPackage tests the package declaration
func TestPackage(t *testing.T) {
	// Create the package
	gt := modelgen.New("main")

	// Generate the code
	data, err := gt.Bytes(true)
	if err != nil {
		t.Error(err)
	}

	assertEqualFile(t, "testdata/t_package.go", string(data))
}

func TestImportSingle(t *testing.T) {
	// Create the package
	gt := modelgen.New("main")

	// Add import
	i := gt.Import()
	i.AddImport("log")
	gt.AddImport(i)

	// Add main func
	f := gt.FuncDecl("main")
	f.AddComment("// main")
	f.AddCallExpr("log.Println", "hello")
	gt.AddFunc(f)

	// Generate the code
	data, err := gt.Bytes(true)
	if err != nil {
		t.Error(err)
	}

	assertEqualFile(t, "testdata/t_import_single.go", string(data))
}

func TestImportDouble(t *testing.T) {
	// Create the package
	gt := modelgen.New("main")

	// Add import
	i := gt.Import()
	i.AddImport("fmt")
	i.AddImport("log")
	gt.AddImport(i)

	// Add main func
	f := gt.FuncDecl("main")
	f.AddComment("// main")
	f.AddCallExpr("log.Println", "hello")
	f.AddCallExpr("fmt.Println", "world")
	gt.AddFunc(f)

	// Generate the code
	data, err := gt.Bytes(true)
	if err != nil {
		t.Error(err)
	}

	assertEqualFile(t, "testdata/t_import_double.go", string(data))
}

func TestReturnNil(t *testing.T) {
	// Create the package
	gt := modelgen.New("main")

	// Add main func
	f := gt.FuncDecl("read")
	f.AddComment("// read")
	f.AddResult("", "error")
	f.AddReturnStmt(nil)
	gt.AddFunc(f)

	// Generate the code
	data, err := gt.Bytes(true)
	if err != nil {
		t.Error(err)
	}

	assertEqualFile(t, "testdata/t_return_nil.go", string(data))
}

func TestReturnErr(t *testing.T) {
	// Create the package
	gt := modelgen.New("main")

	// Add main func
	f := gt.FuncDecl("read")
	f.AddComment("// read")
	f.AddResult("", "error")
	//f.AddAssignStmt("var", "string")
	f.AddDeclStmt("err", "error")
	f.AddReturnStmt(modelgen.Variable("err"))
	gt.AddFunc(f)

	// Generate the code
	data, err := gt.Bytes(true)
	if err != nil {
		t.Error(err)
	}

	assertEqualFile(t, "testdata/t_return_err.go", string(data))
}

func TestGenerateModel(t *testing.T) {
	// Create the package
	gt := modelgen.New("model")

	//gt.AddComment("// Package model handles the loading of models.")

	// Add import
	i := gt.Import()
	i.AddImport("github.com/blue-jay/blueprint/model/note")
	i.AddImport("github.com/blue-jay/blueprint/model/user")
	i.AddImport("github.com/jmoiron/sqlx")
	gt.AddImport(i)

	// Add var
	v := gt.Var()
	v.AddVar("Note", "note.Service", "// Note model")
	v.AddVar("User", "user.Service", "// User model")
	gt.AddVar(v)

	// Add main func
	f := gt.FuncDecl("Load")
	f.AddComment("// Load injects the dependencies for the models")
	f.AddParam("db", "*sqlx.DB")
	f.AddAssignStmt(modelgen.Ident("Note"), token.ASSIGN, "note.Service{db}")
	f.AddAssignStmt(modelgen.Ident("User"), token.ASSIGN, "user.Service{db}")
	gt.AddFunc(f)

	// Generate the code
	data, err := gt.Bytes(true)
	if err != nil {
		t.Error(err)
	}

	final := string(data)

	// Hack to include a package comment at the top
	final = "// Package model handles the loading of models.\n" + final

	assertEqualFile(t, "testdata/t_model.go", final)
}

func zTestAddModel(t *testing.T) {
	// Create the package
	//gt := modelgen.New("model")

	b, err := ioutil.ReadFile("testdata/t_model_add.go")
	if err != nil {
		t.Error(err)
	}

	gt, err := modelgen.ParseFile(string(b))
	if err != nil {
		t.Error(err)
	}

	// Add import
	gt.AddNewImport("github.com/blue-jay/blueprint/model/user")

	// Add var
	gt.AddNewVar("User", "user.Service", "// User model")

	// Add main func
	//f.AddAssignStmt(modelgen.Ident("User"), token.ASSIGN, "user.Service{db}")

	//gt.AddNewAssignStmt("Load", modelgen.Ident("User"), token.ASSIGN, "user.Service{db}")

	// Generate the code
	data, err := gt.Bytes(true)
	if err != nil {
		t.Error(err)
	}

	// Hack to include a package comment at the top
	//addComment := "// Package model handles the loading of models.\n" + string(data)

	//assertEqualFile(t, "testdata/t_model.go", string(data))
	assertEqual(t, "nothing", string(data))
}

func zTestStruct(t *testing.T) {

	b, err := ioutil.ReadFile("testdata/t_model_add.go")
	if err != nil {
		t.Error(err)
	}

	gt, err := modelgen.ParseFile(string(b))
	if err != nil {
		t.Error(err)
	}

	// Copy the struct fields
	sc, err := gt.ReadVar()
	if err != nil {
		log.Println(err)
		return
	}

	// Replace the struct
	err = gt.WriteVar(sc)
	if err != nil {
		log.Println(err)
		return
	}

	// Generate the code
	data, err := gt.Bytes(true)
	if err != nil {
		t.Error(err)
	}

	// Hack to include a package comment at the top
	//addComment := "// Package model handles the loading of models.\n" + string(data)

	assertEqualFile(t, "testdata/t_model_add.go", string(data))
}

// *****************************************************************************
// Helpers
// *****************************************************************************

// assertEqual easily tests expected values
func assertEqual(t *testing.T, expectedValue interface{}, actualValue interface{}) {
	if actualValue != expectedValue {
		t.Errorf("\n got: %v\nwant: %v", actualValue, expectedValue)
	}
}

// assertEqualFile easily tests actual value against file
func assertEqualFile(t *testing.T, expectedFile string, actualValue interface{}) {
	expectedValueByte, err := ioutil.ReadFile(expectedFile)
	if err != nil {
		t.Error(err)
	}

	expectedValue := string(expectedValueByte)

	if actualValue != expectedValue {
		t.Errorf("\n got: %v\nwant: %v", actualValue, expectedValue)
	}
}
