/*
 * Copyright 1999-2018 Alibaba Group.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package downloader contains 2 types of downloader: P2PDownloader,
// DirectDownloader.
// P2PDownloader uses P2P pattern to download files from peers.
// DirectDownloader downloads files from file source directly. It's
// used when P2PDownloader download files failed.
package downloader

// Downloader is the interface to download files
type Downloader interface {
	Run()
}
