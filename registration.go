//just a little code to check how much recycled material goes into the container

//registration
import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	errDuplicateName = "duplicate gadget name: %s"
	errReservedName  = "can't register a gadget of reserved name %s"
	errNoHeader      = "pylons.gadgets file does not start with a valid gadget header: \n%s"
	errBadHeader     = "not a valid gadget header: \n%s"


  	t.Run("Recipe", func(t *testing.T) {
		args := []string{"NewUser0", goodPLR}
		_, err := clitestutil.ExecTestCLICmd(ctx, DevCreate(), args)
		assert.Nil(t, err)
		args = []string{"NewUser0", strings.Replace(goodPLR, "0.0.1", "0.0.2", 1)}
		cmd := DevUpdate()
		_, err = clitestutil.ExecTestCLICmd(ctx, cmd, args)
		assert.Nil(t, err)
	})
