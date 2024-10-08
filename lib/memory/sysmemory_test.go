/*
Copyright 2023 Huawei Cloud Computing Technologies Co., Ltd.

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

package memory

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadSysMemory(t *testing.T) {
	total, available := ReadSysMemory()
	MemUsedPct()
	readSysMemInfo(nil)
	require.NotEmpty(t, total)
	require.NotEmpty(t, available)
}

func BenchmarkReadSysMemory(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		MemUsedPct()
	}
}

func BenchmarkReadSysMemory_Parallel(b *testing.B) {
	b.ReportAllocs()
	var wg sync.WaitGroup
	f := func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			MemUsedPct()
		}
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go f()
	}
	wg.Wait()
}
