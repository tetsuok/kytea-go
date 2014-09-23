/*
Copyright 2014 Tetsuo Kiso

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package kytea provides access to KyTea.
//
// NOTE: This package calls the C APIs of a forked KyTea
// (https://github.com/tetsuok/kytea). This is because the original
// Kytea provides useful C++ APIs, but C APIs are not available.
//
// BUG: This package is in incomplete. Until APIs of C APIs and Go are fixed,
// the compatiblities are not guaranteed.
package kytea

/*
#cgo LDFLAGS: -lkytea -lstdc++

#include "kytea/c.h"
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

type Tagger struct {
	rep    *C.kytea_t
	config *C.kytea_config_t
}

type StringUtil struct {
	rep *C.kytea_stringutil_t
}

type Sentence struct {
	rep *C.kytea_sentence_t
}

// Create creates a Tagger object
// It is the caller's responsiblity to delete the tagger when it is done with it.
func Create(filename string) (*Tagger, error) {
	var rep *C.kytea_t
	name := C.CString(filename)
	defer C.free(unsafe.Pointer(name))
	rep = C.kytea_create()
	if rep == nil {
		return nil, errors.New("Cannot create a KyTea object")
	}
	if ret := C.kytea_read_model(rep, name); ret < 0 {
		return nil, errors.New("Cannot read the model")
	}

	config := C.kytea_get_config(rep)
	if config == nil {
		return nil, errors.New("Cannot get a KyTea config object")
	}
	return &Tagger{rep: rep, config: config}, nil
}

// Destroy frees the allocated C object.
func (t *Tagger) Destroy() {
	C.kytea_destroy(t.rep)
	C.kytea_config_destroy(t.config)
	t.rep = nil
	t.config = nil
}

func (t *Tagger) CalculateWS(s *Sentence) {
	C.kytea_calculate_ws(t.rep, s.rep)
}

func (t *Tagger) CalculateAllTags(s *Sentence) {
	i := C.int(0)
	for i = 0; i < C.kytea_config_get_num_tags(t.config); i++ {
		C.kytea_calculate_tags(t.rep, s.rep, i)
	}
}

// CreateStringUtil creates a StringUtil object.  It is the caller's
// responsiblity to delete the StringUtil object when it is done with
// it.
func (t *Tagger) CreateStringUtil() (*StringUtil, error) {
	rep := C.kytea_get_stringutil(t.rep)
	if rep == nil {
		return nil, errors.New("Cannot get a KyTea stringutil object")
	}
	return &StringUtil{rep}, nil
}

func (u *StringUtil) Destroy() {
	C.kytea_stringutil_destroy(u.rep)
	u.rep = nil
}

// CreateSentence creates a Sentence object.  It is the caller's
// responsiblity to delete the sentence when it is done with it.
func (u *StringUtil) CreateSentence(s string) (*Sentence, error) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	rep := C.kytea_sentence_create(u.rep, cs)
	if rep == nil {
		return nil, errors.New("Cannot create a sentence object")
	}
	return &Sentence{rep}, nil
}

func (s *Sentence) Destroy() {
	C.kytea_sentence_destroy(s.rep)
	s.rep = nil
}

// NumWords returns the number of tokens. This method needs to be called
// after tagger.CalculateWS() is called.
func (s *Sentence) NumWords() int {
	return int(C.kytea_sentence_get_num_words(s.rep))
}

// SurfaceAt returns the surface form at i in the tokenized sentence.
func (u *StringUtil) SurfaceAt(s *Sentence, i int) string {
	cs := C.kytea_sentence_surface_at(u.rep, s.rep, C.int(i))
	defer C.free(unsafe.Pointer(cs))
	return C.GoString(cs)
}

// ReadingAt returns the reading form at i in the tokenized sentence.
// This method needs to be called after tagger.CalculateAllTags() is called.
func (u *StringUtil) ReadingAt(s *Sentence, i int) string {
	cs := C.kytea_sentence_reading_at(u.rep, s.rep, C.int(i))
	defer C.free(unsafe.Pointer(cs))
	return C.GoString(cs)
}
