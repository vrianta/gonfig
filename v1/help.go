package structparser

/**
 * Code here will help to do auto print of arguments the application supports
 */

type help_record struct {
	field       string
	args        string
	env         string
	description string
	def         string
	required    bool
}

func help(help_records []help_record) {

}
