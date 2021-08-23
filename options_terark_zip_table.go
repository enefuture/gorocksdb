package gorocksdb

// #include "rocksdb/c.h"
import "C"

// BlockBasedTableOptions represents block-based table options.
type TerarkZipTableOptions struct {
	c *C.rocksdb_terark_zip_table_options_t
}

// NewDefaultOptions creates the default Options.
func NewTerarkZipTableOptions() *TerarkZipTableOptions {
	return NewNativeTerarkZipTableOptions(C.rocksdb_terark_zip_options_create())
}

func NewNativeTerarkZipTableOptions(c *C.rocksdb_terark_zip_table_options_t) *TerarkZipTableOptions {
	return &TerarkZipTableOptions{c: c}
}

// Destroy deallocates the BlockBasedTableOptions object.
func (opts *TerarkZipTableOptions) Destroy() {
	C.rocksdb_terark_zip_options_destroy(opts.c)
	opts.c = nil
}

func (opts *TerarkZipTableOptions) SetLocalTempDir(tempDir string) {
	C.rocksdb_terark_zip_options_set_local_temp_dir(opts.c, tempDir)
}

func (opts *TerarkZipTableOptions) SetIndexNestLevel(indexNestLevel int32) {
	C.rocksdb_terark_zip_options_set_index_nest_level(opts.c, C.int32_t(indexNestLevel))
}

func (opts *TerarkZipTableOptions) SetSampleRatio(sampleRatio float64) {
	C.rocksdb_terark_zip_options_set_sample_ratio(opts.c, C.double(sampleRatio))
}

func (opts *TerarkZipTableOptions) SetTerarkZipMinLevel(minLevel int) {
	C.rocksdb_terark_zip_options_set_terark_zip_min_level(opts.c, C.int(minLevel))
}
