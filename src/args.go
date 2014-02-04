/**
 * Twist v0.1-dev
 *
 * (c) Harry Lawrence
 *
 * License: MIT
 * 
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package main

import (
	"os"
)

func FilenameIsSet() bool {
	if len(os.Args) < 2 {
		return false
	}
	return true
}

func GetFilenameArg() string {
	return os.Args[1]
}