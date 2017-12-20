// Vision Screening Upload Simulator
// Copyright (C) 2017 Andrew Allen
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"log"
	"os"

	"github.com/achew22/acceptance-testing-vision-upload-server/server"
)

var port = flag.Int("port", 9000, "The port to run the simulator on")

func main() {
	flag.Parse()

	l := log.New(os.Stderr, "", log.LUTC|log.LstdFlags)

	s := server.New(l, *port)
	s.Run()
}
