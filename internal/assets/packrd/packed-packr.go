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
		"b4d945e7ae870935c8d3193b2db2cb13": "1f8b08000000000000ffd2d555d0cecd4c2f4a2c4955082de0e272f409710d52087174f2715548cac94f2a56808838fbfb84fafa299425e694a62a844406b82a6415e7e7252984067bfab94384adacc042d65c5cc886bae497e7717101020000ffffd7d02d5d67000000",
		"ca71285196e32e1f7b766c17e610ab5c": "1f8b08000000000000ff9c92414f83401484effb2be658a23dea8553b5a8240d3585c6dec802cf7623ec92b74bb1fe7a53b655693c18afcccb3733cc4ea7b86ad496a523ac5b21ee57d12c8b90cdee16113a4b6c311100a02a146a6b8995acd1b26a241ff04687eb41adcd5669ec25973bc9939bdb00c93243b25e2cbcde4a6b7bc3151cbdbb0bad64928eaa5c3a38d59075b269d12bb7339dff820fa34904e145baa236c5dfd2ed65ddd1ffadfdade93571ee8d9476d0c64177750da65762d225d9f30f535500a351514d8e504a5bca6a54204ee6d1c617c83db734da3a9647f03239555ba771f288c231112667fb002f4fd12afa8e13a75fa5c231ff98261f96b9c0fb9823fc707666fb354760f1f3a1cc4daf8598af96cfa729e207449b38cd529f3cfc5d1b6c43f1190000ffffdd14254973020000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("migrations", "./migrations")
		b.SetResolver("001_initial.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "ca71285196e32e1f7b766c17e610ab5c"})
		b.SetResolver("002_blob_jsonb.sql", packr.Pointer{ForwardBox: gk, ForwardPath: "b4d945e7ae870935c8d3193b2db2cb13"})
	}()
	return nil
}()
