package core

import (
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/assert"
	"github.com/zokypesch/ceria/helper"
)

func TestElasticCore(t *testing.T) {
	type ExampleElasticTest struct {
		Name string
		Age  string
	}

	config := helper.NewReadConfigService()
	config.Init()

	host := config.GetByName("elastic.host")
	port := config.GetByName("elastic.port")
	stat := config.GetByName("elastic.status")

	if stat == "false" {
		return
	}

	hostElastic := "http://" + host + ":" + port
	elasticTest, errTest := NewServiceElasticCore(&ExampleElasticTest{}, hostElastic)

	if errTest != nil {
		return
	}

	t.Run("Tes remove the current index", func(t *testing.T) {

		assert.NoError(t, errTest)
		if errTest == nil {
			elasticTest.DeleteIndex()
		}

	})

	t.Run("Tes registration Elastic Failure", func(t *testing.T) {
		_, errNew := NewServiceElasticCore(5, "")
		assert.Error(t, errNew)

	})

	t.Run("Passing wrong parameter server", func(t *testing.T) {
		_, err := NewServiceElasticCore(&ExampleElasticTest{}, "http://192.68.1.1:9092")

		assert.Error(t, err)

	})

	t.Run("Tes registration Elastic Success", func(t *testing.T) {
		var exam ExampleElasticTest
		_, err := NewServiceElasticCore(&exam, hostElastic)

		assert.NoError(t, err)

	})

	// Add Document
	t.Run("Tes add document Elastic Failure", func(t *testing.T) {
		var exam ExampleElasticTest
		newElastic, errAssign := NewServiceElasticCore(&exam, hostElastic)
		assert.NoError(t, errAssign)

		err := newElastic.AddDocument("", &ExampleElasticTest{})
		assert.Error(t, err)

	})

	t.Run("Tes add document Elastic Success", func(t *testing.T) {
		var exam ExampleElasticTest
		newElastic, errAssign := NewServiceElasticCore(&exam, hostElastic)
		assert.NoError(t, errAssign)

		if newElastic != nil {
			err := newElastic.AddDocument("1", &ExampleElasticTest{"udin", "30"})
			assert.NoError(t, err)
		}

	})

	// Edit Document
	t.Run("Tes edit document Elastic Failure", func(t *testing.T) {
		var exam ExampleElasticTest
		newElastic, errAssign := NewServiceElasticCore(&exam, hostElastic)
		assert.NoError(t, errAssign)

		err := newElastic.EditDocument("", &ExampleElasticTest{})
		assert.Error(t, err)

	})

	t.Run("Tes edit document Elastic Success", func(t *testing.T) {
		var exam ExampleElasticTest
		newElastic, errAssign := NewServiceElasticCore(&exam, hostElastic)
		assert.NoError(t, errAssign)

		if newElastic != nil {
			err := newElastic.EditDocument("1", &ExampleElasticTest{"paijo", "30"})
			assert.NoError(t, err)
		}

	})

	// Delete Document
	t.Run("Tes delete document Elastic Failure", func(t *testing.T) {
		var exam ExampleElasticTest
		newElastic, err := NewServiceElasticCore(&exam, hostElastic)
		assert.NoError(t, err)

		errAssign := newElastic.DeleteDocument("")
		assert.Error(t, errAssign)

	})

	t.Run("Tes delete document Elastic Success", func(t *testing.T) {
		var exam ExampleElasticTest
		newElastic, err := NewServiceElasticCore(&exam, hostElastic)
		assert.NoError(t, err)

		if newElastic != nil {
			errAssign := newElastic.DeleteDocument("1")
			assert.NoError(t, errAssign)
		}

	})

	// Delete index
	t.Run("Tes delete index Elastic Success", func(t *testing.T) {
		var exam ExampleElasticTest
		newElastic, err := NewServiceElasticCore(&exam, hostElastic)
		assert.NoError(t, err)

		if newElastic != nil {
			errAssign := newElastic.DeleteIndex()
			assert.NoError(t, errAssign)
		}

	})

	t.Run("Tes delete index Elastic Failed", func(t *testing.T) {
		var exam ExampleElasticTest
		newElastic, err := NewServiceElasticCore(&exam, hostElastic)
		assert.NoError(t, err)

		if newElastic != nil {
			newElastic.Index = "Uknow"
			errAssign := newElastic.DeleteIndex()
			assert.Error(t, errAssign)
		}

	})

	//Tes MultipleinsertDocumentByStruct
	t.Run("Tes MultipleinsertDocumentByStruct index Elastic", func(t *testing.T) {

		type EmailCustom struct {
			Email string
		}

		type ExampleElasticTestMulti struct {
			gorm.Model
			Name   string
			Age    string
			Multi  []EmailCustom
			Single EmailCustom
		}

		exam := ExampleElasticTestMulti{
			Model: gorm.Model{ID: 1},
			Name:  "ajui",
			Age:   "30",
			Multi: []EmailCustom{
				EmailCustom{Email: "asd@gmail.com"},
				EmailCustom{Email: "asdfg@gmail.com"},
			},
			Single: EmailCustom{Email: "asdfghijk@gmail.com"},
		}

		newElastic, err := NewServiceElasticCore(&exam, hostElastic)
		assert.NoError(t, err)

		if newElastic != nil {
			errAssign := newElastic.MultipleinsertDocumentByStruct("1", &exam)
			newElastic.MultipleinsertDocumentByStruct("1", 5)
			assert.NoError(t, errAssign)
		}

	})

}
