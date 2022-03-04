package geometry

import "errors"

func CubeVolumn(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	} else {
		return 0, errors.New("zero length edge is now allowed")
	}
}
