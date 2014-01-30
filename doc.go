// Package osxcrossfix solves the common problem of not being
// able to access the system's root certificates when cross-compiling for darwin.
// The error usually encountered when trying this is
//
//     x509: failed to load system roots and no roots provided
//
// This problem is solved by bundling the most common signing authiorities'
// certificates into the binary and injecting them into the http.DefaultClient.
//
// The package has a build constraint and will be completely empty unless
// compiling for darwin while cgo is disabled.
// An underscore import
//
//    import _ "github.com/voxelbrain/osxcrossfix"
//
// is therefore always safe to do.
package osxcrossfix
