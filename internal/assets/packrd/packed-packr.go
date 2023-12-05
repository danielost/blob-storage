// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "b3ea633ad296380d7cebac6c67fda319"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"ca71285196e32e1f7b766c17e610ab5c": "1f8b08000000000000ff4cceb18e82501046e17e9ee22f77b3cb1350a1d01130046a33c8042772b9e43288f8f426d2d09ed37c51843fa77d60133413d1b9ca923a439d9cf20cede0db193f0400daa1d57e96a03c600aea386c78c8f6ffbd4f1e1681c9cb5094358a26cff7710bc226dd950da64e6663376155bbfb652f78fb51e837263a5252bf8e4469555e8e9498e8130000ffff55519262b0000000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("migrations", "./migrations")
		b.SetResolver("001_initial.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "ca71285196e32e1f7b766c17e610ab5c"})
	}()
	return nil
}()
