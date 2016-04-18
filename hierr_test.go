package hierr

import "errors"
import "fmt"

func ExampleError() {
	testcases := []error{
		Errorf(nil, ""),
		Errorf(nil, "simple error"),
		Errorf(nil, "integer: %d", 1),
		Errorf(errors.New("nested"), "top level"),
		Errorf(errors.New("nested"), "top level: %s", "formatting"),
		Errorf(Errorf(errors.New("low level"), "nested"), "top level"),

		Errorf(Errorf("string error", "nested"), "top level"),
		Errorf([]byte("byte"), "top level"),
	}

	for _, test := range testcases {
		fmt.Println()
		fmt.Println("{{{")
		fmt.Println(test.Error())
		fmt.Println("}}}")
	}

	fmt.Println()

	exiter = func(code int) {
		fmt.Println("exit code:", code)
	}

	Fatalf("wow", "critical error")

	// Output:
	//
	// {{{
	//
	// }}}
	//
	// {{{
	// simple error
	// }}}
	//
	// {{{
	// integer: 1
	// }}}
	//
	// {{{
	// top level
	// └─ nested
	// }}}
	//
	// {{{
	// top level: formatting
	// └─ nested
	// }}}
	//
	// {{{
	// top level
	// └─ nested
	//    └─ low level
	// }}}
	//
	// {{{
	// top level
	// └─ nested
	//    └─ string error
	// }}}
	//
	// {{{
	// top level
	// └─ byte
	// }}}
	//
	// critical error
	// └─ wow
	// exit code: 1
}

func ExampleError_Error() {
	BranchDelimiter = "* "
	BranchIndent = 0

	testcases := []error{
		Errorf(Errorf(errors.New("third"), "second"), "top level"),
	}

	for _, test := range testcases {
		fmt.Println()
		fmt.Println("{{{")
		fmt.Println(test.Error())
		fmt.Println("}}}")
	}

	// Output:
	//
	// {{{
	// top level
	// * second
	// * third
	// }}}
}
