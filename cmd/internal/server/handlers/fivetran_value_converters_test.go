package handlers

import (
	"testing"

	fivetransdk "github.com/planetscale/fivetran-source/fivetran_sdk.v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"vitess.io/vitess/go/sqltypes"
)

func TestJSONConverter_EmptyString(t *testing.T) {
	converter, err := GetConverter(fivetransdk.DataType_JSON)
	require.NoError(t, err)

	// Empty string should convert to NULL for BigQuery compatibility
	emptyValue := sqltypes.NewVarChar("")
	result, err := converter(emptyValue)
	require.NoError(t, err)
	require.NotNil(t, result)

	_, ok := result.Inner.(*fivetransdk.ValueType_Null)
	assert.True(t, ok, "empty JSON string should convert to NULL")
}

func TestJSONConverter_ValidJSON(t *testing.T) {
	converter, err := GetConverter(fivetransdk.DataType_JSON)
	require.NoError(t, err)

	// Valid JSON should pass through
	jsonValue := sqltypes.NewVarChar(`{"key": "value"}`)
	result, err := converter(jsonValue)
	require.NoError(t, err)
	require.NotNil(t, result)

	json, ok := result.Inner.(*fivetransdk.ValueType_Json)
	require.True(t, ok)
	assert.Equal(t, `{"key": "value"}`, json.Json)
}
