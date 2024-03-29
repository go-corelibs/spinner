// Copyright (c) 2024  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spinner

import (
	"sync"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("Manual", t, func() {
		s := New(nil, nil)
		So(s, ShouldNotEqual, nil)
		size := len(DefaultSymbols)
		for idx := 0; idx < size; idx++ {
			So(s.String(), ShouldEqual, DefaultSymbols[idx])
			s.Increment()
		}
		So(s.String(), ShouldEqual, DefaultSymbols[0])
		// hit the s.RUnlock() in the if s.ticker != nil else clause
		s.Stop()
	})

	Convey("Automatic", t, func(c C) {
		state := struct {
			i int
			m *sync.RWMutex
		}{
			i: 0,
			m: &sync.RWMutex{},
		}
		s := New(nil, func(symbol string) {
			state.m.Lock()
			defer state.m.Unlock()
			c.So(symbol, ShouldEqual, DefaultSymbols[state.i])
			state.i += 1
			if state.i >= len(DefaultSymbols) {
				state.i = 0
			}
		})
		So(s, ShouldNotEqual, nil)
		go s.Start()
		time.Sleep(time.Second)
		s.Stop()
		// check that it is actually stopped
		state.m.RLock()
		So(s.String(), ShouldEqual, DefaultSymbols[state.i])
		time.Sleep(time.Millisecond * 250)
		So(s.String(), ShouldEqual, DefaultSymbols[state.i])
		state.m.RUnlock()
	})
}
