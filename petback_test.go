package petback
import (
	"fmt"
	"testing"
)

func TestGCHandlerFunc(t *testing.T) {
	data := GCHandlerFunc("string", "GIS", "tesgis")

	fmt.Printf("%+v", data)
}
